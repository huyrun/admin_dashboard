package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/huyrun/admin_dashboard/src/config"
	"github.com/huyrun/go-admin/engine"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Engine struct {
	R      *gin.Engine
	E      *engine.Engine
	DB     *gorm.DB
	Conn   db.Connection
	Config *config.Config
}

func NewEngine(generatorFn func(db *gorm.DB, conn db.Connection) (map[string]table.Generator, error)) (*Engine, error) {
	e := engine.Default()
	r := NewRouter()

	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	err = Migrate(cfg)
	if err != nil {
		return nil, err
	}

	if err = e.AddConfig(cfg.Config).
		Use(r); err != nil {
		return nil, err
	}

	conn := e.PostgresqlConnection()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: conn.GetDB("default")}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	generator, err := generatorFn(gormDB, conn)
	if err != nil {
		return nil, err
	}
	e.AddGenerators(generator)

	return &Engine{
		R:      r,
		E:      e,
		DB:     gormDB,
		Conn:   conn,
		Config: cfg,
	}, nil
}
