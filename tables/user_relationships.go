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

type UserRelationshipsTable struct {
	db   *gorm.DB
	conn db.Connection
}

func NewUserRelationshipsTable(db *gorm.DB, conn db.Connection) (*UserRelationshipsTable, error) {
	return &UserRelationshipsTable{
		db:   db,
		conn: conn,
	}, nil
}

func (t *UserRelationshipsTable) GetUserRelationshipsTable(ctx *context.Context) table.Table {
	userRelationships := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "user_relationships"
	info := userRelationships.GetInfo().HideFilterArea()

	info.AddField("First User ID", "first_user_id", db.Text).FieldAsDetailParam().FieldAsEditParam().FieldAsDeleteParam().FieldSortable().FieldFilterable()
	info.AddField("Second User ID", "second_user_id", db.Text).FieldAsDetailParam().FieldAsEditParam().FieldAsDeleteParam().FieldSortable().FieldFilterable()
	info.AddField("First To Second Status", "first_to_second_status", db.Text)
	info.AddField("Second To First Status", "second_to_first_status", db.Text)
	info.AddField("Are Friends", "are_friends", db.Bool).FieldFilterable()

	info.SetTable(tableName).SetTitle("UserRelationships").SetDescription("User Relationships").AddCSS(cssTableNoWrap)

	formList := userRelationships.GetForm()
	formList.AddField("First_user_id", "first_user_id", db.Text, form.Text)
	formList.AddField("Second_user_id", "second_user_id", db.Text, form.Text)
	formList.AddField("First_to_second_status", "first_to_second_status", db.Text, form.RichText)
	formList.AddField("Second_to_first_status", "second_to_first_status", db.Text, form.RichText)
	formList.AddField("Are_friends", "are_friends", db.Bool, form.Text)

	formList.SetTable(tableName).SetTitle("UserRelationships").SetDescription("User Relationships")

	return userRelationships
}

func (t *UserRelationshipsTable) insert(values form2.Values) error {
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

func (t *UserRelationshipsTable) getDataDetail(param parameter.Parameters) ([]map[string]interface{}, int) {
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

func (t *UserRelationshipsTable) getData(param parameter.Parameters) ([]map[string]interface{}, int) {
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

func (t *UserRelationshipsTable) update(values form2.Values) error {
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

func (t *UserRelationshipsTable) queryFilterFn(param parameter.Parameters, _ db.Connection) (ids []string, stopQuery bool) {
	id := param.GetFieldValue("id")
	u, err := uuid.Parse(id)
	if err != nil {
		return []string{}, false
	}
	uBytes := u[:]
	return []string{string(uBytes)}, true
}

func (t *UserRelationshipsTable) deleteFn(ctx *context.Context) types.DeleteFn {
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
