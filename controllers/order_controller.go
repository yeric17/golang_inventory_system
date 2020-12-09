package controllers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/yeric17/inventory-system/common/apierrors"
	"github.com/yeric17/inventory-system/models"
	"github.com/yeric17/inventory-system/services"
)

var (
	OrderController = orderController{}
)

type orderController struct{}

func (*orderController) GetAll(w http.ResponseWriter, r *http.Request) {
	var queries url.Values
	var limit, page int
	var err error
	queries = r.URL.Query()

	if len(queries) > 0 {
		limit, err = strconv.Atoi(queries["limit"][0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			apierrors.HandleError(w, "[Error] en al pasar la variable limit", err)
			return
		}

		page, err = strconv.Atoi(queries["page"][0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			apierrors.HandleError(w, "[Error] al pasar la variable page", err)
			return
		}
	} else {
		limit = 100
		page = 1
	}
	orders, err := services.OrderServices.GetAll(limit, page)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		apierrors.HandleError(w, "[Error] al solicitar los productos", err)
		return
	}

	err = orders.ToJSON(w, orders)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		apierrors.HandleError(w, "[Error] al transformar a JSON", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (*orderController) Create(w http.ResponseWriter, r *http.Request) {
	var order models.OrderCreate
	var err error
	err = order.FromJSON(r.Body, &order)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		apierrors.HandleError(w, "[Error] Formato JSON no valido", err)
		return
	}

	newOrder, err := services.OrderServices.Create(order)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		apierrors.HandleError(w, "[Error] No se logro crear la orden", err)
		return
	}

	err = newOrder.ToJSON(w, newOrder)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		apierrors.HandleError(w, "[Error] No se retorno order creada", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
