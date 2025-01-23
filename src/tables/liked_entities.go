package tables

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/huyrun/admin_dashboard/src/utils"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	utils2 "github.com/huyrun/go-admin/modules/utils"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type LikedEntities struct {
	db     *gorm.DB
	conn   db.Connection
	user   *User
	entity *Entity
}

func NewLikedEntities(user *User, entity *Entity, db *gorm.DB, conn db.Connection) (*LikedEntities, error) {
	return &LikedEntities{
		db:     db,
		conn:   conn,
		user:   user,
		entity: entity,
	}, nil
}

func (t *LikedEntities) GetLikedEntitiesTable(ctx *context.Context) table.Table {
	likedEntities := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "liked_entities"
	info := likedEntities.GetInfo().SetFilterFormLayout(form.LayoutFilter)
	info.AddField("ID", "id", db.Int8).FieldFilterable().FieldSortable()
	info.AddField("Entity ID", "entity_id", db.Int8).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			return utils.LinkToOtherTable("entities", value.Value)
		})
	info.AddField("User ID", "user_id", db.UUID).FieldSortable().FieldFilterable().FieldDisplay(utils.ParseUserIDToLink)
	info.AddField("Amount", "amount", db.Int2).FieldSortable()

	info.SetTable(tableName).SetTitle("LikedEntities").SetDescription("Liked Entities").AddCSS(utils.CssTableNoWrap)

	formList := likedEntities.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetUpdateFn(t.update)
	formList.SetPostValidator(t.postValidator)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("User ID", "user_id", db.Text, form.Text).FieldDisplay(utils.ParseUserID)
	formList.AddField("Amount", "amount", db.Int8, form.Number).FieldDefault("0")

	formList.SetTable(tableName).SetTitle("LikedEntities").SetDescription("Liked Entities")

	return likedEntities
}

func (t *LikedEntities) postValidator(values form2.Values) error {
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
		return errors.New("user id is required")
	}
	entity, err := t.entity.getByID(entityID)
	if err != nil {
		return err
	}
	if entity == nil {
		return fmt.Errorf("not found entity %s", entityID)
	}

	amountStr := values.Get("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}
	if amount < 0 {
		return errors.New("amount must be greater than zero")
	}

	return nil
}

func (t *LikedEntities) update(values form2.Values) error {
	updateFields := []string{
		"entity_id", "amount",
	}

	m := make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if utils2.InArray(updateFields, k) && len(v) > 0 {
			m[k] = v
		}
	}

	id := values.Get("id")
	userID := values.Get("user_id")
	ulidValue, err := ulid.Parse(userID)
	if err != nil {
		return err
	}
	m["user_id"] = ulidValue
	if err = t.db.Table("liked_entities").Where("id = ?", id).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *LikedEntities) insert(values form2.Values) error {
	insertFields := []string{
		"entity_id", "amount",
	}
	m := make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if utils2.InArray(insertFields, k) && len(v) > 0 {
			m[k] = v
		}
	}

	userID := values.Get("user_id")
	ulidValue, err := ulid.Parse(userID)
	if err != nil {
		return err
	}
	m["user_id"] = ulidValue
	if err = t.db.Table("liked_entities").Create(m).Error; err != nil {
		return err
	}
	return nil
}
