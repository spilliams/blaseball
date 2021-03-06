package apiserver

import (
	"net/http"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (s *Server) GetAllTeams(w http.ResponseWriter, r *http.Request) error {
	teams, err := s.dataStore.GetAllTeams()
	if err != nil {
		return err
	}

	if len(teams.List) == 0 {
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

	return marshalAndWrite(teams, w, r)
}

func (s *Server) GetTeam(w http.ResponseWriter, r *http.Request) error {
	id := getQueryString(r, "id")
	if len(id) != 0 {
		return s.getTeamByID(id, w, r)
	}

	fullname := getQueryString(r, "fullname")
	nickname := getQueryString(r, "nickname")
	if len(id) == 0 && len(fullname) == 0 && len(nickname) == 0 {
		return pkg.NewCodedErrorf(http.StatusBadRequest, "either `id`, `fullname` or `nickname` must be specified in query parameters")
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
		l.Warnf("couldn't fetch team by name: %v", err)
	}

	if team == nil {
		return pkg.NewCodedErrorf(http.StatusNotFound, "no Team found with %s '%s'. Try looking it up by ID?", paramName, paramValue)
	}

	return marshalAndWrite(team, w, r)
}

func (s *Server) getTeamByID(id string, w http.ResponseWriter, r *http.Request) error {
	team, err := s.dataStore.GetTeamByID(id)
	fetchFromRemote := err != nil || team == nil || team.Incomplete()
	if err != nil {
		l := loggerFromRequest(r)
		l.Warnf("couldn't fetch team by id: %v", err)
	}
	if !fetchFromRemote {
		return marshalAndWrite(team, w, r)
	}

	team, err = s.remoteAPI.GetTeamByID(id)
	if err != nil {
		return err
	}
	if team == nil {
		return pkg.NewCodedErrorf(http.StatusNotFound, "no Team found with id '%s'", id)
	}
	if err = s.dataStore.PutTeam(team); err != nil {
		return err
	}

	return marshalAndWrite(team, w, r)
}
