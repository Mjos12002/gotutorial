package patient

import "emr/model/user"

type PatientModel struct {
	Id       int32
	Name     string
	Gender   string
	Category string
	Address  string
	User     user.User
}
