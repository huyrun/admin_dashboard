package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetSavedentitiesTable(ctx *context.Context) table.Table {

	savedEntities := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "saved_entities"
	info := savedEntities.GetInfo().HideFilterArea()

	info.AddField("Entity_id", "entity_id", db.Int8)
	info.AddField("User_id", "user_id", db.Text)
	info.AddField("Type", "type", db.Text)

	info.SetTable(tableName).SetTitle("Savedentities").SetDescription("Savedentities")

	formList := savedEntities.GetForm()
	formList.AddField("Entity_id", "entity_id", db.Int8, form.Text)
	formList.AddField("User_id", "user_id", db.Text, form.Text)
	formList.AddField("Type", "type", db.Text, form.RichText)

	formList.SetTable(tableName).SetTitle("Savedentities").SetDescription("Savedentities")

	return savedEntities
}
