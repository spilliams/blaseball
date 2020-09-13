package pkg

import (
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/spilliams/blaseball/pkg/model"
)

// RemoteDataSession represents a data session with a remote source. This is
// intended to be 1:1 api-compatible with the official Blaseball API.
type RemoteDataSession interface {
	Get(string, url.Values) (*resty.Response, error)

	GetAllDivisions() ([]*model.Division, error)
	GetDivisionByID(string) (*model.Division, error)

	GetPlayersByID([]string) ([]*model.Player, error)

	GetAllTeams() ([]*model.Team, error)
	GetTeamByID(string) (*model.Team, error)
}
