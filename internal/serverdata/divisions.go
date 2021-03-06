package serverdata

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetAllDivisions(showFK bool) (*model.DivisionList, error) {
	q := addShowFKQuery(nil, showFK)
	resp, err := b.get("allDivisions", q)
	if err != nil {
		return nil, err
	}
	var divisions []*model.Division
	err = json.Unmarshal(resp.Body(), &divisions)
	return &model.DivisionList{divisions}, err
}

func (b *BlaseballAPI) GetDivisionByID(id string, showFK bool) (*model.Division, error) {
	q := addShowFKQuery(map[string][]string{"id": {id}}, showFK)
	resp, err := b.get("division", q)

	if err != nil {
		return nil, err
	}
	return divisionFromResponse(resp)
}

func (b *BlaseballAPI) GetDivisionByName(name string, showFK bool) (*model.Division, error) {
	q := addShowFKQuery(map[string][]string{"name": {name}}, showFK)
	resp, err := b.get("division", q)
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
