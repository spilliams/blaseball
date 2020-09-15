package officialdata

import (
	"encoding/json"

	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetAllDivisions() ([]*model.Division, error) {
	resp, err := b.get("allDivisions", nil)
	if err != nil {
		return nil, err
	}
	var divisions []*model.Division
	err = json.Unmarshal(resp.Body(), &divisions)
	return divisions, err
}

func (b *BlaseballAPI) GetDivisionByID(id string) (*model.Division, error) {
	resp, err := b.get("division", map[string][]string{"id": {id}})
	if err != nil {
		return nil, err
	}

	var division *model.Division
	err = json.Unmarshal(resp.Body(), &division)
	return division, err
}
