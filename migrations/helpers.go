package main

import "gorm.io/gorm"

func set_constraints(db *gorm.DB) {
	db.Exec(`
	BEGIN;
	ALTER TABLE module_activities
	DROP CONSTRAINT fk_modules_module_activities;

	ALTER TABLE module_activities
	ADD CONSTRAINT fk_modules_module_activities
		FOREIGN KEY (module_id)
		REFERENCES modules(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE;

	COMMIT;

	BEGIN;
	ALTER TABLE modules
	DROP CONSTRAINT fk_courses_modules;

	ALTER TABLE modules
	ADD CONSTRAINT fk_courses_modules
		FOREIGN KEY (course_id)
		REFERENCES courses(id)
		ON DELETE CASCADE;

	COMMIT;
	`)
}

func set_triggers(db *gorm.DB) {
	db.Exec(`CREATE OR REPLACE FUNCTION update_updated_at_column()
	RETURNS TRIGGER AS $$
	BEGIN
		NEW.updated_at = now();
		RETURN NEW;
	END;
	$$ language 'plpgsql';
	
	CREATE TRIGGER update_course_updated_at BEFORE UPDATE ON courses FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
	CREATE TRIGGER update_module_updated_at BEFORE UPDATE ON modules FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
	CREATE TRIGGER update_activity_updated_at BEFORE UPDATE ON activities FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
	CREATE TRIGGER update_module_activity_updated_at BEFORE UPDATE ON module_activities FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
	`)
}
