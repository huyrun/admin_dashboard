package tables

import (
	"errors"
	"fmt"
	"time"

	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
)

type Activity struct {
	user   *User
	entity *Entity
}

func NewActivity(user *User, entity *Entity) (*Activity, error) {
	return &Activity{
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
	info.AddField("User ID", "user_id", db.UUID).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return linkToOtherTable("users", value.Value)
		})
	info.AddField("Entity ID", "entity_id", db.Int8).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return linkToOtherTable("entities", value.Value)
		})
	info.AddField("Created At", "created_at", db.Timestamptz).FieldSortable().FieldSortable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02 15:04:05")
		})

	info.SetTable(tableName).SetTitle("Activities").SetDescription("Activities").AddCSS(cssTableNoWrap)

	formList := activities.GetForm()
	formList.SetPostValidator(t.postValidator)
	formList.SetPreProcessFn(t.preProcess)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Action", "action", db.Varchar, form.Text)
	formList.AddField("Points", "points", db.Int, form.Number).FieldDefault("0")
	formList.AddField("User ID", "user_id", db.UUID, form.Text)
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)

	formList.SetTable(tableName).SetTitle("Activities").SetDescription("Activities").AddCSS(cssTableNoWrap)

	return activities
}

func (t *Activity) preProcess(values form2.Values) form2.Values {
	switch values.Get(form2.PostTypeKey) {
	case "1": // create
		values.Add("created_at", time.Now().Format(time.RFC3339))
	}
	return values
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
