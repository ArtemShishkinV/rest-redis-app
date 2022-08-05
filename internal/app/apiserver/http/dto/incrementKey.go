package dto

type IncrementKeyRequestDto struct {
	Key string `json:"key"`
	Val int    `json:"val"`
}

type IncrementKeyResponseDto struct {
	Val int `json:"val"`
}
