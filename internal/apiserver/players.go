package apiserver

import (
	"fmt"
	"net/http"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (s *Server) GetAllPlayers(w http.ResponseWriter, r *http.Request) error {
	players, err := s.dataSession.GetAllPlayers()
	if err != nil {
		return err
	}
	return marshalAndWrite(players, w)
}

func (s *Server) GetPlayers(w http.ResponseWriter, r *http.Request) error {
	ids := getQueryStrings(r, "ids")
	name := getQueryString(r, "name")
	if len(ids) == 0 && len(name) == 0 {
		return pkg.NewCodedError(fmt.Errorf("either `ids` or `name` must be specified in query parameters"), http.StatusBadRequest)
	}

	var players []*model.Player
	var err error
	if len(ids) != 0 {
		players, err = s.dataSession.GetPlayersByID(ids)
	} else {
		var player *model.Player
		player, err = s.dataSession.GetPlayerByName(name)
		if player != nil {
			players = append(players, player)
		}
	}
	if err != nil {
		l := loggerFromRequest(r)
		l.Warnf("couldn't fetch players: %v", err)
	}
	if players != nil {
		return marshalAndWrite(players, w)
	}

	if len(ids) == 0 {
		// user asked for it without IDs, but we couldn't find any
		return pkg.NewCodedError(fmt.Errorf("no Player found with name '%s'. Try looking them up by ID?", name), http.StatusNotFound)
	}

	// last resort: fetch from remote
	players, err = s.remoteAPI.GetPlayersByID(ids)
	if err != nil {
		return err
	}
	if err = s.dataSession.PutPlayers(players); err != nil {
		return err
	}
	return marshalAndWrite(players, w)
}
