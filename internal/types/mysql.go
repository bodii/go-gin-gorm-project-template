package types

import "gorm.io/gorm"

type MysqlDBT struct {
	*gorm.DB
}
