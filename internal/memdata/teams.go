package memdata

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataStore) GetAllTeams() ([]*model.Team, error) {
	teams := make([]*model.Team, 0, len(mds.allTeams))
	for _, t := range mds.allTeams {
		teams = append(teams, t)
	}
	return teams, nil
}

func (mds *MemoryDataStore) GetTeamByID(id string) (*model.Team, error) {
	team, ok := mds.allTeams[id]
	if !ok {
		return nil, pkg.NewCodedError(fmt.Errorf("no team with id %s", id), http.StatusNotFound)
	}
	return team, nil
}

func (mds *MemoryDataStore) GetTeamByFullName(name string) (*model.Team, error) {
	for _, t := range mds.allTeams {
		if strings.EqualFold(t.FullName, name) {
			return t, nil
		}
	}
	return nil, pkg.NewCodedError(fmt.Errorf("no team with name %s", name), http.StatusNotFound)
}

func (mds *MemoryDataStore) GetTeamByNickname(name string) (*model.Team, error) {
	for _, t := range mds.allTeams {
		if strings.EqualFold(t.Nickname, name) {
			return t, nil
		}
	}
	return nil, pkg.NewCodedError(fmt.Errorf("no team with name %s", name), http.StatusNotFound)
}

func (mds *MemoryDataStore) PutTeam(team *model.Team) error {
	mds.allTeams[team.ID] = team
	mds.seedPlayers(team.Lineup)
	mds.seedPlayers(team.Rotation)
	mds.seedPlayers(team.Bench)
	mds.seedPlayers(team.Bullpen)
	return nil
}

func (mds *MemoryDataStore) seedTeams(ids []string) {
	for _, id := range ids {
		_, ok := mds.allTeams[id]
		if !ok {
			mds.PutTeam(&model.Team{ID: id})
		}
	}
}
