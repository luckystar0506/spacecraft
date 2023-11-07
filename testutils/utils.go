// testutils/utils.go

package testutils

import (
	"context"
	"log"
	"spacecraft/ent"
	"spacecraft/ent/migrate"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func Setup() *ent.Client {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
