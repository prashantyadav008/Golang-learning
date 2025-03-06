package crudOperation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func CrudOperation() {
	fmt.Println("CRUD Operation")
	fmt.Println("\n --------------------------------- GET Method ---------------------------------")
	PerformGetMethod()

	fmt.Println("\n --------------------------------- POST Method ---------------------------------")
	PerformPostMethod()

	fmt.Println("\n --------------------------------- Update Method ---------------------------------")
	PerformUpdateMethod()

	fmt.Println("\n --------------------------------- Delete Method ---------------------------------")
	PerformDeleteMethod()
}

func PerformGetMethod() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Getting Error in Web Request-", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in Status Code-", res.StatusCode)
		return
	}

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error in New Decoder-", err)
		return
	}

	fmt.Println("Todo Data is - ", todo, "\t Title - ", todo.Title)

	/* or Option to fetch data and data parameters
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in Read Response-", err)
		return
	}
	fmt.Println("URL Response is -", string(data))

	// extract data
	var todo Todo
	err = json.Unmarshal(data, &todo)
	if err != nil {
		fmt.Println("Error in Unmarshal-", err)
		return
	}

	fmt.Println("Todo Data is - ", todo, "\t Title - ", todo.Title)

	*/
}

func PerformPostMethod() {
	var todo Todo = Todo{
		UserId:    19,
		Title:     "Learning Go Language",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in Marshal-", err)
		return
	}
	fmt.Println("Todo Data is - ", string(jsonData))

	jsonReader := strings.NewReader(string(jsonData))
	url := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(url, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Getting Error in Post Request-", err)
		return
	}

	defer res.Body.Close()

	// data, _ := io.ReadAll(res.Body)
	// fmt.Println("Todo Data is - ", string(data))

	// or
	var todoPostData Todo
	err = json.NewDecoder(res.Body).Decode(&todoPostData)
	if err != nil {
		fmt.Println("Error in New Decoder-", err)
		return
	}
	fmt.Println("Post Todo Data is - ", todoPostData)

}

func PerformUpdateMethod() {
	var todo Todo = Todo{
		UserId:    19,
		Title:     "Learning Advance Go Language",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in Marshal-", err)
		return
	}
	fmt.Println("Todo Data is - ", string(jsonData))

	// format into json
	jsonReader := strings.NewReader(string(jsonData))
	url := "https://jsonplaceholder.typicode.com/todos/1"
	// req, err := http.NewRequest("PUT", url, jsonReader)
	req, err := http.NewRequest(http.MethodPut, url, jsonReader)
	if err != nil {
		fmt.Println("Getting Error in Put Request-", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Getting Error in Put Request-", err)
		return
	}

	defer res.Body.Close()

	// data, _ := io.ReadAll(res.Body)
	// fmt.Println("Todo Data is - ", string(data))

	// or
	var todoPostData Todo
	err = json.NewDecoder(res.Body).Decode(&todoPostData)
	if err != nil {
		fmt.Println("Error in New Decoder-", err)
		return
	}
	fmt.Println("Post Todo Data is - ", todoPostData)

}

func PerformDeleteMethod() {

	url := "https://jsonplaceholder.typicode.com/todos/1"
	// req, err := http.NewRequest("DELETE", url, nil)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		fmt.Println("Getting Error in Put Request-", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Getting Error in Put Request-", err)
		return
	}

	defer res.Body.Close()
	fmt.Println("Status Code is - ", res.Status)

}
