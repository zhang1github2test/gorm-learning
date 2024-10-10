package repository

import (
	"errors"
	"github.com/zhang1github2test/gorm-learning/database"
	"github.com/zhang1github2test/gorm-learning/model"
	"gorm.io/gorm"
	"testing"
	"time"
)

type fields struct {
	Db *gorm.DB
}
type args struct {
	user *model.User
}

var field = fields{
	Db: database.GLOBALDB,
}

func TestUserDao_Create(t *testing.T) {
	email := "89954554554@163.com"
	birthday := time.Now().AddDate(-18, 0, 0)

	user := &model.User{
		ID:       100014,
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
	type args struct {
		users     *[]model.User
		batchSize int
	}
	us := []model.User{{Name: "zhangshenglu1"}, {Name: "zhangshenglu2"}, {Name: "zhangshenglu3"}}
	arg := args{
		users:     &us,
		batchSize: 100,
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

func TestUserDao_First(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "First",
			fields: field,
			args: args{
				user: &model.User{},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.First(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("First() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("First() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_Take(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "Take_test",
			fields: field,
			args: args{
				user: &model.User{},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.Take(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Take() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Take() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_Last(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "Last",
			fields: field,
			args: args{
				user: &model.User{},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.Last(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Last() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Last() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_Find(t *testing.T) {
	type args struct {
		users *[]model.User
	}
	us := &[]model.User{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "Find",
			fields: field,
			args: args{
				users: us,
			},
			want:    24,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.Find(tt.args.users)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestUserDao_Scan(t *testing.T) {
	type args struct {
		users *[]model.User
	}
	us := []model.User{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "Scan",
			fields: field,
			args: args{
				users: &us,
			},
			want:    24,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.Scan(tt.args.users)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Scan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_SelectSpecField(t *testing.T) {
	type args struct {
		users *[]model.User
	}
	us := []model.User{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "SelectSpecField",
			fields: field,
			args: args{
				users: &us,
			},
			want:    24,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.SelectSpecField(tt.args.users, "name")
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectSpecField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SelectSpecField() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_StringQuery(t *testing.T) {
	type args struct {
		users *[]model.User
	}
	us := []model.User{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "TestUserDao_StringQuery",
			fields: field,
			args: args{
				users: &us,
			},
			want:    24,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.StringQuery(tt.args.users)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_StructQuery(t *testing.T) {
	type args struct {
		users *[]model.User
		user  *model.User
	}

	us := []model.User{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "id_exist",
			fields: field,
			args: args{
				users: &us,
				user: &model.User{
					ID: 100003,
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.StructQuery(tt.args.users, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StructQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_MapQuery(t *testing.T) {
	type args struct {
		users *[]model.User
	}

	us := []model.User{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "MapQuery",
			fields: field,
			args: args{
				users: &us,
			},
			want:    8,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.MapQuery(tt.args.users)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MapQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_NotQuery(t *testing.T) {

	type args struct {
		users *[]model.User
		user  *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "not_equal",
			fields: field,
			args: args{
				users: &[]model.User{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userDao.NotQuery(tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("NotQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_OrQuery(t *testing.T) {
	type args struct {
		users *[]model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "or_query",
			fields: field,
			args: args{
				users: &[]model.User{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userDao.OrQuery(tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("OrQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_Order(t *testing.T) {
	type args struct {
		users *[]model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "order",
			fields: field,
			args: args{
				users: &[]model.User{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userdao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userdao.Order(tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("Order() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_Group(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "group",
			fields:  field,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userdao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userdao.Group(); (err != nil) != tt.wantErr {
				t.Errorf("Group() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_Having(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "group",
			fields:  field,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userdao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userdao.Having(); (err != nil) != tt.wantErr {
				t.Errorf("Having() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_PageQuery(t *testing.T) {
	type args struct {
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "page_Query",
			fields: field,
			args: args{
				offset: 10,
				limit:  10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userdao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userdao.PageQuery(tt.args.offset, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("PageQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_UpdateSingle(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "updateSingle",
			fields:  field,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userDao.UpdateSingle(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSingle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_Updates(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Updates",
			fields:  field,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userDao.Updates(); (err != nil) != tt.wantErr {
				t.Errorf("Updates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_Delete(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Delete",
			fields:  field,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userDao.Delete(); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_AutoMigrate(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "automigrate",
			fields:  field,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userdao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userdao.AutoMigrate(); (err != nil) != tt.wantErr {
				t.Errorf("AutoMigrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_Transaction(t *testing.T) {
	email := "89954554554@163.com"
	birthday := time.Now().AddDate(-18, 0, 0)

	user := &model.User{
		ID:       100003,
		Name:     "zhangshenglu",
		Email:    &email,
		Age:      18,
		Phone:    "18454612132",
		Birthday: &birthday,
	}

	user2 := &model.User{
		ID:       100004,
		Name:     "zhangshenglu",
		Email:    &email,
		Age:      18,
		Phone:    "18454612132",
		Birthday: &birthday,
	}
	type args struct {
		f    func(tx *gorm.DB, parm interface{}) error
		parm interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "commit",
			fields: field,
			args: args{
				f: func(tx *gorm.DB, parm interface{}) error {
					tx.Save(parm)
					return nil
				},
				parm: user,
			},
			wantErr: false,
		},
		{
			name:   "rollback",
			fields: field,
			args: args{
				f: func(tx *gorm.DB, parm interface{}) error {
					tx.Save(parm)
					return errors.New("模拟发生错误")
				},
				parm: user2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userdao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userdao.Transaction(tt.args.f, tt.args.parm); (err != nil) != tt.wantErr {
				t.Errorf("Transaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_SaveWithTransactionByManual(t *testing.T) {
	email := "89954554554@163.com"
	birthday := time.Now().AddDate(-18, 0, 0)

	user := &model.User{
		ID:       100006,
		Name:     "zhangshenglu",
		Email:    &email,
		Age:      18,
		Phone:    "18454612132",
		Birthday: &birthday,
	}

	user2 := &model.User{
		ID:       100007,
		Name:     "zhangshenglu",
		Email:    &email,
		Age:      18,
		Phone:    "18454612132",
		Birthday: &birthday,
	}
	type args struct {
		user     *model.User
		rollback bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "commit",
			wantErr: false,
			fields:  field,
			args: args{
				user:     user,
				rollback: false,
			},
		},
		{
			name:    "rollback",
			wantErr: true,
			fields:  field,
			args: args{
				user:     user2,
				rollback: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userdao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userdao.SaveWithTransactionByManual(tt.args.user, tt.args.rollback); (err != nil) != tt.wantErr {
				t.Errorf("SaveWithTransactionByManual() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_SingleSessionContext(t *testing.T) {
	user := &model.User{}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "SingleSession",
			fields: field,
			args: args{
				user: user,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userdao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userdao.SingleSessionContext(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("SingleSessionContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_SessionContext(t *testing.T) {
	user := &model.User{}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "SessionContext",
			fields: field,
			args: args{
				user: user,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userdao := &UserDao{
				Db: tt.fields.Db,
			}
			if err := userdao.SessionContext(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("SessionContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
