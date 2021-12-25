package postgres

import (
	"log"
	"reflect"
	"testing"
	"time"

	pb "github.com/Yusuf1999-hub/to-do-service/genproto"
)

func TestTaskRepo_Create(t *testing.T) {
	tests := []struct {
		name  string
		input pb.Task
		want  pb.Task
	}{
		{
			name: "omadli",
			input: pb.Task{
				Id:       "123e4567-e89b-12d3-a456-426614174000",
				Assignee: "salom",
				Title:    "salom of Tashkent",
				Summary:  "The book is very good",
				Deadline: "2020-11-13",
				Status:   "active",
			},
			want: pb.Task{
				Id:        "123e4567-e89b-12d3-a456-426614174000",
				Assignee:  "salom",
				Title:     "salom of Tashkent",
				Summary:   "The book is very good",
				Deadline:  "2020-11-13T00:00:00Z",
				Status:    "active",
				CreatedAt: "",
				UpdatedAt: "",
			},
		},
		{
			name: "additional",
			input: pb.Task{
				Id:       "00000000-0000-0000-0000-000000000002",
				Assignee: "Russia",
				Title:    "History of Russia",
				Summary:  "The book is good",
				Deadline: "2020-11-13",
				Status:   "active",
			},
			want: pb.Task{
				Id:        "00000000-0000-0000-0000-000000000002",
				Assignee:  "Russia",
				Title:     "History of Russia",
				Summary:   "The book is good",
				Deadline:  "2020-11-13T00:00:00Z",
				Status:    "active",
				CreatedAt: "",
				UpdatedAt: "",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Create(tc.input)
			if err != nil {
				log.Fatalf("%s: got: %v", tc.name, err)
			}
			tc.want.CreatedAt = got.CreatedAt
			tc.want.UpdatedAt = got.UpdatedAt

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

/*func TestTaskRepo_Get(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want pb.Task
	}{
		{
			name: "successful",
			id:   "00000000-0000-0000-0000-000000000004",
			want: pb.Task{
				Assignee: "Russia",
				Title:    "History of Russia",
				Summary:  "The book is good",
				Deadline: "2020-11-13T00:00:00Z",
				Status:   "active",
			},
		},
		{
			name: "additional",
			id:   "00000000-0000-0000-0000-000000000003",
			want: pb.Task{
				Assignee: "salom",
				Title:    "salom of Tashkent",
				Summary:  "The book is very good",
				Deadline: "2020-11-13T00:00:00Z",
				Status:   "active",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Get(tc.id)
			if err != nil {
				log.Fatalf("%s: got:%v", tc.name, err)
			}

			got.Id = 0

			if !reflect.DeepEqual(got, tc.want) {
				log.Fatalf("%s: expected:%v got:%v", tc.name, tc.want, got)
			}
		})
	}
}*/

type result struct {
	list  []pb.Task
	count int64
}

func TestTaskRepo_List(t *testing.T) {
	tests := []struct {
		name  string
		page  int64
		limit int64
		want  result
	}{
		{
			name:  "successful",
			page:  1,
			limit: 2,
			want: result{
				list: []pb.Task{
					{
						Id:       "00000000-0000-0000-0000-000000000001",
						Assignee: "salom",
						Title:    "salom of Tashkent",
						Summary:  "The book is very good",
						Deadline: "2020-11-13T00:00:00Z",
						Status:   "active",
					},
					{
						Id:       "00000000-0000-0000-0000-000000000001",
						Assignee: "salom",
						Title:    "salom of Tashkent",
						Summary:  "The book is very good",
						Deadline: "2020-11-13T00:00:00Z",
						Status:   "active",
					},
				},
				count: 4,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, count, err := pgRepo.List(tc.page, tc.limit)
			if err != nil {
				log.Fatalf("%s: got:%v", tc.name, err)
			}
			if tc.want.count == count {
				for i, j := range tc.want.list {
					if j.Assignee != got[i].Assignee || j.Title != got[i].Title || j.Summary != got[i].Summary || j.Deadline != got[i].Deadline || j.Status != got[i].Status {
						log.Fatalf("%s: expected:%v got:%v", tc.name, tc.want.list, got)
					}
				}
			} else {
				log.Fatalf("%s: got:%v", tc.name, err)
			}
		})
	}
}

func TestTaskRepo_Delete(t *testing.T) {
	tests := []struct {
		name string
		id   string
	}{
		{
			name: "birinchi",
			id:   "00000000-0000-0000-0000-000000000002",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := pgRepo.Delete(tc.id)
			if err != nil {
				log.Fatalf("%s: got:%v", tc.name, err)
			}
		})
	}
}

func TestTaskRepo_Update(t *testing.T) {
	tests := []struct {
		name  string
		input pb.Task
	}{
		{
			name: "hullas",
			input: pb.Task{
				Id:       "123e4567-e89b-12d3-a456-426614174000",
				Assignee: "Belgium",
				Title:    "History of Moscow",
				Summary:  "The book is good",
				Deadline: "2021-12-13T00:00:00Z",
				Status:   "active",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Update(tc.input)
			if err != nil {
				log.Fatalf("%s: got:%v", tc.name, err)
			}
			tc.input.CreatedAt = got.CreatedAt
			tc.input.UpdatedAt = got.UpdatedAt

			if !reflect.DeepEqual(got, tc.input) {
				log.Fatalf("%s: excepted: %v got: %v", tc.name, tc.input, got)
			}
		})
	}
}

func TestTaskRepo_ListOverdue(t *testing.T) {
	timer, _ := time.Parse("2006-01-02", "2021-12-31")
	var tests = []struct {
		name  string
		page  int64
		limit int64
		time  time.Time
		want  result
	}{
		{
			name:  "success",
			page:  1,
			limit: 2,
			time:  timer,
			want: result{
				list: []pb.Task{
					{
						Id:        "123e4567-e89b-12d3-a456-426614174000",
						Assignee:  "Belgium",
						Title:     "History of Moscow",
						Summary:   "The book is good",
						Status:    "active",
						CreatedAt: "",
					},
					{
						Id:        "d1672aa7-06ae-46fe-b26d-b1292fe6c4f6",
						Assignee:  "Russia",
						Title:     "History of Belgium",
						Summary:   "The book is nimadur",
						Status:    "active",
						CreatedAt: "",
					},
				},
				count: 2,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, count, err := pgRepo.ListOverdue(tc.page, tc.limit, tc.time)
			if err != nil {
				log.Fatalf("%s: got:%v", tc.name, err)
			}
			if count == tc.want.count {
				for i, j := range tc.want.list {
					if j.Assignee != got[i].Assignee || j.Title != got[i].Title || j.Summary != got[i].Summary || j.Status != got[i].Status {
						log.Fatalf("%s: expected:%v got:%v", tc.name, tc.want.list, got)
					}
				}
			} else {
				log.Fatalf("%s: expected: %d got: %d", tc.name, tc.want.count, count)
			}
		})
	}
}
