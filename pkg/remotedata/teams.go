package remotedata

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetAllTeams() ([]*model.Team, error) {
	resp, err := b.get("allTeams", nil)
	if err != nil {
		return nil, err
	}
	var teams []*model.Team
	if err = json.Unmarshal(resp.Body(), &teams); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
	}

	return teams, nil
}

func (b *BlaseballAPI) GetTeamByID(id string) (*model.Team, error) {
	resp, err := b.get("team", map[string][]string{"id": {id}})
	if err != nil {
		return nil, err
	}
	return teamFromResponse(resp)
}

func (b *BlaseballAPI) GetTeamByFullName(name string) (*model.Team, error) {
	resp, err := b.get("team", map[string][]string{"name": {name}})
	if err != nil {
		return nil, err
	}
	return teamFromResponse(resp)
}

func (b *BlaseballAPI) GetTeamByNickname(name string) (*model.Team, error) {
	resp, err := b.get("team", map[string][]string{"name": {name}})
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
