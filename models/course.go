package models

type Course struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Institution string `json:"institution"`
	CreditHours int    `json:"creditHours"`
	Length      int    `json:"length"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
