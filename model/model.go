package model

type Movie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Language string `json:"language"`
	Genre    string `json:"genre"`
}