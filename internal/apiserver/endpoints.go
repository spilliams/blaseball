package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func marshalAndWrite(obj interface{}, w http.ResponseWriter) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("could not marshal response: %v", err)
	}
	_, err = w.Write(bytes)
	return err
}

func (s *Server) GetDivisions(w http.ResponseWriter, r *http.Request) error {
	// TODO if <something> then also get divisions from external source, put them in data session
	remoteDivisions, err := s.remoteAPI.GetAllDivisions()
	for _, d := range remoteDivisions {
		if err := s.dataSession.PutDivision(d); err != nil {
			return err
		}
	}

	divisions, err := s.dataSession.GetAllDivisions()
	if err != nil {
		return err
	}

	return marshalAndWrite(divisions, w)
}
