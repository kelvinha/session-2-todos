package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"session-2-todos/docs"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server TODOS App server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := chi.NewRouter()

	r.Get("/get-data", Get)
	r.Post("/post-data", PostData)
	r.Delete("/delete-data", DeleteData)
	r.Put("/update-data", UpdateData)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json")))

	log.Println("server listen :8000")
	http.ListenAndServe(":8000", r)
}

// struct
type People struct {
	ID   int    `json:"id,omitempty" example:"1"`
	Name string `json:"name" example:"kelpin h"`
}

var people *[]People

// GET DATA PEOPLE godoc
// New GET DATA PEOPLE
// Endpoint GET DATA PEOPLE
// @Summary GET DATA PEOPLE
// @Description
// @Tags    TODOS
// @Accept  json
// @Produce json
// @Router  /get-data [get]
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	result := map[string]interface{}{
		"status":  http.StatusOK,
		"data":    people,
		"message": nil,
	}

	json.NewEncoder(w).Encode(result)
}

// POST DATA PEOPLE godoc
// New POST DATA PEOPLE
// Endpoint POST DATA PEOPLE
// @Summary POST DATA PEOPLE
// @Description
// @Tags    TODOS
// @Accept  json
// @Produce json
// @Param     Body  body      People  true  "payload"
// @Router  /post-data [post]
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

// DELETE DATA PEOPLE godoc
// New DELETE DATA PEOPLE
// Endpoint DELETE DATA PEOPLE
// @Summary DELETE DATA PEOPLE
// @Description
// @Tags    TODOS
// @Accept  json
// @Produce json
// @Param   id   query     string  true  "id"
// @Router  /delete-data [delete]
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
		} else {
			result := map[string]interface{}{
				"status":  http.StatusNotFound,
				"data":    nil,
				"message": "data tidak tersedia",
			}
			json.NewEncoder(w).Encode(result)
			return
		}
	}

}

// UPDATE DATA PEOPLE godoc
// New UPDATE DATA PEOPLE
// Endpoint UPDATE DATA PEOPLE
// @Summary UPDATE DATA PEOPLE
// @Description
// @Tags    TODOS
// @Accept  json
// @Produce json
// @Param   id   query     string  true  "id"
// @Param   Body  body      People  true  "payload"
// @Router  /update-data [put]
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
			result := map[string]interface{}{
				"status":  http.StatusOK,
				"data":    nil,
				"message": "data berhasil di update",
			}
			json.NewEncoder(w).Encode(result)
			return
		} else {
			result := map[string]interface{}{
				"status":  http.StatusNotFound,
				"data":    nil,
				"message": "data tidak tersedia",
			}
			json.NewEncoder(w).Encode(result)
			return
		}
	}
}
