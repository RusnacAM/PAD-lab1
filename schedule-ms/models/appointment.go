package models

type Appointment struct {
	ID           string `gorm:"apptid" gorm:"primarykey"`
	PatientName  string
	DoctorName   string
	ApptDateTime string
	ApptType     string
}
