from functools import wraps
from datetime import datetime, timedelta
import time
from flask import jsonify
import grpc
import logging
import os

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s,%(msecs)d %(levelname)s: %(message)s",
    datefmt="%H:%M:%S",
)

def circuit_breaker(f):
    failure_timestamps = []
    delay = 0.5 * 3.5 * 1000
    threshold = 3

    @wraps(f)
    def circuit_breaker_wrapper(*args, **kwargs):
        try:
            result = f(*args, **kwargs)
            return result
        except grpc.RpcError as e:
            failure_time = round(time.time()*1000)
            failure_timestamps.append(failure_time)
            if len(failure_timestamps) > threshold:
                delta_time = failure_timestamps[2] - failure_timestamps[0]
                if delta_time <= delay:
                    logging.info("The service accessed is currently down!")
                    return jsonify({'message': 'The service accessed is currently down and cannot be reached!'}), 503
            response = jsonify({'message': 'update appointment timed out'})
            return response, 408
    return circuit_breaker_wrapper