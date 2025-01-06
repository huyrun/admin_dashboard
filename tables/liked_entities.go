package tables

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/parameter"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
	"gorm.io/gorm"
	"net/url"
	"regexp"
)

type LikedEntitiesTable struct {
	db   *gorm.DB
	conn db.Connection
}

func NewLikedEntitiesTable(db *gorm.DB, conn db.Connection) (*LikedEntitiesTable, error) {
	return &LikedEntitiesTable{
		db:   db,
		conn: conn,
	}, nil
}

func (t *LikedEntitiesTable) GetLikedEntitiesTable(ctx *context.Context) table.Table {
	tableConfig := table.DefaultConfigWithDriver("postgresql")
	tableConfig.PrimaryKey = table.PrimaryKey{
		Type: db.UUID,
		Name: "user_id",
	}
	likedEntities := table.NewDefaultTable(ctx, tableConfig)
	tableName := "liked_entities"
	info := likedEntities.GetInfo().SetFilterFormLayout(form.LayoutFilter).SetGetDataFn(t.getData)
	info.SetDeleteFn(t.deleteFn(ctx))
	info.AddField("Entity ID", "entity_id", db.Int8).FieldAsDetailParam().FieldAsEditParam().FieldAsDeleteParam().FieldSortable().FieldFilterable()
	info.AddField("User ID", "user_id", db.UUID).FieldAsDetailParam().FieldAsEditParam().FieldAsDeleteParam().FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			valueByte := []byte(value.Value)
			u, err := uuid.FromBytes(valueByte)
			if err != nil {
				return linkToOtherTable("users", value.Value)
			}
			return linkToOtherTable("users", u.String())
		})
	info.AddField("Amount", "amount", db.Int2).FieldSortable()

	info.SetTable(tableName).SetTitle("LikedEntities").SetDescription("Liked Entities").AddCSS(cssTableNoWrap)

	formList := likedEntities.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetUpdateFn(t.update)
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("User ID", "user_id", db.Text, form.Text)
	formList.AddField("Amount", "amount", db.Int2, form.Number).FieldDefault("10")

	formList.SetTable(tableName).SetTitle("LikedEntities").SetDescription("Liked Entities")

	likedEntities.GetDetailFromInfo().SetTable(tableName).SetTitle("LikedEntities").
		SetDescription("Liked Entities").SetGetDataFn(t.getDataDetail)

	return likedEntities
}

func (t *LikedEntitiesTable) insert(values form2.Values) error {
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

	if err := t.db.Table("liked_entities").Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *LikedEntitiesTable) getDataDetail(param parameter.Parameters) ([]map[string]interface{}, int) {
	var keyPrefix string
	if ok, err := regexp.MatchString(`\/.*\/info\/.*\/edit`, param.URLPath); err == nil && ok {
		keyPrefix = "__goadmin_edit"
	} else if ok, err = regexp.MatchString(`\/.*\/info\/.*\/detail`, param.URLPath); err == nil && ok {
		keyPrefix = "__goadmin_detail"
	}
	userID := param.GetFieldValue(fmt.Sprintf("%s_user_id", keyPrefix))
	u, err := uuid.Parse(userID)
	if err != nil {
		return []map[string]interface{}{}, 0
	}
	entityID := param.GetFieldValue(fmt.Sprintf("%s_entity_id", keyPrefix))
	query := `select entity_id, encode(user_id, 'hex')::uuid as user_id, amount
from liked_entities
where entity_id = ?
and user_id = decode(?, 'hex')
order by user_id desc, entity_id desc
limit 1;`
	res, err := t.conn.Query(query, entityID, hex.EncodeToString(u[:]))
	if err != nil {
		return []map[string]interface{}{}, 0
	}
	return res, len(res)
}

func (t *LikedEntitiesTable) getData(param parameter.Parameters) ([]map[string]interface{}, int) {
	query := `select entity_id, encode(user_id, 'hex')::uuid as user_id, amount
from liked_entities
order by user_id desc, entity_id desc
offset ? limit ?;`
	res, err := t.conn.Query(query, (param.PageInt-1)*param.PageSizeInt, param.PageInt*param.PageSizeInt)
	if err != nil {
		return []map[string]interface{}{}, 0
	}
	return res, len(res)
}

func (t *LikedEntitiesTable) update(values form2.Values) error {
	var m = make(map[string]interface{})
	var previousUserID, previousEntityID string
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
		if k == form2.PreviousKey {
			link := values.Get(k)
			parsedURL, err := url.Parse(link)
			if err != nil {
				return err
			}
			urlQuery := parsedURL.Query()
			previousUserID = urlQuery.Get("__goadmin_edit_user_id")
			previousEntityID = urlQuery.Get("__goadmin_edit_entity_id")
			continue
		}
		if k != form2.TokenKey && len(v) > 0 {
			m[k] = v
			continue
		}
	}

	u, err := uuid.Parse(previousUserID)
	if err != nil {
		return err
	}

	if err := t.db.Table("liked_entities").
		Where("user_id = decode(?, 'hex') and entity_id = ?", hex.EncodeToString(u[:]), previousEntityID).
		Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *LikedEntitiesTable) queryFilterFn(param parameter.Parameters, _ db.Connection) (ids []string, stopQuery bool) {
	id := param.GetFieldValue("id")
	u, err := uuid.Parse(id)
	if err != nil {
		return []string{}, false
	}
	uBytes := u[:]
	return []string{string(uBytes)}, true
}

func (t *LikedEntitiesTable) deleteFn(ctx *context.Context) types.DeleteFn {
	return func(ids []string) error {
		if ok, err := regexp.MatchString(`\/.*\/delete\/.*`, ctx.Request.URL.Path); err != nil && !ok {
			return nil
		}

		var userID, entityID string
		userID = ctx.Query("__goadmin_delete_user_id")
		if userID == "" {
			userID = ctx.Query("__goadmin_detail_user_id")
		}

		entityID = ctx.Query("__goadmin_delete_entity_id")
		if entityID == "" {
			entityID = ctx.Query("__goadmin_detail_entity_id")
		}

		u, err := uuid.Parse(userID)
		if err != nil {
			return err
		}

		result := t.db.Table("liked_entities").
			Where("user_id = decode(?, 'hex') and entity_id = ?", hex.EncodeToString(u[:]), entityID).
			Delete(nil)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("no record found to delete")
		}

		return nil
	}
}
