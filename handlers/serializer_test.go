package handlers

import (
	"testing"
	"upcourse/internal/mocks"
)

func TestSerializeCourse(t *testing.T) {
	t.Run("serializes a course with modules and moduleActivities", func(t *testing.T) {
		course := mocks.FullCourse()
		defer teardown()

		serializedCourse := SerializeCourse(*course)

		assertResponseValue(t, serializedCourse.Type, "course", "course type")
		assertResponseValue(t, serializedCourse.ID, course.ID, "Id")

		modules, ok := serializedCourse.Relationships.(map[string][]SerializedResource)
		if !ok {
			t.Errorf("error converting course relationships to map")
		}
		firstModule := modules["modules"][0]
		firstMockModule := course.Modules[0]
		assertResponseValue(t, firstModule.Type, "module", "module type")
		assertResponseValue(t, firstModule.ID, firstMockModule.ID, "first module ID")

		moduleActivities, ok := firstModule.Relationships.(map[string][]SerializedResource)
		if !ok {
			t.Errorf("error converting module relationships to map")
		}

		firstModuleActivity := moduleActivities["moduleActivities"][0]
		firstMockModActivity := firstMockModule.ModuleActivities[0]
		assertResponseValue(t, firstModuleActivity.Type, "moduleActivity", "moduleActivity type")
		assertResponseValue(t, firstModuleActivity.ID, firstMockModActivity.ID, "first module ID")

		activity, ok := firstModuleActivity.Relationships.(map[string]SerializedResource)
		if !ok {
			t.Errorf("error converting moduleActivity relationships to map")
		}
		firstActivity := activity["activity"]
		firstMockActivity := firstMockModActivity.Activity
		assertResponseValue(t, firstActivity.Type, "activity", "activity type")
		assertResponseValue(t, firstActivity.ID, firstMockActivity.ID, "first activity ID")
	})

	t.Run("it serializes a course without module activities", func(t *testing.T) {
		course := mocks.CourseList()[0]
		defer teardown()

		serializedCourse := SerializeCourse(course)

		assertResponseValue(t, serializedCourse.Type, "course", "course type")
		assertResponseValue(t, serializedCourse.ID, course.ID, "Id")

		modules, ok := serializedCourse.Relationships.(map[string][]SerializedResource)
		if !ok {
			t.Errorf("error converting course relationships to map")
		}
		firstModule := modules["modules"][0]
		firstMockModule := course.Modules[0]
		assertResponseValue(t, firstModule.Type, "module", "module type")
		assertResponseValue(t, firstModule.ID, firstMockModule.ID, "first module ID")

		if firstModule.Relationships != nil {
			t.Errorf("got %v, expected nil for module relationships", firstModule.Relationships)
		}
	})
}

func TestSerializeModule(t *testing.T) {
	t.Run("serializes a module with moduleActivities", func(t *testing.T) {
		module := mocks.Module()
		defer teardown()

		serializedModule := SerializeModule(module)

		assertResponseValue(t, serializedModule.Type, "module", "module type")
		assertResponseValue(t, serializedModule.ID, module.ID, "Id")

		moduleActivities, ok := serializedModule.Relationships.(map[string][]SerializedResource)
		if !ok {
			t.Errorf("error converting module relationships to map")
		}
		firstModuleActivity := moduleActivities["moduleActivities"][0]
		firstMockModActivity := module.ModuleActivities[0]
		assertResponseValue(t, firstModuleActivity.Type, "moduleActivity", "moduleActivity type")
		assertResponseValue(t, firstModuleActivity.ID, firstMockModActivity.ID, "first module ID")

		activity, ok := firstModuleActivity.Relationships.(map[string]SerializedResource)
		if !ok {
			t.Errorf("error converting moduleActivity relationships to map")
		}
		firstActivity := activity["activity"]
		firstMockActivity := firstMockModActivity.Activity
		assertResponseValue(t, firstActivity.Type, "activity", "activity type")
		assertResponseValue(t, firstActivity.ID, firstMockActivity.ID, "first activity ID")
	})

	t.Run("serializes a module without moduleActivities", func(t *testing.T) {
		module := mocks.SimpleModule()
		defer teardown()

		serializedModule := SerializeModule(module)

		assertResponseValue(t, serializedModule.Type, "module", "module type")
		assertResponseValue(t, serializedModule.ID, module.ID, "Id")

		if serializedModule.Relationships != nil {
			t.Errorf("got %v, expected nil for module relationships", serializedModule.Relationships)
		}
	})
}

func TestSerializeModuleActivity(t *testing.T) {
	moduleActivity := mocks.Module().ModuleActivities[0]

	serializedModuleActivity := SerializeModuleActivity(moduleActivity)

	assertResponseValue(t, serializedModuleActivity.Type, "moduleActivity", "moduleActivity type")
	assertResponseValue(t, serializedModuleActivity.ID, moduleActivity.ID, "first module ID")

	activity, ok := serializedModuleActivity.Relationships.(map[string]SerializedResource)
	if !ok {
		t.Errorf("error converting moduleActivity relationships to map")
	}
	firstActivity := activity["activity"]
	firstMockActivity := moduleActivity.Activity
	assertResponseValue(t, firstActivity.Type, "activity", "activity type")
	assertResponseValue(t, firstActivity.ID, firstMockActivity.ID, "first activity ID")
}

func TestSerializeActivity(t *testing.T) {
	activity := mocks.DefaultActivities()[0]

	serializedActivity := SerializeActivity(activity)

	assertResponseValue(t, serializedActivity.Type, "activity", "activity type")
	assertResponseValue(t, serializedActivity.ID, activity.ID, "first activity ID")
}
