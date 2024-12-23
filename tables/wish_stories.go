package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWishstoriesTable(ctx *context.Context) table.Table {

	wishStories := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "wish_stories"
	info := wishStories.GetInfo().HideFilterArea()

	info.AddField("Created_at", "created_at", db.Timestamptz)
	info.AddField("Updated_at", "updated_at", db.Timestamptz)
	info.AddField("Entity_id", "entity_id", db.Int8)
	info.AddField("Body", "body", db.Varchar)
	info.AddField("Image", "image", db.Varchar)
	info.AddField("Status", "status", db.Varchar)

	info.SetTable(tableName).SetTitle("Wishstories").SetDescription("Wishstories")

	formList := wishStories.GetForm()
	formList.AddField("Created_at", "created_at", db.Timestamptz, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamptz, form.Datetime)
	formList.AddField("Entity_id", "entity_id", db.Int8, form.Text)
	formList.AddField("Body", "body", db.Varchar, form.Text)
	formList.AddField("Image", "image", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Varchar, form.Text)

	formList.SetTable(tableName).SetTitle("Wishstories").SetDescription("Wishstories")

	return wishStories
}
