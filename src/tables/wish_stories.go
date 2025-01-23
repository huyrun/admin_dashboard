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
	"github.com/huyrun/go-admin/template"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type WishStory struct {
	db       *gorm.DB
	conn     db.Connection
	entity   *Entity
	statuses utils.StatusMap
}

func NewWishStory(entity *Entity, db *gorm.DB, conn db.Connection) (*WishStory, error) {
	statuses := make(utils.StatusMap)
	statuses.Set("published", "Published", utils.SaffronYellow, utils.NavyBlue)

	return &WishStory{
		db:       db,
		conn:     conn,
		entity:   entity,
		statuses: statuses,
	}, nil
}

func (t *WishStory) GetWishStoryTable(ctx *context.Context) table.Table {
	tableConfig := table.DefaultConfigWithDriver("postgresql")
	tableConfig.PrimaryKey = table.PrimaryKey{
		Type: db.Int8,
		Name: "entity_id",
	}
	wishStories := table.NewDefaultTable(ctx, tableConfig)
	tableName := "wish_stories"
	info := wishStories.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("Entity ID", "entity_id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("Body", "body", db.Text)
	info.AddField("Status", "status", db.Text).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldSortable().
		FieldFilterOptions(t.statuses.ToFieldOptions()).FieldDisplay(t.statuses.ToFieldDisplay)
	info.AddField("Image", "image", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "" {
				return value.Value
			}
			return template.Default().Image().WithModal().SetSrc(template.HTML(value.Value)).GetContent()
		})
	info.AddField("Created At", "created_at", db.Timestamptz).FieldSortable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02 15:04:05")
		})
	info.AddField("UpdatedAt", "updated_at", db.Timestamptz).FieldSortable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02 15:04:05")
		})

	info.SetTable(tableName).SetTitle("WishStories").SetDescription("WishS tories").AddCSS(utils.CssTableNoWrap)

	formList := wishStories.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetUpdateFn(t.update)
	formList.SetPostValidator(t.postValidator)
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("Body", "body", db.Varchar, form.TextArea)
	formList.AddField("Image", "image", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Tinyint, form.Switch).FieldOptions(t.statuses.ToFieldOptions())

	formList.SetTable(tableName).SetTitle("WishStories").SetDescription("Wish Stories")

	return wishStories
}

func (t *WishStory) postValidator(values form2.Values) error {
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

func (t *WishStory) update(values form2.Values) error {
	updateFields := []string{
		"entity_id", "body", "image", "status", "updated_at",
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
	m["updated_at"] = time.Now()
	if err = t.db.Table("wishes").Where("id = ?", id).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *WishStory) insert(values form2.Values) error {
	insertFields := []string{
		"entity_id", "body", "image", "status", "created_at", "updated_at",
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
	if err = t.db.Table("wishes").Create(m).Error; err != nil {
		return err
	}
	return nil
}
