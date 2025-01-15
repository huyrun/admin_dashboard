package tables

import (
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types/form"
	"gorm.io/gorm"
)

type Tag struct {
	db   *gorm.DB
	conn db.Connection
}

func NewTag(db *gorm.DB, conn db.Connection) (*Tag, error) {
	return &Tag{
		db:   db,
		conn: conn,
	}, nil
}

func (t *Tag) GetTagsTable(ctx *context.Context) table.Table {
	tags := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "tags"
	info := tags.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("Tag Name", "tag_name", db.Text).FieldSortable().FieldFilterable()

	info.SetTable(tableName).SetTitle("Tags").SetDescription("Tags").AddCSS(cssTableNoWrap)

	formList := tags.GetForm()
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Tag Name", "tag_name", db.Text, form.RichText)

	formList.SetTable(tableName).SetTitle("Tags").SetDescription("Tags")

	return tags
}

func (t *Tag) getByTagName(tagName string) (map[string]interface{}, error) {
	query := `select id
from tags
where tag_name = ?
order by id desc
limit 1;`
	res, err := t.conn.Query(query, tagName)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res[0], nil
}
