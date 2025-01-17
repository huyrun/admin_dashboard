package tables

import (
	"fmt"
	"time"

	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/color"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
)

func GetGooseDbVersionTable(ctx *context.Context) table.Table {
	gooseDbVersion := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "goose_db_version"
	info := gooseDbVersion.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("VersionID", "version_id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("IsApplied", "is_applied", db.Bool).
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldSortable().FieldFilterOptions(types.FieldOptions{
		{Value: "true", Text: "Yes"},
		{Value: "false", Text: "No"},
	}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "false" {
				return fmt.Sprintf(`<span class="label" style="background-color: %s; color: %s;">No</span>`, color.Gray, color.White)
			}
			if value.Value == "true" {
				return fmt.Sprintf(`<span class="label" style="background-color: %s;  color: %s;">Yes</span>`, color.RoyalBlue, color.White)
			}
			return fmt.Sprintf(`<span class="label" style="text-decoration: line-through; background-color: %s; color: %s;">Unknown</span>`, color.Red, color.Black)
		})
	info.AddField("TimeStamp", "tstamp", db.Timestamp).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02 15:04:05")
		})

	info.SetTable(tableName).SetTitle("GooseDbVersion").SetDescription("Goose DB Version").AddCSS(cssTableNoWrap)

	formList := gooseDbVersion.GetForm()
	formList.AddField("ID", "id", db.Int8, form.Default).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("VersionID", "version_id", db.Int8, form.Text)
	formList.AddField("IsApplied", "is_applied", db.Bool, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "Yes", Value: "true"},
			{Text: "No", Value: "false"},
		}).FieldDefault("true")
	formList.SetPreProcessFn(func(values form2.Values) form2.Values {
		values.Add("TimeStamp", time.Now().Format(time.RFC3339))
		return values
	})

	formList.SetTable(tableName).SetTitle("GooseDbVersion").SetDescription("Goose DB Version")

	return gooseDbVersion
}
