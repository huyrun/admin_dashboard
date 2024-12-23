package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetLikedentitiesTable(ctx *context.Context) table.Table {

	likedEntities := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "liked_entities"
	info := likedEntities.GetInfo().HideFilterArea()

	info.AddField("Entity_id", "entity_id", db.Int8)
	info.AddField("User_id", "user_id", db.Text)
	info.AddField("Amount", "amount", db.Int2)

	info.SetTable(tableName).SetTitle("Likedentities").SetDescription("Likedentities")

	formList := likedEntities.GetForm()
	formList.AddField("Entity_id", "entity_id", db.Int8, form.Text)
	formList.AddField("User_id", "user_id", db.Text, form.Text)
	formList.AddField("Amount", "amount", db.Int2, form.Text)

	formList.SetTable(tableName).SetTitle("Likedentities").SetDescription("Likedentities")

	return likedEntities
}
