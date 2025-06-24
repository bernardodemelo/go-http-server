package dtos

type RollerCoaster struct {
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Height   float64 `json:"height"`
	Speed    float64 `json:"speed"`
}
