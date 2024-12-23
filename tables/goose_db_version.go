package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetGoosedbversionTable(ctx *context.Context) table.Table {

	gooseDbVersion := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "goose_db_version"
	info := gooseDbVersion.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int4).
		FieldFilterable()
	info.AddField("Version_id", "version_id", db.Int8)
	info.AddField("Is_applied", "is_applied", db.Bool)
	info.AddField("Tstamp", "tstamp", db.Timestamp)

	info.SetTable(tableName).SetTitle("Goosedbversion").SetDescription("Goosedbversion")

	formList := gooseDbVersion.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default)
	formList.AddField("Version_id", "version_id", db.Int8, form.Text)
	formList.AddField("Is_applied", "is_applied", db.Bool, form.Text)
	formList.AddField("Tstamp", "tstamp", db.Timestamp, form.Datetime)

	formList.SetTable(tableName).SetTitle("Goosedbversion").SetDescription("Goosedbversion")

	return gooseDbVersion
}
