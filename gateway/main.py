import logging
import grpc
from flask import Flask, request, jsonify
from scheduler_svc import scheduler_pb2
from scheduler_svc import scheduler_pb2_grpc
from staff_svc import staff_records_pb2
from staff_svc import staff_records_pb2_grpc
from google.protobuf.json_format import MessageToDict

app = Flask(__name__)

scheduler_stub = ""
staff_stub = ""

def init():
    scheduler_channel = grpc.insecure_channel('localhost:3030')
    staff_channel = grpc.insecure_channel('localhost:8080')
    global scheduler_stub 
    global staff_stub
    scheduler_stub = scheduler_pb2_grpc.SchedulerStub(scheduler_channel)
    staff_stub = staff_records_pb2_grpc.StaffRecordsStub(staff_channel)

@app.route("/appointment", methods=['GET'])
def get_appointments():
    try:
        response = scheduler_stub.GetAppt(scheduler_pb2.GetAppointments(), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'get appointments timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/appointment", methods=['POST'])
def create_appointment():
    data = request.get_json()
    appt_data = data["appointment"]
    try:
        response = scheduler_stub.CreateAppt(scheduler_pb2.CreateAppointment(appointment={
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
        response = scheduler_stub.UpdateAppt(scheduler_pb2.UpdateAppointment(appointment={
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
        response = scheduler_stub.DeleteAppt(scheduler_pb2.DeleteAppointment(apptID=req_id), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'delete appointment timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/staff", methods=['GET'])
def get_staff():
    try:
        response = staff_stub.Get(staff_records_pb2.GetStaffRecords(), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'get staff timed out'})
        return response, 408
    print(response)
    return MessageToDict(response)

@app.route("/staff/availability", methods=['GET'])
def get_staff_availability():
    data = request.get_json()
    req_id = data["staffID"]
    try:
        response = staff_stub.GetAvailability(staff_records_pb2.GetStaffAvailability(staffID=req_id), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'get staff availability timed out'})
        return response, 408
    return MessageToDict(response)

@app.route("/staff", methods=['POST'])
def create_staff():
    data = request.get_json()
    staff_data = data["staffRecord"]
    try:
        response = staff_stub.Create(staff_records_pb2.CreateStaff(staff={
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
        response = staff_stub.Update(staff_records_pb2.UpdateStaff(staffRecord={
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
        response = staff_stub.Delete(staff_records_pb2.DeleteStaff(staffID=req_id), timeout=0.5)
    except grpc.RpcError as e:
        response = jsonify({'message': 'delete staff timed out'})
        return response, 408
    return MessageToDict(response)


if __name__ == "__main__":
    logging.basicConfig()
    init()
    app.run(host='localhost', port=5050)
