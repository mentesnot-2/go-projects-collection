package main


import (
	"github.com/mentesnot-2/task_management_with_mongodb/router"

)
func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}