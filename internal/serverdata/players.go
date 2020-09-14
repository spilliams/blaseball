package serverdata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetAllPlayers(showFK bool) (*model.PlayerList, error) {
	q := addShowFKQuery(nil, showFK)
	resp, err := b.get("allPlayers", q)
	if err != nil {
		return nil, err
	}
	var players []*model.Player
	if err = json.Unmarshal(resp.Body(), &players); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
	}

	return &model.PlayerList{players}, nil
}

func (b *BlaseballAPI) GetPlayerByID(id string, showFK bool) (*model.Player, error) {
	players, err := b.GetPlayersByID([]string{id}, showFK)
	if err != nil {
		return nil, err
	}
	if players == nil || len(players.List) == 0 {
		return nil, pkg.NewCodedError(fmt.Errorf("no Player with id %s", id), http.StatusNotFound)
	}
	return players.List[0], nil
}

func (b *BlaseballAPI) GetPlayerByName(name string, showFK bool) (*model.Player, error) {
	q := addShowFKQuery(map[string][]string{"name": {name}}, showFK)
	resp, err := b.get("players", q)
	if err != nil {
		return nil, err
	}
	var player *model.Player
	if err = json.Unmarshal(resp.Body(), &player); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
	}

	return player, nil
}

const chunkSize = 199

func (b *BlaseballAPI) GetPlayersByID(ids []string, showFK bool) (*model.PlayerList, error) {
	chunkedIDs := chunk(ids, chunkSize)
	bodies := make([][]byte, 0, len(chunkedIDs))
	b.Debugf("calling in %d separate chunks of size %d", len(chunkedIDs), chunkSize)
	for _, chunk := range chunkedIDs {
		q := addShowFKQuery(map[string][]string{"ids": {strings.Join(chunk, ",")}}, showFK)
		resp, err := b.get("players", q)
		if err != nil {
			return nil, err
		}
		bodies = append(bodies, resp.Body())
	}

	var players []*model.Player
	for _, body := range bodies {
		var respPlayers []*model.Player
		if err := json.Unmarshal(body, &respPlayers); err != nil {
			return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
		}
		players = append(players, respPlayers...)
	}

	return &model.PlayerList{players}, nil
}

func chunk(items []string, chunkSize int) (chunks [][]string) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
