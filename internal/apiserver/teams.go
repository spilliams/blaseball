package apiserver

import (
	"fmt"
	"net/http"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (s *Server) GetTeams(w http.ResponseWriter, r *http.Request) error {
	teams, err := s.dataStore.GetAllTeams()
	if err != nil {
		return err
	}

	if len(teams) == 0 {
		remoteTeams, err := s.remoteAPI.GetAllTeams()
		if err != nil {
			return err
		}
		for _, t := range remoteTeams {
			if err := s.dataStore.PutTeam(t); err != nil {
				return err
			}
		}
		teams, err = s.dataStore.GetAllTeams()
		if err != nil {
			return err
		}
	}

	return marshalAndWrite(teams, w)
}

func (s *Server) GetTeam(w http.ResponseWriter, r *http.Request) error {
	id := getQueryString(r, "id")
	if len(id) != 0 {
		return s.getTeamByID(id, w, r)
	}

	fullname := getQueryString(r, "fullname")
	nickname := getQueryString(r, "nickname")
	if len(id) == 0 && len(fullname) == 0 && len(nickname) == 0 {
		return pkg.NewCodedError(fmt.Errorf("either `id`, `fullname` or `nickname` must be specified in query parameters"), http.StatusBadRequest)
	}

	var team *model.Team
	var err error
	paramName := "full name"
	paramValue := fullname
	if len(fullname) != 0 {
		team, err = s.dataStore.GetTeamByFullName(fullname)
	} else {
		team, err = s.dataStore.GetTeamByNickname(nickname)
		paramName = "nickname"
		paramValue = nickname
	}

	if err != nil {
		l := loggerFromRequest(r)
		l.Warn("couldn't fetch team by name: %v", err)
	}

	if team == nil {
		return pkg.NewCodedError(fmt.Errorf("no Team found with %s '%s'. Try looking it up by ID?", paramName, paramValue), http.StatusNotFound)
	}

	return marshalAndWrite(team, w)
}

func (s *Server) getTeamByID(id string, w http.ResponseWriter, r *http.Request) error {
	team, err := s.dataStore.GetTeamByID(id)
	fetchFromRemote := err != nil || team == nil || team.Incomplete()
	if err != nil {
		l := loggerFromRequest(r)
		l.Warn("couldn't fetch team by id: %v", err)
	}
	if !fetchFromRemote {
		return marshalAndWrite(team, w)
	}

	team, err = s.remoteAPI.GetTeamByID(id)
	if err != nil {
		return err
	}
	if err = s.dataStore.PutTeam(team); err != nil {
		return err
	}

	return marshalAndWrite(team, w)
}
