package tables

import (
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types/form"
)

func GetWishesTable(ctx *context.Context) table.Table {

	wishes := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "wishes"
	info := wishes.GetInfo().HideFilterArea()

	info.AddField("Created_at", "created_at", db.Timestamptz)
	info.AddField("Updated_at", "updated_at", db.Timestamptz)
	info.AddField("Entity_id", "entity_id", db.Int8)
	info.AddField("Type", "type", db.Varchar)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Description", "description", db.Varchar)
	info.AddField("Story", "story", db.Varchar)
	info.AddField("Price", "price", db.Int8)
	info.AddField("Currency", "currency", db.Varchar)
	info.AddField("Category_id", "category_id", db.Int8)
	info.AddField("Visible_by", "visible_by", db.Int8)
	info.AddField("Image", "image", db.Varchar)
	info.AddField("User_id", "user_id", db.Text)
	info.AddField("Status", "status", db.Varchar)

	info.SetTable(tableName).SetTitle("Wishes").SetDescription("Wishes").AddCSS(cssTableNoWrap)

	formList := wishes.GetForm()
	formList.AddField("Created_at", "created_at", db.Timestamptz, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamptz, form.Datetime)
	formList.AddField("Entity_id", "entity_id", db.Int8, form.Text)
	formList.AddField("Type", "type", db.Varchar, form.Text)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.AddField("Story", "story", db.Varchar, form.Text)
	formList.AddField("Price", "price", db.Int8, form.Text)
	formList.AddField("Currency", "currency", db.Varchar, form.Text)
	formList.AddField("Category_id", "category_id", db.Int8, form.Text)
	formList.AddField("Visible_by", "visible_by", db.Int8, form.Text)
	formList.AddField("Image", "image", db.Varchar, form.Text)
	formList.AddField("User_id", "user_id", db.Text, form.Text)
	formList.AddField("Status", "status", db.Varchar, form.Text)

	formList.SetTable(tableName).SetTitle("Wishes").SetDescription("Wishes")

	return wishes
}
