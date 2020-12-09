package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yeric17/inventory-system/common/apierrors"
	"github.com/yeric17/inventory-system/models"
	"github.com/yeric17/inventory-system/services"
)

var (
	ProductController = productController{}
)

type productController struct{}

func (productController) GetAll(w http.ResponseWriter, r *http.Request) {
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

	products, err := services.ProductServices.GetAll(int32(limit), int32(page))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		apierrors.HandleError(w, "[Error] al solicitar los productos", err)
		return
	}

	err = products.ToJSON(w, products)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		apierrors.HandleError(w, "[Error] al transformar a JSON", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (productController) Update(w http.ResponseWriter, r *http.Request) {
	var prodUpdate models.ProductUpdate
	var err error

	err = prodUpdate.DecodeJSON.FromJSON(r.Body, &prodUpdate)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		apierrors.HandleError(w, "[Error] Formato JSON no valido", err)
		return
	}

	prod, err := services.ProductServices.Update(prodUpdate)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		apierrors.HandleError(w, "[Error] Formato JSON no valido", err)
		return
	}
	fmt.Printf("Producto actualizado: %+v", prod)
	err = prod.ToJSON(w, prod)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		apierrors.HandleError(w, "[Error] Error al transformar a JSON", err)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (productController) Delete(w http.ResponseWriter, r *http.Request) {
	var prod models.Product
	var err error
	vars := mux.Vars(r)

	prodID, err := strconv.Atoi(vars["id"])

	fmt.Println(prodID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		apierrors.HandleError(w, "[Error] Id no valido, debe ser numérico", err)
		return
	}

	prod, err = services.ProductServices.Delete(uint64(prodID))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		apierrors.HandleError(w, "[Error] Id no encontrado", err)
		return
	}

	fmt.Printf("Producto eliminado: %+v", prod)
	err = prod.ToJSON(w, prod)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		apierrors.HandleError(w, "[Error] Error al transformar a JSON", err)
		return
	}
	w.WriteHeader(http.StatusOK)

}

// func (*productController) GetFiltered(w http.ResponseWriter, r *http.Request) {
// 	var queries url.Values
// 	var limit, page int
// 	var err error
// 	queries = r.URL.Query()

// 	if len(queries) > 0 {
// 		limit, err = strconv.Atoi(queries["limit"][0])
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			apierrors.HandleError(w, "[Error] en al pasar la variable limit", err)
// 			return
// 		}

// 		page, err = strconv.Atoi(queries["page"][0])
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			apierrors.HandleError(w, "[Error] al pasar la variable page", err)
// 			return
// 		}
// 	} else {
// 		limit = 100
// 		page = 1
// 	}
// 	prodFilter := models.ProductFilters{}

// 	name, nameOK := queries["name"]
// 	if nameOK {
// 		prodFilter.Name = name[0]
// 	}

// 	maxPrice, maxPriceOK := queries["max-price"]
// 	if maxPriceOK {
// 		prodFilter.MaxPrice, _ = strconv.ParseFloat(maxPrice[0], 64)
// 	}

// 	minPrice, minPriceOK := queries["min-price"]
// 	if minPriceOK {
// 		prodFilter.MaxPrice, _ = strconv.ParseFloat(minPrice[0], 64)
// 	}

// 	minQuantity, minQuantityOK := queries["min-quantity"]
// 	if minQuantityOK {
// 		prodFilter.MinQuantity, _ = strconv.Atoi(minQuantity[0])
// 	}

// 	maxQuantity, maxQuantityOK := queries["max-quantity"]
// 	if maxQuantityOK {
// 		prodFilter.MinQuantity, _ = strconv.Atoi(maxQuantity[0])
// 	}

// }
