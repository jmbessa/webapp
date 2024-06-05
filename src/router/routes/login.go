package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var loginRoute = []Route{
	{
		URI:                   "/",
		Method:                http.MethodGet,
		Function:              controllers.LoadLoginScreen,
		RequireAuthentication: false,
	},
	{
		URI:                   "/login",
		Method:                http.MethodGet,
		Function:              controllers.LoadLoginScreen,
		RequireAuthentication: false,
	},
}
