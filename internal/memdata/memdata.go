package memdata

import (
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/pkg/model"
)

// MemoryDataStore represents an in-memory set of blaseball data
type MemoryDataStore struct {
	divisions              map[string]*model.Division
	gameStatsheets         map[string]*model.GameStatsheet
	players                map[string]*model.Player
	playerSeasonStatsheets map[string]*model.PlayerSeasonStatsheet
	seasonStatsheets       map[string]*model.SeasonStatsheet
	teams                  map[string]*model.Team
	teamStatsheets         map[string]*model.TeamStatsheet
}

// NewStore returns a new, empty in-memory data session
func NewStore() internal.DataStorageSession {
	return &MemoryDataStore{
		divisions:              map[string]*model.Division{},
		gameStatsheets:         map[string]*model.GameStatsheet{},
		players:                map[string]*model.Player{},
		playerSeasonStatsheets: map[string]*model.PlayerSeasonStatsheet{},
		seasonStatsheets:       map[string]*model.SeasonStatsheet{},
		teams:                  map[string]*model.Team{},
		teamStatsheets:         map[string]*model.TeamStatsheet{},
	}
}
