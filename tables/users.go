package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUsersTable(ctx *context.Context) table.Table {

	users := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "users"
	info := users.GetInfo().HideFilterArea()

	info.AddField("Created_at", "created_at", db.Timestamptz)
	info.AddField("Updated_at", "updated_at", db.Timestamptz)
	info.AddField("Id", "id", db.Text).
		FieldFilterable()
	info.AddField("Username", "username", db.Varchar)
	info.AddField("First_name", "first_name", db.Varchar)
	info.AddField("Last_name", "last_name", db.Varchar)
	info.AddField("Email", "email", db.Varchar)
	info.AddField("Role", "role", db.Varchar)
	info.AddField("Password_hash", "password_hash", db.Text)
	info.AddField("Age", "age", db.Int2)
	info.AddField("Dob", "dob", db.Date)
	info.AddField("Sex", "sex", db.Int8)
	info.AddField("Country", "country", db.Varchar)
	info.AddField("City", "city", db.Varchar)
	info.AddField("Points", "points", db.Int2)
	info.AddField("Avatar_url", "avatar_url", db.Varchar)
	info.AddField("Google_sub", "google_sub", db.Varchar)
	info.AddField("Fb_id", "fb_id", db.Varchar)
	info.AddField("Status", "status", db.Int2)

	info.SetTable(tableName).SetTitle("Users").SetDescription("Users")

	formList := users.GetForm()
	formList.AddField("Created_at", "created_at", db.Timestamptz, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamptz, form.Datetime)
	formList.AddField("Id", "id", db.Text, form.Default)
	formList.AddField("Username", "username", db.Varchar, form.Text)
	formList.AddField("First_name", "first_name", db.Varchar, form.Text)
	formList.AddField("Last_name", "last_name", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Role", "role", db.Varchar, form.Text)
	formList.AddField("Password_hash", "password_hash", db.Text, form.RichText)
	formList.AddField("Age", "age", db.Int2, form.Text)
	formList.AddField("Dob", "dob", db.Date, form.Datetime)
	formList.AddField("Sex", "sex", db.Int8, form.Text)
	formList.AddField("Country", "country", db.Varchar, form.Text)
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("Points", "points", db.Int2, form.Text)
	formList.AddField("Avatar_url", "avatar_url", db.Varchar, form.Text)
	formList.AddField("Google_sub", "google_sub", db.Varchar, form.Text)
	formList.AddField("Fb_id", "fb_id", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Int2, form.Text)

	formList.SetTable(tableName).SetTitle("Users").SetDescription("Users")

	return users
}
