package tables

import (
	"errors"
	"fmt"

	"github.com/huyrun/admin_dashboard/src/utils"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
)

type EntityTag struct {
	entity *Entity
	tag    *Tag
}

func NewEntityTags(entity *Entity, tag *Tag) (*EntityTag, error) {
	return &EntityTag{
		entity: entity,
		tag:    tag,
	}, nil
}

func (t *EntityTag) GetEntityTagTable(ctx *context.Context) table.Table {
	entityTags := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "entity_tags"
	info := entityTags.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldFilterable().FieldSortable()
	info.AddField("Entity ID", "entity_id", db.Int8).FieldFilterable().FieldSortable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return utils.LinkToOtherTable("entities", value.Value)
		})
	info.AddField("Tag ID", "tag_id", db.Text).FieldFilterable().FieldSortable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return utils.LinkToOtherTable("tags", value.Value)
		})

	info.SetTable(tableName).SetTitle("EntityTags").SetDescription("Entity Tags").AddCSS(utils.CssTableNoWrap)

	formList := entityTags.GetForm()
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("Tag ID", "tag_id", db.Text, form.Text)

	formList.SetTable(tableName).SetTitle("EntityTags").SetDescription("Entity Tags")

	return entityTags
}

func (t *EntityTag) postValidator(values form2.Values) error {
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

	tagID := values.Get("tag_id")
	if tagID == "" {
		return errors.New("tag id is required")
	}
	tag, err := t.tag.getByTagName(tagID)
	if err != nil {
		return err
	}
	if tag == nil {
		return fmt.Errorf("not found tag %s", tagID)
	}

	return nil
}
