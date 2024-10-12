package callback

import (
	"gorm.io/gorm"
)

func Register(db *gorm.DB) {
	db.Callback().Query().After("gorm:after_query").Register("customer:decrypt_query", Decrypt)
	db.Callback().Create().Before("gorm:before_create").Register("customer:encrypt_create", Encrypt)
	db.Callback().Update().Before("gorm:before_update").Register("customer:encrypt_update", Encrypt)
}

var DATA_KEY = []byte("0123456789123456")

type DecryptCreate interface {
	Decrypt(db *gorm.DB) error
}

// 只针对
func Decrypt(db *gorm.DB) {
	if db.Error == nil && db.Statement.Schema != nil && !db.Statement.SkipHooks {
		callMethod(db, func(value interface{}, tx *gorm.DB) (called bool) {
			if i, ok := value.(DecryptCreate); ok {
				called = true
				db.AddError(i.Decrypt(tx))
			}

			return called
		})
	}
}
