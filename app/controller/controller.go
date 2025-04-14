package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/app/model"
	"myapp/app/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// reuable
func getUserId(userIdParam string) (int64, error) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, userErr
	}
	return userId, nil
}

func GetStdu(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Student{StdId: stdId}
	getErr := s.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "student not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, s)

}
