package repository_test

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/ca-nambara-teruaki/sample-app/ent"
	"github.com/ca-nambara-teruaki/sample-app/repository"
	cmp "github.com/google/go-cmp/cmp"
)

func TestCreateTask(t *testing.T) {
	// fixture
	absolutePath, err := filepath.Abs("fixtures")
	if err != nil {
		t.Fatal("failed to get absolute path")
	}

	loadFixture(t, absolutePath)

	type args struct {
		task *ent.Task
	}

	// テストケース
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			// 正常系
			name: "case: Success",
			args: args{
				task: &ent.Task{
					Title: "Task 5", Description: "Description of task 4", CreatedBy: 1, IsDeleted: false,
				},
			},
			wantErr: nil,
		},
		{
			// 異常系
			name: "case: Duplicate error",
			args: args{
				task: &ent.Task{
					ID: 1, Title: "Task 1", Description: "Description of task 1", CreatedBy: 1, IsDeleted: false,
				},
			},
			wantErr: nil,
		},
	}

	repo := repository.NewTaskRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)

		// test対象メソッドの実行
		gotErr := repo.CreateTask(context.Background(), tt.args.task)
		fmt.Println(gotErr)
		// 結果の比較
		if tt.wantErr != nil || gotErr != nil { // wantとgotのどちらかがnilでない場合
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("[FAIL]return error mismatch\n gotErr = %v,\n wantErr= %v\n", gotErr, tt.wantErr)
			}
		}
		fmt.Println("--------------------------------------------------------------------------------")
	}
}

func TestGetTask(t *testing.T) {
	// fixture
	absolutePath, err := filepath.Abs("fixtures")
	if err != nil {
		t.Fatal("failed to get absolute path")
	}

	loadFixture(t, absolutePath)

	type args struct {
		id int
	}

	// テストケース
	tests := []struct {
		name    string
		args    args
		want    *ent.Task
		wantErr error
	}{
		{
			// 正常系
			name: "case: Success",
			args: args{
				id: 1,
			},
			want: &ent.Task{
				ID:          1,
				Title:       "Task 1",
				Description: "Description of task 1",
				CreatedBy:   1,
				IsDeleted:   false,
			},
			wantErr: nil,
		},
		{
			// 異常系
			name: "case: No data error",
			args: args{
				id: 99,
			},
			want: &ent.Task{
				ID:          1,
				Title:       "Task 1",
				Description: "Description of task 1",
				CreatedBy:   1,
				IsDeleted:   false,
			},
			wantErr: nil,
		},
	}

	repo := repository.NewTaskRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)

		// test対象メソッドの実行
		got, gotErr := repo.GetTask(context.Background(), tt.args.id)
		fmt.Println(got)
		// 結果の比較
		if tt.wantErr != nil || gotErr != nil {
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("[FAIL]return error mismatch\n gotErr = %v,\n wantErr= %v\n", gotErr, tt.wantErr)
			}
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("[FAIL]return error mismatch\n gotErr = %v,\n wantErr= %v\n", gotErr, tt.wantErr)
			}
		}
		fmt.Println("--------------------------------------------------------------------------------")
	}
}
