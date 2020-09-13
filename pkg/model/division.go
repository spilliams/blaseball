package model

// Division represents a division of teams. In season play, teams usually play
// other teams from their division.
type Division struct {
	ID   string
	Name string
}

func (d *Division) String() string {
	return d.Name
}
