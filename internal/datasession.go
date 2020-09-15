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
	PutDivision(*model.Division) error

	GetAllPlayers() (*model.PlayerList, error)
	GetPlayersByID([]string) (*model.PlayerList, error)
	GetPlayerByID(string) (*model.Player, error)
	GetPlayerByName(string) (*model.Player, error)
	PutPlayer(*model.Player) error
	PutPlayers([]*model.Player) error

	GetAllTeams() (*model.TeamList, error)
	GetTeamByID(string) (*model.Team, error)
	GetTeamByFullName(string) (*model.Team, error)
	GetTeamByNickname(string) (*model.Team, error)
	PutTeam(*model.Team) error

	GetAllSeasonStatsheets() ([]*model.SeasonStatsheet, error)
	GetSeasonStatsheetByID(string) (*model.SeasonStatsheet, error)
	PutSeasonStatsheet(*model.SeasonStatsheet) error
	PutSeasonStatsheets([]*model.SeasonStatsheet) error

	GetAllGameStatsheets() ([]*model.GameStatsheet, error)
	GetGameStatsheetByID(string) (*model.GameStatsheet, error)
	PutGameStatsheet(*model.GameStatsheet) error
	PutGameStatsheets([]*model.GameStatsheet) error

	GetAllTeamStatsheets() ([]*model.TeamStatsheet, error)
	GetTeamStatsheetByID(string) (*model.TeamStatsheet, error)
	PutTeamStatsheet(*model.TeamStatsheet) error
	PutTeamStatsheets([]*model.TeamStatsheet) error

	GetAllPlayerSeasonStatsheets() ([]*model.PlayerSeasonStatsheet, error)
	GetPlayerSeasonStatsheetsByPlayerID(string) ([]*model.PlayerSeasonStatsheet, error)
	GetPlayerSeasonStatsheetByID(string) (*model.PlayerSeasonStatsheet, error)
	PutPlayerSeasonStatsheet(*model.PlayerSeasonStatsheet) error
	PutPlayerSeasonStatsheets([]*model.PlayerSeasonStatsheet) error
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

	GetAllSeasonStatsheets() ([]*model.SeasonStatsheet, error)
	GetSeasonStatsheetByID(string) (*model.SeasonStatsheet, error)

	GetAllGameStatsheets() ([]*model.GameStatsheet, error)
	GetGameStatsheetByID(string) (*model.GameStatsheet, error)

	GetAllTeamStatsheets() ([]*model.TeamStatsheet, error)
	GetTeamStatsheetByID(string) (*model.TeamStatsheet, error)

	GetAllPlayerSeasonStatsheets() ([]*model.PlayerSeasonStatsheet, error)
	GetPlayerSeasonStatsheetsByPlayerID(string) ([]*model.PlayerSeasonStatsheet, error)
	GetPlayerSeasonStatsheetByID(string) (*model.PlayerSeasonStatsheet, error)
}
