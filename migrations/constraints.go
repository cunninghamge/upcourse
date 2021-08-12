package main

import "upcourse/models"

func (m GormMigration) setConstraints() error {
	if !m.db.Migrator().HasIndex(&models.ModuleActivity{}, "index_module_activities_on_activities_modules") {
		if err := m.createIndexes(); err != nil {
			return err
		}
		if err := m.setOnDelete(); err != nil {
			return err
		}
		if err := m.setTriggers(); err != nil {
			return err
		}
	}
	return nil
}

func (m GormMigration) createIndexes() error {
	return m.db.Exec(`CREATE UNIQUE INDEX index_module_activities_on_activities_modules ON module_activities (module_id, activity_id);`).Error
}

func (m GormMigration) setOnDelete() error {
	return m.db.Exec(`
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

func (m GormMigration) setTriggers() error {
	return m.db.Exec(`CREATE OR REPLACE FUNCTION update_updated_at_column()
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
	`).Error
}
