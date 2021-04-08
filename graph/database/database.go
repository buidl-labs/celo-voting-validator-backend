package database

import (
	"github.com/go-pg/pg/v10"
)

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
