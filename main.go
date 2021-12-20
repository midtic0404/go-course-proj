package main

import (
	"fmt"
	"net/http"

	"github.com/midtic0404/webservice/controllers"
)

func main() {
	fmt.Println("Running service..")
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
