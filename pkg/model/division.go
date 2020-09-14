package model

import (
	"encoding/json"
	"fmt"
)

// DivisionList represents a list of Divisions
type DivisionList struct {
	List []*Division
}

func (dl *DivisionList) Unforbid() {
	for _, d := range dl.List {
		d.Unforbid()
	}
}

func (dl *DivisionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(dl.List)
}

// Division represents a division of teams. In season play, teams usually play
// other teams from their division.
type Division struct {
	ID      string
	Name    string
	TeamIDs []string `json:"teams"`
}

func (d *Division) String() string {
	return fmt.Sprintf("<D %s (%d teams)>", d.Name, len(d.TeamIDs))
}

func (d *Division) Unforbid() {
	return
}
