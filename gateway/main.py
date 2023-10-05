import logging
import grpc
from scheduler_svc import scheduler_pb2
from scheduler_svc import scheduler_pb2_grpc
from flask import Flask, jsonify



def run():
    channel = grpc.insecure_channel('localhost:3030')

    stub = scheduler_pb2_grpc.SchedulerStub(channel)
    response = stub.GetAppt(scheduler_pb2.GetAppointments(), timeout=0.5) 
        # response = stub.CreateAppt(scheduler_pb2.CreateAppointment(appointment={"patientName": "grgsr"}))
        # response = stub.UpdateAppt(scheduler_pb2.UpdateAppointment(appointment={"apptID": "736dc925-5d5b-45ee-b546-d888e0a2de92", "patientName": "new test name"}))
        # response = stub.DeleteAppt(scheduler_pb2.DeleteAppointment(apptID="736dc925-5d5b-45ee-b546-d888e0a2de92"))
    print(response)

if __name__ == "__main__":
    logging.basicConfig()
    run()