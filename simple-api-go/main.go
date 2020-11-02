package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)
// definir strutura do json
type event struct {
	ID string `json:"ID"`
	Title string `json:"Title"`
	Description string `json:"Description"`
}
// array de eventos
type allEvents [] event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}
// w do tipo response
// r do tipo request
func homeLink(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome")
}

func createEvent(w http.ResponseWriter, r *http.Request){
	var newEvent event
	// pegar dados do body da requisição
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w,"error")
	}
	// Associar o body que está em ascii para string
	json.Unmarshal(reqBody,&newEvent)
	events = append(events,newEvent)
	w.WriteHeader(http.StatusCreated)
	// resposta da requisição
	json.NewEncoder(w).Encode(newEvent)
}
func getEvents( w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func getOneEvent( w http.ResponseWriter, r *http.Request){
	eventID := mux.Vars(r)["id"]
	for _, singleEvent := range events{
		if singleEvent.ID == eventID{
			json.NewEncoder(w).Encode(singleEvent)
		}else{
			fmt.Fprint(w,"error")
		}
	}
}


func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event",createEvent).Methods("POST")
	router.HandleFunc("/event",getEvents).Methods("GET")
	router.HandleFunc("/event/{id}",getOneEvent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}