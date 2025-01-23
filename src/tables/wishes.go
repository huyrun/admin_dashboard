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

type Wish struct {
	db       *gorm.DB
	conn     db.Connection
	user     *User
	entity   *Entity
	category *Category
	statuses utils.StatusMap
}

func NewWish(user *User, entity *Entity, category *Category, db *gorm.DB, conn db.Connection) (*Wish, error) {
	statuses := make(utils.StatusMap)
	statuses.Set("new", "New", utils.BrightBlue, utils.White)
	statuses.Set("completed", "Completed", utils.SaffronYellow, utils.NavyBlue)

	return &Wish{
		db:       db,
		conn:     conn,
		user:     user,
		entity:   entity,
		category: category,
		statuses: statuses,
	}, nil
}

func (t *Wish) GetWishTable(ctx *context.Context) table.Table {
	wishes := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "wishes"
	info := wishes.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("User ID", "user_id", db.UUID).FieldSortable().FieldFilterable().FieldDisplay(utils.ParseUserIDToLink)
	info.AddField("Entity ID", "entity_id", db.Int8).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return utils.LinkToOtherTable("entities", value.Value)
		})
	info.AddField("Type", "type", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Title", "title", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Story", "story", db.Text)
	info.AddField("Price", "price", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("Currency", "currency", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Category ID", "category_id", db.Int8).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return utils.LinkToOtherTable("categories", value.Value)
		})
	info.AddField("Visible By", "visible_by", db.Int8)
	info.AddField("Image", "image", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return template.Default().Image().WithModal().SetSrc(template.HTML(value.Value)).GetContent()
		})
	info.AddField("Status", "status", db.Text).FieldSortable().FieldFilterable(types.FilterType{FormType: form.SelectSingle}).
		FieldFilterOptions(t.statuses.ToFieldOptions()).FieldDisplay(t.statuses.ToFieldDisplay)
	info.AddField("Description", "description", db.Text)
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

	info.SetTable(tableName).SetTitle("Wishes").SetDescription("Wishes").AddCSS(utils.CssTableNoWrap)

	formList := wishes.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetUpdateFn(t.update)
	formList.SetPostValidator(t.postValidator)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("User ID", "user_id", db.Text, form.Text).FieldDisplay(utils.ParseUserID)
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("Type", "type", db.Varchar, form.Text)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.TextArea)
	formList.AddField("Story", "story", db.Varchar, form.TextArea)
	formList.AddField("Price", "price", db.Int8, form.Number).FieldDisplay(utils.CastToNumber)
	formList.AddField("Currency", "currency", db.Varchar, form.Text)
	formList.AddField("Category ID", "category_id", db.Int8, form.Text)
	formList.AddField("Visible By", "visible_by", db.Int8, form.Number).FieldDisplay(utils.CastToNumber)
	formList.AddField("Image", "image", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Tinyint, form.SelectSingle).FieldDefault("new").FieldOptions(t.statuses.ToFieldOptions())

	formList.SetTable(tableName).SetTitle("Wishes").SetDescription("Wishes")

	return wishes
}

func (t *Wish) postValidator(values form2.Values) error {
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

	categoryID := values.Get("category_id")
	if categoryID == "" {
		return errors.New("category id is required")
	}
	category, err := t.category.getByID(categoryID)
	if err != nil {
		return err
	}
	if category == nil {
		return fmt.Errorf("not found category %s", categoryID)
	}

	return nil
}

func (t *Wish) update(values form2.Values) error {
	updateFields := []string{
		"entity_id", "type", "title", "description", "story", "currency", "category_id", "visible_by", "image", "status", "updated_at",
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

func (t *Wish) insert(values form2.Values) error {
	insertFields := []string{
		"entity_id", "type", "title", "description", "story", "currency", "category_id", "visible_by", "image", "status", "created_at", "updated_at",
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
