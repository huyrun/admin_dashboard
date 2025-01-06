package tables

import (
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types/form"
)

func GetEntitytagsTable(ctx *context.Context) table.Table {
	tableConfig := table.DefaultConfigWithDriver("postgresql")
	tableConfig.PrimaryKey = table.PrimaryKey{
		Type: db.Int8,
		Name: "entity_id",
	}
	entityTags := table.NewDefaultTable(ctx, tableConfig)
	tableName := "entity_tags"
	info := entityTags.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("Entity ID", "entity_id", db.Int8).FieldFilterable().FieldSortable()
	info.AddField("Tag ID", "tag_id", db.Text).FieldFilterable().FieldSortable()

	info.SetTable(tableName).SetTitle("EntityTags").SetDescription("EntityTags").AddCSS(cssTableNoWrap)

	formList := entityTags.GetForm()
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("Tag ID", "tag_id", db.Text, form.RichText)

	formList.SetTable(tableName).SetTitle("EntityTags").SetDescription("EntityTags")

	return entityTags
}
