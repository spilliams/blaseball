package apiserver

import (
	"fmt"
	"net/http"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (s *Server) GetDivisions(w http.ResponseWriter, r *http.Request) error {
	divisions, err := s.dataStore.GetAllDivisions()
	if err != nil {
		return err
	}

	if len(divisions) == 0 {
		remoteDivisions, err := s.remoteAPI.GetAllDivisions()
		if err != nil {
			return err
		}
		for _, d := range remoteDivisions {
			if err := s.dataStore.PutDivision(d); err != nil {
				return err
			}
		}
		divisions, err = s.dataStore.GetAllDivisions()
		if err != nil {
			return err
		}
	}

	return marshalAndWrite(divisions, w)
}

func (s *Server) GetDivision(w http.ResponseWriter, r *http.Request) error {
	id := getQueryString(r, "id")
	name := getQueryString(r, "name")
	if len(id) == 0 && len(name) == 0 {
		return pkg.NewCodedError(fmt.Errorf("either `id` or `name` must be specified in query parameters"), http.StatusBadRequest)
	}

	var div *model.Division
	var err error
	if len(id) != 0 {
		div, err = s.dataStore.GetDivisionByID(id)
	} else {
		div, err = s.dataStore.GetDivisionByName(name)
	}
	if err != nil {
		l := loggerFromRequest(r)
		l.Warn("couldn't fetch division: %v", err)
	}
	if div != nil {
		return marshalAndWrite(div, w)
	}

	if len(id) == 0 {
		// user asked for it without ID, but we have no such ID in data
		return pkg.NewCodedError(fmt.Errorf("no Division found with name '%s'. Try looking it up by ID?", name), http.StatusNotFound)
	}

	// last resort: fetch from remote
	div, err = s.remoteAPI.GetDivisionByID(id)
	if err != nil {
		return err
	}
	if err = s.dataStore.PutDivision(div); err != nil {
		return err
	}
	return marshalAndWrite(div, w)
}
