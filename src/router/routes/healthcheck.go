package routes

import (
	"fmt"
	"net/http"
)

var healthcheckRoute = []Route{
	{
		URI:    "/",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Working!!!")
		},
		AuthRequired: false,
	},
}
