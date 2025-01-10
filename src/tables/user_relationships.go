package tables

import (
	"errors"
	"fmt"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/color"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
)

type UserRelationships struct {
	user   *User
	entity *Entity
}

func NewUserRelationships(user *User, entity *Entity) (*UserRelationships, error) {
	return &UserRelationships{
		user:   user,
		entity: entity,
	}, nil
}

func (t *UserRelationships) GetUserRelationshipsTable(ctx *context.Context) table.Table {
	userRelationships := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "user_relationships"
	info := userRelationships.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("First User ID", "first_user_id", db.Text).FieldAsDetailParam().FieldAsEditParam().FieldAsDeleteParam().FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return linkToOtherTable("users", value.Value)
		})
	info.AddField("Second User ID", "second_user_id", db.Text).FieldAsDetailParam().FieldAsEditParam().FieldAsDeleteParam().FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return linkToOtherTable("users", value.Value)
		})
	info.AddField("First To Second Status", "first_to_second_status", db.Text).
		FieldDisplay(func(value types.FieldModel) interface{} {
			switch value.Value {
			case "subscribed":
				return "Subscribed"
			case "not_subscribed":
				return "Not Subscribed"
			default:
				return "Unknown"
			}
		})
	info.AddField("Second To First Status", "second_to_first_status", db.Text).
		FieldDisplay(func(value types.FieldModel) interface{} {
			switch value.Value {
			case "subscribed":
				return "Subscribed"
			case "not_subscribed":
				return "Not Subscribed"
			default:
				return "Unknown"
			}
		})
	info.AddField("Are Friends", "are_friends", db.Bool).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			switch value.Value {
			case "true":
				return fmt.Sprintf(`<span class="label" style="background-color: %s;  color: %s;">Yes</span>`, color.Yellow, color.Black)
			case "false":
				return fmt.Sprintf(`<span class="label" style="background-color: %s;  color: %s;">No</span>`, color.SkyBlue, color.Black)
			default:
				return fmt.Sprintf(`<span class="label" style="background-color: %s;  color: %s;">Unknown</span>`, color.Gray, color.White)
			}
		})

	info.SetTable(tableName).SetTitle("UserRelationships").SetDescription("User Relationships").AddCSS(cssTableNoWrap)

	formList := userRelationships.GetForm()
	formList.SetPostValidator(t.postValidator)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("First User ID", "first_user_id", db.Text, form.Text)
	formList.AddField("Second User ID", "second_user_id", db.Text, form.Text)
	formList.AddField("First To Second Status", "first_to_second_status", db.Text, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "Subscribed", Value: "subscribed"},
			{Text: "Not subscribed", Value: "not_subscribed"},
		})
	formList.AddField("Second To First Status", "second_to_first_status", db.Text, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "Subscribed", Value: "subscribed"},
			{Text: "Not subscribed", Value: "not_subscribed"},
		})
	formList.SetTable(tableName).SetTitle("UserRelationships").SetDescription("User Relationships")

	return userRelationships
}

func (t *UserRelationships) postValidator(values form2.Values) error {
	firstUserID := values.Get("first_user_id")
	if firstUserID == "" {
		return errors.New("first user id is required")
	}
	firstUser, err := t.user.getByID(firstUserID)
	if err != nil {
		return err
	}
	if firstUser == nil {
		return fmt.Errorf("not found first user %s", firstUserID)
	}

	secondUserID := values.Get("second_user_id")
	if secondUserID == "" {
		return errors.New("second user id is required")
	}
	secondUser, err := t.user.getByID(secondUserID)
	if err != nil {
		return err
	}
	if secondUser == nil {
		return fmt.Errorf("not found second user %s", secondUserID)
	}

	return nil
}
