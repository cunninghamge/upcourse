package models

type Module struct {
	Id               int              `json:"id"`
	Name             string           `json:"name"`
	Number           int              `json:"number"`
	CourseId         int              `json:"courseId"`
	CreatedAt        string           `json:"-"`
	UpdatedAt        string           `json:"-"`
	Course           Course           `json:"-"`
	ModuleActivities []ModuleActivity `json:"moduleActivities"`
}
