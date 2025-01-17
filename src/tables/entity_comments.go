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
	"github.com/oklog/ulid/v2"
)

type EntityCommentsTable struct {
	user   *User
	entity *Entity
}

func NewEntityComments(user *User, entity *Entity) (*EntityCommentsTable, error) {
	return &EntityCommentsTable{
		user:   user,
		entity: entity,
	}, nil
}

func (t *EntityCommentsTable) GetEntityCommentsTable(ctx *context.Context) table.Table {
	tableConfig := table.DefaultConfigWithDriver("postgresql")
	tableConfig.PrimaryKey = table.PrimaryKey{
		Type: db.Int8,
		Name: "comment_no",
	}
	entityComments := table.NewDefaultTable(ctx, tableConfig)
	tableName := "entity_comments"
	info := entityComments.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("Comment No", "comment_no", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("Entity ID", "entity_id", db.Int8).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return linkToOtherTable("entities", value.Value)
		})
	info.AddField("User ID", "user_id", db.UUID).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			var id ulid.ULID
			err := id.UnmarshalBinary([]byte(value.Value))
			if err != nil {
				return linkToOtherTable("users", value.Value)
			}
			return linkToOtherTable("users", id.String())
		})

	info.AddField("Comment", "comment", db.Varchar).FieldFilterable()
	info.AddField("Created At", "created_at", db.Timestamptz).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02 15:04:05")
		})
	info.AddField("Updated At", "updated_at", db.Timestamptz).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02 15:04:05")
		})

	info.SetTable(tableName).SetTitle("EntityComments").SetDescription("Entity Comments").AddCSS(cssTableNoWrap)

	formList := entityComments.GetForm()
	formList.SetPostValidator(t.postValidator)
	formList.SetPreProcessFn(t.preProcess)
	formList.AddField("Comment No", "comment_no", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("User ID", "user_id", db.Text, form.Text)
	formList.AddField("Comment", "comment", db.Varchar, form.Text)

	formList.SetTable(tableName).SetTitle("EntityComments").SetDescription("Entity Comments").AddCSS(cssTableNoWrap)

	return entityComments
}

func (t *EntityCommentsTable) preProcess(values form2.Values) form2.Values {
	switch values.Get(form2.PostTypeKey) {
	case "0": // update
		values.Add("updated_at", time.Now().Format(time.RFC3339))
	case "1": // create
		values.Add("created_at", time.Now().Format(time.RFC3339))
		values.Add("updated_at", time.Now().Format(time.RFC3339))
	}
	return values
}

func (t *EntityCommentsTable) postValidator(values form2.Values) error {
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
		return errors.New("user id is required")
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
