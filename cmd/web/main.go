package main

import (
	"database/sql"
	"flag"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mvahedii/gorem/internal/handlers"
	"github.com/mvahedii/gorem/internal/repositories/database"
	v1 "github.com/mvahedii/gorem/internal/services/v1"
	"github.com/mvahedii/gorem/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP Network Port")

	dsn := flag.String("dsn", "dbadmin:pass@/gorem?parseTime=true", "MySQL data source name")

	flag.Parse()

	db, err := gorm.Open(mysql.Open(*dsn), &gorm.Config{})

	if err != nil {
		utils.ErrLog.Fatal(err)
	}

	// defer db.Close()

	wordRepository := database.NewWordRepository(db)
	wordService := v1.NewWordService(wordRepository)
	wordHandler := handlers.NewWordHandler(wordService)

	srv := handlers.NewHTTPServer(*wordHandler, addr)

	utils.InfoLog.Print("Server Starting...", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		utils.ErrLog.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
