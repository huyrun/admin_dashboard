package tables

import (
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types/form"
)

func GetTagsTable(ctx *context.Context) table.Table {

	tags := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "tags"
	info := tags.GetInfo().HideFilterArea()

	info.AddField("Tag_name", "tag_name", db.Text)

	info.SetTable(tableName).SetTitle("Tags").SetDescription("Tags").AddCSS(cssTableNoWrap)

	formList := tags.GetForm()
	formList.AddField("Tag_name", "tag_name", db.Text, form.RichText)

	formList.SetTable(tableName).SetTitle("Tags").SetDescription("Tags")

	return tags
}
