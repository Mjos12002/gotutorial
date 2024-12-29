package main

import (
	"emr/controller/patient"
	"fmt"
)

func main() {
	pat := patient.SetPatient()
	fmt.Println(pat.Address)
	fmt.Println(pat.Category)
	fmt.Println(pat.Gender)
	fmt.Println(pat.Id)
	fmt.Println(pat.Name)
	fmt.Println(pat.User.DOB)
	fmt.Println(pat.User.Gender)
	fmt.Println(pat.User.Name)
	fmt.Println(pat.User.Id)
}
