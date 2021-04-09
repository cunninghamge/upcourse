package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table modules...")
		_, err := db.Exec(`CREATE TABLE modules(
id SERIAL,
name VARCHAR(255),
number INT,
course_id INT NOT NULL,
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
PRIMARY KEY(id),
CONSTRAINT fk_course
	FOREIGN KEY(course_id)
		REFERENCES courses(id)
)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table modules...")
		_, err := db.Exec(`DROP TABLE modules`)
		return err
	})
}
