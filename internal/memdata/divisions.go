package memdata

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataStore) GetAllDivisions() (*model.DivisionList, error) {
	divisions := make([]*model.Division, 0, len(mds.allDivisions))
	for _, d := range mds.allDivisions {
		divisions = append(divisions, d)
	}
	return &model.DivisionList{divisions}, nil
}

func (mds *MemoryDataStore) GetDivisionByID(id string) (*model.Division, error) {
	division, ok := mds.allDivisions[id]
	if !ok {
		return nil, pkg.NewCodedError(fmt.Errorf("no Division with id %s", id), http.StatusNotFound)
	}
	return division, nil
}

func (mds *MemoryDataStore) GetDivisionByName(name string) (*model.Division, error) {
	for _, d := range mds.allDivisions {
		if strings.EqualFold(d.Name, name) {
			return d, nil
		}
	}
	return nil, pkg.NewCodedError(fmt.Errorf("no Division with name %s", name), http.StatusNotFound)
}

func (mds *MemoryDataStore) PutDivision(div *model.Division) error {
	mds.allDivisions[div.ID] = div
	mds.seedTeams(div.TeamIDs)
	return nil
}
