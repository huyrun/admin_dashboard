package tables

import (
	"errors"
	"strconv"
	"time"

	"github.com/huyrun/admin_dashboard/embed"
	"github.com/huyrun/admin_dashboard/src/utils"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	utils2 "github.com/huyrun/go-admin/modules/utils"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/parameter"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template"
	"github.com/huyrun/go-admin/template/types"
	"github.com/huyrun/go-admin/template/types/form"
	"github.com/oklog/ulid/v2"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type User struct {
	db         *gorm.DB
	conn       db.Connection
	countries  []*Country
	countryMap map[string]*Country
	statuses   utils.StatusMap
}

type Country struct {
	Name string `json:"name" yaml:"name"`
	Code string `json:"code" yaml:"code"`
}

func NewUser(db *gorm.DB, conn db.Connection) (*User, error) {
	var countries []*Country
	err := yaml.Unmarshal(embed.CountriesData, &countries)
	if err != nil {
		return nil, err
	}

	countryMap := make(map[string]*Country)
	for _, c := range countries {
		countryMap[c.Code] = c
	}

	statuses := make(utils.StatusMap)
	statuses.Set("1", "Active", utils.SaffronYellow, utils.NavyBlue)
	statuses.Set("0", "Inactive", utils.DarkGray, utils.SaffronYellow)

	return &User{
		db:         db,
		conn:       conn,
		countries:  countries,
		countryMap: countryMap,
		statuses:   statuses,
	}, nil
}

func (t *User) GetUsersTable(ctx *context.Context) table.Table {
	tableConfig := table.DefaultConfigWithDriver("postgresql")
	tableConfig.PrimaryKey = table.PrimaryKey{
		Type: db.UUID,
		Name: "id",
	}
	users := table.NewDefaultTable(ctx, tableConfig)
	tableName := "users"
	info := users.GetInfo().SetFilterFormLayout(form.LayoutFilter).SetSortField("created_at")

	info.AddField("ID", "id", db.UUID).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			var id ulid.ULID
			err := id.UnmarshalBinary([]byte(value.Value))
			if err != nil {
				return value.Value
			}
			return id.String()
		})
	info.AddField("User Name", "username", db.Varchar).FieldSortable()
	info.AddField("First Name", "first_name", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Last Name", "last_name", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Email", "email", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Role", "role", db.Varchar).FieldSortable()
	info.AddField("Password Hash", "password_hash", db.Varchar)
	info.AddField("Age", "age", db.Int2).FieldSortable()
	info.AddField("DOB", "dob", db.Date).FieldSortable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02")
		})
	info.AddField("Sex", "sex", db.Tinyint).FieldSortable().FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "👨 Men"
		}
		if model.Value == "1" {
			return "👩 Women"
		}
		return "unknown"
	})
	info.AddField("Country", "country", db.Varchar).FieldSortable().FieldFilterable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			c, _ := t.countryMap[value.Value]
			if c == nil {
				return value.Value
			}
			return c.Name
		})
	info.AddField("City", "city", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Points", "points", db.Int).FieldSortable()
	info.AddField("Avatar URL", "avatar_url", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "" {
				return value.Value
			}
			return template.Default().Image().WithModal().SetSrc(template.HTML(value.Value)).GetContent()
		})
	info.AddField("Google Sub", "google_sub", db.Varchar)
	info.AddField("FbID", "fb_id", db.Varchar)
	info.AddField("Status", "status", db.Tinyint).
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldSortable().FieldFilterOptions(t.statuses.ToFieldOptions()).
		FieldDisplay(t.statuses.ToFieldDisplay)
	info.AddField("Created At", "created_at", db.Timestamptz).FieldSortable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02 15:04:05")
		})
	info.AddField("UpdatedAt", "updated_at", db.Timestamptz).FieldSortable().
		FieldDisplay(func(value types.FieldModel) interface{} {
			v, err := time.Parse(time.RFC3339, value.Value)
			if err != nil {
				return value.Value
			}
			return v.Format("2006-01-02 15:04:05")
		})

	info.SetTable(tableName).SetTitle("Users").SetDescription("Users").AddCSS(utils.CssTableNoWrap)

	formList := users.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetUpdateFn(t.update)
	formList.AddField("User Name", "username", db.Varchar, form.Text)
	formList.AddField("First Name", "first_name", db.Varchar, form.Text)
	formList.AddField("Last Name", "last_name", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Role", "role", db.Varchar, form.Text)
	formList.AddField("Password Hash", "password_hash", db.Text, form.Text)
	formList.AddField("Age", "age", db.Int2, form.Number).FieldDisplay(utils.CastToNumber)
	formList.AddField("Dob", "dob", db.Date, form.Date).FieldDefault(time.Now().Format("2006-01-02"))
	formList.AddField("Sex", "sex", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "👨 Men", Value: "0"},
			{Text: "👩 Women", Value: "1"},
		}).FieldDefault("0")
	formList.AddField("Country", "country", db.Varchar, form.SelectSingle).
		FieldInputWidth(4).FieldOptions(t.countryList())
	formList.AddField("City", "city", db.Varchar, form.Text).FieldInputWidth(4)
	formList.AddField("Points", "points", db.Int, form.Number).FieldDisplay(utils.CastToNumber)
	formList.AddField("Avatar URL", "avatar_url", db.Varchar, form.Text)
	formList.AddField("Google Sub", "google_sub", db.Varchar, form.Text)
	formList.AddField("FbID", "fb_id", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Tinyint, form.Switch).FieldDefault("1").FieldOptions(t.statuses.ToFieldOptions())

	formList.SetTable(tableName).SetTitle("Users").SetDescription("Users")

	users.GetDetailFromInfo().SetGetDataFn(t.getDetail)
	users.GetInfo().SetDeleteFn(t.delete)

	return users
}

func (t *User) getDetail(param parameter.Parameters) ([]map[string]interface{}, int) {
	ulidString := param.PK()
	ulidValue, err := ulid.Parse(ulidString)
	if err != nil {
		return nil, 0
	}

	var result []map[string]interface{}
	err = t.db.Raw("SELECT * FROM users WHERE id = ?", ulidValue).Scan(&result).Error
	if err != nil {
		return nil, 0
	}

	newResult := make([]map[string]interface{}, 0, len(result))
	for _, row := range result {
		for k, v := range row {
			switch v.(type) {
			case []uint8:
				row[k] = string(v.([]uint8))
			case time.Time:
				row[k] = v.(time.Time).Format(time.RFC3339)
			default:
			}
		}
		newResult = append(newResult, row)
	}

	return newResult, len(newResult)
}

func (t *User) delete(ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	ulids := make([]ulid.ULID, 0, len(ids))
	for _, id := range ids {
		ulidValue, err := ulid.Parse(id)
		if err != nil {
			return err
		}
		ulids = append(ulids, ulidValue)
	}

	result := t.db.Exec("DELETE FROM users WHERE id IN ?", ulids)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (t *User) update(values form2.Values) error {
	updateFields := []string{
		"username", "first_name", "last_name", "email", "role", "password_hash", "age", "dob", "sex",
		"country", "city", "points", "avatar_url", "google_sub", "fb_id", "status", "updated_at",
	}

	m := make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if utils2.InArray(updateFields, k) && len(v) > 0 {
			m[k] = v
		}
	}

	id := values.Get("id")
	ulidValue, err := ulid.Parse(id)
	if err != nil {
		return err
	}
	m["updated_at"] = time.Now()
	if err = t.db.Table("users").Where("id = ?", ulidValue).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *User) insert(values form2.Values) error {
	insertFields := []string{
		"username", "first_name", "last_name", "email", "role", "password_hash", "age", "dob", "sex",
		"country", "city", "points", "avatar_url", "google_sub", "fb_id", "status", "created_at", "updated_at",
	}
	m := make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if utils2.InArray(insertFields, k) && len(v) > 0 {
			m[k] = v
		}
	}

	m["id"] = ulid.Make()
	timeNow := time.Now()
	m["created_at"] = timeNow
	m["updated_at"] = timeNow
	if err := t.db.Table("users").Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (t *User) countryList() types.FieldOptions {
	fieldOptions := types.FieldOptions{}
	for _, c := range t.countries {
		fieldOptions = append(fieldOptions, types.FieldOption{Text: c.Name, Value: c.Code})
	}
	return fieldOptions
}

func (t *User) getByID(id string) (map[string]interface{}, error) {
	ulidValue, err := ulid.Parse(id)
	if err != nil {
		return nil, err
	}

	query := `select *
from users
where id = ?
order by id desc
limit 1;`
	res, err := t.conn.Query(query, ulidValue)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res[0], nil
}

func (t *User) postValidator(values form2.Values) error {
	pointsStr := values.Get("points")
	points, err := strconv.Atoi(pointsStr)
	if err != nil {
		return err
	}
	if points < 0 {
		return errors.New("points must be greater than zero")
	}

	ageStr := values.Get("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return err
	}
	if age < 0 {
		return errors.New("age must be greater than zero")
	}

	return nil
}
