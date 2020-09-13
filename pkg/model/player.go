package model

type Player struct {
	ID            string
	Name          string
	Deceased      bool
	PeanutAllergy bool

	// Baserunning attributes
	BaseThirst     float64
	Continuation   float64
	GroundFriction float64
	Indulgence     float64
	LaserLikeness  float64

	// Defense attributes
	Anticapitalism float64
	Chasiness      float64
	Omniscience    float64
	Tenaciousness  float64
	Watchfulness   float64

	// Hitting attributes
	Buoyancy      float64
	Divinity      float64
	Martyrdom     float64
	Moxie         float64
	Musclitude    float64
	Patheticism   float64
	Thwackability float64
	Tragicness    float64

	// Pitching attributes
	Coldness         float64
	Overpowerment    float64
	Ruthlessness     float64
	Shakespearianism float64
	Suppression      float64
	Unthwackability  float64
	TotalFingers     int

	// Other attributes
	Cinnamon       float64
	Pressurization float64
	Soul           int
	Bat            string
	Armor          string
	Ritual         string
	Coffee         int
	Blood          int
	Fate           int

	WeekAttributes      []string `json:"weekAttr"`
	GameAttributes      []string `json:"gameAttr"`
	SeasonAttributes    []string `json:"seasAttr"`
	PermanentAttributes []string `json:"permAttr"`
}
