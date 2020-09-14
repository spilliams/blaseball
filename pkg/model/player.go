package model

// Player represents a player of Blaseball
type Player struct {
	ID            string
	Name          string
	Deceased      bool // fk only if necromancy?
	PeanutAllergy bool // fk

	// Baserunning attributes
	BaseThirst     float64 // fk
	Continuation   float64 // fk
	GroundFriction float64 // fk
	Indulgence     float64 // fk
	LaserLikeness  float64 // fk

	// Defense attributes
	Anticapitalism float64 // fk
	Chasiness      float64 // fk
	Omniscience    float64 // fk
	Tenaciousness  float64 // fk
	Watchfulness   float64 // fk

	// Hitting attributes
	Buoyancy      float64 // fk
	Divinity      float64 // fk
	Martyrdom     float64 // fk
	Moxie         float64 // fk
	Musclitude    float64 // fk
	Patheticism   float64 // fk
	Thwackability float64 // fk
	Tragicness    float64 // fk

	// Pitching attributes
	Coldness         float64 // fk
	Overpowerment    float64 // fk
	Ruthlessness     float64 // fk
	Shakespearianism float64 // fk
	Suppression      float64 // fk
	Unthwackability  float64 // fk
	TotalFingers     int     // fk

	// Other attributes
	Cinnamon       float64 // fk
	Pressurization float64 // fk
	Soul           int     // soulscream? or fk
	Bat            string
	Armor          string
	Ritual         string
	Coffee         int
	Blood          int
	Fate           int

	WeekAttributes      []string `json:"weekAttr"` // fk
	GameAttributes      []string `json:"gameAttr"` // fk
	SeasonAttributes    []string `json:"seasAttr"` // fk
	PermanentAttributes []string `json:"permAttr"` // fk
}

func (p *Player) Incomplete() bool {
	return len(p.Name) == 0
}
