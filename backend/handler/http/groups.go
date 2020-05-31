package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/backend/handler"
	"github.com/backend/model"
	"github.com/backend/repository"
	"github.com/backend/repository/user"
	"github.com/go-chi/chi"
)

type Groups struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewGroupsHandler(conn *sql.DB) *Groups {
	//log.Printf("Passing user handler an sql connection to get Groups struct in handler/http/user.Groups")
	return &Groups{
		repo: user.NewGroupsRepository(conn),
	}
}

func (user *Groups) GetHTTPHandler() []*handler.HTTPHandler {
	log.Printf("Http Handler Redirecting to object based on path ..groups.")
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "groups/{id}", Func: user.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "groups", Func: user.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "groups/{id}", Func: user.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "groups/{id}", Func: user.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "groups", Func: user.GetAll},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPatch, Path: "groups/{id}", Func: user.Update},
	}
}
func (user *Groups) GetByID(w http.ResponseWriter, r *http.Request) {
	//log.Printf("Handler")
	var usr interface{}
	//log.Printf("Creating Groups object and Parsing and getting urlpattern")
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = user.repo.GetByID(r.Context(), id)
		break
	}
	log.Printf("Writing response in JSON")
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (user *Groups) Create(w http.ResponseWriter, r *http.Request) {
	//log.Printf("Handler")
	//log.Printf("Creating usr object defined in model")
	var usr model.Groups
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		//log.Printf("Now calling Create function Create function with req.context and usr object")

		_, err = user.repo.Create(r.Context(), usr)

		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (user *Groups) Update(w http.ResponseWriter, r *http.Request) {
	//log.Printf("Handler")
	//log.Printf("PATCH/PUT call")
	//log.Printf("Creating usr object defined in handler/http")
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	//log.Printf("Update by Id called")
	usr := model.Groups{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	//log.Printf("Decoding request in usr object")
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		usr.UpdatedBy = 0
		//log.Printf("Now calling Update function with req.context and usr object")
		iUsr, err = user.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.Groups)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (user *Groups) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}
		log.Printf("Now calling  delete with req.context and usr object")
		_, err = user.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "Groups deleted successfully"
		break
	}
	log.Printf("Writing response to JSON")
	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (user *Groups) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("Calling Getall Function")
	usrs, err := user.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
