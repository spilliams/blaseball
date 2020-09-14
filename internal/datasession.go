package internal

import (
	"github.com/spilliams/blaseball/pkg/model"
)

// DataStorageSession represents a session with a data store, such as the
// server built from this codebase might use to store data
type DataStorageSession interface {
	GetAllDivisions() (*model.DivisionList, error)
	GetDivisionByID(string) (*model.Division, error)
	GetDivisionByName(string) (*model.Division, error)

	GetAllPlayers() (*model.PlayerList, error)
	GetPlayersByID([]string) (*model.PlayerList, error)
	GetPlayerByID(string) (*model.Player, error)
	GetPlayerByName(string) (*model.Player, error)

	GetAllTeams() (*model.TeamList, error)
	GetTeamByID(string) (*model.Team, error)
	GetTeamByFullName(string) (*model.Team, error)
	GetTeamByNickname(string) (*model.Team, error)

	PutDivision(*model.Division) error
	PutPlayer(*model.Player) error
	PutPlayers([]*model.Player) error
	PutTeam(*model.Team) error
}

// ServerDataSession represents a data session with the server built from this
// codebase. It is meant to be 1:1 API-compatible with the HTTP interface of
// this server.
type ServerDataSession interface {
	GetAllDivisions(bool) (*model.DivisionList, error)
	GetDivisionByID(string, bool) (*model.Division, error)
	GetDivisionByName(string, bool) (*model.Division, error)

	GetAllPlayers(bool) (*model.PlayerList, error)
	GetPlayersByID([]string, bool) (*model.PlayerList, error)
	GetPlayerByID(string, bool) (*model.Player, error)
	GetPlayerByName(string, bool) (*model.Player, error)

	GetAllTeams(bool) (*model.TeamList, error)
	GetTeamByID(string, bool) (*model.Team, error)
	GetTeamByFullName(string, bool) (*model.Team, error)
	GetTeamByNickname(string, bool) (*model.Team, error)
}
