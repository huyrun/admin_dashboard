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

type EntityCommentsTable struct {
	db     *gorm.DB
	conn   db.Connection
	user   *User
	entity *Entity
}

func NewEntityComments(user *User, entity *Entity, db *gorm.DB, conn db.Connection) (*EntityCommentsTable, error) {
	return &EntityCommentsTable{
		db:     db,
		conn:   conn,
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
			return utils.LinkToOtherTable("entities", value.Value)
		})
	info.AddField("User ID", "user_id", db.UUID).FieldSortable().FieldFilterable().FieldDisplay(utils.ParseUserIDToLink)

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

	info.SetTable(tableName).SetTitle("EntityComments").SetDescription("Entity Comments").AddCSS(utils.CssTableNoWrap)

	formList := entityComments.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetUpdateFn(t.update)
	formList.SetPostValidator(t.postValidator)
	formList.AddField("Comment No", "comment_no", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("User ID", "user_id", db.Text, form.Text).FieldDisplay(utils.ParseUserID)
	formList.AddField("Comment", "comment", db.Varchar, form.TextArea)

	formList.SetTable(tableName).SetTitle("EntityComments").SetDescription("Entity Comments").AddCSS(utils.CssTableNoWrap)

	return entityComments
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

func (t *EntityCommentsTable) update(values form2.Values) error {
	updateFields := []string{
		"entity_id", "comment_no", "comment", "updated_at",
	}

	m := make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if utils2.InArray(updateFields, k) && len(v) > 0 {
			m[k] = v
		}
	}

	commentNo := values.Get("comment_no")
	userID := values.Get("user_id")
	ulidValue, err := ulid.Parse(userID)
	if err != nil {
		return err
	}
	m["user_id"] = ulidValue
	m["updated_at"] = time.Now()
	if err = t.db.Table("entity_comments").Where("comment_no = ?", commentNo).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *EntityCommentsTable) insert(values form2.Values) error {
	insertFields := []string{
		"entity_id", "comment_no", "comment", "created_at", "updated_at",
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
	timeNow := time.Now()
	m["created_at"] = timeNow
	m["updated_at"] = timeNow
	if err = t.db.Table("entity_comments").Create(m).Error; err != nil {
		return err
	}
	return nil
}
