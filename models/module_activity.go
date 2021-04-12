package models

type ModuleActivity struct {
	Id         int      `json:"id"`
	Input      int      `json:"input"`
	Notes      string   `json:"notes"`
	ModuleId   int      `json:"-"`
	ActivityId int      `json:"-"`
	CreatedAt  string   `json:"-"`
	UpdatedAt  string   `json:"-"`
	Activity   Activity `json:"activity"`
}
