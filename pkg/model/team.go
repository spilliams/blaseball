package model

import (
	"encoding/json"
	"fmt"
)

// TeamList represents a list of Teams
type TeamList struct {
	List []*Team
}

func (tl *TeamList) Unforbid() {
	for _, t := range tl.List {
		t.Unforbid()
	}
}

func (tl *TeamList) MarshalJSON() ([]byte, error) {
	return json.Marshal(tl.List)
}

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
	Bullpen  []string `blase:"fk"`
	Bench    []string `blase:"fk"`

	ShameRuns      int `blase:"fk"`
	TotalShames    int `blase:"fk"`
	TotalShamings  int `blase:"fk"`
	SeasonShames   int `blase:"fk"`
	SeasonShamings int `blase:"fk"`
	Championships  int `blase:"fk"`
	RotationSlot   int `blase:"fk"`

	WeekAttributes      []string `json:"weekAttr" blase:"fk"`
	GameAttributes      []string `json:"gameAttr" blase:"fk"`
	SeasonAttributes    []string `json:"seasAttr" blase:"fk"`
	PermanentAttributes []string `json:"permAttr" blase:"fk"`

	showForbiddenKnowledge bool
}

func (t *Team) String() string {
	return fmt.Sprintf("<T %s>", t.FullName)
}

func (t *Team) Incomplete() bool {
	return len(t.FullName) == 0
}

func (t *Team) Unforbid() {
	t.showForbiddenKnowledge = true
}

func (t *Team) MarshalJSON() ([]byte, error) {
	if t.showForbiddenKnowledge {
		return t.marshalAll()
	}
	return t.marshalUnforbidden()
}

func (t *Team) marshalUnforbidden() ([]byte, error) {
	unfk := map[string]interface{}{
		"ID":             t.ID,
		"FullName":       t.FullName,
		"Nickname":       t.Nickname,
		"Shorthand":      t.Shorthand,
		"Emoji":          t.Emoji,
		"Slogan":         t.Slogan,
		"Location":       t.Location,
		"MainColor":      t.MainColor,
		"SecondaryColor": t.SecondaryColor,
		"Lineup":         t.Lineup,
		"Rotation":       t.Rotation,
	}
	return json.Marshal(unfk)
}

func (t *Team) marshalAll() ([]byte, error) {
	all := map[string]interface{}{
		"ID":                  t.ID,
		"FullName":            t.FullName,
		"Nickname":            t.Nickname,
		"Shorthand":           t.Shorthand,
		"Emoji":               t.Emoji,
		"Slogan":              t.Slogan,
		"Location":            t.Location,
		"MainColor":           t.MainColor,
		"SecondaryColor":      t.SecondaryColor,
		"Lineup":              t.Lineup,
		"Rotation":            t.Rotation,
		"Bullpen":             t.Bullpen,
		"Bench":               t.Bench,
		"ShameRuns":           t.ShameRuns,
		"TotalShames":         t.TotalShames,
		"TotalShamings":       t.TotalShamings,
		"SeasonShames":        t.SeasonShames,
		"SeasonShamings":      t.SeasonShamings,
		"Championships":       t.Championships,
		"RotationSlot":        t.RotationSlot,
		"WeekAttributes":      t.WeekAttributes,
		"GameAttributes":      t.GameAttributes,
		"SeasonAttributes":    t.SeasonAttributes,
		"PermanentAttributes": t.PermanentAttributes,
	}
	return json.Marshal(all)
}
