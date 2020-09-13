package model

import "fmt"

// Team represents a single Blaseball team. Teams contain players, and play
// games against other teams.
type Team struct {
	ID                  string
	FullName            string
	Nickname            string
	Shorthand           string
	Emoji               string
	Slogan              string
	Location            string
	MainColor           string
	SecondaryColor      string
	Lineup              []string
	Rotation            []string
	Bullpen             []string
	Bench               []string
	ShameRuns           int
	TotalShames         int
	TotalShamings       int
	SeasonShames        int
	SeasonShamings      int
	Championships       int
	RotationSlot        int
	WeekAttributes      []string `json:"weekAttr"`
	GameAttributes      []string `json:"gameAttr"`
	SeasonAttributes    []string `json:"seasAttr"`
	PermanentAttributes []string `json:"permAttr"`
}

func (t *Team) String() string {
	return fmt.Sprintf("<T %s>", t.FullName)
}
