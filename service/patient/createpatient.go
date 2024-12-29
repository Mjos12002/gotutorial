package patient

import (
	"emr/model/patient"
	"fmt"
)

func CreatePatient(patient patient.PatientModel) {
	fmt.Printf("%s, %s", patient.Address, patient.User.DOB)
}
