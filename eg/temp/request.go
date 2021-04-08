package temp

import (
	"github.com/zngue/go_helper/pkg"
	"gorm.io/gorm"
)

type Request struct {
	pkg.CommonRequest
}

func (r *Request) Common(db *gorm.DB) *gorm.DB {
	tx := r.Init(db, *r)
	return tx
}
