package xk6firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

type App struct {
	vu  modules.VU
	app *firebase.App
}

func (app *App) Auth() *goja.Object {
	rt := app.vu.Runtime()

	client, err := app.app.Auth(context.Background())
	if err != nil {
		common.Throw(rt, err)
	}

	return rt.ToValue(&Auth{
		vu:   app.vu,
		auth: client,
	}).ToObject(rt)
}
