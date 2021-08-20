package models

import (
	"reflect"
	"testing"
	db "upcourse/database"
)

func TestGetActivities(t *testing.T) {
	defaultActivities := defaultActivities()

	t.Run("returns a list of the default activities", func(t *testing.T) {
		activities, err := GetActivities()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		for i := 0; i < len(activities); i++ {
			got := activities[i]
			want := defaultActivities[i]
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v for response activities[%d]", got, want, i)
			}
		}
	})

	t.Run("does not include custom activities", func(t *testing.T) {
		db.Conn.Create(&Activity{Custom: true})
		defer teardown()

		activities, err := GetActivities()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(activities) != len(defaultActivities) {
			t.Errorf("got %d want %d for number of results", len(activities), len(defaultActivities))
		}
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		forceError()
		defer clearError()

		activities, err := GetActivities()

		if activities != nil {
			t.Errorf("unexpected response: %v", activities)
		}
		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

// fetches the list of core activities from database
func defaultActivities() []*Activity {
	var activities []*Activity
	db.Conn.Select("id, name, description, metric, multiplier").Where("custom=false").Find(&activities)

	return activities
}
