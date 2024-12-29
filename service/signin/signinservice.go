package signin

import (
	"emr/model/user"
	"fmt"
)

func Signin() {
	jean := user.User{Name: "Joseph", Gender: "Male", DOB: "2000-20-22", Id: 1}
	fmt.Printf("%s who is %s born on %s with ID %d has logged in", jean.Name, jean.Gender, jean.DOB, jean.Id)
}
