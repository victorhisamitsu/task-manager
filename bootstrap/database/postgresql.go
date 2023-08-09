package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Hitsa/task-manager/internal/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Teste struct {
	bun.BaseModel `bun:"customers"`
	CustomerID    string `bun:"customer_id"`
}

func NewDb() {
	connString := "postgres://postgres:smarters123@localhost:5432/meu_banco?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connString)))
	conexao, _ := sqldb.Conn((context.Background()))
	err := conexao.PingContext(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	db := bun.NewDB(sqldb, pgdialect.New())

	teste := new(Teste)
	err = db.NewSelect().Model(teste).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	// Select first 10 testes.
	var testes []Teste
	err = db.NewSelect().Model(&testes).Limit(10).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(testes)

}

// Criar as tabelas do DB
func CreateTables(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().IfNotExists().Model((*models.Note)(nil)).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = db.NewCreateTable().IfNotExists().Model((*models.Task)(nil)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Conexão DB
func ConexaoDb() *bun.DB {
	dsn := "postgres://postgres:smarters123@localhost:5432/taskmanager?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}
