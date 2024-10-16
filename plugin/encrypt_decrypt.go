package plugin

import (
	"github.com/zhang1github2test/gorm-learning/callback"
	"gorm.io/gorm"
)

type Encrypt struct {
}

func (encrypt *Encrypt) Name() string {
	return "my_customize:encrypt_plugin"
}
func (encrypt *Encrypt) Initialize(db *gorm.DB) error {
	callback.Register(db)
	return nil
}
