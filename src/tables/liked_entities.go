package tables

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
	"github.com/oklog/ulid/v2"
)

type LikedEntities struct {
	user   *User
	entity *Entity
}

func NewLikedEntities(user *User, entity *Entity) (*LikedEntities, error) {
	return &LikedEntities{
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
			return linkToOtherTable("entities", value.Value)
		})
	info.AddField("User ID", "user_id", db.UUID).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			var id ulid.ULID
			err := id.UnmarshalBinary([]byte(value.Value))
			if err != nil {
				return linkToOtherTable("users", value.Value)
			}
			return linkToOtherTable("users", id.String())
		})
	info.AddField("Amount", "amount", db.Int2).FieldSortable()

	info.SetTable(tableName).SetTitle("LikedEntities").SetDescription("Liked Entities").AddCSS(cssTableNoWrap)

	formList := likedEntities.GetForm()
	formList.SetPostValidator(t.postValidator)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Entity ID", "entity_id", db.Int8, form.Text)
	formList.AddField("User ID", "user_id", db.Text, form.Text)
	formList.AddField("Amount", "amount", db.Int2, form.Number).FieldDefault("0")

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
