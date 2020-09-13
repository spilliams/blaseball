package memdata

import (
	"fmt"

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
	return nil, fmt.Errorf("no division with id %s", id)
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
