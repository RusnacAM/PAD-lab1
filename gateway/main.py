import logging
import grpc
import requests
import os
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
from circuit_breaker import circuit_breaker

app = Flask(__name__)

serv_discovery = "http://service-discovery:5051/route"
cache = ExpiringDict(max_len=100, max_age_seconds=10)
limiter = Limiter(
    get_remote_address,
    app=app,
    default_limits=["5/second", "100 per hour", "250 per day"]

)
lock = Lock()

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

prev_svc = 0
def get_svc(svc_type):
    try:
        response = requests.get(serv_discovery, timeout=float(os.environ["TIMEOUT"]))
    except requests.exceptions.Timeout:
        response = "service discovery timed out"
        return response
    services = response.json()
    svc_stub = ""
    if svc_type in services:
        global prev_svc
        svc_list = services[svc_type]
        curr_svc = prev_svc % len(svc_list)
        channel = grpc.insecure_channel(svc_list[curr_svc])
        
        if svc_type == "scheduler_svc":
            svc_stub = scheduler_pb2_grpc.SchedulerStub(channel)
        elif svc_type == "staff_svc":
            svc_stub = staff_records_pb2_grpc.StaffRecordsStub(channel)

        prev_svc += 1
    return svc_stub


@app.route("/appointment", methods=['GET'])
@circuit_breaker
@limiter.limit("10 per minute", override_defaults=False)
@cache_req
def get_appointments():
    try:
        stub = get_svc("scheduler_svc")
        response = stub.GetAppt(scheduler_pb2.GetAppointments(), timeout=float(os.environ["TIMEOUT"]))
    except grpc.RpcError as e:
        raise grpc.RpcError
    return MessageToDict(response)

@app.route("/appointment", methods=['POST'])
@circuit_breaker
@limiter.limit("10 per minute", override_defaults=False)
def create_appointment():
    data = request.get_json()
    appt_data = data["appointment"]
    try:
        stub = get_svc("scheduler_svc")
        response = stub.CreateAppt(scheduler_pb2.CreateAppointment(appointment={
            "patientName": appt_data["patientName"],
            "staffID": appt_data["staffID"],
            "apptDateTime": appt_data["apptDateTime"],
            "apptType": appt_data["apptType"]
        }), timeout=float(os.environ["TIMEOUT"]))
    except grpc.RpcError as e:
        raise grpc.RpcError
    return MessageToDict(response)

@app.route("/appointment", methods=['PUT'])
@circuit_breaker
@limiter.limit("10 per minute", override_defaults=False)
def update_appointment():
    data = request.get_json()
    appt_data = data["appointment"]
    try:
        stub = get_svc("scheduler_svc")
        response = stub.UpdateAppt(scheduler_pb2.UpdateAppointment(appointment={
            "apptID": appt_data["apptID"],
            "patientName": appt_data["patientName"],
            "staffID": appt_data["staffID"],
            "apptDateTime": appt_data["apptDateTime"],
            "apptType": appt_data["apptType"]
        }), timeout=float(os.environ["TIMEOUT"]))
    except grpc.RpcError as e:
        raise grpc.RpcError
    return MessageToDict(response)

@app.route("/appointment", methods=['DELETE'])
@circuit_breaker
@limiter.limit("10 per minute", override_defaults=False)
def delete_appointment():
    data = request.get_json()
    req_id = data["apptID"]
    try:
        stub = get_svc("scheduler_svc")
        response = stub.DeleteAppt(scheduler_pb2.DeleteAppointment(apptID=req_id), timeout=float(os.environ["TIMEOUT"]))
    except grpc.RpcError as e:
        raise grpc.RpcError
    return MessageToDict(response)

@app.route("/staff", methods=['GET'])
@circuit_breaker
@limiter.limit("10 per minute", override_defaults=False)
@cache_req
def get_staff():
    try:
        stub = get_svc("staff_svc")
        response = stub.Get(staff_records_pb2.GetStaffRecords(), timeout=float(os.environ["TIMEOUT"]))
    except grpc.RpcError as e:
        raise grpc.RpcError
    return MessageToDict(response)

@app.route("/staff/availability", methods=['GET'])
@circuit_breaker
@limiter.limit("10 per minute", override_defaults=False)
def get_staff_availability():
    data = request.get_json()
    req_id = data["staffID"]
    try:
        stub = get_svc("staff_svc")
        response = stub.GetAvailability(staff_records_pb2.GetStaffAvailability(staffID=req_id), timeout=float(os.environ["TIMEOUT"]))
    except grpc.RpcError as e:
        raise grpc.RpcError
    return MessageToDict(response)

@app.route("/staff", methods=['POST'])
@circuit_breaker
@limiter.limit("10 per minute", override_defaults=False)
def create_staff():
    data = request.get_json()
    staff_data = data["staff"]
    try:
        stub = get_svc("staff_svc")
        response = stub.Create(staff_records_pb2.CreateStaff(staff={
            "name": staff_data["name"],
            "jobTitle": staff_data["jobTitle"],
            "department": staff_data["department"],
            "isAvailable": staff_data["isAvailable"]
        }), timeout=float(os.environ["TIMEOUT"]))
    except grpc.RpcError as e:
        raise grpc.RpcError
    return MessageToDict(response)

@app.route("/staff", methods=['PUT'])
@circuit_breaker
@limiter.limit("10 per minute", override_defaults=False)
def update_staff():
    data = request.get_json()
    staff_data = data["staff"]
    print(staff_data["isAvailable"])
    print(type(staff_data["isAvailable"]))
    try:
        stub = get_svc("staff_svc")
        response = stub.Update(staff_records_pb2.UpdateStaff(staffRecord={
            "staffID": staff_data["staffID"],
            "name": staff_data["name"],
            "jobTitle": staff_data["jobTitle"],
            "department": staff_data["department"],
            "isAvailable": staff_data["isAvailable"]
        }), timeout=float(os.environ["TIMEOUT"]))
    except grpc.RpcError as e:
        raise grpc.RpcError
    return MessageToDict(response)

@app.route("/staff", methods=['DELETE'])
@circuit_breaker
@limiter.limit("10 per minute", override_defaults=False)
def delete_staff():
    data = request.get_json()
    req_id = data["staffID"]
    try:
        stub = get_svc("staff_svc")
        response = stub.Delete(staff_records_pb2.DeleteStaff(staffID=req_id), timeout=float(os.environ["TIMEOUT"]))
    except grpc.RpcError as e:
        raise grpc.RpcError
    return MessageToDict(response)

@app.route("/health", methods=['GET'])
def get_health():
    resp = jsonify(status="SERVING")
    resp.status_code = 200
    return resp

@app.errorhandler(429)
def ratelimit_handler(e):
  return "Too many requests. You have exceeded your rate-limit.", 429

if __name__ == "__main__":
    logging.basicConfig()
    app.run(host='0.0.0.0', port=5050)
