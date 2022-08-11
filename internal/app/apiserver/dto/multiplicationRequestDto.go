package dto

type MultiplicationRequestDto struct {
	Multipliers []MultiplierRequestDto `json:"array"`
}

type MultiplierRequestDto struct {
	A   string `json:"a"`
	B   string `json:"b"`
	Key string `json:"key"`
}
