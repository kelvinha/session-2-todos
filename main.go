package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/get-data", Get)
	http.HandleFunc("/post-data", PostData)
	http.HandleFunc("/delete-data", DeleteData)
	http.HandleFunc("/update-data", UpdateData)

	log.Println("server listen :8000")
	http.ListenAndServe(":8000", nil)
}

// struct
type People struct {
	ID   int
	Name string
}

var people *[]People

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if len(*people) == 0 {
		people = nil
	}
	json.NewEncoder(w).Encode(people)
}

func PostData(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// define var
	var (
		todoPost People
		save     []People
	)

	json.NewDecoder(r.Body).Decode(&todoPost)
	save = append(save, todoPost)
	people = &save
	w.Write([]byte("data berhasil di post"))
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	var dataPeople []People
	dataPeople = *people

	for index, data := range dataPeople {
		if data.ID == id {
			dataPeople = append(dataPeople[:index], dataPeople[index+1:]...)
			people = &dataPeople
			w.Write([]byte("data berhasil dihapus"))
		}
	}

}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	var dataPeople []People
	dataPeople = *people

	var peoplePost People

	json.NewDecoder(r.Body).Decode(&peoplePost)
	for index, data := range dataPeople {
		if data.ID == id {
			dataPeople = append(dataPeople[:index], dataPeople[index+1:]...)
			dataPeople = append(dataPeople, peoplePost)
			people = &dataPeople
		}
	}
	w.Write([]byte("data berhasil di update"))
}
