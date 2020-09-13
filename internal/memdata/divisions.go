package memdata

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataSession) GetAllDivisions() ([]*model.Division, error) {
	divisions := make([]*model.Division, 0, len(mds.allDivisions))
	for _, d := range mds.allDivisions {
		divisions = append(divisions, d)
	}
	return divisions, nil
}

func (mds *MemoryDataSession) GetDivisionByID(id string) (*model.Division, error) {
	division, ok := mds.allDivisions[id]
	if !ok {
		return nil, pkg.NewCodedError(fmt.Errorf("no Division with id %s", id), http.StatusNotFound)
	}
	return division, nil
}

func (mds *MemoryDataSession) GetDivisionByName(name string) (*model.Division, error) {
	for _, d := range mds.allDivisions {
		if strings.EqualFold(d.Name, name) {
			return d, nil
		}
	}
	return nil, pkg.NewCodedError(fmt.Errorf("no Division with name %s", name), http.StatusNotFound)
}

func (mds *MemoryDataSession) PutDivision(div *model.Division) error {
	mds.allDivisions[div.ID] = div
	// TODO: make sure there are team entries for all team ids in the division?
	// follow-up: make sure when fetching a team, if it only has an ID then it's stale
	return nil
}
