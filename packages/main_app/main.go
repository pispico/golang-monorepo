package main

import (
	"fmt"

	"github.com/pispico/golang-monorepo/packages/main_app/router"
)

func main() {
	fmt.Println("Run main application router")
	_ = router.GetEngine().Run(":8080")
}

func Check() {
	fmt.Println("Check")
}
