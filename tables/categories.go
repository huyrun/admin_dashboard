package tables

import (
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types/form"
)

func GetCategoriesTable(ctx *context.Context) table.Table {
	categories := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "categories"
	info := categories.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("Name", "name", db.Varchar).FieldSortable().FieldFilterable()

	info.SetTable(tableName).SetTitle("Categories").SetDescription("Categories").AddCSS(cssTableNoWrap)

	formList := categories.GetForm()
	formList.AddField("ID", "id", db.Int8, form.Number).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Name", "name", db.Varchar, form.Text)

	formList.SetTable(tableName).SetTitle("Categories").SetDescription("Categories").AddCSS(cssTableNoWrap)

	return categories
}
