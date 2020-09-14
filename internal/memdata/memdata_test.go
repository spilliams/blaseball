package memdata

import (
	"testing"

	"github.com/spilliams/blaseball/pkg/model"
)

func TestIdempotency(t *testing.T) {
	ds := NewStore()
	teamName := "abc def"
	team := &model.Team{ID: "12345", FullName: teamName}
	err := ds.PutTeam(team)
	if err != nil {
		t.Fatal(err)
	}

	got, err := ds.GetTeamByFullName(teamName)
	if err != nil {
		t.Fatal(err)
	}
	got.FullName = "def abc"
	// expect one to still exist in storage
	if _, err = ds.GetTeamByFullName(teamName); err != nil {
		t.Fatal(err)
	}
}
