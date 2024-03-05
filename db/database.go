package db

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

func Conection() (db *sqlx.DB) {
	dsn := "root:root@tcp(localhost:3306)/tg4?parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		slog.Error(fmt.Sprintf(" Erro ao conectar no banco: %v", err))
	}

	slog.Info("Conexão estabelecida com sucesso!")

	return db
}
