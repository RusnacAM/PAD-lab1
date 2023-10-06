from flask import Flask, request, jsonify
import logging

app = Flask(__name__)

routes = dict()

@app.route("/route", methods=['POST'])
def register_route():
    data = request.get_json()
    if data["svc"] in routes:
        routes[data["svc"]].append(data["route"])
        response = jsonify({'message': 'route registered successfully'})
    else:
        routes[data["svc"]] = data["route"]
        response = jsonify({'message': 'route registered successfully'})

    return response, 200

@app.route("/route", methods=['GET'])
def get_routes():
    if len(routes) == 0:
        return jsonify({'message': 'No services registered'}), 209
    else:
        return jsonify(routes), 200

if __name__ == "__main__":
    logging.basicConfig()
    app.run(host='localhost', port=5051)