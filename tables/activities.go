package tables

import (
	"github.com/google/uuid"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/parameter"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	db   *gorm.DB
	conn db.Connection
}

func NewActivityTable(db *gorm.DB, conn db.Connection) (*Activity, error) {
	return &Activity{
		db:   db,
		conn: conn,
	}, nil
}

func (t *Activity) GetActivitiesTable(ctx *context.Context) table.Table {
	activities := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "activities"
	info := activities.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("Action", "action", db.Varchar)
	info.AddField("Points", "points", db.Int).FieldSortable()
	info.AddField("User ID", "user_id", db.UUID).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			valueByte := []byte(value.Value)
			u, err := uuid.FromBytes(valueByte)
			if err != nil {
				return linkToOtherTable("users", value.Value)
			}
			return linkToOtherTable("users", u.String())
		})
	info.AddField("Entity ID", "entity_id", db.Int8).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return linkToOtherTable("entities", value.Value)
		})
	info.AddField("Created At", "created_at", db.Timestamptz).FieldSortable().FieldSortable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02 15:04:05")
		})

	info.SetTable(tableName).SetTitle("Activities").SetDescription("Activities").AddCSS(cssTableNoWrap)

	formList := activities.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetPreProcessFn(t.preProcess)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Action", "action", db.Varchar, form.Text)
	formList.AddField("Points", "points", db.Int, form.Number).FieldDefault("0")
	formList.AddField("User ID", "user_id", db.UUID, form.Text)
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)

	formList.SetTable(tableName).SetTitle("Activities").SetDescription("Activities").AddCSS(cssTableNoWrap)

	activities.GetDetailFromInfo().SetTable(tableName).SetTitle("Activities").
		SetDescription("Activities").SetGetDataFn(t.getDataDetail)

	return activities
}

func (t *Activity) preProcess(values form2.Values) form2.Values {
	switch values.Get(form2.PostTypeKey) {
	case "1": // create
		values.Add("created_at", time.Now().Format(time.RFC3339))
	}
	return values
}

func (t *Activity) insert(values form2.Values) error {
	var m = make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if k == "user_id" {
			u, err := uuid.Parse(v)
			if err != nil {
				return err
			}
			m["user_id"] = u[:]
			continue
		}
		if (k != form2.PreviousKey && k != form2.TokenKey) && len(v) > 0 {
			m[k] = v
			continue
		}
	}

	if err := t.db.Table("activities").Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *Activity) getDataDetail(param parameter.Parameters) ([]map[string]interface{}, int) {
	id := param.GetFieldValue(parameter.PrimaryKey)
	query := `select id, action, points, encode(user_id, 'hex')::uuid as user_id, entity_id, created_at
from activities
where id = ?
order by id desc
limit 1;`
	res, err := t.conn.Query(query, id)
	if err != nil {
		return []map[string]interface{}{}, 0
	}
	return res, 0
}
