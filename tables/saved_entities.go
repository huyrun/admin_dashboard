package tables

import (
	"errors"
	"fmt"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
)

type SavedEntitiesTable struct {
	user   *User
	entity *Entity
}

func NewSavedEntities(user *User, entity *Entity) (*SavedEntitiesTable, error) {
	return &SavedEntitiesTable{
		user:   user,
		entity: entity,
	}, nil
}

func (t *SavedEntitiesTable) GetSavedEntitiesTable(ctx *context.Context) table.Table {
	savedEntities := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "saved_entities"
	info := savedEntities.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldFilterable().FieldSortable()
	info.AddField("Entity ID", "entity_id", db.Int8).FieldAsDetailParam().FieldAsEditParam().FieldAsDeleteParam().FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return linkToOtherTable("entities", value.Value)
		})
	info.AddField("User ID", "user_id", db.UUID).FieldAsDetailParam().FieldAsEditParam().FieldAsDeleteParam().FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return linkToOtherTable("users", value.Value)
		})
	info.AddField("Type", "type", db.Text).FieldSortable()

	info.SetTable(tableName).SetTitle("SavedEntities").SetDescription("Saved Entities").AddCSS(cssTableNoWrap)

	formList := savedEntities.GetForm()
	formList.SetPostValidator(t.postValidator)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("User ID", "user_id", db.Text, form.Text)
	formList.AddField("Type", "type", db.Text, form.Text).FieldDefault("wish")

	formList.SetTable(tableName).SetTitle("SavedEntities").SetDescription("Saved Entities")

	return savedEntities
}

func (t *SavedEntitiesTable) postValidator(values form2.Values) error {
	userID := values.Get("user_id")
	if userID == "" {
		return errors.New("user id is required")
	}
	user, err := t.user.getByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("not found user %s", userID)
	}

	return nil
}
