package product

import (
	"net/http"

	"github.com/HarshitPG/ecommerce_api_go/cmd/types"
	"github.com/HarshitPG/ecommerce_api_go/cmd/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	Store types.ProductStore
}

func NewHandler(Store types.ProductStore) *Handler {
	return &Handler{Store: Store}
}


func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/products",h.handleCreateProduct).Methods(http.MethodGet)
	router.HandleFunc("/products",h.handleCreateProduct).Methods(http.MethodPost)
}


func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request){
	ps,err := h.Store.GetProducts()
	if err!=nil{
		utils.WriteError(w, http.StatusInternalServerError,err)
		return
	}
	utils.WriteJSON(w, http.StatusOK,ps)
}
