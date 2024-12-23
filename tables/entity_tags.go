package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetEntitytagsTable(ctx *context.Context) table.Table {

	entityTags := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "entity_tags"
	info := entityTags.GetInfo().HideFilterArea()

	info.AddField("Entity_id", "entity_id", db.Int8)
	info.AddField("Tag_id", "tag_id", db.Text)

	info.SetTable(tableName).SetTitle("Entitytags").SetDescription("Entitytags")

	formList := entityTags.GetForm()
	formList.AddField("Entity_id", "entity_id", db.Int8, form.Text)
	formList.AddField("Tag_id", "tag_id", db.Text, form.RichText)

	formList.SetTable(tableName).SetTitle("Entitytags").SetDescription("Entitytags")

	return entityTags
}
