// Package database this package made to manage database actions and connections on top of an ORM
// this package is just a sample change it to whatever you need
// checkout this link https://gitea.com/xorm/xorm
// I would recommend to use EngineGroup (this is just a demo)
package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/mhkarimi1383/goAPIBaseProject/configuration"
	"github.com/mhkarimi1383/goAPIBaseProject/logger"
	"github.com/mhkarimi1383/goAPIBaseProject/types"
	"xorm.io/xorm"
)

var (
	engine *xorm.Engine
)

func init() {
	var err error = nil
	cfg, _ := configuration.GetConfig()
	engine, err = xorm.NewEngine(cfg.DatabaseDriver, "file:db.sqlite?cache=shared")
	if err != nil {
		logger.Fatalf(true, "problem while creating database engine %v", err)
	}
}

func SyncDatabase() error {
	return engine.Sync(new(types.User))
}

func CreateUser(users ...types.User) (int64, error) {
	var s []any
	for i, v := range users {
		s[i] = &v
	}
	return engine.Insert(s...)
}

func GetUser(user types.User, whereStatement string, columns string, descColumns string) (bool, error) {
	if descColumns != "" {
		return engine.Where(whereStatement).Cols(columns).Desc(descColumns).Get(&user)
	}
	return engine.Where(whereStatement).Cols(columns).Get(&user)
}
