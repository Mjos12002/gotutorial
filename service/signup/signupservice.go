package signup

import (
	"emr/model/user"
	"fmt"
)

func Signup() {
	user := user.User{
		Name:   "Joseph",
		DOB:    "2022-20-2",
		Id:     234,
		Gender: "Male",
	}

	fmt.Printf("A %s named %s born on %s with ID %d has signed out", user.Gender, user.Name, user.DOB, user.Id)
}
