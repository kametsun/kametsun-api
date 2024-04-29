package wishItem

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type WishItem struct{
	ID string	`json:"id"`
	ImageURL string	`json:"image_url"`
	ItemURL string	`json:"item_url"`
	ItemName string	`json:"item_name"`
	CreatedAt time.Time	`json:"created_at"`
}

// WishItem一覧を取得する
func GetWishItems(db *sql.DB) ([]WishItem, error){
	const sql = `
		SELECT id, image_url, item_url, item_name, created_at
		FROM wish_items
		ORDER BY created_at DESC
	`

	rows, err := db.Query(sql)
	if err != nil {
		log.Printf("WishItemの取得に失敗しました。: %v", err)
		return nil, err
	}
	defer rows.Close()

	var wishItems []WishItem
	for rows.Next() {
		var item WishItem
		err := rows.Scan(&item.ID, &item.ImageURL, &item.ItemURL, &item.ItemName, &item.CreatedAt)
		if err != nil {
			log.Printf("WishItemのスキャンに失敗しました。: %v", err)
			return nil, err
		}
		wishItems = append(wishItems, item)
	}
	return wishItems, nil
}

func NewWishItem(imageURL, itemURL, itemName string) WishItem {
	return WishItem{
		ID: uuid.New().String(),
		ImageURL: imageURL,
		ItemURL: itemURL,
		ItemName: itemName,
		CreatedAt: time.Now(),
	}
}

func CreateWishItemTable(db *sql.DB) {
	const sql = `
	CREATE TABLE IF NOT EXISTS wish_items (
		id UUID PRIMARY KEY,
		image_url TEXT,
		item_url TEXT,
		item_name TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(sql)
	if err != nil {
		log.Printf("wish_itemsテーブルの作成に失敗しました: %v",err)
	}

	log.Println("wish_tables is OK.")
}

func CreateWishItem(db *sql.DB, item WishItem) error {
	const sql = `
		INSERT INTO wish_items (id, image_url, item_url, item_name, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	item.ID = uuid.New().String()
	item.CreatedAt = time.Now()

	_, err := db.Exec(sql, item.ID, item.ImageURL, item.ItemURL, item.ItemName, item.CreatedAt)
	if err != nil {
		log.Printf("欲しい物リストに追加できませんでした。: %v", err)
		return err
	}

	log.Println("ほしい物リストに追加されました。")
	return nil
}
