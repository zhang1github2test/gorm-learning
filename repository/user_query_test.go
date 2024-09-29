package repository

import (
	"github.com/zhang1github2test/gorm-learning/database"
	"github.com/zhang1github2test/gorm-learning/model"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestUserDao_Create(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		user *model.User
	}
	field := fields{
		Db: database.GLOBALDB,
	}
	email := "89954554554@163.com"
	birthday := time.Now().AddDate(-18, 0, 0)

	user := &model.User{
		ID:       100003,
		Name:     "zhangshenglu",
		Email:    &email,
		Age:      18,
		Birthday: &birthday,
	}

	arg := args{
		user: user,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "createFirst",
			fields:  field,
			args:    arg,
			want:    1,
			wantErr: false,
		},
		{
			name:    "createAgain",
			fields:  field,
			args:    arg,
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usedao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := usedao.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_Save(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		user *model.User
	}

	field := fields{
		Db: database.GLOBALDB,
	}
	email := "89954554554@163.com"
	birthday := time.Now().AddDate(-18, 0, 0)

	user := &model.User{
		ID:       100002,
		Name:     "zhangshenglu",
		Email:    &email,
		Age:      18,
		Birthday: &birthday,
	}

	arg := args{
		user: user,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "SaveFirst",
			fields:  field,
			args:    arg,
			want:    1,
			wantErr: false,
		},
		{
			name:    "SaveAgain",
			fields:  field,
			args:    arg,
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usedao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := usedao.Save(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_CreateInBatches(t *testing.T) {
	type fields struct {
		Db *gorm.DB
	}
	type args struct {
		users     *[]model.User
		batchSize int
	}
	us := []model.User{{ID: 100003,
		Name: "zhangshenglu1"}, {Name: "zhangshenglu2"}, {Name: "zhangshenglu3"}}
	arg := args{
		users:     &us,
		batchSize: 2,
	}
	field := fields{
		Db: database.GLOBALDB,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "batchSave",
			fields:  field,
			args:    arg,
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.CreateInBatches(tt.args.users, tt.args.batchSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateInBatches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateInBatches() got = %v, want %v", got, tt.want)
			}
		})
	}
}
