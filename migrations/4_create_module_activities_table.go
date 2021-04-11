package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table module activities...")
		_, err := db.Exec(`CREATE TABLE module_activities(
id SERIAL, 
input INT,
notes VARCHAR(255),
module_id INT NOT NULL,
activity_id INT NOT NULL,
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
PRIMARY KEY(id),
CONSTRAINT fk_module
	FOREIGN KEY(module_id)
		REFERENCES modules(id),
CONSTRAINT fk_activity
	FOREIGN KEY(activity_id)
		REFERENCES activities(id)
)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table module activities...")
		_, err := db.Exec(`DROP TABLE module_activities`)
		return err
	})
}
