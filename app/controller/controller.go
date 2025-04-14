package controller

import (
	"encoding/json"
	"myapp/app/model"
	"myapp/app/utils/httpResp"
	"net/http"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// create variable type Student
	var stud model.Student

	// read the request body and create a decoder object
	decoder := json.NewDecoder(r.Body)

	// store the json object data to stud variable
	if err := decoder.Decode(&stud); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid json Body")
		return
	}

	// defer the closing of request body until the function returns
	defer r.Body.Close()

	// call the Create() using student object, stud
	saveErr := stud.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}

	// no error
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "Student Added"})
}
