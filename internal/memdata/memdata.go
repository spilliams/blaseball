package memdata

import (
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/pkg/model"
)

// MemoryDataSession represents an in-memory set of blaseball data
type MemoryDataSession struct {
	allDivisions []*model.Division
	allTeams     []*model.Team
}

// NewSession returns a new, empty in-memory data session
func NewSession() internal.LocalDataSession {
	return &MemoryDataSession{
		allDivisions: []*model.Division{},
		allTeams:     []*model.Team{},
	}
}
