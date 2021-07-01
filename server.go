package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/buidl-labs/celo-voting-validator-backend/graph"
	"github.com/buidl-labs/celo-voting-validator-backend/graph/database"
	"github.com/buidl-labs/celo-voting-validator-backend/graph/generated"
	"github.com/buidl-labs/celo-voting-validator-backend/graph/model"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	DB_URL := os.Getenv("DB_URL")
	log.Println(DB_URL)
	if DB_URL == "" {
		log.Fatal("Please provide a DB url.")
	}

	opts, err := pg.ParseURL(DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	DB := database.New(opts)

	defer DB.Close()

	// DB.AddQueryHook(pgdebug.DebugHook{
	// 	Verbose: true,
	// })

	ctx := context.Background()
	if err := DB.Ping(ctx); err != nil {
		log.Println(err)
	}

	// DropAllTables(DB)
	// CreateAllTables(DB)

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	config := generated.Config{Resolvers: &graph.Resolver{
		DB: DB,
	}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	router.Handle("/", playground.Handler("CVVT", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func DropAllTables(DB *pg.DB) {
	qs := []string{
		"drop table epochs",
		"drop table validators",
		"drop table validator_stats",
		"drop table validator_groups",
		"drop table validator_group_stats",
	}

	for _, q := range qs {
		_, err := DB.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

func CreateAllTables(DB *pg.DB) {
	models := []interface{}{
		(*model.Epoch)(nil),
		(*model.ValidatorGroup)(nil),
		(*model.ValidatorGroupStats)(nil),
		(*model.Validator)(nil),
		(*model.ValidatorStats)(nil),
	}

	for _, model := range models {
		err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			log.Print(err)
		}
	}
}
