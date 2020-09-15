package apiserver

import (
	"net/http"

	"github.com/spilliams/blaseball/pkg"
)

func (s *Server) GetAllGameStatsheets(w http.ResponseWriter, r *http.Request) error {
	sheets, err := s.dataStore.GetAllGameStatsheets()
	if err != nil {
		return err
	}

	incompleteSheets := make([]string, 0, len(sheets))
	for _, sheet := range sheets {
		if sheet.Incomplete() {
			incompleteSheets = append(incompleteSheets, sheet.ID)
		}
	}
	if len(incompleteSheets) == 0 {
		return marshalAndWrite(sheets, w, r)
	}

	l := loggerFromRequest(r)
	l.Infof("fetching %d complete game statesheets", len(incompleteSheets))
	completeSheets, err := s.remoteAPI.GetGameStatsheetsByID(incompleteSheets)
	if err != nil {
		return err
	}
	if err := s.dataStore.PutGameStatsheets(completeSheets); err != nil {
		return err
	}
	sheets, err = s.dataStore.GetAllGameStatsheets()
	if err != nil {
		return err
	}

	return marshalAndWrite(sheets, w, r)
}

func (s *Server) GetGameStatsheet(w http.ResponseWriter, r *http.Request) error {
	id := getQueryString(r, "id")
	if len(id) == 0 {
		return pkg.NewCodedErrorf(http.StatusBadRequest, "`id` must be specified in query parameters")
	}
	sheet, err := s.dataStore.GetGameStatsheetByID(id)
	fetchFromRemote := err != nil || sheet == nil || sheet.Incomplete()
	if err != nil {
		l := loggerFromRequest(r)
		l.Warnf("couldn't fetch game statsheet: %v", err)
	}
	if !fetchFromRemote {
		return marshalAndWrite(sheet, w, r)
	}

	sheets, err := s.remoteAPI.GetGameStatsheetsByID([]string{id})
	if err != nil {
		return err
	}
	if sheets == nil || len(sheets) == 0 {
		return pkg.NewCodedErrorf(http.StatusNotFound, "no Game Statsheet found with id '%s'", id)
	}
	if err = s.dataStore.PutGameStatsheet(sheets[0]); err != nil {
		return err
	}

	return marshalAndWrite(sheets[0], w, r)
}

func (s *Server) GetAllPlayerSeasonStatsheets(w http.ResponseWriter, r *http.Request) error {
	sheets, err := s.dataStore.GetAllPlayerSeasonStatsheets()
	if err != nil {
		return err
	}

	incompleteSheets := make([]string, 0, len(sheets))
	for _, sheet := range sheets {
		if sheet.Incomplete() {
			incompleteSheets = append(incompleteSheets, sheet.ID)
		}
	}
	if len(incompleteSheets) == 0 {
		return marshalAndWrite(sheets, w, r)
	}

	l := loggerFromRequest(r)
	l.Infof("fetching %d complete player season statesheets", len(incompleteSheets))
	completeSheets, err := s.remoteAPI.GetPlayerSeasonStatsheetsByID(incompleteSheets)
	if err != nil {
		return err
	}
	if err := s.dataStore.PutPlayerSeasonStatsheets(completeSheets); err != nil {
		return err
	}
	sheets, err = s.dataStore.GetAllPlayerSeasonStatsheets()
	if err != nil {
		return err
	}

	return marshalAndWrite(sheets, w, r)
}

func (s *Server) GetPlayerSeasonStatsheets(w http.ResponseWriter, r *http.Request) error {
	playerID := getQueryString(r, "playerID")
	if len(playerID) == 0 {
		return pkg.NewCodedErrorf(http.StatusBadRequest, "`playerID` must be specified in query parameters")
	}
	sheets, err := s.dataStore.GetPlayerSeasonStatsheetsByPlayerID(playerID)
	if err != nil {
		return err
	}
	return marshalAndWrite(sheets, w, r)
}

func (s *Server) GetPlayerSeasonStatsheet(w http.ResponseWriter, r *http.Request) error {
	id := getQueryString(r, "id")
	if len(id) == 0 {
		return pkg.NewCodedErrorf(http.StatusBadRequest, "`id` must be specified in query parameters")
	}
	sheet, err := s.dataStore.GetPlayerSeasonStatsheetByID(id)
	fetchFromRemote := err != nil || sheet == nil || sheet.Incomplete()
	if err != nil {
		l := loggerFromRequest(r)
		l.Warnf("couldn't fetch player season statsheet: %v", err)
	}
	if !fetchFromRemote {
		return marshalAndWrite(sheet, w, r)
	}

	sheets, err := s.remoteAPI.GetPlayerSeasonStatsheetsByID([]string{id})
	if err != nil {
		return err
	}
	if len(sheets) == 0 {
		return pkg.NewCodedErrorf(http.StatusNotFound, "no Player Season Statsheet found with id '%s'", id)
	}
	if err = s.dataStore.PutPlayerSeasonStatsheet(sheets[0]); err != nil {
		return err
	}
	return marshalAndWrite(sheets[0], w, r)
}

func (s *Server) GetAllSeasonStatsheets(w http.ResponseWriter, r *http.Request) error {
	sheets, err := s.dataStore.GetAllSeasonStatsheets()
	if err != nil {
		return err
	}

	incompleteSheets := make([]string, 0, len(sheets))
	for _, sheet := range sheets {
		if sheet.Incomplete() {
			incompleteSheets = append(incompleteSheets, sheet.ID)
		}
	}
	if len(incompleteSheets) == 0 {
		return marshalAndWrite(sheets, w, r)
	}

	l := loggerFromRequest(r)
	l.Infof("fetching %d complete season statesheets", len(incompleteSheets))
	completeSheets, err := s.remoteAPI.GetSeasonStatsheetsByID(incompleteSheets)
	if err != nil {
		return err
	}
	if err := s.dataStore.PutSeasonStatsheets(completeSheets); err != nil {
		return err
	}
	sheets, err = s.dataStore.GetAllSeasonStatsheets()
	if err != nil {
		return err
	}

	return marshalAndWrite(sheets, w, r)
}

func (s *Server) GetSeasonStatsheet(w http.ResponseWriter, r *http.Request) error {
	id := getQueryString(r, "id")
	if len(id) == 0 {
		return pkg.NewCodedErrorf(http.StatusBadRequest, "`id` must be specified in query parameters")
	}
	sheet, err := s.dataStore.GetSeasonStatsheetByID(id)
	fetchFromRemote := err != nil || sheet == nil || sheet.Incomplete()
	if err != nil {
		l := loggerFromRequest(r)
		l.Warnf("couldn't fetch season statsheet: %v", err)
	}
	if !fetchFromRemote {
		return marshalAndWrite(sheet, w, r)
	}

	sheets, err := s.remoteAPI.GetSeasonStatsheetsByID([]string{id})
	if err != nil {
		return err
	}
	if len(sheets) == 0 {
		return pkg.NewCodedErrorf(http.StatusNotFound, "no Season Statsheet found with id '%s'", id)
	}
	if err = s.dataStore.PutSeasonStatsheet(sheets[0]); err != nil {
		return err
	}
	return marshalAndWrite(sheets[0], w, r)
}

func (s *Server) GetAllTeamStatsheets(w http.ResponseWriter, r *http.Request) error {
	sheets, err := s.dataStore.GetAllTeamStatsheets()
	if err != nil {
		return err
	}

	incompleteSheets := make([]string, 0, len(sheets))
	for _, sheet := range sheets {
		if sheet.Incomplete() {
			incompleteSheets = append(incompleteSheets, sheet.ID)
		}
	}
	if len(incompleteSheets) == 0 {
		return marshalAndWrite(sheets, w, r)
	}

	l := loggerFromRequest(r)
	l.Infof("fetching %d complete team statesheets", len(incompleteSheets))
	completeSheets, err := s.remoteAPI.GetTeamStatsheetsByID(incompleteSheets)
	if err != nil {
		return err
	}
	if err := s.dataStore.PutTeamStatsheets(completeSheets); err != nil {
		return err
	}
	sheets, err = s.dataStore.GetAllTeamStatsheets()
	if err != nil {
		return err
	}

	return marshalAndWrite(sheets, w, r)
}

func (s *Server) GetTeamStatsheet(w http.ResponseWriter, r *http.Request) error {
	id := getQueryString(r, "id")
	if len(id) == 0 {
		return pkg.NewCodedErrorf(http.StatusBadRequest, "`id` must be specified in query parameters")
	}
	sheet, err := s.dataStore.GetTeamStatsheetByID(id)
	fetchFromRemote := err != nil || sheet == nil || sheet.Incomplete()
	if err != nil {
		l := loggerFromRequest(r)
		l.Warnf("couldn't fetch team statsheet: %v", err)
	}
	if !fetchFromRemote {
		return marshalAndWrite(sheet, w, r)
	}

	sheets, err := s.remoteAPI.GetTeamStatsheetsByID([]string{id})
	if err != nil {
		return err
	}
	if len(sheets) == 0 {
		return pkg.NewCodedErrorf(http.StatusNotFound, "no Team Statsheet found with id '%s'", id)
	}
	if err = s.dataStore.PutTeamStatsheet(sheets[0]); err != nil {
		return err
	}
	return marshalAndWrite(sheets[0], w, r)
}
