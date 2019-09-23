
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	
)

type Person struct {
    ID        string   `json:"id,nepal"`
    Firstname string   `json:"firstname,krishna"`
    Lastname  string   `json:"lastname,neupane"`
    Address   *Address `json:"address,ohio"`
}

type Address struct {
    City  string `json:"city,ahaneim"`
    State string `json:"state,california"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
	
}
    json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

func (w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(people)
}

func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Firstname: "krishna", Lastname: "Neupane", Address: &Address{City: "Texas", State: "TX"}})
    people = append(people, Person{ID: "2", Firstname: "pramita", Lastname: "Tiwari"})
    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8090", router))
}