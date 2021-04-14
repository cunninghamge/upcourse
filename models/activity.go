package models

type Activity struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Metric      string `json:"metric"`
	Multiplier  int    `json:"multiplier"`
	Custom      bool   `json:"-"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
}

type ActivityTotals struct {
	Id         int    `json:"activityId"`
	Name       string `json:"activityName"`
	ModuleId   int    `json:"moduleId"`
	ModuleName string `json:"moduleName"`
	Minutes    int    `json:"minutes"`
}
