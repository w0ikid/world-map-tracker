package models
type Country struct {
	ISO  string `json:"alpha-2"`
	Name string `json:"name"`
}