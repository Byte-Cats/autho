package main

import (
	"github.com/byte-cats/autho/db"
	"github.com/byte-cats/autho/routing"
	"github.com/byte-cats/autho/serve"
)

func main() {
	defer db.DB.Close()
	db.InitDB()

	serve.Serve(routing.InitRoutes())
}
