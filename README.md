# PAD Lab1

## How To Run
1. First obtain the images from https://hub.docker.com/repositories/rusnacam
2. Clone the respective repository by doing ```git clone <repo>```
3. Run the command ```docker-compose up```

All the endpoints defined below can be accessed through the gateway at ```http://localhost:5050/``` after ensuring the container is running properly.

When accessing the gateway, you should first create some staff records if you don't have any, since scheduling appointments
is dependent on staff availability.

If you want to run the unit test, use the command in the makefile, from the staff directory: ```make unit-test```
## Application Suitability

1. Healthcare and patient service will always be a relevant topic, in countries with or without free healthcare, you will still need a sort of booking system for appointments
2. A healthcare application for appointment scheduling can't be solved any other way than distributed systems since it is related to the entire population of one country, which cannot be serviced by only one machine.
3. Improved accessibility and availability for patients, to schedule appointments from any location, like the comfort of their own home. Which in turn leads to ↓
4. Reduced waiting time for all. A distributed system can efficiently distribute healthcare requests across facilities and providers.
5. More efficient resource utilization, such a system can better allocate healthcare resources such as free rooms, doctors…
6. Redundancy and reliability,if one server or facility experiences an issue, the system can automatically redirect patients to other available locations for scheduling.
7. Scalability, easily increasing the size of the system with increased appointments and users.
8. Integration with other systems such as electronic health records and billing systems
9. Data security
10. Real-time updates, cancellations, rescheduling…


## Service Boundaries

1. Service for appointment scheduling, includes create/update/read/delete appt
2. Second service for staff records, which will include their availability, based on  which the appointments will be made
3. ![System Architecture](https://cdn.discordapp.com/attachments/758662311287980075/1166452303079952454/Screenshot_2023-10-24_at_22.04.54.png?ex=654a8a5c&is=6538155c&hm=8a1772262e6a3e330018ec97589359bd1cd5ac9ff263af85aec5555a22965dcf&)

## Technology Stack

1. Languages: Microservices in Golang, Gateway in Python
2. Communication through gRPC
3. DB postgres

## Data Management Design

Scheduler Endpoints:
1. POST ```/appointment```
   1. Request Body: 
   ```
   {
        "appointment": {
            "patientName": "Abigail White",
            "staffID": "56f85ca6-5ba6-43ea-9fcd-7d1231cec857",
            "apptDateTime": "2023-11-06 11:45 AM",
            "apptType": "Follow-up Consultation"
        }
   }
   ```
   2. Response Body:
   ```
   {
        "apptID": "6d7ca4ad-e9e5-4982-8058-0230172dd31b",
        "message": "The appointment was created successfully"
    }
   ```
2. GET ```/appointment```
   1. Response Body:
   ```
   {
    "appointments": [
        {
            "apptDateTime": "2023-11-06 11:45 AM",
            "apptID": "cdc670ce-f0de-405e-8535-6e52baf2a7b3",
            "apptType": "General check-up",
            "patientName": "Abigail White",
            "staffID": "56f85ca6-5ba6-43ea-9fcd-7d1231cec857"
        },
        {
            "apptDateTime": "2023-11-06 11:45 AM",
            "apptID": "6d7ca4ad-e9e5-4982-8058-0230172dd31b",
            "apptType": "General check-up",
            "patientName": "TEST PATIENT",
            "staffID": "56f85ca6-5ba6-43ea-9fcd-7d1231cec857"
        }
    ]
    }
   ```
3. PUT ```/appointment```
    1. Request Body:
   ```
   {
        "appointment": {
            "apptID": "cdc670ce-f0de-405e-8535-6e52baf2a7b3"
            "patientName": "Abigail White",
            "staffID": "56f85ca6-5ba6-43ea-9fcd-7d1231cec857",
            "apptDateTime": "2023-11-06 11:45 AM",
            "apptType": "Follow-up Consultation"
        }
   }
   ```
    2. Response Body:
   ```
   {
    "appointment": {
        "apptDateTime": "2023-11-06 11:45 AM",
        "apptType": "General check-up",
        "patientName": "NEW NAME",
        "staffID": "56f85ca6-5ba6-43ea-9fcd-7d1231cec857"
    },
    "message": "Appointment updated successfully."
    }
   ```
4. DELETE ```/appointment```
   1. Request Body:
   ```
   {
    "staffID": "16b5e5ea-bea5-4561-b10d-71d7f15ab7c3"
    }
   ```
   2. Response Body:
   ```
   {
    "apptID": "6d7ca4ad-e9e5-4982-8058-0230172dd31b",
    "message": "Appointment deleted successfully."
    }
   ```

Staff Records Endpoints:
1. POST ```/staff```
    1. Request Body:
   ```
   {
        "staff": {
            "name": "Linda Mary",
            "jobTitle": "Cardiologist",
            "department": "Cardiology",
            "isAvailable": true
        }
    }
   ```
    2. Response Body:
   ```
   {
    "message": "Staff record created successfully",
    "staffID": "affcfcca-56af-4621-81b2-35d013300dfb"
    }
   ```
2. GET  ```/staff```
   1. Response Body:
   ```
   {
    "staffRecords": [
        {
            "department": "Cardiology",
            "isAvailable": true,
            "jobTitle": "Cardiologist",
            "name": "Linda Mary",
            "staffID": "56f85ca6-5ba6-43ea-9fcd-7d1231cec857"
        },
        {
            "department": "Cardiology",
            "isAvailable": true,
            "jobTitle": "Cardiologist",
            "name": "TEST NAME",
            "staffID": "affcfcca-56af-4621-81b2-35d013300dfb"
        }
    ]
    }
   ```
3. GET ```/staff/availability```
   1. Request Body
   ```
   {
    "staffID": "affcfcca-56af-4621-81b2-35d013300dfb"
    }
   ```
   2. Response Body
   ```
   {
    "isAvailable": true,
    "staffID": "affcfcca-56af-4621-81b2-35d013300dfb"
    }
   ```
4. PUT ```/staff```
   1. Request Body
   ```
   {
    "staff": {
        "staffID": "56f85ca6-5ba6-43ea-9fcd-7d1231cec857",
        "name": "Sebastian Alcott",
        "jobTitle": "Cardiologist",
        "department": "Cardiology",
        "isAvailable": true
    }
    }
   ```
   2. Response Body
   ```
   {
    "message": "staff record successfully updated",
    "staffID": "56f85ca6-5ba6-43ea-9fcd-7d1231cec857"
    }
   ```
5. DELETE ```/staff```
   1. Request Body
   ```
   {
    "staffID": "affcfcca-56af-4621-81b2-35d013300dfb"
    }
   ```
   2. Response Body
   ```
   {
    "message": "record successfully deleted",
    "staffID": "affcfcca-56af-4621-81b2-35d013300dfb"
    }
   ```
   

## Deployment and Scaling

Containerization with docker