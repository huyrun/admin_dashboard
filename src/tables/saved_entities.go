package tables

import (
	"errors"
	"fmt"

	"github.com/huyrun/admin_dashboard/src/utils"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	utils2 "github.com/huyrun/go-admin/modules/utils"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type SavedEntitiesTable struct {
	db     *gorm.DB
	conn   db.Connection
	user   *User
	entity *Entity
}

func NewSavedEntities(user *User, entity *Entity, db *gorm.DB, conn db.Connection) (*SavedEntitiesTable, error) {
	return &SavedEntitiesTable{
		db:     db,
		conn:   conn,
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
			return utils.LinkToOtherTable("entities", value.Value)
		})
	info.AddField("User ID", "user_id", db.UUID).FieldAsDetailParam().FieldAsEditParam().FieldAsDeleteParam().FieldSortable().FieldFilterable().FieldDisplay(utils.ParseUserIDToLink)
	info.AddField("Type", "type", db.Text).FieldSortable()

	info.SetTable(tableName).SetTitle("SavedEntities").SetDescription("Saved Entities").AddCSS(utils.CssTableNoWrap)

	formList := savedEntities.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetUpdateFn(t.update)
	formList.SetPostValidator(t.postValidator)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("User ID", "user_id", db.Text, form.Text).FieldDisplay(utils.ParseUserID)
	formList.AddField("Type", "type", db.Text, form.Text)

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

func (t *SavedEntitiesTable) update(values form2.Values) error {
	updateFields := []string{
		"entity_id", "type",
	}

	m := make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if utils2.InArray(updateFields, k) && len(v) > 0 {
			m[k] = v
		}
	}

	id := values.Get("id")
	userID := values.Get("user_id")
	ulidValue, err := ulid.Parse(userID)
	if err != nil {
		return err
	}
	m["user_id"] = ulidValue
	if err = t.db.Table("saved_entities").Where("id = ?", id).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *SavedEntitiesTable) insert(values form2.Values) error {
	insertFields := []string{
		"entity_id", "type",
	}
	m := make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if utils2.InArray(insertFields, k) && len(v) > 0 {
			m[k] = v
		}
	}

	userID := values.Get("user_id")
	ulidValue, err := ulid.Parse(userID)
	if err != nil {
		return err
	}
	m["user_id"] = ulidValue
	if err = t.db.Table("saved_entities").Create(m).Error; err != nil {
		return err
	}
	return nil
}
