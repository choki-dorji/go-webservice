package controller

import (
	"encoding/json"
	"myapp/app/model"
	"net/http"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// create variable type Student
	var stud model.Student

	// read the request body and create a decoder object
	decoder := json.NewDecoder(r.Body)

	// store the json object data to stud variable
	if err := decoder.Decode(&stud); err != nil {
		w.Write([]byte("Invalid json data"))
		return
	}

	// defer the closing of request body until the function returns
	defer r.Body.Close()

	// call the Create() using student object, stud
	saveErr := stud.Create()
	if saveErr != nil {
		w.Write([]byte("Database error"))
		return
	}

	// no error
	w.Write([]byte("response success."))
}
