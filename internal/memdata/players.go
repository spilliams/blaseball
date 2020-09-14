package memdata

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataStore) GetAllPlayers() (*model.PlayerList, error) {
	players := make([]*model.Player, 0, len(mds.allPlayers))
	for _, p := range mds.allPlayers {
		players = append(players, p)
	}
	return &model.PlayerList{players}, nil
}

func (mds *MemoryDataStore) GetPlayersByID(ids []string) (*model.PlayerList, error) {
	players := make([]*model.Player, 0, len(ids))
	for _, id := range ids {
		player, ok := mds.allPlayers[id]
		if !ok {
			continue
		}
		players = append(players, player)
	}
	if len(players) == 0 {
		return nil, pkg.NewCodedError(fmt.Errorf("no players found with ids %v", ids), http.StatusNotFound)
	}
	return &model.PlayerList{players}, nil
}

func (mds *MemoryDataStore) GetPlayerByID(id string) (*model.Player, error) {
	player, ok := mds.allPlayers[id]
	if !ok {
		return nil, pkg.NewCodedError(fmt.Errorf("no Player with id %s", id), http.StatusNotFound)
	}
	return player, nil
}

func (mds *MemoryDataStore) GetPlayerByName(name string) (*model.Player, error) {
	for _, p := range mds.allPlayers {
		if strings.EqualFold(p.Name, name) {
			return p, nil
		}
	}
	return nil, pkg.NewCodedError(fmt.Errorf("no player with name %s", name), http.StatusNotFound)
}

func (mds *MemoryDataStore) PutPlayer(p *model.Player) error {
	mds.allPlayers[p.ID] = p
	return nil
}

func (mds *MemoryDataStore) PutPlayers(players []*model.Player) error {
	for _, p := range players {
		if err := mds.PutPlayer(p); err != nil {
			return err
		}
	}
	return nil
}

func (mds *MemoryDataStore) seedPlayers(ids []string) {
	for _, id := range ids {
		_, ok := mds.allPlayers[id]
		if !ok {
			mds.PutPlayer(&model.Player{ID: id})
		}
	}
}
