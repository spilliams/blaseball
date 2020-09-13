package model

import "fmt"

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
