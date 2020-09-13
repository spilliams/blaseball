package api

import (
	"encoding/json"
	"fmt"

	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetAllDivisions() ([]*model.Division, error) {
	resp, err := b.get("allDivisions")
	if err != nil {
		return nil, err
	}
	var divisions []*model.Division
	if err = json.Unmarshal(resp.Body(), &divisions); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
	}

	return divisions, nil
}

func (b *BlaseballAPI) GetDivisionByID(id string) (*model.Division, error) {
	resp, err := b.get(fmt.Sprintf("division?id=%s", id))
	if err != nil {
		return nil, err
	}
	var division *model.Division
	if err = json.Unmarshal(resp.Body(), &division); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
	}

	return division, nil
}
