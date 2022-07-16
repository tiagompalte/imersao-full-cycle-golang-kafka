package repository

import (
	"database/sql"
	"imersao-full-cycle/entity"
)

type CourseMySqlRepository struct {
	Db *sql.DB
}

func (c CourseMySqlRepository) Insert(course entity.Course) error {
	stmt, err := c.Db.Prepare(`insert into courses(id, name, description, status) values (?,?,?,?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status)
	if err != nil {
		return err
	}
	
	return nil
}
