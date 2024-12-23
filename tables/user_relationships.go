package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUserrelationshipsTable(ctx *context.Context) table.Table {

	userRelationships := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "user_relationships"
	info := userRelationships.GetInfo().HideFilterArea()

	info.AddField("First_user_id", "first_user_id", db.Text)
	info.AddField("Second_user_id", "second_user_id", db.Text)
	info.AddField("First_to_second_status", "first_to_second_status", db.Text)
	info.AddField("Second_to_first_status", "second_to_first_status", db.Text)
	info.AddField("Are_friends", "are_friends", db.Bool)

	info.SetTable(tableName).SetTitle("Userrelationships").SetDescription("Userrelationships")

	formList := userRelationships.GetForm()
	formList.AddField("First_user_id", "first_user_id", db.Text, form.Text)
	formList.AddField("Second_user_id", "second_user_id", db.Text, form.Text)
	formList.AddField("First_to_second_status", "first_to_second_status", db.Text, form.RichText)
	formList.AddField("Second_to_first_status", "second_to_first_status", db.Text, form.RichText)
	formList.AddField("Are_friends", "are_friends", db.Bool, form.Text)

	formList.SetTable(tableName).SetTitle("Userrelationships").SetDescription("Userrelationships")

	return userRelationships
}
