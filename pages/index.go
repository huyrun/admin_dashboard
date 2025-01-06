package pages

import (
	"github.com/huyrun/go-admin/context"
	"github.com/huyrun/go-admin/template/types"
)

func GetWelcome(ctx *context.Context) (types.Panel, error) {
	return types.Panel{
		Content: `<img src="/static/assets/dist/img/welcome.png" alt="Welcome Image" style="width:100%; height:100%;">`,
	}, nil
}
