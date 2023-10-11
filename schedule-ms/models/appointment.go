package models

type Appointment struct {
	ApptID       string `gorm:"apptid" gorm:"primarykey"`
	PatientName  string
	StaffID      string
	ApptDateTime string
	ApptType     string
}
