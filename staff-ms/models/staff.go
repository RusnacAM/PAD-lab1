package models

type StaffRecord struct {
	StaffID     string `gorm:"staffid" gorm:"primarykey"`
	Name        string
	JobTitle    string
	Department  string
	IsAvailable bool
}
