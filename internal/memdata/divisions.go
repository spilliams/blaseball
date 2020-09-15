package memdata

import (
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataStore) GetAllDivisions() (*model.DivisionList, error) {
	divisions := make([]*model.Division, 0, len(mds.divisions))
	for _, d := range mds.divisions {
		copy := *d
		divisions = append(divisions, &copy)
	}
	return &model.DivisionList{List: divisions}, nil
}

func (mds *MemoryDataStore) GetDivisionByID(id string) (*model.Division, error) {
	division, ok := mds.divisions[id]
	if !ok {
		return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no Division with id %s", id)
	}
	copy := *division
	return &copy, nil
}

func (mds *MemoryDataStore) GetDivisionByName(name string) (*model.Division, error) {
	for _, d := range mds.divisions {
		if strings.EqualFold(d.Name, name) {
			copy := *d
			return &copy, nil
		}
	}
	return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no Division with name %s", name)
}

func (mds *MemoryDataStore) PutDivision(div *model.Division) error {
	mds.divisions[div.ID] = div
	return mds.seedTeams(div.TeamIDs)
}
