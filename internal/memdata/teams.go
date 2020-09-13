package memdata

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataSession) GetAllTeams() ([]*model.Team, error) {
	teams := make([]*model.Team, 0, len(mds.allTeams))
	for _, t := range mds.allTeams {
		teams = append(teams, t)
	}
	return teams, nil
}

func (mds *MemoryDataSession) GetTeamByID(id string) (*model.Team, error) {
	team, ok := mds.allTeams[id]
	if !ok {
		return nil, pkg.NewCodedError(fmt.Errorf("no team with id %s", id), http.StatusNotFound)
	}
	return team, nil
}

func (mds *MemoryDataSession) GetTeamByFullName(name string) (*model.Team, error) {
	for _, t := range mds.allTeams {
		if strings.EqualFold(t.FullName, name) {
			return t, nil
		}
	}
	return nil, pkg.NewCodedError(fmt.Errorf("no team with name %s", name), http.StatusNotFound)
}

func (mds *MemoryDataSession) GetTeamByNickname(name string) (*model.Team, error) {
	for _, t := range mds.allTeams {
		if strings.EqualFold(t.Nickname, name) {
			return t, nil
		}
	}
	return nil, pkg.NewCodedError(fmt.Errorf("no team with name %s", name), http.StatusNotFound)
}

func (mds *MemoryDataSession) PutTeam(team *model.Team) error {
	mds.allTeams[team.ID] = team
	return nil
}
