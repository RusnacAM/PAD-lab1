package models

type Appointment struct {
	ApptID       string `gorm:"apptid" gorm:"primarykey"`
	PatientName  string
	DoctorName   string
	ApptDateTime string
	ApptType     string
}
