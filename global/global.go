package global

import (
	"github.com/jmoiron/sqlx"
)

const (
	SERVER_NAME string = "wow-admin"
)

var (
	DB *sqlx.DB
)
