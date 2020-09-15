package memdata

import (
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataStore) GetAllPlayers() (*model.PlayerList, error) {
	players := make([]*model.Player, 0, len(mds.players))
	for _, p := range mds.players {
		copy := *p
		players = append(players, &copy)
	}
	return &model.PlayerList{List: players}, nil
}

func (mds *MemoryDataStore) GetPlayersByID(ids []string) (*model.PlayerList, error) {
	players := make([]*model.Player, 0, len(ids))
	for _, id := range ids {
		player, ok := mds.players[id]
		if !ok {
			continue
		}
		copy := *player
		players = append(players, &copy)
	}
	if len(players) == 0 {
		return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no players found with ids %v", ids)
	}
	return &model.PlayerList{List: players}, nil
}

func (mds *MemoryDataStore) GetPlayerByID(id string) (*model.Player, error) {
	player, ok := mds.players[id]
	if !ok {
		return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no Player with id %s", id)
	}
	copy := *player
	return &copy, nil
}

func (mds *MemoryDataStore) GetPlayerByName(name string) (*model.Player, error) {
	for _, p := range mds.players {
		if strings.EqualFold(p.Name, name) {
			copy := *p
			return &copy, nil
		}
	}
	return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no player with name %s", name)
}

func (mds *MemoryDataStore) PutPlayer(p *model.Player) error {
	mds.players[p.ID] = p
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

func (mds *MemoryDataStore) seedPlayers(ids []string) error {
	for _, id := range ids {
		_, ok := mds.players[id]
		if !ok {
			if err := mds.PutPlayer(&model.Player{ID: id}); err != nil {
				return err
			}
		}
	}
	return nil
}
