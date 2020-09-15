package officialdata

import (
	"encoding/json"

	"github.com/spilliams/blaseball/pkg/model"
)

func (b *BlaseballAPI) GetSeasonStatsheetsByID(ids []string) ([]*model.SeasonStatsheet, error) {
	bodies, err := b.getAll("seasonStatsheets", "ids", ids)
	if err != nil {
		return nil, err
	}

	var seasonStats []*model.SeasonStatsheet
	for _, body := range bodies {
		var respStats []*model.SeasonStatsheet
		if err := json.Unmarshal(body, &respStats); err != nil {
			return nil, err
		}
		seasonStats = append(seasonStats, respStats...)
	}

	return seasonStats, nil
}

func (b *BlaseballAPI) GetGameStatsheetsByID(ids []string) ([]*model.GameStatsheet, error) {
	bodies, err := b.getAll("gameStatsheets", "ids", ids)
	if err != nil {
		return nil, err
	}

	var gameStats []*model.GameStatsheet
	for _, body := range bodies {
		var respStats []*model.GameStatsheet
		if err := json.Unmarshal(body, &respStats); err != nil {
			return nil, err
		}
		gameStats = append(gameStats, respStats...)
	}

	return gameStats, nil
}

func (b *BlaseballAPI) GetTeamStatsheetsByID(ids []string) ([]*model.TeamStatsheet, error) {
	bodies, err := b.getAll("teamStatsheets", "ids", ids)
	if err != nil {
		return nil, err
	}

	var teamStats []*model.TeamStatsheet
	for _, body := range bodies {
		var respStats []*model.TeamStatsheet
		if err := json.Unmarshal(body, &respStats); err != nil {
			return nil, err
		}
		teamStats = append(teamStats, respStats...)
	}

	return teamStats, nil
}

func (b *BlaseballAPI) GetPlayerSeasonStatsheetsByID(ids []string) ([]*model.PlayerSeasonStatsheet, error) {
	bodies, err := b.getAll("playerStatsheets", "ids", ids)
	if err != nil {
		return nil, err
	}

	var playerStats []*model.PlayerSeasonStatsheet
	for _, body := range bodies {
		var respStats []*model.PlayerSeasonStatsheet
		if err := json.Unmarshal(body, &respStats); err != nil {
			return nil, err
		}
		playerStats = append(playerStats, respStats...)
	}

	return playerStats, nil
}
