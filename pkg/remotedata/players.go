package remotedata

import (
	"encoding/json"
	"fmt"

	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetPlayersByID(ids []string) ([]*model.Player, error) {
	resp, err := b.Get("players", map[string][]string{"ids": ids})
	if err != nil {
		return nil, err
	}
	var players []*model.Player
	if err = json.Unmarshal(resp.Body(), &players); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
	}

	return players, nil
}
