package main

import (
	"encoding/json"
	"fmt"
)

type Carrer struct {
	Role     string `json:"role"`
	Pay      int    `json:"pay"`
	Location string `json:"location"`
}

type Person struct {
	Name   string
	Age    int
	Carrer Carrer
}

func main() {

	carrer := Carrer{
		Role:     "Backend Engineer",
		Pay:      90000,
		Location: "Nairobi, Kenya",
	}
	person := Person{
		Name:   "Jimmie",
		Age:    23,
		Carrer: carrer,
	}
	jimmie, err := json.Marshal(
		// Person{
		// 	Name: "Jimmy",
		// 	Age:  25,
		// 	Carrer: Carrer{
		// 		Role:     "Android Engineer",
		// 		Pay:      90000,
		// 		Location: "Nairobi, Kenya",
		// 	},
		// },

		person,
	)

	if err != nil {
		fmt.Printf("Failed to Marshal data :%s ", err.Error())
		return
	}

	fmt.Println("Person : ", string(jimmie))

}
