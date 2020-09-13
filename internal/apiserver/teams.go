package apiserver

import (
	"fmt"
	"net/http"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (s *Server) GetTeams(w http.ResponseWriter, r *http.Request) error {
	teams, err := s.dataSession.GetAllTeams()
	if err != nil {
		return err
	}

	if len(teams) == 0 {
		// TODO or if divisions are stale?
		remoteTeams, err := s.remoteAPI.GetAllTeams()
		if err != nil {
			return err
		}
		for _, t := range remoteTeams {
			if err := s.dataSession.PutTeam(t); err != nil {
				return err
			}
		}
		teams, err = s.dataSession.GetAllTeams()
		if err != nil {
			return err
		}
	}

	return marshalAndWrite(teams, w)
}

func (s *Server) GetTeam(w http.ResponseWriter, r *http.Request) error {
	id := getQueryString(r, "id")
	fullname := getQueryString(r, "fullname")
	nickname := getQueryString(r, "nickname")
	if len(id) == 0 && len(fullname) == 0 && len(nickname) == 0 {
		return pkg.NewCodedError(fmt.Errorf("either `id`, `fullname` or `nickname` must be specified in query parameters"), http.StatusBadRequest)
	}

	var team *model.Team
	var err error
	paramName := "full name"
	paramValue := fullname
	if len(id) != 0 {
		team, err = s.dataSession.GetTeamByID(id)
	} else if len(fullname) != 0 {
		team, err = s.dataSession.GetTeamByFullName(fullname)
	} else {
		team, err = s.dataSession.GetTeamByNickname(nickname)
		paramName = "nickname"
		paramValue = nickname
	}
	if err != nil {
		l := loggerFromRequest(r)
		l.Warn("couldn't fetch team: %v", err)
	}
	if team != nil {
		return marshalAndWrite(team, w)
	}

	if len(id) == 0 {
		// user asked for it without ID, but we have no such ID in data
		return pkg.NewCodedError(fmt.Errorf("no Team found with %s '%s'. Try looking it up by ID?", paramName, paramValue), http.StatusNotFound)
	}

	// last resort: fetch from remote
	team, err = s.remoteAPI.GetTeamByID(id)
	if err != nil {
		return err
	}
	if err = s.dataSession.PutTeam(team); err != nil {
		return err
	}
	return marshalAndWrite(team, w)
}
