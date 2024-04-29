package controllers

import (
	"database/sql"
	wishItem "kametsun-api/models/WishItem"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WishItemController struct {
	db *sql.DB
}

func NewWishItemController(db *sql.DB) *WishItemController{
	return &WishItemController{db: db}
}

// WishItem一覧を取得する
func (w *WishItemController) GetWishItems(c *gin.Context){
	wishItems, err := wishItem.GetWishItems(w.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wishItems)
}

// WishItemを登録する
func (w *WishItemController) Create(c *gin.Context){
	var item wishItem.WishItem
	if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// バリデーション
	// 商品名は必ず入力する
    if item.ItemName == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "商品名が入力されていません。"})
        return
    }

	newItem := wishItem.NewWishItem(item.ImageURL, item.ItemURL, item.ItemName)

	err := wishItem.CreateWishItem(w.db, newItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, newItem)
}
