package pkg

import "github.com/spilliams/blaseball/pkg/model"

// RemoteDataSession represents a data session with a remote data source (such
// as the official Blaseball API).
type RemoteDataSession interface {
	DataSessionDivisionsRead
	DataSessionTeamsRead
}

// DataSessionDivisionsRead represents a data session that can read things
// about Divisions.
type DataSessionDivisionsRead interface {
	GetAllDivisions() ([]*model.Division, error)
	GetDivisionByID(string) (*model.Division, error)
	GetDivisionByName(string) (*model.Division, error)
}

type DataSessionTeamsRead interface {
	GetAllTeams() ([]*model.Team, error)
	GetTeamByID(string) (*model.Team, error)
	GetTeamByFullName(string) (*model.Team, error)
	GetTeamByNickname(string) (*model.Team, error)
}
