package bio

import "database/sql"

type CreateExperience struct {
	ImageURL    string `form:"image"`
	JobTitle    string `form:"job_title" validate:"required"`
	JobLocation string `form:"job_location"`
	SkillUsed   string `form:"skill_used"`
	StartDate   string `form:"start_date" validate:"required"`
	EndDate     string `form:"end_date"`
	Description string `form:"description"`
}

type ExperienceDB struct {
	ID          sql.NullString `db:"id"`
	UserID      sql.NullString `db:"user_id"`
	ImageURL    sql.NullString `db:"image_url"`
	JobTitle    sql.NullString `db:"job_title"`
	JobLocation sql.NullString `db:"job_location"`
	SkillUsed   sql.NullString `db:"skill_used"`
	StartDate   sql.NullString `db:"start_date"`
	EndDate     sql.NullString `db:"end_date"`
	Description sql.NullString `db:"description"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
}

type UpdateExperience struct {
	ImageURL    string `form:"image"`
	JobTitle    string `form:"job_title"`
	JobLocation string `form:"job_location"`
	SkillUsed   string `form:"skill_used"`
	StartDate   string `form:"start_date"`
	EndDate     string `form:"end_date"`
	Description string `form:"description"`
}
