package remotedata

import (
	"encoding/json"
	"fmt"

	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetAllTeams() ([]*model.Team, error) {
	resp, err := b.Get("allTeams", nil)
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
	resp, err := b.Get("team", map[string][]string{"id": {id}})
	if err != nil {
		return nil, err
	}
	var team *model.Team
	err = json.Unmarshal(resp.Body(), &team)
	return team, err
}
