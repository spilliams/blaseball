package officialdata

import (
	"encoding/json"

	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetAllTeams() ([]*model.Team, error) {
	resp, err := b.get("allTeams", nil)
	if err != nil {
		return nil, err
	}
	var teams []*model.Team
	err = json.Unmarshal(resp.Body(), &teams)
	return teams, err
}

func (b *BlaseballAPI) GetTeamByID(id string) (*model.Team, error) {
	resp, err := b.get("team", map[string][]string{"id": {id}})
	if err != nil {
		return nil, err
	}
	var team *model.Team
	err = json.Unmarshal(resp.Body(), &team)
	return team, err
}
