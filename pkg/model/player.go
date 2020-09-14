package model

import "encoding/json"

// PlayerList represents a list of Players
type PlayerList struct {
	List []*Player
}

func (pl *PlayerList) Unforbid() {
	for _, p := range pl.List {
		p.Unforbid()
	}
}

func (pl *PlayerList) MarshalJSON() ([]byte, error) {
	return json.Marshal(pl.List)
}

// Player represents a player of Blaseball
type Player struct {
	ID            string
	Name          string
	Deceased      bool `blase:"fk"`
	PeanutAllergy bool `blase:"fk"`

	// Baserunning attributes
	BaseThirst     float64 `blase:"fk"`
	Continuation   float64 `blase:"fk"`
	GroundFriction float64 `blase:"fk"`
	Indulgence     float64 `blase:"fk"`
	LaserLikeness  float64 `blase:"fk"`

	// Defense attributes
	Anticapitalism float64 `blase:"fk"`
	Chasiness      float64 `blase:"fk"`
	Omniscience    float64 `blase:"fk"`
	Tenaciousness  float64 `blase:"fk"`
	Watchfulness   float64 `blase:"fk"`

	// Hitting attributes
	Buoyancy      float64 `blase:"fk"`
	Divinity      float64 `blase:"fk"`
	Martyrdom     float64 `blase:"fk"`
	Moxie         float64 `blase:"fk"`
	Musclitude    float64 `blase:"fk"`
	Patheticism   float64 `blase:"fk"`
	Thwackability float64 `blase:"fk"`
	Tragicness    float64 `blase:"fk"`

	// Pitching attributes
	Coldness         float64 `blase:"fk"`
	Overpowerment    float64 `blase:"fk"`
	Ruthlessness     float64 `blase:"fk"`
	Shakespearianism float64 `blase:"fk"`
	Suppression      float64 `blase:"fk"`
	Unthwackability  float64 `blase:"fk"`
	TotalFingers     int     `blase:"fk"`

	// Other attributes
	Cinnamon       float64 `blase:"fk"`
	Pressurization float64 `blase:"fk"`
	Soul           int     `blase:"fk"`
	Bat            string
	Armor          string
	Ritual         string
	Coffee         int
	Blood          int
	Fate           int

	WeekAttributes      []string `json:"weekAttr" blase:"fk"`
	GameAttributes      []string `json:"gameAttr" blase:"fk"`
	SeasonAttributes    []string `json:"seasAttr" blase:"fk"`
	PermanentAttributes []string `json:"permAttr" blase:"fk"`

	showForbiddenKnowledge bool
}

func (p *Player) Incomplete() bool {
	return len(p.Name) == 0
}

func (p *Player) Unforbid() {
	p.showForbiddenKnowledge = true
}

func (p *Player) MarshalJSON() ([]byte, error) {
	// TODO: use the field tag (`blase:"fk"`) to reflect this out? or too clever?
	if p.showForbiddenKnowledge {
		return p.marshalAll()
	}
	return p.marshalUnforbidden()
}

func (p *Player) marshalUnforbidden() ([]byte, error) {
	unfk := map[string]interface{}{
		"ID":     p.ID,
		"Name":   p.Name,
		"Soul":   p.Soul,
		"Bat":    p.Bat,
		"Armor":  p.Armor,
		"Ritual": p.Ritual,
		"Coffee": p.Coffee,
		"Blood":  p.Blood,
		"Fate":   p.Fate,
	}
	return json.Marshal(unfk)
}

func (p *Player) marshalAll() ([]byte, error) {
	all := map[string]interface{}{
		"ID":                  p.ID,
		"Name":                p.Name,
		"Deceased":            p.Deceased,
		"PeanutAllergy":       p.PeanutAllergy,
		"BaseThirst":          p.BaseThirst,
		"Continuation":        p.Continuation,
		"GroundFriction":      p.GroundFriction,
		"Indulgence":          p.Indulgence,
		"LaserLikeness":       p.LaserLikeness,
		"Anticapitalism":      p.Anticapitalism,
		"Chasiness":           p.Chasiness,
		"Omniscience":         p.Omniscience,
		"Tenaciousness":       p.Tenaciousness,
		"Watchfulness":        p.Watchfulness,
		"Buoyancy":            p.Buoyancy,
		"Divinity":            p.Divinity,
		"Martyrdom":           p.Martyrdom,
		"Moxie":               p.Moxie,
		"Musclitude":          p.Musclitude,
		"Patheticism":         p.Patheticism,
		"Thwackability":       p.Thwackability,
		"Tragicness":          p.Tragicness,
		"Coldness":            p.Coldness,
		"Overpowerment":       p.Overpowerment,
		"Ruthlessness":        p.Ruthlessness,
		"Shakespearianism":    p.Shakespearianism,
		"Suppression":         p.Suppression,
		"Unthwackability":     p.Unthwackability,
		"TotalFingers":        p.TotalFingers,
		"Cinnamon":            p.Cinnamon,
		"Pressurization":      p.Pressurization,
		"Soul":                p.Soul,
		"Bat":                 p.Bat,
		"Armor":               p.Armor,
		"Ritual":              p.Ritual,
		"Coffee":              p.Coffee,
		"Blood":               p.Blood,
		"Fate":                p.Fate,
		"WeekAttributes":      p.WeekAttributes,
		"GameAttributes":      p.GameAttributes,
		"SeasonAttributes":    p.SeasonAttributes,
		"PermanentAttributes": p.PermanentAttributes,
	}
	return json.Marshal(all)
}
