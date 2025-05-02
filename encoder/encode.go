package encoder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Dog struct {
	Breed          string `json:"breed"`
	Age            int    `json:"age"`
	Name           string `json:"name"`
	FavouriteTreat string `json:"favourite_treat"`
}

func Encode() {
	newDog := Dog{
		Breed:          "Pug",
		Age:            2,
		Name:           "Lex",
		FavouriteTreat: "Bone",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		err := encoder.Encode(newDog)
		if err != nil {
			fmt.Printf("Failed to encode : %v\n", err.Error())
			return
		}

		fmt.Println("Encoding success")

	})

	fmt.Println("Server started")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Printf("Failed to start : %v\n", err.Error())
		return
	}

}
