package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetEntitycommentsTable(ctx *context.Context) table.Table {

	entityComments := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "activities"
	info := entityComments.GetInfo().HideFilterArea()

	info.AddField("Created_at", "created_at", db.Timestamptz)
	info.AddField("Updated_at", "updated_at", db.Timestamptz)
	info.AddField("Entity_id", "entity_id", db.Int8)
	info.AddField("User_id", "user_id", db.Text)
	info.AddField("Comment_no", "comment_no", db.Int8)
	info.AddField("Comment", "comment", db.Varchar)

	info.SetTable(tableName).SetTitle("Entitycomments").SetDescription("Entitycomments")

	formList := entityComments.GetForm()
	formList.AddField("Created_at", "created_at", db.Timestamptz, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamptz, form.Datetime)
	formList.AddField("Entity_id", "entity_id", db.Int8, form.Text)
	formList.AddField("User_id", "user_id", db.Text, form.Text)
	formList.AddField("Comment_no", "comment_no", db.Int8, form.Text)
	formList.AddField("Comment", "comment", db.Varchar, form.Text)

	formList.SetTable(tableName).SetTitle("Entitycomments").SetDescription("Entitycomments")

	return entityComments
}
