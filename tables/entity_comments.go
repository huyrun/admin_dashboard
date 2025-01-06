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

type EntityCommentsTable struct {
	db   *gorm.DB
	conn db.Connection
}

func NewEntityCommentsTable(db *gorm.DB, conn db.Connection) (*EntityCommentsTable, error) {
	return &EntityCommentsTable{
		db:   db,
		conn: conn,
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
			return linkToOtherTable("entities", value.Value)
		})
	info.AddField("User ID", "user_id", db.UUID).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			valueByte := []byte(value.Value)
			u, err := uuid.FromBytes(valueByte)
			if err != nil {
				return linkToOtherTable("users", value.Value)
			}
			return linkToOtherTable("users", u.String())
		})

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

	info.SetTable(tableName).SetTitle("EntityComments").SetDescription("Entity Comments").AddCSS(cssTableNoWrap)

	formList := entityComments.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetPreProcessFn(t.preProcess)
	formList.AddField("Comment No", "comment_no", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("User ID", "user_id", db.Text, form.Text)
	formList.AddField("Comment", "comment", db.Varchar, form.Text)

	formList.SetTable(tableName).SetTitle("EntityComments").SetDescription("Entity Comments").AddCSS(cssTableNoWrap)

	entityComments.GetDetailFromInfo().SetTable(tableName).SetTitle("EntityComments").
		SetDescription("Entity Comments").SetGetDataFn(t.getDataDetail)

	return entityComments
}

func (t *EntityCommentsTable) queryFilterFn(param parameter.Parameters, _ db.Connection) (ids []string, stopQuery bool) {
	id := param.GetFieldValue("id")
	u, err := uuid.Parse(id)
	if err != nil {
		return []string{}, false
	}
	uBytes := u[:]
	return []string{string(uBytes)}, true
}

func (t *EntityCommentsTable) preProcess(values form2.Values) form2.Values {
	switch values.Get(form2.PostTypeKey) {
	case "0": // update
		values.Add("updated_at", time.Now().Format(time.RFC3339))
	case "1": // create
		values.Add("created_at", time.Now().Format(time.RFC3339))
		values.Add("updated_at", time.Now().Format(time.RFC3339))
	}
	return values
}

func (t *EntityCommentsTable) insert(values form2.Values) error {
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

	if err := t.db.Table("entity_comments").Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *EntityCommentsTable) getDataDetail(param parameter.Parameters) ([]map[string]interface{}, int) {
	commentNo := param.GetFieldValue(parameter.PrimaryKey)
	query := `select comment_no, entity_id, encode(user_id, 'hex')::uuid as user_id, comment, created_at, updated_at
from entity_comments
where comment_no = ?
order by comment_no desc
limit 1;`
	res, err := t.conn.Query(query, commentNo)
	if err != nil {
		return []map[string]interface{}{}, 0
	}
	return res, 0
}
