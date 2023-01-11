package models

type Website struct {
	ID   int    `json:"id"`
    URL  string `json:"url"`
	Host string `json:"host"`
	Category string `json:"category"`
	Protected bool `json:"protected"`
}