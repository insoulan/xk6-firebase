package xk6firebase

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

type Auth struct {
	vu   modules.VU
	auth *auth.Client
}

func (client *Auth) CreateUser(params map[string]any) *goja.Object {
	rt := client.vu.Runtime()

	user := &auth.UserToCreate{}

	for k, v := range params {
		switch k {
		case "disabled":
			user.Disabled(v.(bool))
		case "displayName":
			user.DisplayName(v.(string))
		case "email":
			user.Email(v.(string))
		case "emailVerified":
			user.EmailVerified(v.(bool))
		case "password":
			user.Password(v.(string))
		case "phoneNumber":
			user.PhoneNumber(v.(string))
		case "photoUrl":
			user.PhotoURL(v.(string))
		case "uid":
			user.UID(v.(string))
		}
	}

	rec, err := client.auth.CreateUser(context.Background(), user)
	if err != nil {
		common.Throw(rt, err)
	}

	return rt.ToValue(rec).ToObject(rt)
}
