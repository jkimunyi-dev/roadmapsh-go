package marshal_and_unmarshal

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type (
	Person struct {
		Name    string
		Age     int
		Address Address
		Pets    []Pet
	}

	Address struct {
		Line1  string
		Line2  string
		Postal int
	}

	Pet struct {
		Name  string
		Kind  string
		Age   int
		Color string
	}
)

func ConvertToJson() {
	data, err := os.ReadFile("complex-data.json")
	if err != nil {
		fmt.Printf("Failed to read file : %v\n", err.Error())
		return
	}

	var person Person

	json.Unmarshal(data, &person)

	fmt.Printf("Name : %s\n", person.Name)
	fmt.Printf("Age : %d\n", person.Age)
	fmt.Printf("Address : %v\n", person.Address)
	fmt.Printf("Pets : %v\n", person.Pets)
	personType := reflect.TypeOf(person).Kind()
	fmt.Printf(personType.String())
}
