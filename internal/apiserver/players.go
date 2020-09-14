package apiserver

import (
	"fmt"
	"net/http"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (s *Server) GetAllPlayers(w http.ResponseWriter, r *http.Request) error {
	players, err := s.dataStore.GetAllPlayers()
	if err != nil {
		return err
	}

	incompletePlayerIDs := make([]string, 0, len(players))
	for _, p := range players {
		if p.Incomplete() {
			incompletePlayerIDs = append(incompletePlayerIDs, p.ID)
		}
	}
	if len(incompletePlayerIDs) > 0 {
		l := loggerFromRequest(r)
		l.Infof("fetching %d complete player records", len(incompletePlayerIDs))
		completePlayers, err := s.remoteAPI.GetPlayersByID(incompletePlayerIDs)
		if err != nil {
			return err
		}
		s.dataStore.PutPlayers(completePlayers)
		players, err = s.dataStore.GetAllPlayers()
		if err != nil {
			return err
		}
	}

	return marshalAndWrite(players, w)
}

func (s *Server) GetPlayers(w http.ResponseWriter, r *http.Request) error {
	ids := getQueryStrings(r, "ids")
	name := getQueryString(r, "name")
	if len(ids) != 0 {
		return s.getPlayersByID(ids, w, r)
	}
	if len(name) == 0 {
		return pkg.NewCodedError(fmt.Errorf("either `ids` or `name` must be specified in query parameters"), http.StatusBadRequest)
	}

	player, err := s.dataStore.GetPlayerByName(name)
	if err != nil {
		l := loggerFromRequest(r)
		l.Warnf("couldn't fetch player by name: %v", err)
	}
	if player == nil {
		return pkg.NewCodedError(fmt.Errorf("no Player found with name '%s'. Try looking them up by ID?", name), http.StatusNotFound)
	}

	return marshalAndWrite([]*model.Player{player}, w)
}

func (s *Server) getPlayersByID(ids []string, w http.ResponseWriter, r *http.Request) error {
	players, err := s.dataStore.GetPlayersByID(ids)
	incompletePlayerIDs := make([]string, 0, len(ids))
	// if there was an error, fetch them all from remote. Otherwise only fetch
	// the incomplete ones
	if err != nil {
		l := loggerFromRequest(r)
		l.Warnf("couldn't fetch players by id: %v", err)
		incompletePlayerIDs = append(incompletePlayerIDs, ids...)
	} else {
		for _, p := range players {
			if p.Incomplete() {
				incompletePlayerIDs = append(incompletePlayerIDs, p.ID)
			}
		}
	}

	if len(incompletePlayerIDs) > 0 {
		l := loggerFromRequest(r)
		l.Infof("fetching %d complete player records", len(incompletePlayerIDs))
		completePlayers, err := s.remoteAPI.GetPlayersByID(incompletePlayerIDs)
		if err != nil {
			return err
		}
		s.dataStore.PutPlayers(completePlayers)
		players, err = s.dataStore.GetPlayersByID(ids)
		if err != nil {
			return err
		}
	}

	return marshalAndWrite(players, w)
}
