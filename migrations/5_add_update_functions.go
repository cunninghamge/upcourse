package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("setting update functions...")
		_, err := db.Exec(`CREATE OR REPLACE FUNCTION update_updated_at_column()
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
		return err
	}, func(db migrations.DB) error {
		fmt.Println("removing update functions...")
		_, err := db.Exec(`DROP TRIGGER update_updated_at_column ON courses;
DROP TRIGGER update_updated_at_column ON modules;
DROP TRIGGER update_updated_at_column ON activities;
DROP TRIGGER update_updated_at_column ON module_activities;
`)
		return err
	})
}
