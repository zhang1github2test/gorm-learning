package callback

import (
	"gorm.io/gorm"
	"reflect"
)

type EncryptHook interface {
	Encrypt(db *gorm.DB) error
}

func Encrypt(db *gorm.DB) {
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
					decrypt, err := AesEncrypt([]byte(ecryptStr), DATA_KEY)
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
