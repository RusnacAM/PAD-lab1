from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Appointment(_message.Message):
    __slots__ = ["apptID", "patientName", "staffID", "apptDateTime", "apptType"]
    APPTID_FIELD_NUMBER: _ClassVar[int]
    PATIENTNAME_FIELD_NUMBER: _ClassVar[int]
    STAFFID_FIELD_NUMBER: _ClassVar[int]
    APPTDATETIME_FIELD_NUMBER: _ClassVar[int]
    APPTTYPE_FIELD_NUMBER: _ClassVar[int]
    apptID: str
    patientName: str
    staffID: str
    apptDateTime: str
    apptType: str
    def __init__(self, apptID: _Optional[str] = ..., patientName: _Optional[str] = ..., staffID: _Optional[str] = ..., apptDateTime: _Optional[str] = ..., apptType: _Optional[str] = ...) -> None: ...

class CreateAppointment(_message.Message):
    __slots__ = ["appointment"]
    APPOINTMENT_FIELD_NUMBER: _ClassVar[int]
    appointment: Appointment
    def __init__(self, appointment: _Optional[_Union[Appointment, _Mapping]] = ...) -> None: ...

class CreateResponse(_message.Message):
    __slots__ = ["apptID", "status", "error"]
    APPTID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    apptID: str
    status: int
    error: str
    def __init__(self, apptID: _Optional[str] = ..., status: _Optional[int] = ..., error: _Optional[str] = ...) -> None: ...

class GetAppointments(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GetResponse(_message.Message):
    __slots__ = ["appointments", "status", "error"]
    APPOINTMENTS_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    appointments: _containers.RepeatedCompositeFieldContainer[Appointment]
    status: int
    error: str
    def __init__(self, appointments: _Optional[_Iterable[_Union[Appointment, _Mapping]]] = ..., status: _Optional[int] = ..., error: _Optional[str] = ...) -> None: ...

class UpdateAppointment(_message.Message):
    __slots__ = ["appointment"]
    APPOINTMENT_FIELD_NUMBER: _ClassVar[int]
    appointment: Appointment
    def __init__(self, appointment: _Optional[_Union[Appointment, _Mapping]] = ...) -> None: ...

class UpdateResponse(_message.Message):
    __slots__ = ["appointment", "status", "error"]
    APPOINTMENT_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    appointment: Appointment
    status: int
    error: str
    def __init__(self, appointment: _Optional[_Union[Appointment, _Mapping]] = ..., status: _Optional[int] = ..., error: _Optional[str] = ...) -> None: ...

class DeleteAppointment(_message.Message):
    __slots__ = ["apptID"]
    APPTID_FIELD_NUMBER: _ClassVar[int]
    apptID: str
    def __init__(self, apptID: _Optional[str] = ...) -> None: ...

class DeleteResponse(_message.Message):
    __slots__ = ["apptID", "status", "error"]
    APPTID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    apptID: str
    status: int
    error: str
    def __init__(self, apptID: _Optional[str] = ..., status: _Optional[int] = ..., error: _Optional[str] = ...) -> None: ...

class HealthCheckRequest(_message.Message):
    __slots__ = ["service"]
    SERVICE_FIELD_NUMBER: _ClassVar[int]
    service: str
    def __init__(self, service: _Optional[str] = ...) -> None: ...

class HealthCheckResponse(_message.Message):
    __slots__ = ["status"]
    class ServingStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = []
        UNKNOWN: _ClassVar[HealthCheckResponse.ServingStatus]
        SERVING: _ClassVar[HealthCheckResponse.ServingStatus]
        NOT_SERVING: _ClassVar[HealthCheckResponse.ServingStatus]
    UNKNOWN: HealthCheckResponse.ServingStatus
    SERVING: HealthCheckResponse.ServingStatus
    NOT_SERVING: HealthCheckResponse.ServingStatus
    STATUS_FIELD_NUMBER: _ClassVar[int]
    status: HealthCheckResponse.ServingStatus
    def __init__(self, status: _Optional[_Union[HealthCheckResponse.ServingStatus, str]] = ...) -> None: ...
