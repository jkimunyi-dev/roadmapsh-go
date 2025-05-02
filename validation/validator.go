package validation

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"required,gt=18,max=99"`
}

func Validator() {
	input := `{
		"username": "jimmie",
		"password": "123456",
		"email":    "jimmie@gmail.com",
		"age":      23

	}`

	var user User

	err := json.Unmarshal([]byte(input), &user)
	if err != nil {
		fmt.Printf("Failed to unmarshal : %v\n", err.Error())
		return
	}

	fmt.Printf("User before validation : %v\n", user)

	err = validator.New().Struct(user)
	if err != nil {
		fmt.Printf("Validation failed : %v\n", err.Error())
		return
	}

}
