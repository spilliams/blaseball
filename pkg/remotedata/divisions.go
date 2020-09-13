package remotedata

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetAllDivisions() ([]*model.Division, error) {
	resp, err := b.get("allDivisions", nil)
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
	resp, err := b.get("division", map[string][]string{"id": {id}})

	if err != nil {
		return nil, err
	}
	return divisionFromResponse(resp)
}

func (b *BlaseballAPI) GetDivisionByName(name string) (*model.Division, error) {
	resp, err := b.get("division", map[string][]string{"name": {name}})
	if err != nil {
		return nil, err
	}
	return divisionFromResponse(resp)
}

func divisionFromResponse(resp *resty.Response) (*model.Division, error) {
	var division *model.Division
	err := json.Unmarshal(resp.Body(), &division)
	return division, err
}
