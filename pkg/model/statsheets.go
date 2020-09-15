package model

type SeasonStatsheet struct {
	ID        string
	TeamStats []string
}

func (ss *SeasonStatsheet) Incomplete() bool {
	return ss.TeamStats == nil || len(ss.TeamStats) == 0
}

type GameStatsheet struct {
	ID                   string
	HomeTeamRunsByInning []int
	AwayTeamRunsByInning []int
	AwayTeamTotalBatters int
	HomeTeamTotalBatters int
	AwayTeamStats        string
	HomeTeamStats        string
}

func (gs *GameStatsheet) Incomplete() bool {
	return gs.AwayTeamStats == ""
}

type TeamStatsheet struct {
	ID          string
	PlayerStats []string
	GamesPlayed int
	Wins        int
	Losses      int
	Name        string
	TeamID      string
}

func (ts *TeamStatsheet) Incomplete() bool {
	return ts.TeamID == ""
}

type PlayerSeasonStatsheet struct {
	ID             string
	PlayerID       string
	TeamID         string
	Team           string
	Name           string
	AtBats         int
	CaughtStealing int
	Doubles        int
	EarnedRuns     int
	GroundIntoDP   int
	Hits           int
	HitsAllowed    int
	HomeRuns       int
	Losses         int
	OutsRecorded   int
	RBIs           int
	Runs           int
	StolenBases    int
	Strikeouts     int
	Struckouts     int
	Triples        int
	Walks          int
	WalksIssued    int
	Wins           int
}

func (ps *PlayerSeasonStatsheet) Incomplete() bool {
	return ps.PlayerID == ""
}
