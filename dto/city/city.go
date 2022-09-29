package city

type RequestCity struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}
type ResponseCity struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type RequestUpdateCity struct {
	Name string `json:"name" `
}
