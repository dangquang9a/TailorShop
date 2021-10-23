package api

import (
	db "TailorShop/db/sqlc"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCustomerRequest struct {
	FullName    string `json:"fullName" binding:"required"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

func (server *Server) createCustomer(ctx *gin.Context) {
	var req createCustomerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorReponse(err))
		return
	}

	arg := db.CreateCustomerParams{
		FullName:    req.FullName,
		Address:     sql.NullString{String: req.Address, Valid: true},
		PhoneNumber: req.PhoneNumber,
	}

	customer, err := server.store.CreateCustomer(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorReponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

type getCustomerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCustomer(ctx *gin.Context) {

	var req getCustomerRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorReponse(err))
		return
	}

	customer, err := server.store.GetCustomer(ctx, int32(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorReponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorReponse(err))
		return
	}
	ctx.JSON(http.StatusOK, customer)

}

type listcustomerRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCustomer(ctx *gin.Context) {
	var req listcustomerRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorReponse(err))
		return
	}

	arg := db.ListCustomersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	customers, err := server.store.ListCustomers(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorReponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)

}

func errorReponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
