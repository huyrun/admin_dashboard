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
)

type WishStory struct {
	entity *Entity
}

func NewWishStory(entity *Entity) (*WishStory, error) {
	return &WishStory{
		entity: entity,
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
	info.AddField("Status", "status", db.Text).
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldSortable().FieldFilterOptions(types.FieldOptions{
		{Value: "completed", Text: "Completed"},
		{Value: "draft", Text: "Draft"},
	}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "draft" {
				return fmt.Sprintf(`<span class="label" style="background-color: %s; color: %s;">Draft</span>`, color.Gray, color.White)
			}
			if value.Value == "completed" {
				return fmt.Sprintf(`<span class="label" style="background-color: %s;  color: %s;">Completed</span>`, color.Yellow, color.Black)
			}
			return fmt.Sprintf(`<span class="label" style="text-decoration: line-through; background-color: %s; color: %s;">Unknown</span>`, color.Red, color.Black)
		})
	info.AddField("Image", "image", db.Varchar).FieldCopyable().
		FieldDisplay(func(value types.FieldModel) interface{} {
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

	info.SetTable(tableName).SetTitle("WishStories").SetDescription("WishS tories").AddCSS(cssTableNoWrap)

	formList := wishStories.GetForm()
	formList.SetPreProcessFn(t.preProcess)
	formList.SetPostValidator(t.postValidator)
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Body", "body", db.Varchar, form.RichText)
	formList.AddField("Image", "image", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Tinyint, form.Switch).FieldDefault("draft").
		FieldOptions(types.FieldOptions{
			{Text: "Completed", Value: "completed"},
			{Text: "Draft", Value: "draft"},
		})

	formList.SetTable(tableName).SetTitle("WishStories").SetDescription("WishS tories")

	return wishStories
}

func (t *WishStory) preProcess(values form2.Values) form2.Values {
	switch values.Get(form2.PostTypeKey) {
	case "0": // update
		values.Add("updated_at", time.Now().Format(time.RFC3339))
	case "1": // create
		values.Add("created_at", time.Now().Format(time.RFC3339))
		values.Add("updated_at", time.Now().Format(time.RFC3339))
	}
	return values
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
