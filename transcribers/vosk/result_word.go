package vosk

type resultWord struct {
	Confidence float64 `json:"conf"`
	EndTime    float64 `json:"end"`
	StartTime  float64 `json:"start"`
	Word       string  `json:"word"`
}
