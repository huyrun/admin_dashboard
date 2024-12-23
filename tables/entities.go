package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetEntitiesTable(ctx *context.Context) table.Table {

	entities := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "entities"
	info := entities.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int8).
		FieldFilterable()

	info.SetTable(tableName).SetTitle("Entities").SetDescription("Entities")

	formList := entities.GetForm()
	formList.AddField("Id", "id", db.Int8, form.Default)

	formList.SetTable(tableName).SetTitle("Entities").SetDescription("Entities")

	return entities
}
