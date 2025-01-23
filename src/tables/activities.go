package tables

import (
	"errors"
	"fmt"
	"time"

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

type Activity struct {
	db     *gorm.DB
	conn   db.Connection
	user   *User
	entity *Entity
}

func NewActivity(user *User, entity *Entity, db *gorm.DB, conn db.Connection) (*Activity, error) {
	return &Activity{
		db:     db,
		conn:   conn,
		user:   user,
		entity: entity,
	}, nil
}

func (t *Activity) GetActivitiesTable(ctx *context.Context) table.Table {
	activities := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "activities"
	info := activities.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("Action", "action", db.Varchar)
	info.AddField("Points", "points", db.Int).FieldSortable()
	info.AddField("User ID", "user_id", db.UUID).FieldSortable().FieldFilterable().FieldDisplay(utils.ParseUserIDToLink)
	info.AddField("Entity ID", "entity_id", db.Int8).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return utils.LinkToOtherTable("entities", value.Value)
		})
	info.AddField("Created At", "created_at", db.Timestamptz).FieldSortable().FieldSortable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02 15:04:05")
		})

	info.SetTable(tableName).SetTitle("Activities").SetDescription("Activities").AddCSS(utils.CssTableNoWrap)

	formList := activities.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetUpdateFn(t.update)
	formList.SetPostValidator(t.postValidator)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Action", "action", db.Varchar, form.Text)
	formList.AddField("Points", "points", db.Int, form.Number).FieldDefault("0")
	formList.AddField("User ID", "user_id", db.UUID, form.Text).FieldDisplay(utils.ParseUserID)
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)

	formList.SetTable(tableName).SetTitle("Activities").SetDescription("Activities").AddCSS(utils.CssTableNoWrap)

	return activities
}

func (t *Activity) postValidator(values form2.Values) error {
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

	entityID := values.Get("entity_id")
	if entityID == "" {
		return errors.New("entity id is required")
	}
	entity, err := t.entity.getByID(entityID)
	if err != nil {
		return err
	}
	if entity == nil {
		return fmt.Errorf("not found entity %s", entityID)
	}

	return nil
}

func (t *Activity) update(values form2.Values) error {
	updateFields := []string{
		"action", "points", "entity_id",
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
	if err = t.db.Table("activities").Where("id = ?", id).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *Activity) insert(values form2.Values) error {
	insertFields := []string{
		"action", "points", "entity_id",
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
	m["created_at"] = time.Now()
	if err = t.db.Table("activities").Create(m).Error; err != nil {
		return err
	}
	return nil
}
