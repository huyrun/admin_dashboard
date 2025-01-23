package tables

import (
	"github.com/huyrun/admin_dashboard/src/utils"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types/form"
	"gorm.io/gorm"
)

type Entity struct {
	db   *gorm.DB
	conn db.Connection
}

func NewEntity(db *gorm.DB, conn db.Connection) (*Entity, error) {
	return &Entity{
		db:   db,
		conn: conn,
	}, nil
}

func (t *Entity) GetEntitiesTable(ctx *context.Context) table.Table {
	entities := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "entities"
	info := entities.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()

	info.SetTable(tableName).SetTitle("Entities").SetDescription("Entities").AddCSS(utils.CssTableNoWrap)

	formList := entities.GetForm()
	formList.AddField("Id", "id", db.Int8, form.Number).FieldDisplayButCanNotEditWhenUpdate()

	formList.SetTable(tableName).SetTitle("Entities").SetDescription("Entities").AddCSS(utils.CssTableNoWrap)

	return entities
}

func (t *Entity) getByID(id string) (map[string]interface{}, error) {
	query := `select id
from entities
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
