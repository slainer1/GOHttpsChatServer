package main

import (
	"log"
	"main/internal/db"
	"main/internal/env"
	"main/internal/store"
)

// //////////////////////////////////////////////
// // To run main,
// // Use cmd 'air' in terminal to run while in 'social' dir
// // or start by using go run /cmd/api/*.go
// //////////////////////////////////////////////

const version = "0.0.2"

//	@title			GopherSocial API
//	@description	API for GopherSocial, a social network for gophers
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @desciription
func main() {
	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgresql://:Nathansno1@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Printf("DB Conntection Pool Established")
	store := store.NewStorage(db)
	app := &application{config: cfg, store: store}

	mux := app.mount()
	log.Fatal(app.run(mux))

}

// Package Validator
//func main() {
//	post := Post{}
//	t := reflect.TypeOf(post)
//	field, _ := t.FieldByName("Post")
//	validateTag := field.Tag.Get("validate")
//	validateRules := strings.Split(validateTag, ",")
//	fmt.Println("Validate Tags", validateTag)
//	fmt.Println("Validate Rules", validateRules)
//}
