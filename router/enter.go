package router

import (
	"blog/server/admin"
	"blog/server/api"
)

var (
	adminApiGroup = admin.AdminApi{}
	apiGroup = api.Api{}
)