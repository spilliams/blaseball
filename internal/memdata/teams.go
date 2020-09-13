package memdata

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataSession) GetAllTeams() ([]*model.Team, error) {
	return mds.allTeams, nil
}

func (mds *MemoryDataSession) GetTeamByID(id string) (*model.Team, error) {
	for _, t := range mds.allTeams {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, pkg.NewCodedError(fmt.Errorf("no team with id %s", id), http.StatusNotFound)
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

func (mds *MemoryDataSession) PutTeam(n *model.Team) error {
	for i, t := range mds.allTeams {
		if t.ID == n.ID {
			mds.allTeams[i] = n
			return nil
		}
	}
	mds.allTeams = append(mds.allTeams, n)
	return nil
}
