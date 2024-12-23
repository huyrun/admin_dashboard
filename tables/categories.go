package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCategoriesTable(ctx *context.Context) table.Table {
	
	categories := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "categories"
	info := categories.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int8).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)

	info.SetTable(tableName).SetTitle("Categories").SetDescription("Categories")

	formList := categories.GetForm()
	formList.AddField("Id", "id", db.Int8, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)

	formList.SetTable(tableName).SetTitle("Categories").SetDescription("Categories")

	return categories
}
