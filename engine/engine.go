package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/huyrun/go-admin/engine"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"gorm.io/gorm"
	"project/models"
	"project/pages"
)

type Engine struct {
	R    *gin.Engine
	E    *engine.Engine
	DB   *gorm.DB
	Conn db.Connection
}

func NewEngine(generatorFn func(db *gorm.DB, conn db.Connection) (map[string]table.Generator, error)) (*Engine, error) {
	e := engine.Default()
	r := NewRouter()

	if err := e.AddConfigFromYAML("./config.yml").
		Use(r); err != nil {
		return nil, err
	}

	e.HTML("GET", "/admin", pages.GetWelcome)

	conn := e.PostgresqlConnection()
	gormDB, err := models.NewDB(conn)
	if err != nil {
		return nil, err
	}

	generator, err := generatorFn(gormDB, conn)
	if err != nil {
		return nil, err
	}
	e.AddGenerators(generator)

	return &Engine{
		R:    r,
		E:    e,
		DB:   gormDB,
		Conn: conn,
	}, nil
}
