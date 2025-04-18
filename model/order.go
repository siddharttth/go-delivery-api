package model

type OrderRequest struct {
	Products map[string]int `json:"products"`
}
