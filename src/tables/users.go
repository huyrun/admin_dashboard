package tables

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/huyrun/admin_dashboard/embed"
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/modules/db"
	"github.com/huyrun/go-admin/modules/utils"
	form2 "github.com/huyrun/go-admin/plugins/admin/modules/form"
	"github.com/huyrun/go-admin/plugins/admin/modules/table"
	"github.com/huyrun/go-admin/template"
	"github.com/huyrun/go-admin/template/color"
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
}

type Country struct {
	Name string `json:"name" yaml:"name"`
	Code string `json:"code" yaml:"code"`
}

var userFields = []string{
	"username", "first_name", "last_name", "email", "role", "password_hash", "age", "dob", "sex",
	"country", "city", "points", "avatar_url", "google_sub", "fb_id", "status",
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

	return &User{
		db:         db,
		conn:       conn,
		countries:  countries,
		countryMap: countryMap,
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

	info.AddField("ID", "id", db.UUID).FieldSortable().FieldFilterable()
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
			return template.Default().Image().WithModal().SetSrc(template.HTML(value.Value)).GetContent()
		})
	info.AddField("Google Sub", "google_sub", db.Varchar)
	info.AddField("FbID", "fb_id", db.Varchar)
	info.AddField("Status", "status", db.Tinyint).
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldSortable().FieldFilterOptions(types.FieldOptions{
		{Value: "1", Text: "Active"},
		{Value: "0", Text: "Inactive"},
	}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "0" {
				return fmt.Sprintf(`<span class="label" style="background-color: %s; color: %s;">Inactive</span>`, color.Gray, color.White)
			}
			if value.Value == "1" {
				return fmt.Sprintf(`<span class="label" style="background-color: %s;  color: %s;">Active</span>`, color.Yellow, color.Black)
			}
			return fmt.Sprintf(`<span class="label" style="text-decoration: line-through; background-color: %s; color: %s;">Unknown</span>`, color.Red, color.Black)
		})
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

	info.SetTable(tableName).SetTitle("Users").SetDescription("Users").AddCSS(cssTableNoWrap)

	formList := users.GetForm()
	formList.SetInsertFn(t.insert)
	formList.SetPreProcessFn(t.preProcess)
	formList.AddField("User Name", "username", db.Varchar, form.Text)
	formList.AddField("First Name", "first_name", db.Varchar, form.Text)
	formList.AddField("Last Name", "last_name", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Role", "role", db.Varchar, form.Text)
	formList.AddField("Password Hash", "password_hash", db.Text, form.Text).FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Age", "age", db.Int2, form.Number).FieldDefault("18")
	formList.AddField("Dob", "dob", db.Date, form.Date).FieldDefault(time.Now().Format("2006-01-02"))
	formList.AddField("Sex", "sex", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "👨 Men", Value: "0"},
			{Text: "👩 Women", Value: "1"},
		}).FieldDefault("0")
	formList.AddField("Country", "country", db.Varchar, form.SelectSingle).
		FieldInputWidth(4).FieldOptions(t.countryList())
	formList.AddField("City", "city", db.Varchar, form.Text).FieldInputWidth(4)
	formList.AddField("Points", "points", db.Int, form.Number).FieldDefault("0")
	formList.AddField("Avatar URL", "avatar_url", db.Varchar, form.Text)
	formList.AddField("Google Sub", "google_sub", db.Varchar, form.Text)
	formList.AddField("FbID", "fb_id", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Tinyint, form.Switch).FieldDefault("1").
		FieldOptions(types.FieldOptions{
			{Text: "Active", Value: "1"},
			{Text: "Inactive", Value: "0"},
		})

	formList.SetTable(tableName).SetTitle("Users").SetDescription("Users")

	return users
}

func (t *User) preProcess(values form2.Values) form2.Values {
	switch values.Get(form2.PostTypeKey) {
	case "0": // update
		values.Add("updated_at", time.Now().Format(time.RFC3339))
	case "1": // create
		values.Add("created_at", time.Now().Format(time.RFC3339))
		values.Add("updated_at", time.Now().Format(time.RFC3339))
	}
	return values
}

func (t *User) insert(values form2.Values) error {
	m := make(map[string]interface{})
	for k := range values {
		v := values.Get(k)
		if utils.InArray(userFields, k) && len(v) > 0 {
			m[k] = v
		}
	}

	m["id"] = ulid.Make().String()
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
	query := `select id, username, first_name, last_name, email, role, password_hash, 
       age, dob ,sex , country , city, points, avatar_url, google_sub, fb_id, status, created_at, updated_at
from users
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
