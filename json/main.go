package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Producer struct {
	Name string `json:"name"`
}

type Product struct {
	Title    string   `json:"title"`
	Producer Producer `json:"producer"`
}

type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

type Address struct {
	Street string `json:"street"`
}

func main() {
	// Marshalling
	producer := Producer{Name: "Somethinc"}
	product := Product{Title: "Moisturizer", Producer: producer}

	byteArr, err := json.MarshalIndent(product, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArr))

	// Unmarshalling
	apiUrl := "http://jsonplaceholder.typicode.com/users"

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(users)
}
