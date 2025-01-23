package utils

import (
	"fmt"
	"html/template"

	template2 "github.com/huyrun/go-admin/template"
	"github.com/huyrun/go-admin/template/types"
	"github.com/oklog/ulid/v2"
)

const (
	BrightBlue    = "#0063EC"
	SaffronYellow = "#FFD154"
	DarkGray      = "#3D3D3D"
	MintGreen     = "#68D69D"
	NavyBlue      = "#002759"
	DeepPurple    = "#401D83"
)

var CssTableNoWrap = template.CSS(`table{white-space: nowrap;}`)

func LinkToOtherTable(tableName, value string) template.HTML {
	return template2.Default().
		Link().
		SetURL(fmt.Sprintf("/admin/info/%s/detail?__goadmin_detail_pk=%s", tableName, value)).
		SetContent(template.HTML(value)).
		OpenInNewTab().
		SetTabTitle(template.HTML(fmt.Sprintf("%s Detail(%s)", CapitalizeFirst(tableName), value))).
		GetContent()
}

func ParseUserID(value types.FieldModel) interface{} {
	var id ulid.ULID
	err := id.UnmarshalBinary([]byte(value.Value))
	if err != nil {
		return value.Value
	}
	return id.String()
}

func ParseUserIDToLink(value types.FieldModel) interface{} {
	var id ulid.ULID
	err := id.UnmarshalBinary([]byte(value.Value))
	if err != nil {
		return LinkToOtherTable("users", value.Value)
	}
	return LinkToOtherTable("users", id.String())
}

func ToLabel(text, backgroundColor, color string) string {
	return fmt.Sprintf(`<span class="label" style="background-color: %s; color: %s;">%s</span>`, backgroundColor, color, text)
}

type Status struct {
	Value           string
	Text            string
	BackgroundColor string
	Color           string
}

type StatusMap map[string]Status

func (s StatusMap) Set(value, text, backgroundColor, color string) {
	s[value] = Status{value, text, backgroundColor, color}
}

func (s StatusMap) ToFieldOptions() types.FieldOptions {
	var fieldOptions types.FieldOptions
	for value, status := range s {
		fieldOptions = append(fieldOptions, types.FieldOption{Value: value, Text: status.Text})
	}
	return fieldOptions
}

func (s StatusMap) ToFieldDisplay(value types.FieldModel) interface{} {
	if status, ok := s[value.Value]; ok {
		return ToLabel(status.Text, status.BackgroundColor, status.Color)
	}
	return ToLabel(value.Value, MintGreen, DeepPurple)
}
