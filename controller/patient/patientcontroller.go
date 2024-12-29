package patient

import (
	"emr/model/patient"
	"emr/model/user"
)

func SetPatient() patient.PatientModel {
	user := user.User{
		Name:   "Juliane",
		Id:     123,
		Gender: "Female",
		DOB:    "2000-09-20",
	}
	Apatient := patient.PatientModel{
		Id:       1,
		Name:     "Jean Pierre",
		Gender:   "Male",
		Category: "Own",
		Address:  "Kigali",
		User:     user,
	}
	return Apatient
}
