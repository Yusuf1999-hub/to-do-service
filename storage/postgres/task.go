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
	var id int64
	err := r.db.QueryRow(`
		INSERT INTO tasks(assignee, title, summary, deadline, status)
		VALUES($1, $2, $3, $4, $5) returning id`,
		task.Assignee,
		task.Title,
		task.Summary,
		task.Deadline,
		task.Status,
	).Scan(&id)
	if err != nil {
		return pb.Task{}, err
	}

	task, err = r.Get(id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Get(id int64) (pb.Task, error) {
	var task pb.Task
	err := r.db.Get(&task, `
		SELECT id, assignee, title, summary, deadline, status FROM tasks
		WHERE id=$1`, id)
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
				FROM tasks LIMIT $1 OFFSET $2`, limit, offset)
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
	result, err := r.db.Exec(`UPDATE tasks SET assignee=$1, title=$2, summary=$3, deadline=$4, status=$5 
						WHERE id=$6`,
		&task.Assignee,
		&task.Title,
		&task.Summary,
		&task.Deadline,
		&task.Status,
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

func (r *taskRepo) Delete(id int64) error {
	result, err := r.db.Exec(`DELETE FROM tasks WHERE id=$1`, id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *taskRepo) ListOverdue(time time.Time) ([]*pb.Task, error) {
	var tasks []*pb.Task

	err := r.db.Select(&tasks, `
				SELECT id, assignee, title, summary, deadline, status 
				FROM tasks WHERE deadline >= $1`, time)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
