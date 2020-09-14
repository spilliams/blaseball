package model

import "fmt"

// Team represents a single Blaseball team. Teams contain players, and play
// games against other teams.
type Team struct {
	ID             string
	FullName       string
	Nickname       string
	Shorthand      string
	Emoji          string
	Slogan         string
	Location       string
	MainColor      string
	SecondaryColor string

	Lineup   []string
	Rotation []string
	Bullpen  []string // fk
	Bench    []string // fk

	ShameRuns      int // fk
	TotalShames    int // fk
	TotalShamings  int // fk
	SeasonShames   int // fk
	SeasonShamings int // fk
	Championships  int // fk
	RotationSlot   int // fk

	WeekAttributes      []string `json:"weekAttr"` // fk
	GameAttributes      []string `json:"gameAttr"` // fk
	SeasonAttributes    []string `json:"seasAttr"` // fk
	PermanentAttributes []string `json:"permAttr"` // fk
}

func (t *Team) String() string {
	return fmt.Sprintf("<T %s>", t.FullName)
}

func (t *Team) Incomplete() bool {
	return len(t.FullName) == 0
}
