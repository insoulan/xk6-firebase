package xk6firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"google.golang.org/api/option"
)

type (
	RootModule     struct{}
	ModuleInstance struct {
		vu modules.VU
	}
)

var (
	_ modules.Module   = &RootModule{}
	_ modules.Instance = &ModuleInstance{}
)

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{vu: vu}
}

func (inst *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"App": inst.NewApp,
		},
	}
}

func (inst *ModuleInstance) NewApp(call goja.ConstructorCall) *goja.Object {
	rt := inst.vu.Runtime()

	var filename string
	err := rt.ExportTo(call.Arguments[0], &filename)
	if err != nil {
		common.Throw(rt, err)
	}

	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(filename))
	if err != nil {
		common.Throw(rt, err)
	}

	return rt.ToValue(&App{
		vu:  inst.vu,
		app: app,
	}).ToObject(rt)
}
