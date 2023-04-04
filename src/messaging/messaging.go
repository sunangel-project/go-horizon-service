package messaging

const IN_Q = "spots"
const GROUP = "horizon-service"

const OUT_Q = "horizons"
const ERR_Q = "error"

type PartSubMessage struct {
	Id uint `json:"id"`
	Of uint `json:"of"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type SpotSubMessage struct {
	Dir  float64  `json:"dir"`
	Kind string   `json:"kind"`
	Loc  Location `json:"loc"`
}

type SpotMessage struct {
	Part PartSubMessage `json:"part"`
	Spot SpotSubMessage `json:"spot"`
}
