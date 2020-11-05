package controllers

import (
	"MiCasa-API/internal/db"
	"MiCasa-API/internal/models"
	"MiCasa-API/pkg/logging"
	"github.com/gin-gonic/gin"
	"time"

	_ "MiCasa-API/docs"
)

// ReadAllDocumentHandler ...
// @Summary 全てのドキュメントを取得する
// @Produce  json
// @Success 200 {object} []models.Connection
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/db/document/readall [get]
func ReadAllDocumentHandler(ctx *gin.Context) {
	documents, err := db.FindAll()
	logging.PrintErorWithGinContext(err, ctx)
	logging.StatusOK(err, ctx, documents)
}

// CreateHandler ...
// @Summary 新規のドキュメントを作成する
// @Accept	 json
// @Produce  json
// @Param create body models.Connection true "Connection"
// @Success 200 {object} models.InsertOneResult
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/db/document/create [post]
func CreateHandler(ctx *gin.Context) {
	var params models.Connection
	err := ctx.BindJSON(&params)
	logging.PrintErorWithGinContext(err, ctx)
	params.CreatedAt = time.Now().Format(time.RFC3339)

	createResult, err := db.InsertRecord(params)
	logging.PrintErorWithGinContext(err, ctx)
	logging.StatusOK(err, ctx, createResult.InsertedID)
}

// ReadHandler ...
// @Summary 指定したIDのドキュメントを取得する
// @Accept	 json
// @Produce  json
// @Param read body models.Connection{uuid} true "UUID"
// @Success 200 {object} models.Connection
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/db/document/read [post]
func ReadHandler(ctx *gin.Context) {
	var params models.Connection
	err := ctx.BindJSON(&params)
	logging.PrintErorWithGinContext(err, ctx)

	findResult, err := db.FindByID(params.UUID)
	logging.PrintErorWithGinContext(err, ctx)
	logging.StatusOK(err, ctx, findResult)
}

// UpdateHandler ...
// @Summary 既存のドキュメントを更新する
// @Accept	 json
// @Produce  json
// @Param read body models.Connection{uuid} true "UUID"
// @Success 200 {object} models.UpdateResult
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/db/document/update [post]
func UpdateHandler(ctx *gin.Context) {
	var params models.Connection
	err := ctx.BindJSON(&params)
	logging.PrintErorWithGinContext(err, ctx)
	updateResult, err := db.UpdateByID(params)
	logging.PrintErorWithGinContext(err, ctx)
	logging.StatusOK(err, ctx, updateResult)
}

// DeleteHandler ...
// @Summary 既存のドキュメントを削除する
// @Accept	 json
// @Produce  json
// @Param delete body models.Connection{uuid} true "UUID"
// @Success 200 {object} models.DeleteResult
// @Failure 400 {object} models.ErrorResponse "UUID Not found"
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/db/document/delete [delete]
func DeleteHandler(ctx *gin.Context) {
	var params models.Connection
	err := ctx.BindJSON(&params)
	logging.PrintErorWithGinContext(err, ctx)

	deleteResult, err := db.DeleteByID(params.UUID)
	logging.PrintErorWithGinContext(err, ctx)
	if deleteResult.DeletedCount > 0 {
		logging.StatusOK(err, ctx, deleteResult)
	} else {
		logging.StatusBadRequest(err, ctx, "Not found")
	}
}
