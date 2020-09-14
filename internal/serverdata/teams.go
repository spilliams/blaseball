package serverdata

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetAllTeams(showFK bool) (*model.TeamList, error) {
	q := addShowFKQuery(nil, showFK)
	resp, err := b.get("allTeams", q)
	if err != nil {
		return nil, err
	}
	var teams []*model.Team
	if err = json.Unmarshal(resp.Body(), &teams); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
	}

	return &model.TeamList{teams}, nil
}

func (b *BlaseballAPI) GetTeamByID(id string, showFK bool) (*model.Team, error) {
	q := addShowFKQuery(map[string][]string{"id": {id}}, showFK)
	resp, err := b.get("team", q)
	if err != nil {
		return nil, err
	}
	return teamFromResponse(resp)
}

func (b *BlaseballAPI) GetTeamByFullName(name string, showFK bool) (*model.Team, error) {
	q := addShowFKQuery(map[string][]string{"name": {name}}, showFK)
	resp, err := b.get("team", q)
	if err != nil {
		return nil, err
	}
	return teamFromResponse(resp)
}

func (b *BlaseballAPI) GetTeamByNickname(name string, showFK bool) (*model.Team, error) {
	q := addShowFKQuery(map[string][]string{"name": {name}}, showFK)
	resp, err := b.get("team", q)
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
