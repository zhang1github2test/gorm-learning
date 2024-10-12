package callback

import (
	"gorm.io/gorm"
	"reflect"
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
			reflectValue := reflect.ValueOf(value)
			typeofRe := reflect.TypeOf(value)
			if typeofRe.Kind() != reflect.Ptr {
				called = false
				return called
			}
			reflectValue = reflectValue.Elem()
			typeofRe = typeofRe.Elem()

			// Iterate over the fields
			for i := 0; i < typeofRe.NumField(); i++ {
				field := typeofRe.Field(i)

				// Get the "encryption" tag
				encryptionTag := field.Tag.Get("encryption")

				// Check if the encryption tag is set to "true"
				if encryptionTag == "true" {
					ecryptStr := reflectValue.Field(i).String()
					decrypt, err := AesDecrypt(ecryptStr, DATA_KEY)
					if err != nil {
						db.AddError(err)
					}
					reflectValue.Field(i).SetString(decrypt)
				}
			}
			return called
		})
	}
}
