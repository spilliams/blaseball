package memdata

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/pkg/model"
)

// MemoryDataSession represents an in-memory set of blaseball data
type MemoryDataSession struct {
	allDivisions map[string]*model.Division
	allTeams     map[string]*model.Team
	allPlayers   map[string]*model.Player
}

// NewSession returns a new, empty in-memory data session
func NewSession() internal.LocalDataSession {
	return &MemoryDataSession{
		allDivisions: map[string]*model.Division{},
		allTeams:     map[string]*model.Team{},
		allPlayers:   map[string]*model.Player{},
	}
}

func (mds *MemoryDataSession) Get(path string, queryParams url.Values) (*resty.Response, error) {
	return nil, fmt.Errorf("no-op")
}
