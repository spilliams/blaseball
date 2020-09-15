package memdata

import (
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataStore) GetAllTeams() (*model.TeamList, error) {
	teams := make([]*model.Team, 0, len(mds.teams))
	for _, t := range mds.teams {
		copy := *t
		teams = append(teams, &copy)
	}
	return &model.TeamList{List: teams}, nil
}

func (mds *MemoryDataStore) GetTeamByID(id string) (*model.Team, error) {
	team, ok := mds.teams[id]
	if !ok {
		return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no team with id %s", id)
	}
	copy := *team
	return &copy, nil
}

func (mds *MemoryDataStore) GetTeamByFullName(name string) (*model.Team, error) {
	for _, t := range mds.teams {
		if strings.EqualFold(t.FullName, name) {
			copy := *t
			return &copy, nil
		}
	}
	return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no team with name %s", name)
}

func (mds *MemoryDataStore) GetTeamByNickname(name string) (*model.Team, error) {
	for _, t := range mds.teams {
		if strings.EqualFold(t.Nickname, name) {
			copy := *t
			return &copy, nil
		}
	}
	return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no team with name %s", name)
}

func (mds *MemoryDataStore) PutTeam(team *model.Team) error {
	mds.teams[team.ID] = team
	playerLists := [][]string{team.Lineup, team.Rotation, team.Bench, team.Bullpen}
	players := make([]string, 0, len(team.Lineup)+len(team.Rotation)+len(team.Bench)+len(team.Bullpen))
	for _, list := range playerLists {
		players = append(players, list...)
	}
	return mds.seedPlayers(players)
}

func (mds *MemoryDataStore) seedTeams(ids []string) error {
	for _, id := range ids {
		_, ok := mds.teams[id]
		if !ok {
			if err := mds.PutTeam(&model.Team{ID: id}); err != nil {
				return err
			}
		}
	}
	return nil
}
