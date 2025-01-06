package tables

import (
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types/form"
)

func GetEntitiesTable(ctx *context.Context) table.Table {
	entities := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "entities"
	info := entities.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()

	info.SetTable(tableName).SetTitle("Entities").SetDescription("Entities").AddCSS(cssTableNoWrap)

	formList := entities.GetForm()
	formList.AddField("Id", "id", db.Int8, form.Number).FieldDisplayButCanNotEditWhenUpdate()

	formList.SetTable(tableName).SetTitle("Entities").SetDescription("Entities").AddCSS(cssTableNoWrap)

	return entities
}
