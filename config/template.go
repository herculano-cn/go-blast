package config

import "github.com/gofiber/template/handlebars/v2"

// TODO: Add template engine logic here
func SetEngine() *handlebars.Engine {
	engine := handlebars.New("./views", ".hbs")
	return engine
}
