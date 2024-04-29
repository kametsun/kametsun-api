package utils

import (
	"database/sql"
	"fmt"
	wishitem "kametsun-api/models/WishItem"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// データベース接続をする
func InitDataBase() *sql.DB{
	MODE := os.Getenv("MODE")
	if MODE == "" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(fmt.Errorf("環境変数を読み込みできませんでした: %v", err))
		}
		fmt.Println("環境変数を読み込みました。")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		panic("環境変数 DB_URL が設定されていません。")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(fmt.Errorf("データベース接続に失敗しました: %v", err))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Errorf("データベースへの接続確認に失敗しました: %v", err))
	}
	fmt.Println("データベース接続に成功しました。")
	createTables(db)
	return db
}

// 必要なテーブルはここに追加する
func createTables(db *sql.DB){
	wishitem.CreateWishItemTable(db)
}
