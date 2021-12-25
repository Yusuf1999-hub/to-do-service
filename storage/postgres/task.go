package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/Yusuf1999-hub/to-do-service/genproto"
)

type taskRepo struct {
	db *sqlx.DB
}

// NewTaskRepo ...
func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task pb.Task) (pb.Task, error) {
	err := r.db.QueryRow(`
		INSERT INTO tasks(id, assignee, title, summary, deadline, status, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8) returning id`,
		task.Id,
		task.Assignee,
		task.Title,
		task.Summary,
		task.Deadline,
		task.Status,
		time.Now().UTC(),
		time.Now().UTC(),
	).Scan(&task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	task, err = r.Get(task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Get(id string) (pb.Task, error) {
	var task pb.Task
	err := r.db.QueryRow(`
		SELECT id, assignee, title, summary, deadline, status, created_at, updated_at FROM tasks
		WHERE id=$1 and deleted_at is NULL`, id).Scan(
		&task.Id,
		&task.Assignee,
		&task.Title,
		&task.Summary,
		&task.Deadline,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) List(page, limit int64) ([]*pb.Task, int64, error) {
	var (
		tasks []*pb.Task
		count int64
	)

	offset := (page - 1) * limit
	err := r.db.Select(&tasks, `
				SELECT id, assignee, title, summary, deadline, status 
				FROM tasks WHERE deleted_at IS NULL ORDER BY id ASC LIMIT $1 OFFSET $2 `, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	err = r.db.QueryRow(`SELECT count(*) FROM tasks`).Scan(&count)

	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}

func (r *taskRepo) Update(task pb.Task) (pb.Task, error) {
	result, err := r.db.Exec(`UPDATE tasks SET assignee=$1, title=$2, summary=$3, deadline=$4, status=$5, updated_at = $6 
						WHERE id=$7`,
		&task.Assignee,
		&task.Title,
		&task.Summary,
		&task.Deadline,
		&task.Status,
		time.Now().UTC(),
		&task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Task{}, sql.ErrNoRows
	}

	task, err = r.Get(task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, err
}

func (r *taskRepo) Delete(id string) error {
	currentTime := time.Now().UTC()

	result, err := r.db.Exec(`UPDATE tasks SET deleted_at = $1  WHERE id=$2`, currentTime, id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *taskRepo) ListOverdue(page, limit int64, timer time.Time) ([]*pb.Task, int64, error) {
	var (
		count int64
		tasks []*pb.Task
	)
	offset := (page - 1) * limit

	rows, err := r.db.Queryx(`
				SELECT id, assignee, title, summary, deadline, status, created_at, updated_at
				FROM tasks WHERE deadline < $1 AND deleted_at IS NULL ORDER BY id ASC LIMIT $2 OFFSET $3`, timer, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var task pb.Task

		err = rows.Scan(
			&task.Id,
			&task.Assignee,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM tasks WHERE deadline < $1 AND deleted_at IS NULL`, timer).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}
