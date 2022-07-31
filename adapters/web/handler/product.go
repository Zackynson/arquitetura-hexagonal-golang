package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Zackynson/go-hexagonal/adapters/dto"
	"github.com/Zackynson/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandler(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/products/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/products", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	router.Handle("/products/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/products/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(errorJSON(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorJSON(err.Error()))
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(errorJSON(err.Error()))
			return
		}

		updatedProduct, err := service.Enable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorJSON(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorJSON(err.Error()))
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(errorJSON(err.Error()))
			return
		}

		updatedProduct, err := service.Disable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorJSON(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorJSON(err.Error()))
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDTO dto.Product

		err := json.NewDecoder(r.Body).Decode(&productDTO)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorJSON(err.Error()))
			return
		}

		createdProduct, err := service.Create(productDTO.Name, productDTO.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorJSON(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(createdProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorJSON(err.Error()))
			return
		}
	})
}
