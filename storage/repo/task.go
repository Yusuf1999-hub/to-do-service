package repo

import (
	pb "github.com/Yusuf1999-hub/to-do-service/genproto"
	"time"
)

// TaskStorageI ...
type TaskStorageI interface {
	Create(pb.Task) (pb.Task, error)
	Get(id int64) (pb.Task, error)
	List(page, limit int64) ([]*pb.Task, int64, error)
	Update(pb.Task) (pb.Task, error)
	Delete(id int64) error
	ListOverdue(time time.Time) ([]*pb.Task, error)
}
