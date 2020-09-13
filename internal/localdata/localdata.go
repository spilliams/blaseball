package localdata

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
	"github.com/spilliams/blaseball/pkg/remotedata"
)

type BlaseballAPI struct {
	pkg.RemoteDataSession
}

func NewAPI(blaseURL string, blasePath string, logLevel logrus.Level) internal.ReadableDataSession {
	b := remotedata.NewAPI(blaseURL, blasePath, logLevel)
	return &BlaseballAPI{b}
}

func (b *BlaseballAPI) GetDivisionByName(name string) (*model.Division, error) {
	resp, err := b.Get("division", map[string][]string{"name": {name}})
	if err != nil {
		return nil, err
	}
	var division *model.Division
	err = json.Unmarshal(resp.Body(), &division)
	return division, err
}

func (b *BlaseballAPI) GetAllPlayers() ([]*model.Player, error) {
	resp, err := b.Get("allPlayers", nil)
	if err != nil {
		return nil, err
	}
	var players []*model.Player
	if err = json.Unmarshal(resp.Body(), &players); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
	}

	return players, nil
}

func (b *BlaseballAPI) GetPlayerByID(id string) (*model.Player, error) {
	players, err := b.GetPlayersByID([]string{id})
	if err != nil {
		return nil, err
	}
	if players == nil || len(players) == 0 {
		return nil, pkg.NewCodedError(fmt.Errorf("no Player with id %s", id), http.StatusNotFound)
	}
	return players[0], nil
}

func (b *BlaseballAPI) GetPlayerByName(name string) (*model.Player, error) {
	resp, err := b.Get("players", map[string][]string{"name": {name}})
	if err != nil {
		return nil, err
	}
	var player *model.Player
	if err = json.Unmarshal(resp.Body(), &player); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
	}

	return player, nil
}

func (b *BlaseballAPI) GetTeamByFullName(name string) (*model.Team, error) {
	resp, err := b.Get("team", map[string][]string{"name": {name}})
	if err != nil {
		return nil, err
	}
	return teamFromResponse(resp)
}

func (b *BlaseballAPI) GetTeamByNickname(name string) (*model.Team, error) {
	resp, err := b.Get("team", map[string][]string{"name": {name}})
	if err != nil {
		return nil, err
	}
	return teamFromResponse(resp)
}

func teamFromResponse(resp *resty.Response) (*model.Team, error) {
	var team *model.Team
	err := json.Unmarshal(resp.Body(), &team)
	return team, err
}
