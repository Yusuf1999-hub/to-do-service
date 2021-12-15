package postgres

import (
    "reflect"
    "testing"

    pb "github.com/Yusuf1999-hub/to-do-service/genproto"
)

func TestTaskRepo_Create(t *testing.T) {
    tests := []struct {
        name    string
        input   pb.Task
        want    pb.Task
        wantErr bool
    }{
        {
            name: "successful",
            input: pb.Task{
                Assignee: "asdf",
				Title: "sadf",
				Summary: "fsda",
            },
            want: pb.Task{
                
            },
            wantErr: false,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got, err := pgRepo.Create(tc.input)
            if err != nil {
                t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
            }
            got.Id=0
            if !reflect.DeepEqual(tc.want, got) {
                t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
            }
        })
    }
}