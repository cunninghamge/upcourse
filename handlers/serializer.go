package handlers

import "upcourse/models"

type SerializedResource struct {
	Type          string      `json:"type"`
	ID            int         `json:"id"`
	Attributes    interface{} `json:"attributes"`
	Relationships interface{} `json:"relationships,omitempty"`
}

func SerializeCourse(c models.Course) SerializedResource {
	sr := SerializedResource{
		Type:       "course",
		ID:         c.ID,
		Attributes: c,
		Relationships: map[string][]SerializedResource{
			"modules": func(modules []models.Module) []SerializedResource {
				var r []SerializedResource
				for _, m := range modules {
					r = append(r, SerializeModule(m))
				}
				return r
			}(c.Modules),
		},
	}

	return sr
}

func SerializeModule(m models.Module) SerializedResource {
	sr := SerializedResource{
		Type: "module",
		ID:   m.ID,
		Attributes: map[string]interface{}{
			"name":   m.Name,
			"number": m.Number,
		},
	}

	if len(m.ModuleActivities) > 0 {
		sr.Relationships = func() map[string][]SerializedResource {
			return map[string][]SerializedResource{
				"moduleActivities": func(modActivities []models.ModuleActivity) []SerializedResource {
					var r []SerializedResource
					for _, ma := range modActivities {
						r = append(r, SerializeModuleActivity(ma))
					}
					return r
				}(m.ModuleActivities),
			}
		}()
	}

	return sr
}

func SerializeModuleActivity(ma models.ModuleActivity) SerializedResource {
	return SerializedResource{
		Type:       "moduleActivity",
		ID:         ma.ID,
		Attributes: ma,
		Relationships: map[string]SerializedResource{
			"activity": SerializeActivity(ma.Activity),
		},
	}
}

func SerializeActivity(a models.Activity) SerializedResource {
	return SerializedResource{
		Type:       "activity",
		ID:         a.ID,
		Attributes: a,
	}
}
