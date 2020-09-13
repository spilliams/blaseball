package apiserver

import (
	"net/http"
)

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
