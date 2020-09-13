package api

import (
	"encoding/json"
	"fmt"

	"github.com/spilliams/blaseball/pkg/model"
)

// ListAllDivisions returns a list of all team divisions
func ListAllDivisions() ([]*model.Division, error) {
	url := "https://www.blaseball.com/database/allDivisions"
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	var divisions []*model.Division
	if err = json.Unmarshal(resp.Body(), &divisions); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response: %v", err)
	}

	return divisions, nil
}
