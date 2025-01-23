package tables

import (
	"github.com/huyrun/admin_dashboard/src/utils"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types/form"
	"gorm.io/gorm"
)

type Category struct {
	db   *gorm.DB
	conn db.Connection
}

func NewCategory(db *gorm.DB, conn db.Connection) (*Category, error) {
	return &Category{
		db:   db,
		conn: conn,
	}, nil
}

func (t *Category) GetCategoriesTable(ctx *context.Context) table.Table {
	categories := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "categories"
	info := categories.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("Name", "name", db.Varchar).FieldSortable().FieldFilterable()

	info.SetTable(tableName).SetTitle("Categories").SetDescription("Categories").AddCSS(utils.CssTableNoWrap)

	formList := categories.GetForm()
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Name", "name", db.Varchar, form.Text)

	formList.SetTable(tableName).SetTitle("Categories").SetDescription("Categories").AddCSS(utils.CssTableNoWrap)

	return categories
}

func (t *Category) getByID(id string) (map[string]interface{}, error) {
	query := `select id, name
from categories
where id = ?
order by id desc
limit 1;`
	res, err := t.conn.Query(query, id)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res[0], nil
}
