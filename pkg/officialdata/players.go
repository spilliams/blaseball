package officialdata

import (
	"encoding/json"

	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetPlayersByID(ids []string) ([]*model.Player, error) {
	bodies, err := b.getAll("players", "ids", ids)
	if err != nil {
		return nil, err
	}

	var players []*model.Player
	for _, body := range bodies {
		var respPlayers []*model.Player
		if err := json.Unmarshal(body, &respPlayers); err != nil {
			return nil, err
		}
		players = append(players, respPlayers...)
	}

	return players, nil
}
