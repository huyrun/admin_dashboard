package tables

import (
	"errors"
	"fmt"

	"github.com/huyrun/admin_dashboard/src/utils"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	utils2 "github.com/huyrun/go-admin/modules/utils"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template/color"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type UserRelationships struct {
	db        *gorm.DB
	conn      db.Connection
	user      *User
	entity    *Entity
	statuses1 utils.StatusMap
	statuses2 utils.StatusMap
}

func NewUserRelationships(user *User, entity *Entity, db *gorm.DB, conn db.Connection) (*UserRelationships, error) {
	statuses1 := make(utils.StatusMap)
	statuses1.Set("subscribed", "Subscribed", utils.SaffronYellow, color.Blue)

	statuses2 := make(utils.StatusMap)
	statuses2.Set("true", "Yes", utils.SaffronYellow, utils.NavyBlue)
	statuses2.Set("false", "No", utils.DarkGray, utils.SaffronYellow)

	return &UserRelationships{
		db:        db,
		conn:      conn,
		user:      user,
		entity:    entity,
		statuses1: statuses1,
		statuses2: statuses2,
	}, nil
}

func (t *UserRelationships) GetUserRelationshipsTable(ctx *context.Context) table.Table {
	userRelationships := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))
	tableName := "user_relationships"
	info := userRelationships.GetInfo().SetFilterFormLayout(form.LayoutFilter)

	info.AddField("ID", "id", db.Int8).FieldSortable().FieldFilterable()
	info.AddField("First User ID", "first_user_id", db.Text).FieldSortable().FieldFilterable().FieldDisplay(utils.ParseUserIDToLink)
	info.AddField("Second User ID", "second_user_id", db.Text).FieldSortable().FieldFilterable().FieldDisplay(utils.ParseUserIDToLink)
	info.AddField("First To Second Status", "first_to_second_status", db.Text).FieldDisplay(t.statuses1.ToFieldDisplay)
	info.AddField("Second To First Status", "second_to_first_status", db.Text).FieldDisplay(t.statuses1.ToFieldDisplay)
	info.AddField("Are Friends", "are_friends", db.Bool).FieldSortable().FieldFilterable(types.FilterType{FormType: form.SelectSingle}).
		FieldFilterOptions(t.statuses2.ToFieldOptions()).FieldDisplay(t.statuses2.ToFieldDisplay)

	info.SetTable(tableName).SetTitle("UserRelationships").SetDescription("User Relationships").AddCSS(utils.CssTableNoWrap)

	formList := userRelationships.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetUpdateFn(t.update)
	formList.SetPostValidator(t.postValidator)
	formList.AddField("ID", "id", db.Int8, form.Text).FieldDisableWhenCreate().FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("First User ID", "first_user_id", db.Text, form.Text).FieldDisplay(utils.ParseUserID)
	formList.AddField("Second User ID", "second_user_id", db.Text, form.Text).FieldDisplay(utils.ParseUserID)
	formList.AddField("First To Second Status", "first_to_second_status", db.Text, form.SelectSingle).FieldOptions(t.statuses1.ToFieldOptions())
	formList.AddField("Second To First Status", "second_to_first_status", db.Text, form.SelectSingle).FieldOptions(t.statuses1.ToFieldOptions())
	formList.AddField("Are Friends", "are_friends", db.Bool, form.Text).FieldDisplayButCanNotEditWhenUpdate().FieldDisplay(t.statuses2.ToFieldDisplay)
	formList.SetTable(tableName).SetTitle("UserRelationships").SetDescription("User Relationships")

	return userRelationships
}

func (t *UserRelationships) postValidator(values form2.Values) error {
	firstUserID := values.Get("first_user_id")
	if firstUserID == "" {
		return errors.New("first user id is required")
	}
	firstUser, err := t.user.getByID(firstUserID)
	if err != nil {
		return err
	}
	if firstUser == nil {
		return fmt.Errorf("not found first user %s", firstUserID)
	}

	secondUserID := values.Get("second_user_id")
	if secondUserID == "" {
		return errors.New("second user id is required")
	}
	secondUser, err := t.user.getByID(secondUserID)
	if err != nil {
		return err
	}
	if secondUser == nil {
		return fmt.Errorf("not found second user %s", secondUserID)
	}

	if firstUserID == secondUserID {
		return fmt.Errorf("the first user ID and the second user ID are the same %s", firstUserID)
	}

	return nil
}

func (t *UserRelationships) update(values form2.Values) error {
	updateFields := []string{
		"first_user_id", "second_user_id", "first_to_second_status", "second_to_first_status",
	}

	m := make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if utils2.InArray(updateFields, k) && len(v) > 0 {
			m[k] = v
		}
	}

	id := values.Get("id")
	firstUserID := values.Get("first_user_id")
	ulidValue, err := ulid.Parse(firstUserID)
	if err != nil {
		return err
	}
	m["first_user_id"] = ulidValue

	secondUserID := values.Get("second_user_id")
	ulidValue, err = ulid.Parse(secondUserID)
	if err != nil {
		return err
	}
	m["second_user_id"] = ulidValue
	if err = t.db.Table("user_relationships").Where("id = ?", id).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *UserRelationships) insert(values form2.Values) error {
	insertFields := []string{
		"first_user_id", "second_user_id", "first_to_second_status", "second_to_first_status",
	}
	m := make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if utils2.InArray(insertFields, k) && len(v) > 0 {
			m[k] = v
		}
	}

	firstUserID := values.Get("first_user_id")
	ulidValue, err := ulid.Parse(firstUserID)
	if err != nil {
		return err
	}
	m["first_user_id"] = ulidValue

	secondUserID := values.Get("second_user_id")
	ulidValue, err = ulid.Parse(secondUserID)
	if err != nil {
		return err
	}
	m["second_user_id"] = ulidValue
	if err = t.db.Table("user_relationships").Create(m).Error; err != nil {
		return err
	}
	return nil
}
