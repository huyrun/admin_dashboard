package tables

import (
	"errors"
	"fmt"
	"time"

	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template"
	"github.com/huyrun/go-admin/template/color"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
	"github.com/oklog/ulid/v2"
)

type Wish struct {
	user     *User
	entity   *Entity
	category *Category
}

func NewWish(user *User, entity *Entity, category *Category) (*Wish, error) {
	return &Wish{
		user:     user,
		entity:   entity,
		category: category,
	}, nil
}

func (t *Wish) GetWishTable(ctx *context.Context) table.Table {
	wishes := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "wishes"
	info := wishes.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("User ID", "user_id", db.UUID).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			var id ulid.ULID
			err := id.UnmarshalBinary([]byte(value.Value))
			if err != nil {
				return linkToOtherTable("users", value.Value)
			}
			return linkToOtherTable("users", id.String())
		})
	info.AddField("Entity ID", "entity_id", db.Int8).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return linkToOtherTable("entities", value.Value)
		})
	info.AddField("Type", "type", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Title", "title", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Story", "story", db.Text)
	info.AddField("Price", "price", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("Currency", "currency", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Category ID", "category_id", db.Int8).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return linkToOtherTable("categories", value.Value)
		})
	info.AddField("Visible By", "visible_by", db.Int8)
	info.AddField("Image", "image", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return template.Default().Image().WithModal().SetSrc(template.HTML(value.Value)).GetContent()
		})
	info.AddField("Status", "status", db.Text).FieldSortable().FieldFilterable().
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldSortable().FieldFilterOptions(types.FieldOptions{
		{Value: "new", Text: "New"},
		{Value: "deactivated", Text: "Deactivated"},
	}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "new" {
				return fmt.Sprintf(`<span class="label" style="background-color: %s; color: %s;">Deactivated</span>`, color.Gray, color.White)
			}
			if value.Value == "deactivated" {
				return fmt.Sprintf(`<span class="label" style="background-color: %s;  color: %s;">New</span>`, color.Yellow, color.Black)
			}
			return fmt.Sprintf(`<span class="label" style="text-decoration: line-through; background-color: %s; color: %s;">Unknown</span>`, color.Red, color.Black)
		})
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

	info.SetTable(tableName).SetTitle("Wishes").SetDescription("Wishes").AddCSS(cssTableNoWrap)

	formList := wishes.GetForm()
	formList.SetPostValidator(t.postValidator)
	formList.SetPreProcessFn(t.preProcess)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("User ID", "user_id", db.Text, form.Text)
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("Type", "type", db.Varchar, form.Text)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.RichText)
	formList.AddField("Story", "story", db.Varchar, form.RichText)
	formList.AddField("Price", "price", db.Int8, form.Number).FieldDefault("0")
	formList.AddField("Currency", "currency", db.Varchar, form.Number).FieldDefault("0")
	formList.AddField("Category ID", "category_id", db.Int8, form.Text)
	formList.AddField("Visible By", "visible_by", db.Int8, form.Number)
	formList.AddField("Image", "image", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Tinyint, form.SelectSingle).FieldDefault("new").
		FieldOptions(types.FieldOptions{
			{Value: "new", Text: "New"},
			{Value: "deactivated", Text: "Deactivated"},
		})

	formList.SetTable(tableName).SetTitle("Wishes").SetDescription("Wishes")

	return wishes
}

func (t *Wish) preProcess(values form2.Values) form2.Values {
	switch values.Get(form2.PostTypeKey) {
	case "0": // update
		values.Add("updated_at", time.Now().Format(time.RFC3339))
	case "1": // create
		values.Add("created_at", time.Now().Format(time.RFC3339))
		values.Add("updated_at", time.Now().Format(time.RFC3339))
	}
	return values
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
