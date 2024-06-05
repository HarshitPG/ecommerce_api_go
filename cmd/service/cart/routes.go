package cart

import (
	"fmt"
	"net/http"

	"github.com/HarshitPG/ecommerce_api_go/cmd/service/auth"
	"github.com/HarshitPG/ecommerce_api_go/cmd/types"
	"github.com/HarshitPG/ecommerce_api_go/cmd/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	
	productStore types.ProductStore
	userStore types.UserStore
	orderStore types.OrderStore
}

func NewHandler(orderStore types.OrderStore, productStore types.ProductStore,userStore types.UserStore) *Handler {
	return &Handler{orderStore: orderStore, productStore: productStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/cart/checkout",auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request){
	userID:=auth.GetUserIDFromContext(r.Context())
	var cart types.CartCheckoutPayload
	if err:= utils.ParseJSON(r,&cart); err!=nil{
		utils.WriteError(w, http.StatusBadRequest,err)
		return
	}

	if err:= utils.Validate.Struct(cart); err!=nil{
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest,fmt.Errorf("invalid payload: %v",errors))
		return
	}

	productIDs, err := getCartItemsIDs(cart.Items)
	if err !=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}
	ps,err:= h.productStore.GetProductsByID(productIDs)
	if err !=nil {
		utils.WriteError(w, http.StatusInternalServerError,err)
		return
	}

	orderID, totalPrice,err:=h.createOrder(ps,cart.Items,userID)
	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}
	utils.WriteJSON(w,http.StatusOK,map[string]any{
		"total_price":totalPrice,
		"order_id":orderID,
	})

}