import logging
import grpc
import requests
from flask import Flask, request, jsonify
from functools import wraps
from threading import Lock
from scheduler_svc import scheduler_pb2
from scheduler_svc import scheduler_pb2_grpc
from staff_svc import staff_records_pb2
from staff_svc import staff_records_pb2_grpc
from google.protobuf.json_format import MessageToDict
from expiringdict import ExpiringDict

app = Flask(__name__)

serv_discovery = "http://localhost:5051/route"
cache = ExpiringDict(max_len=100, max_age_seconds=10)
lock = Lock()
svc_stub = {}

def cache_req(f):
    @wraps(f)
    def cache_wrapper(*args):
        key = str(f)
        with lock:
            global cache
            result = cache.get(key)
            if result is None:
                result = f()
                cache[key] = result
            return result
    return cache_wrapper


def init():
    global svc_stub
    try:
        response = requests.get(serv_discovery, timeout=0.5)
    except requests.exceptions.Timeout:
        response = "service discovery timed out"
        return response
    services = response.json()
    for service in services:
        channel = grpc.insecure_channel(services[service])
        if service == "scheduler_svc":
            svc_stub[service] = scheduler_pb2_grpc.SchedulerStub(channel)
        elif service == "staff_svc":
            svc_stub[service] = staff_records_pb2_grpc.StaffRecordsStub(channel)


@app.route("/appointment", methods=['GET'])
@cache_req
def get_appointments():
    try:
        # print(request.endpoint)
        response = svc_stub["scheduler_svc"].GetAppt(scheduler_pb2.GetAppointments(), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'get appointments timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/appointment", methods=['POST'])
def create_appointment():
    data = request.get_json()
    appt_data = data["appointment"]
    try:
        response = svc_stub["scheduler_svc"].CreateAppt(scheduler_pb2.CreateAppointment(appointment={
            "patientName": appt_data["patientName"],
            "doctorName": appt_data["doctorName"],
            "apptDateTime": appt_data["apptDateTime"],
            "apptType": appt_data["apptType"]
        }), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'create appointment timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/appointment", methods=['PUT'])
def update_appointment():
    data = request.get_json()
    appt_data = data["appointment"]
    try:
        response = svc_stub["scheduler_svc"].UpdateAppt(scheduler_pb2.UpdateAppointment(appointment={
            "apptID": appt_data["apptID"],
            "patientName": appt_data["patientName"],
            "doctorName": appt_data["doctorName"],
            "apptDateTime": appt_data["apptDateTime"],
            "apptType": appt_data["apptType"]
        }), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'update appointment timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/appointment", methods=['DELETE'])
def delete_appointment():
    data = request.get_json()
    req_id = data["apptID"]
    try:
        response = svc_stub["scheduler_svc"].DeleteAppt(scheduler_pb2.DeleteAppointment(apptID=req_id), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'delete appointment timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/staff", methods=['GET'])
@cache_req
def get_staff():
    try:
        response = svc_stub["staff_svc"].Get(staff_records_pb2.GetStaffRecords(), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'get staff timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/staff/availability", methods=['GET'])
def get_staff_availability():
    data = request.get_json()
    req_id = data["staffID"]
    try:
        response = svc_stub["staff_svc"].GetAvailability(staff_records_pb2.GetStaffAvailability(staffID=req_id), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'get staff availability timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/staff", methods=['POST'])
def create_staff():
    data = request.get_json()
    staff_data = data["staffRecord"]
    try:
        response = svc_stub["staff_svc"].Create(staff_records_pb2.CreateStaff(staff={
            "name": staff_data["name"],
            "jobTitle": staff_data["jobTitle"],
            "department": staff_data["department"],
            "isAvailable": staff_data["isAvailable"]
        }), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'create staff timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/staff", methods=['PUT'])
def update_staff():
    data = request.get_json()
    staff_data = data["staffRecord"]
    print(staff_data["isAvailable"])
    print(type(staff_data["isAvailable"]))
    try:
        response = svc_stub["staff_svc"].Update(staff_records_pb2.UpdateStaff(staffRecord={
            "staffID": staff_data["staffID"],
            "name": staff_data["name"],
            "jobTitle": staff_data["jobTitle"],
            "department": staff_data["department"],
            "isAvailable": staff_data["isAvailable"]
        }), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'update staff timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/staff", methods=['DELETE'])
def delete_staff():
    data = request.get_json()
    req_id = data["staffID"]
    try:
        response = svc_stub["staff_svc"].Delete(staff_records_pb2.DeleteStaff(staffID=req_id), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'delete staff timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/health", methods=['GET'])
def get_health():
    resp = jsonify(status="SERVING")
    resp.status_code = 200
    return resp


if __name__ == "__main__":
    logging.basicConfig()
    init()
    app.run(host='localhost', port=5050)
