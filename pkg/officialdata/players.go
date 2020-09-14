package officialdata

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spilliams/blaseball/pkg/model"
)

// Asking for /players?ids=... with 272+ ids will return a 414 URI Too Long.
// Asking for 200-271 ids will return a 431 Request Header Fields Too Large.
const chunkSize = 199

func (b *BlaseballAPI) GetPlayersByID(ids []string) ([]*model.Player, error) {
	chunkedIDs := chunk(ids, chunkSize)
	bodies := make([][]byte, 0, len(chunkedIDs))
	b.Debugf("calling in %d separate chunks of size %d", len(chunkedIDs), chunkSize)
	for _, chunk := range chunkedIDs {
		resp, err := b.get("players", map[string][]string{"ids": {strings.Join(chunk, ",")}})
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

	return players, nil
}

func chunk(items []string, chunkSize int) (chunks [][]string) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
