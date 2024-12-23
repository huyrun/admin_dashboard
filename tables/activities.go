package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetActivitiesTable(ctx *context.Context) table.Table {
	
	activities := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "activities"
	info := activities.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int8).
		FieldFilterable()
	info.AddField("Action", "action", db.Varchar)
	info.AddField("Points", "points", db.Int2)
	info.AddField("User_id", "user_id", db.Text)
	info.AddField("Entity_id", "entity_id", db.Int8)
	info.AddField("Created_at", "created_at", db.Timestamptz)

	info.SetTable(tableName).SetTitle("Activities").SetDescription("Activities")

	formList := activities.GetForm()
	formList.AddField("Id", "id", db.Int8, form.Default)
	formList.AddField("Action", "action", db.Varchar, form.Text)
	formList.AddField("Points", "points", db.Int2, form.Text)
	formList.AddField("User_id", "user_id", db.Text, form.Text)
	formList.AddField("Entity_id", "entity_id", db.Int8, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamptz, form.Datetime)

	formList.SetTable(tableName).SetTitle("Activities").SetDescription("Activities")

	return activities
}
