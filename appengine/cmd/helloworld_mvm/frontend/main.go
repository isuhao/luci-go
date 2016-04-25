// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package main implements HTTP server that handles requests to default
// module.
package main

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"google.golang.org/appengine"

	"github.com/luci/gae/service/info"
	"github.com/luci/luci-go/appengine/gaeauth/server"
	"github.com/luci/luci-go/appengine/gaemiddleware"
	"github.com/luci/luci-go/server/auth"
	"github.com/luci/luci-go/server/middleware"
	"github.com/luci/luci-go/server/templates"
)

// templateBundle is used to render HTML templates. It provides a base args
// passed to all templates.
var templateBundle = &templates.Bundle{
	Loader:    templates.FileSystemLoader("templates"),
	DebugMode: appengine.IsDevAppServer(),
	DefaultArgs: func(c context.Context) (templates.Args, error) {
		loginURL, err := auth.LoginURL(c, "/")
		if err != nil {
			return nil, err
		}
		logoutURL, err := auth.LogoutURL(c, "/")
		if err != nil {
			return nil, err
		}
		isAdmin, err := auth.IsMember(c, "administrators")
		if err != nil {
			return nil, err
		}
		return templates.Args{
			"AppVersion":  strings.Split(info.Get(c).VersionID(), ".")[0],
			"IsAnonymous": auth.CurrentIdentity(c) == "anonymous:anonymous",
			"IsAdmin":     isAdmin,
			"User":        auth.CurrentUser(c),
			"LoginURL":    loginURL,
			"LogoutURL":   logoutURL,
		}, nil
	},
}

// base is the root of the middleware chain.
func base(h middleware.Handler) httprouter.Handle {
	methods := auth.Authenticator{
		&server.OAuth2Method{Scopes: []string{server.EmailScope}},
		server.CookieAuth,
		&server.InboundAppIDAuthMethod{},
	}
	h = auth.Use(h, methods)
	h = templates.WithTemplates(h, templateBundle)
	return gaemiddleware.BaseProd(h)
}

//// Routes.

func main() {
	router := httprouter.New()
	server.InstallHandlers(router, base)
	router.GET("/", base(auth.Authenticate(indexPage)))
	http.DefaultServeMux.Handle("/", router)

	appengine.Main()
}

//// Handlers.

func indexPage(c context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templates.MustRender(c, w, "pages/index.html", nil)
}
