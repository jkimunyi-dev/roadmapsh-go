package marshal_and_unmarshal

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ConvertToJson() {
	// struct to json

	jimmie := Person{
		Name: "Jimmie",
		Age:  25,
		Address: Address{
			Line1:  "Block 78 Woodgrove Avenue 5",
			Line2:  "Unit #05-111",
			Postal: "654378",
		},
		Pets: []Pet{
			{
				Name:  "Lex",
				Kind:  "Dog",
				Age:   4,
				Color: "Gray",
			},
			{
				Name:  "Faye",
				Kind:  "Cat",
				Age:   6,
				Color: "Orange",
			},
		},
	}

	jimmieJson, err := json.Marshal(jimmie)
	if err != nil {
		fmt.Printf("Failed to marshal : %v\n", err.Error())
		return
	}

	fmt.Println(string(jimmieJson))
	fmt.Println(reflect.TypeOf(jimmieJson))
}
