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
from flask_limiter import Limiter
from flask_limiter.util import get_remote_address
from google.protobuf.json_format import MessageToDict
from expiringdict import ExpiringDict

app = Flask(__name__)

serv_discovery = "http://service-discovery:5051/route"
cache = ExpiringDict(max_len=100, max_age_seconds=10)
limiter = Limiter(
    get_remote_address,
    app=app,
    default_limits=["5/second", "100 per hour", "250 per day"]

)
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


def get_svc():
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
    print(svc_stub)


@app.route("/appointment", methods=['GET'])
@limiter.limit("10 per minute", override_defaults=False)
@cache_req
def get_appointments():
    try:
        get_svc()
        response = svc_stub["scheduler_svc"].GetAppt(scheduler_pb2.GetAppointments(), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'get appointments timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/appointment", methods=['POST'])
@limiter.limit("10 per minute", override_defaults=False)
def create_appointment():
    data = request.get_json()
    appt_data = data["appointment"]
    try:
        get_svc()
        response = svc_stub["scheduler_svc"].CreateAppt(scheduler_pb2.CreateAppointment(appointment={
            "patientName": appt_data["patientName"],
            "staffID": appt_data["staffID"],
            "apptDateTime": appt_data["apptDateTime"],
            "apptType": appt_data["apptType"]
        }), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'create appointment timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/appointment", methods=['PUT'])
@limiter.limit("10 per minute", override_defaults=False)
def update_appointment():
    data = request.get_json()
    appt_data = data["appointment"]
    try:
        get_svc()
        response = svc_stub["scheduler_svc"].UpdateAppt(scheduler_pb2.UpdateAppointment(appointment={
            "apptID": appt_data["apptID"],
            "patientName": appt_data["patientName"],
            "staffID": appt_data["staffID"],
            "apptDateTime": appt_data["apptDateTime"],
            "apptType": appt_data["apptType"]
        }), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'update appointment timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/appointment", methods=['DELETE'])
@limiter.limit("10 per minute", override_defaults=False)
def delete_appointment():
    data = request.get_json()
    req_id = data["apptID"]
    try:
        get_svc()
        response = svc_stub["scheduler_svc"].DeleteAppt(scheduler_pb2.DeleteAppointment(apptID=req_id), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'delete appointment timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/staff", methods=['GET'])
@limiter.limit("10 per minute", override_defaults=False)
@cache_req
def get_staff():
    try:
        get_svc()
        response = svc_stub["staff_svc"].Get(staff_records_pb2.GetStaffRecords(), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'get staff timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/staff/availability", methods=['GET'])
@limiter.limit("10 per minute", override_defaults=False)
def get_staff_availability():
    data = request.get_json()
    req_id = data["staffID"]
    try:
        get_svc()
        response = svc_stub["staff_svc"].GetAvailability(staff_records_pb2.GetStaffAvailability(staffID=req_id), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'get staff availability timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/staff", methods=['POST'])
@limiter.limit("10 per minute", override_defaults=False)
def create_staff():
    data = request.get_json()
    staff_data = data["staff"]
    try:
        get_svc()
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
@limiter.limit("10 per minute", override_defaults=False)
def update_staff():
    data = request.get_json()
    staff_data = data["staff"]
    print(staff_data["isAvailable"])
    print(type(staff_data["isAvailable"]))
    try:
        get_svc()
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
@limiter.limit("10 per minute", override_defaults=False)
def delete_staff():
    data = request.get_json()
    req_id = data["staffID"]
    try:
        get_svc()
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

@app.errorhandler(429)
def ratelimit_handler(e):
  return "Too many requests. You have exceeded your rate-limit."

if __name__ == "__main__":
    logging.basicConfig()
    app.run(host='0.0.0.0', port=5050)
