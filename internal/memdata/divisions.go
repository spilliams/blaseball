package memdata

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataSession) GetAllDivisions() ([]*model.Division, error) {
	return mds.allDivisions, nil
}

func (mds *MemoryDataSession) GetDivisionByID(id string) (*model.Division, error) {
	for _, d := range mds.allDivisions {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, pkg.NewCodedError(fmt.Errorf("no division with id %s", id), http.StatusNotFound)
}

func (mds *MemoryDataSession) GetDivisionByName(name string) (*model.Division, error) {
	for _, d := range mds.allDivisions {
		if strings.EqualFold(d.Name, name) {
			return d, nil
		}
	}
	return nil, pkg.NewCodedError(fmt.Errorf("no division with name %s", name), http.StatusNotFound)
}

func (mds *MemoryDataSession) PutDivision(n *model.Division) error {
	for i, d := range mds.allDivisions {
		if d.ID == n.ID {
			mds.allDivisions[i] = n
			return nil
		}
	}
	mds.allDivisions = append(mds.allDivisions, n)
	return nil
}
