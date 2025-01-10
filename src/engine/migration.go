package engine

import (
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/huyrun/admin_dashboard/src/config"
)

const (
	baseMigrationFolder = "file://./schema/migration"
)

func Migrate(cfg *config.Config) error {
	d := cfg.Databases["default"]
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", d.User, d.Pwd, d.Host, d.Port, d.Name, strings.TrimSpace(d.ParamStr()))
	mg, err := migrate.New(baseMigrationFolder, databaseUrl)
	if err != nil {
		return err
	}
	defer func(mg *migrate.Migrate) {
		err, _ := mg.Close()
		if err != nil {
			panic(err)
		}
	}(mg)

	return mg.Up()
}
