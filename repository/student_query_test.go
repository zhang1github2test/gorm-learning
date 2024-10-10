package repository

import (
	"fmt"
	"github.com/zhang1github2test/gorm-learning/model"
	"testing"
	"time"
)

func TestStudentDao_Create(t *testing.T) {
	birthday := time.Now().AddDate(-18, 0, 0)
	s := &model.Student{
		ID:       100014,
		Name:     "zhangshenglu",
		Age:      18,
		Birthday: &birthday,
	}
	type args struct {
		student *model.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "create",
			fields: field,
			args: args{
				student: s,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stDao := &StudentDao{
				Db: tt.fields.Db,
			}
			if err := stDao.Create(tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStudentDao_First(t *testing.T) {
	s := &model.Student{}
	type args struct {
		student *model.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "create",
			fields: field,
			args: args{
				student: s,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stDao := &StudentDao{
				Db: tt.fields.Db,
			}
			if err := stDao.First(tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("First() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	fmt.Println("测试完成")
}
