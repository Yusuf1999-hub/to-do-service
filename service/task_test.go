package service

import (
	"context"
	"fmt"
	pb "github.com/Yusuf1999-hub/to-do-service/genproto"
	"reflect"
	"testing"
)

func TestTaskService_Create(t *testing.T) {
	tests := []struct {
		name  string
		input pb.Task
		want  pb.Task
	}{
		{
			name: "successful",
			input: pb.Task{
				Assignee: "asdf",
				Title:    "nimadur",
				Summary:  "asf",
				Deadline: "2021-12-13T00:10:00Z",
				Status:   "active",
			},
			want: pb.Task{
				Assignee: "asdf",
				Title:    "nimadur",
				Summary:  "asf",
				Deadline: "2021-12-13T00:10:00Z",
				Status:   "active",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.input)
			got, err := task.Create(context.Background(), &tc.input)
			fmt.Println(got)
			if err != nil {
				t.Error("failed to create task", err)
			}
			tc.want.Id = got.Id
			tc.want.CreatedAt = (*got).CreatedAt
			tc.want.UpdatedAt = got.UpdatedAt

			if !reflect.DeepEqual(*got, tc.want) {
				t.Fatalf("%s: expected: %v  got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskService_Get(t *testing.T) {
	tests := []struct {
		name string
		id   pb.ByIdReq
		want pb.Task
	}{
		{
			name: "successful",
			id: pb.ByIdReq{
				Id: "123e4567-e89b-12d3-a456-426614174000",
			},
			want: pb.Task{
				Id:        "123e4567-e89b-12d3-a456-426614174000",
				Assignee:  "Belgium",
				Title:     "History of Moscow",
				Summary:   "The book is good",
				Deadline:  "2021-12-13T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-21T08:25:03.12633Z",
				UpdatedAt: "2021-12-21T08:26:23.909723Z",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := task.Get(context.Background(), &tc.id)
			if err != nil {
				t.Fatalf("Failed to get the task %v", err)
			}

			if !reflect.DeepEqual(*got, tc.want) {
				t.Fatalf("%s: expected: %v  got: %v", tc.name, tc.want, *got)
			}
		})
	}
}

func TestTaskService_List(t *testing.T) {
	tests := []struct {
		name  string
		input pb.ListReq
		want  []*pb.Task
		count int64
	}{
		{
			name: "successful",
			input: pb.ListReq{
				Page:  3,
				Limit: 2,
			},
			want: []*pb.Task{
				{
					Id:       "d1672aa7-06ae-46fe-b26d-b1292fe6c4f6",
					Assignee: "asdf",
					Title:    "nimadur",
					Summary:  "asf",
					Deadline: "2021-12-13T00:10:00Z",
					Status:   "active",
				},
			},
			count: 7,
		},
	}
	for _, tc := range tests {
		got, err := task.List(context.Background(), &tc.input)
		if err != nil {
			t.Errorf("Failed to list the task %v", err)
		}
		if tc.count == got.Count {
			for i, _ := range got.Tasks {
				if !reflect.DeepEqual(tc.want[i], got.Tasks[i]) {
					t.Fatalf("%s: expected: %v  got: %v", tc.name, tc.want, *got)
				} else {
					t.Fatalf("%s expected:%d got count:%d", tc.name, tc.count, got.Count)
				}
			}
		}
	}
}

func TestTaskService_Update(t *testing.T) {
	tests := []struct {
		name  string
		input pb.Task
	}{
		{
			name: "successful",
			input: pb.Task{
				Id:       "d1672aa7-06ae-46fe-b26d-b1292fe6c4f6",
				Assignee: "Russia",
				Title:    "History of Belgium",
				Summary:  "The book is nimadur",
				Deadline: "2021-07-29T00:00:00Z",
				Status:   "active",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := task.Update(context.Background(), &tc.input)
			if err != nil {
				t.Error("Failed to update the task")
			}
			tc.input.UpdatedAt = got.UpdatedAt
			tc.input.CreatedAt = (*got).CreatedAt
			if !reflect.DeepEqual(*got, tc.input) {
				t.Errorf("%s expected:%v got:%v", tc.name, tc.input, got)
			}
		})
	}
}

func TestTaskService_Delete(t *testing.T) {
	tests := []struct {
		name string
		id   pb.ByIdReq
		want pb.EmptyResp
	}{
		{
			name: "successful",
			id: pb.ByIdReq{
				Id: "00000000-0000-0000-0000-000000000001",
			},
		},
		{
			name: "second",
			id: pb.ByIdReq{
				Id: "00000000-0000-0000-0000-000000000002",
			},
		},
	}
	for _, tc := range tests {
		got, err := task.Delete(context.Background(), &tc.id)
		if err != nil {
			t.Error("Failed to delete the task")
		}
		if !reflect.DeepEqual(*got, tc.want) {
			t.Errorf("%s  got:%v", tc.name, got)
		}
	}
}
func TestTaskService_ListOverdue(t *testing.T) {
	var tests = []struct {
		name  string
		input pb.ListOverReq
		want  []pb.Task
		count int64
	}{
		{
			name: "successful",
			input: pb.ListOverReq{
				Page:  1,
				Limit: 2,
				Time:  "2021-12-31",
			},
			want: []pb.Task{
				{
					Id:       "123e4567-e89b-12d3-a456-426614174000",
					Assignee: "Belgium",
					Title:    "History of Moscow",
					Summary:  "The book is good",
					Status:   "active",
				},
				{
					Id:       "d1672aa7-06ae-46fe-b26d-b1292fe6c4f6",
					Assignee: "Russia",
					Title:    "History of Belgium",
					Summary:  "The book is nimadur",
					Status:   "active",
				},
			},
			count: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := task.ListOverdue(context.Background(), &tc.input)
			fmt.Println(got)

			if err != nil {
				t.Error("Failed to listOverReq the task")
			}
		})

	}
}
