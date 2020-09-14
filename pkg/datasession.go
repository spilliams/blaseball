package pkg

import (
	"github.com/spilliams/blaseball/pkg/model"
)

// OfficialDataSession represents a data session with the official Blaseball
// API (https://www.blaseball.com/database/).
type OfficialDataSession interface {
	GetAllDivisions() ([]*model.Division, error)
	GetDivisionByID(string) (*model.Division, error)

	GetPlayersByID([]string) ([]*model.Player, error)

	GetAllTeams() ([]*model.Team, error)
	GetTeamByID(string) (*model.Team, error)
}
