package models

type Property struct {
	Id       string `json:"id"`
	Area     string `json:"area"`
	Owner    string `json:"owner"`
	Location string `json:"location"`
}
