package postgres

import (
	"log"
	"reflect"
	"testing"

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
				Assignee: "Uzbekistan",
				Title:    "History of Tashkent",
				Summary:  "The book is very good",
				Deadline: "2020-12-13",
				Status:   "active",
			},
			want: pb.Task{
				Assignee: "Uzbekistan",
				Title:    "History of Tashkent",
				Summary:  "The book is very good",
				Deadline: "2020-12-13T00:00:00Z",
				Status:   "active",
			},
		},
		{
			name: "additional",
			input: pb.Task{
				Assignee: "Russia",
				Title:    "History of Russia",
				Summary:  "The book is good",
				Deadline: "2020-12-13",
				Status:   "active",
			},
			want: pb.Task{
				Assignee: "Russia",
				Title:    "History of Russia",
				Summary:  "The book is good",
				Deadline: "2020-12-13T00:00:00Z",
				Status:   "active",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Create(tc.input)
			if err != nil {
				log.Fatalf("%s: got: %v", tc.name, err)
			}
			got.Id = 0

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Get(t *testing.T) {
	tests := []struct {
		name string
		id   int64
		want pb.Task
	}{
		{
			name: "successful",
			id:   21,
			want: pb.Task{
				Assignee: "Russia",
				Title:    "History of Russia",
				Summary:  "The book is good",
				Deadline: "2020-12-13T00:00:00Z",
				Status:   "active",
			},
		},
		{
			name: "additional",
			id:   22,
			want: pb.Task{
				Assignee: "Uzbekistan",
				Title:    "History of Tashkent",
				Summary:  "The book is very good",
				Deadline: "2020-12-13T00:00:00Z",
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
}

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
			page:  11,
			limit: 2,
			want: result{
				list: []pb.Task{
					{
						Id:       21,
						Assignee: "Russia",
						Title:    "History of Russia",
						Summary:  "The book is good",
						Deadline: "2020-12-13T00:00:00Z",
						Status:   "active",
					},
					{
						Id:       22,
						Assignee: "Uzbekistan",
						Title:    "History of Tashkent",
						Summary:  "The book is very good",
						Deadline: "2020-12-13T00:00:00Z",
						Status:   "active",
					},
				},
				count: 24,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, count, err := pgRepo.List(tc.page, tc.limit)
			if err != nil {
				log.Fatalf("%s: got:%v", tc.name, err)
			}
			if tc.want.count == int64(count) {
				for i, j := range tc.want.list {
					if j.Assignee != got[i].Assignee || j.Title != got[i].Title || j.Summary != got[i].Summary || j.Deadline != got[i].Deadline || j.Status != got[i].Status {
						log.Fatalf("%s: expected:%v got:%v", tc.name, tc.want, got)
					}
				}
			} else {
				log.Fatalf("%s: got:%v", tc.name, err)
			}
		})
	}
}

/*
func TestTaskRepo_Delete(t *testing.T) {
	tests := []struct {
		id int64,
	}
}
*/
