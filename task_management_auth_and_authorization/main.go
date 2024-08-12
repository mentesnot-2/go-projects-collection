package main



import (
	"github.com/mentesnot-2/task_management_auth_and_authorization/router"
)



func  main() {
	r := router.SetupRouter()
	r.Run(":8080")
}