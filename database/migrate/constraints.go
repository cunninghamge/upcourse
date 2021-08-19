package main

import (
	db "upcourse/database"
	"upcourse/models"
)

func setConstraints() error {
	if !db.Conn.Migrator().HasIndex(&models.ModuleActivity{}, "index_module_activities_on_activities_modules") {
		if err := createIndexes(); err != nil {
			return err
		}
		if err := setOnDelete(); err != nil {
			return err
		}
		if err := setTriggers(); err != nil {
			return err
		}
	}
	return nil
}

// TODO add unique index for module number within a course
// TODO allow users to change module numbers as a way of changing order?
func createIndexes() error {
	return db.Conn.Exec(`CREATE UNIQUE INDEX index_module_activities_on_activities_modules ON module_activities (module_id, activity_id);`).Error
}

func setOnDelete() error {
	return db.Conn.Exec(`
		BEGIN;
		ALTER TABLE module_activities
		DROP CONSTRAINT fk_modules_module_activities;
		ALTER TABLE module_activities
		ADD CONSTRAINT fk_modules_module_activities
			FOREIGN KEY (module_id)
			REFERENCES modules(id)
			ON DELETE CASCADE
			ON UPDATE CASCADE;
		ALTER TABLE module_activities
		ADD CONSTRAINT fk_activities_module_activities
			FOREIGN KEY (activity_id)
			REFERENCES activities(id)
			ON DELETE CASCADE;
		ALTER TABLE modules
		DROP CONSTRAINT fk_courses_modules;
		ALTER TABLE modules
		ADD CONSTRAINT fk_courses_modules
			FOREIGN KEY (course_id)
			REFERENCES courses(id)
			ON DELETE CASCADE;
		COMMIT;
	`).Error
}

func setTriggers() error {
	return db.Conn.Exec(`CREATE OR REPLACE FUNCTION update_updated_at_column()
	RETURNS TRIGGER AS $$
	BEGIN
		NEW.updated_at = now();
		RETURN NEW;
	END;
	$$ language 'plpgsql';
	
	DROP TRIGGER IF EXISTS updated_course_updated_at ON courses;
	DROP TRIGGER IF EXISTS updated_activity_updated_at ON activities;
	DROP TRIGGER IF EXISTS updated_module_activity_updated_at ON module_activities;
	DROP TRIGGER IF EXISTS updated_module_updated_at ON modules;
	CREATE TRIGGER update_course_updated_at BEFORE UPDATE ON courses FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
	CREATE TRIGGER update_module_updated_at BEFORE UPDATE ON modules FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
	CREATE TRIGGER update_activity_updated_at BEFORE UPDATE ON activities FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
	CREATE TRIGGER update_module_activity_updated_at BEFORE UPDATE ON module_activities FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
	`).Error
}
