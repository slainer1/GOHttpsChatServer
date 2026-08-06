package main

import (
	_ "main/docs"
	"main/internal/db"
	"main/internal/env"
	"main/internal/store"
	"time"

	"go.uber.org/zap"
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

//	@host	localhost:8080

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@desciription				JWT Authorization header using the Bearer scheme

func main() {
	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgresql://user:Nathansno1@localhost:5432/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp: time.Hour * 24 * 3,
		},
	}
	//Logger

	//flushes buffered log entries
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()
	//Database

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("DB Conntection Pool Established")

	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))

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
