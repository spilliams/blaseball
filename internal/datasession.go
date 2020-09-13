package internal

import (
	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

// LocalDataSession represents a session with a local data source (such as a
// database or in-memory store)
type LocalDataSession interface {
	ReadableDataSession
	writableDataSession
}

// ReadableDataSession represents a data session that is readable. This is
// meant to be 1:1 api-compatible with the web server built from this codebase.
type ReadableDataSession interface {
	pkg.RemoteDataSession

	GetDivisionByName(string) (*model.Division, error)
	GetTeamByFullName(string) (*model.Team, error)
	GetTeamByNickname(string) (*model.Team, error)
	GetAllPlayers() ([]*model.Player, error)
	GetPlayerByID(string) (*model.Player, error)
	GetPlayerByName(string) (*model.Player, error)
}

// writableDataSession represents a data session that is writable.
type writableDataSession interface {
	PutDivision(*model.Division) error
	PutPlayer(*model.Player) error
	PutPlayers([]*model.Player) error
	PutTeam(*model.Team) error
}
