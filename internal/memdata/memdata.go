package memdata

import (
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/pkg/model"
)

// MemoryDataStore represents an in-memory set of blaseball data
type MemoryDataStore struct {
	allDivisions map[string]*model.Division
	allTeams     map[string]*model.Team
	allPlayers   map[string]*model.Player
}

// NewSession returns a new, empty in-memory data session
func NewSession() internal.DataStorageSession {
	return &MemoryDataStore{
		allDivisions: map[string]*model.Division{},
		allTeams:     map[string]*model.Team{},
		allPlayers:   map[string]*model.Player{},
	}
}
