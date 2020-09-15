package memdata

import (
	"net/http"

	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

func (mds *MemoryDataStore) GetAllSeasonStatsheets() ([]*model.SeasonStatsheet, error) {
	stats := make([]*model.SeasonStatsheet, 0, len(mds.seasonStatsheets))
	for _, s := range mds.seasonStatsheets {
		copy := *s
		stats = append(stats, &copy)
	}
	return stats, nil
}

func (mds *MemoryDataStore) GetSeasonStatsheetByID(id string) (*model.SeasonStatsheet, error) {
	sheet, ok := mds.seasonStatsheets[id]
	if !ok {
		return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no Season Statsheet with id %s", id)
	}
	copy := *sheet
	return &copy, nil
}

func (mds *MemoryDataStore) PutSeasonStatsheet(sheet *model.SeasonStatsheet) error {
	mds.seasonStatsheets[sheet.ID] = sheet
	return mds.seedTeamStatsheets(sheet.TeamStats)
}

func (mds *MemoryDataStore) PutSeasonStatsheets(sheets []*model.SeasonStatsheet) error {
	for _, sheet := range sheets {
		if err := mds.PutSeasonStatsheet(sheet); err != nil {
			return err
		}
	}
	return nil
}

func (mds *MemoryDataStore) GetAllGameStatsheets() ([]*model.GameStatsheet, error) {
	stats := make([]*model.GameStatsheet, 0, len(mds.gameStatsheets))
	for _, s := range mds.gameStatsheets {
		copy := *s
		stats = append(stats, &copy)
	}
	return stats, nil
}

func (mds *MemoryDataStore) GetGameStatsheetByID(id string) (*model.GameStatsheet, error) {
	sheet, ok := mds.gameStatsheets[id]
	if !ok {
		return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no Game Statsheet with id %s", id)
	}
	copy := *sheet
	return &copy, nil
}

func (mds *MemoryDataStore) PutGameStatsheet(sheet *model.GameStatsheet) error {
	mds.gameStatsheets[sheet.ID] = sheet
	// TODO: verify that these are indeed team sheets
	// return mds.seedTeamStatsheets([]string{sheet.AwayTeamStats, sheet.HomeTeamStats})
	return nil
}

func (mds *MemoryDataStore) PutGameStatsheets(sheets []*model.GameStatsheet) error {
	for _, sheet := range sheets {
		if err := mds.PutGameStatsheet(sheet); err != nil {
			return err
		}
	}
	return nil
}

func (mds *MemoryDataStore) GetAllTeamStatsheets() ([]*model.TeamStatsheet, error) {
	stats := make([]*model.TeamStatsheet, 0, len(mds.teamStatsheets))
	for _, s := range mds.teamStatsheets {
		copy := *s
		stats = append(stats, &copy)
	}
	return stats, nil
}

func (mds *MemoryDataStore) GetTeamStatsheetByID(id string) (*model.TeamStatsheet, error) {
	sheet, ok := mds.teamStatsheets[id]
	if !ok {
		return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no Team Statsheet with id %s", id)
	}
	copy := *sheet
	return &copy, nil
}

func (mds *MemoryDataStore) PutTeamStatsheet(sheet *model.TeamStatsheet) error {
	mds.teamStatsheets[sheet.ID] = sheet
	return mds.seedPlayerSeasonStatsheets(sheet.PlayerStats)
}

func (mds *MemoryDataStore) PutTeamStatsheets(sheets []*model.TeamStatsheet) error {
	for _, sheet := range sheets {
		if err := mds.PutTeamStatsheet(sheet); err != nil {
			return err
		}
	}
	return nil
}

func (mds *MemoryDataStore) seedTeamStatsheets(ids []string) error {
	for _, id := range ids {
		_, ok := mds.teamStatsheets[id]
		if !ok {
			if err := mds.PutTeamStatsheet(&model.TeamStatsheet{ID: id}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (mds *MemoryDataStore) GetAllPlayerSeasonStatsheets() ([]*model.PlayerSeasonStatsheet, error) {
	stats := make([]*model.PlayerSeasonStatsheet, 0, len(mds.playerSeasonStatsheets))
	for _, s := range mds.playerSeasonStatsheets {
		copy := *s
		stats = append(stats, &copy)
	}
	return stats, nil
}

func (mds *MemoryDataStore) GetPlayerSeasonStatsheetsByPlayerID(playerID string) ([]*model.PlayerSeasonStatsheet, error) {
	stats := make([]*model.PlayerSeasonStatsheet, 0)
	for _, sheet := range mds.playerSeasonStatsheets {
		if sheet.PlayerID != playerID {
			continue
		}
		copy := *sheet
		stats = append(stats, &copy)
	}
	return stats, nil
}

func (mds *MemoryDataStore) GetPlayerSeasonStatsheetByID(id string) (*model.PlayerSeasonStatsheet, error) {
	sheet, ok := mds.playerSeasonStatsheets[id]
	if !ok {
		return nil, pkg.NewCodedErrorf(http.StatusNotFound, "no Player Season Statsheet with id %s", id)
	}
	copy := *sheet
	return &copy, nil
}

func (mds *MemoryDataStore) PutPlayerSeasonStatsheet(sheet *model.PlayerSeasonStatsheet) error {
	mds.playerSeasonStatsheets[sheet.ID] = sheet
	return nil
}

func (mds *MemoryDataStore) PutPlayerSeasonStatsheets(sheets []*model.PlayerSeasonStatsheet) error {
	for _, sheet := range sheets {
		if err := mds.PutPlayerSeasonStatsheet(sheet); err != nil {
			return err
		}
	}
	return nil
}

func (mds *MemoryDataStore) seedPlayerSeasonStatsheets(ids []string) error {
	for _, id := range ids {
		_, ok := mds.playerSeasonStatsheets[id]
		if !ok {
			if err := mds.PutPlayerSeasonStatsheet(&model.PlayerSeasonStatsheet{ID: id}); err != nil {
				return err
			}
		}
	}
	return nil
}
