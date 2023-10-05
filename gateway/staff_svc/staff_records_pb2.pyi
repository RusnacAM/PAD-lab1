from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class StaffRecord(_message.Message):
    __slots__ = ["staffID", "name", "jobTitle", "department", "isAvailable"]
    STAFFID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    JOBTITLE_FIELD_NUMBER: _ClassVar[int]
    DEPARTMENT_FIELD_NUMBER: _ClassVar[int]
    ISAVAILABLE_FIELD_NUMBER: _ClassVar[int]
    staffID: str
    name: str
    jobTitle: str
    department: str
    isAvailable: bool
    def __init__(self, staffID: _Optional[str] = ..., name: _Optional[str] = ..., jobTitle: _Optional[str] = ..., department: _Optional[str] = ..., isAvailable: bool = ...) -> None: ...

class CreateStaff(_message.Message):
    __slots__ = ["staff"]
    STAFF_FIELD_NUMBER: _ClassVar[int]
    staff: StaffRecord
    def __init__(self, staff: _Optional[_Union[StaffRecord, _Mapping]] = ...) -> None: ...

class CreateResponse(_message.Message):
    __slots__ = ["staffID", "message", "error"]
    STAFFID_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    staffID: str
    message: str
    error: str
    def __init__(self, staffID: _Optional[str] = ..., message: _Optional[str] = ..., error: _Optional[str] = ...) -> None: ...

class GetStaffRecords(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GetResponse(_message.Message):
    __slots__ = ["staffRecords", "error"]
    STAFFRECORDS_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    staffRecords: _containers.RepeatedCompositeFieldContainer[StaffRecord]
    error: str
    def __init__(self, staffRecords: _Optional[_Iterable[_Union[StaffRecord, _Mapping]]] = ..., error: _Optional[str] = ...) -> None: ...

class GetStaffAvailability(_message.Message):
    __slots__ = ["staffID"]
    STAFFID_FIELD_NUMBER: _ClassVar[int]
    staffID: str
    def __init__(self, staffID: _Optional[str] = ...) -> None: ...

class GetAvailabilityResponse(_message.Message):
    __slots__ = ["staffID", "isAvailable", "error"]
    STAFFID_FIELD_NUMBER: _ClassVar[int]
    ISAVAILABLE_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    staffID: str
    isAvailable: bool
    error: str
    def __init__(self, staffID: _Optional[str] = ..., isAvailable: bool = ..., error: _Optional[str] = ...) -> None: ...

class UpdateStaff(_message.Message):
    __slots__ = ["staffRecord"]
    STAFFRECORD_FIELD_NUMBER: _ClassVar[int]
    staffRecord: StaffRecord
    def __init__(self, staffRecord: _Optional[_Union[StaffRecord, _Mapping]] = ...) -> None: ...

class UpdateResponse(_message.Message):
    __slots__ = ["staffID", "message", "error"]
    STAFFID_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    staffID: str
    message: str
    error: str
    def __init__(self, staffID: _Optional[str] = ..., message: _Optional[str] = ..., error: _Optional[str] = ...) -> None: ...

class DeleteStaff(_message.Message):
    __slots__ = ["staffID"]
    STAFFID_FIELD_NUMBER: _ClassVar[int]
    staffID: str
    def __init__(self, staffID: _Optional[str] = ...) -> None: ...

class DeleteResponse(_message.Message):
    __slots__ = ["staffID", "message", "error"]
    STAFFID_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    staffID: str
    message: str
    error: str
    def __init__(self, staffID: _Optional[str] = ..., message: _Optional[str] = ..., error: _Optional[str] = ...) -> None: ...

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
