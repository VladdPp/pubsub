package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type subStruct struct {
	Pubsubname string `json:"pubsubname,omitempty"`
	Topic      string `json:"topic,omitempty"`
	Route      string `json:"route,omitempty"`
}

func subscribe(w http.ResponseWriter, r *http.Request) {
	testvalueA := subStruct{"pubsub", "A", "A"}
	testvalueB := subStruct{"pubsub", "B", "B"}

	pubSlice := [2]subStruct{testvalueA, testvalueB}

	pubSliceJSON, err := json.Marshal(pubSlice)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(pubSliceJSON)

}
func aSubscriber(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("A: status: 200")
	strBody := string(body)
	s := strings.Split(strBody, ",\"")
	fmt.Println(strings.TrimSuffix(s[6], "}"))
}
func bSubscriber(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("B: status: 200")
	strBody := string(body)
	s := strings.Split(strBody, ",\"")
	fmt.Println(strings.TrimSuffix(s[6], "}"))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/dapr/subscribe", subscribe).Methods("GET", "OPTIONS")
	router.HandleFunc("/A", aSubscriber).Methods("POST", "OPTIONS")
	router.HandleFunc("/B", bSubscriber).Methods("POST", "OPTIONS")
	log.Println("Starting server at :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
