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
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

var people *[]People

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	result := map[string]interface{}{
		"status":  http.StatusOK,
		"data":    people,
		"message": nil,
	}

	json.NewEncoder(w).Encode(result)
}

func PostData(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// define var
	var (
		todoPost      People
		save, getData []People
	)
	json.NewDecoder(r.Body).Decode(&todoPost)
	if people == nil {
		todoPost.ID = 1
		save = append(save, todoPost)
		people = &save
	} else {
		getData = *people
		todoPost.ID += 1
		getData = append(getData, todoPost)
		people = &getData
	}

	result := map[string]interface{}{
		"status":  http.StatusOK,
		"data":    nil,
		"message": "data berhasil di post",
	}

	json.NewEncoder(w).Encode(result)
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	if people == nil {
		result := map[string]interface{}{
			"status":  http.StatusNotFound,
			"data":    nil,
			"message": "id not found",
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	var dataPeople []People
	dataPeople = *people

	for index, data := range dataPeople {
		if data.ID == id {
			dataPeople = append(dataPeople[:index], dataPeople[index+1:]...)
			people = &dataPeople
			result := map[string]interface{}{
				"status":  http.StatusOK,
				"data":    nil,
				"message": "data berhasil di hapus",
			}
			json.NewEncoder(w).Encode(result)
		}
	}

}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	if people == nil {
		result := map[string]interface{}{
			"status":  http.StatusNotFound,
			"data":    nil,
			"message": "id not found",
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	var dataPeople []People
	dataPeople = *people

	var peoplePost People

	json.NewDecoder(r.Body).Decode(&peoplePost)
	for index, data := range dataPeople {
		if data.ID == id {
			dataPeople[index].Name = peoplePost.Name
			people = &dataPeople
		} else {
			http.Error(w, "Error", http.StatusBadRequest)
			return
		}
	}
	json.NewEncoder(w).Encode("data berhasil di update")
}
