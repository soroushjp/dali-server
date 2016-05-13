package main

import (
	"strconv"

	"github.com/soroushjp/dali-server/context"
	"github.com/soroushjp/dali-server/handlers"
)

func main() {
	// use default environment
	app, err := context.NewAppContext()
	if err != nil {
		panic(err)
	}
	eng := handlers.NewEngine(app)

	eng.Run(":" + strconv.Itoa(int(app.Env.Port)))
}
