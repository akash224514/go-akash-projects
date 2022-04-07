package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Patient struct {
	ID      string `json : "id"`
	Isbn    string `json: "isbn"`
	Disease string `json:"disease"`

	Personal_Info *Personal_Info `json:"personal_info"`
}

type Personal_Info struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var patients []Patient

func getPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(patients) //Putting all data inside patients slice into the json file

}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range patients {
		if item.ID == params["id"] {
			patients = append(patients[:index], patients[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(patients)
}

func getPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range patients {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var patient Patient
	_ = json.NewDecoder(r.Body).Decode(&patient)
	patient.ID = strconv.Itoa(rand.Intn(1000000000))
	patients = append(patients, patient)
	json.NewEncoder(w).Encode(patient)
}

func updatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json") //SET CONTENT TYPE
	params := mux.Vars(r)                              //Set Paramaters

	for index, item := range patients {
		if item.ID == params["id"] {
			patients = append(patients[:index], patients[index+1:]...)

			var patient Patient
			_ = json.NewDecoder(r.Body).Decode(&patient)

			//patient.ID == params["id"]
			patients = append(patients, patient)
			json.NewEncoder(w).Encode(patient)
		}
	}

}

func main() {

	r := mux.NewRouter()

	patients = append(patients, Patient{ID: "1", Isbn: "1232", Disease: "Flue", Personal_Info: &Personal_Info{Firstname: "Sudhir", Lastname: "Moghe"}})
	patients = append(patients, Patient{ID: "2", Isbn: "56765788", Disease: "Fever", Personal_Info: &Personal_Info{Firstname: "Pankaj", Lastname: "Tripathi"}})
	r.HandleFunc("/patients", getPatients).Methods("GET")
	r.HandleFunc("/patients/{id}", getPatient).Methods("GET")
	r.HandleFunc("/patients", createPatient).Methods("POST")
	r.HandleFunc("/patients/{id}", updatePatient).Methods("PUT")
	r.HandleFunc("/patients/{id}", deletePatient).Methods("DELETE")

	fmt.Printf("Starting server...")

	log.Fatal(http.ListenAndServe(":8000", r))
	//log.fatal(http.ListenAndServ("", r))
}
