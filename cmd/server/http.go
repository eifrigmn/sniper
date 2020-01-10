package server

import (
	"context"
	"net/http"

	"sniper/cmd/server/hook"
	"sniper/dao/user"
	user_v1 "sniper/rpc/user/v1"
	"sniper/server/userserver1"

	"github.com/bilibili/twirp"
)

var hooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewLog(),
)

var loginHooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewCheckLogin(),
	hook.NewLog(),
)

func initMux(mux *http.ServeMux) {
	dao := user.New(context.Background(), user.DBName)
	server := &userserver1.Server{Dao: dao}
	handler := user_v1.NewUserServer(server, hooks)
	mux.Handle(user_v1.UserPathPrefix, handler)
}

func initInternalMux(mux *http.ServeMux) {
}
