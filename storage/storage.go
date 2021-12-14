package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/Yusuf1999-hub/to-do-service/storage/postgres"
	"github.com/Yusuf1999-hub/to-do-service/storage/repo"
)

// IStorage ...
type IStorage interface {
	Task() repo.TaskStorageI
}

type storagePg struct {
	db       *sqlx.DB
	taskRepo repo.TaskStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		taskRepo: postgres.NewTaskRepo(db),
	}
}

func (s storagePg) Task() repo.TaskStorageI {
	return s.taskRepo
}
