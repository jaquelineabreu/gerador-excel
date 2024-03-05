package main

import (
	"fmt"
	"gerador-excel/db"
	"gerador-excel/internal"
	"gerador-excel/repository"
	"log"
	"log/slog"
	"net/http"
	"os"

	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	slog.Info("Conectando ao banco..")
	db := db.Conection()

	slog.Info("Obtendo dados do banco sem canal...")
	allocations, err := repository.SelectAllAllocationKeyReports(db)
	if err != nil {
		slog.Error(fmt.Sprintf("Erro ao obter dados do banco: %v", err))
	}

	slog.Info("Gerando planilha sem canal...")
	if err := internal.NewStreamWriterSemCanal(allocations); err != nil {
		slog.Error(fmt.Sprintf("Falha ao gerar a planilha: %v", err))
		os.Exit(1)
	}

}
