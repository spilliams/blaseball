package serverdata

import (
	"encoding/json"

	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetAllSeasonStatsheets() ([]*model.SeasonStatsheet, error) {
	resp, err := b.get("allSeasonStatsheets", nil)
	if err != nil {
		return nil, err
	}
	var sheets []*model.SeasonStatsheet
	err = json.Unmarshal(resp.Body(), &sheets)
	return sheets, err
}

func (b *BlaseballAPI) GetSeasonStatsheetByID(id string) (*model.SeasonStatsheet, error) {
	resp, err := b.get("seasonStatsheet", map[string][]string{"id": {id}})
	if err != nil {
		return nil, err
	}
	var sheet *model.SeasonStatsheet
	err = json.Unmarshal(resp.Body(), &sheet)
	return sheet, err
}

func (b *BlaseballAPI) GetAllGameStatsheets() ([]*model.GameStatsheet, error) {
	resp, err := b.get("allGameStatsheets", nil)
	if err != nil {
		return nil, err
	}
	var sheets []*model.GameStatsheet
	err = json.Unmarshal(resp.Body(), &sheets)
	return sheets, err
}

func (b *BlaseballAPI) GetGameStatsheetByID(id string) (*model.GameStatsheet, error) {
	resp, err := b.get("gameStatsheet", map[string][]string{"id": {id}})
	if err != nil {
		return nil, err
	}
	var sheet *model.GameStatsheet
	err = json.Unmarshal(resp.Body(), &sheet)
	return sheet, err
}

func (b *BlaseballAPI) GetAllTeamStatsheets() ([]*model.TeamStatsheet, error) {
	resp, err := b.get("allTeamStatsheets", nil)
	if err != nil {
		return nil, err
	}
	var sheets []*model.TeamStatsheet
	err = json.Unmarshal(resp.Body(), &sheets)
	return sheets, err
}

func (b *BlaseballAPI) GetTeamStatsheetByID(id string) (*model.TeamStatsheet, error) {
	resp, err := b.get("teamStatsheet", map[string][]string{"id": {id}})
	if err != nil {
		return nil, err
	}
	var sheet *model.TeamStatsheet
	err = json.Unmarshal(resp.Body(), &sheet)
	return sheet, err
}

func (b *BlaseballAPI) GetAllPlayerSeasonStatsheets() ([]*model.PlayerSeasonStatsheet, error) {
	resp, err := b.get("allPlayerSeasonStatsheets", nil)
	if err != nil {
		return nil, err
	}
	var sheets []*model.PlayerSeasonStatsheet
	err = json.Unmarshal(resp.Body(), &sheets)
	return sheets, err
}

func (b *BlaseballAPI) GetPlayerSeasonStatsheetsByPlayerID(playerID string) ([]*model.PlayerSeasonStatsheet, error) {
	resp, err := b.get("playerSeasonStatsheets", map[string][]string{"playerID": {playerID}})
	if err != nil {
		return nil, err
	}
	var sheets []*model.PlayerSeasonStatsheet
	err = json.Unmarshal(resp.Body(), &sheets)
	return sheets, err
}

func (b *BlaseballAPI) GetPlayerSeasonStatsheetByID(id string) (*model.PlayerSeasonStatsheet, error) {
	resp, err := b.get("playerSeasonStatsheet", map[string][]string{"id": {id}})
	if err != nil {
		return nil, err
	}
	var sheet *model.PlayerSeasonStatsheet
	err = json.Unmarshal(resp.Body(), &sheet)
	return sheet, err
}
