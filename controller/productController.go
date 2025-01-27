package controller

import (
	"encoding/json"
	"go_test/internal/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiConfig struct {
	DB *database.Queries
}

func (apiConfig *ApiConfig) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := apiConfig.DB.ListAllProducts(r.Context())
	if err != nil {
		log.Printf("Failed to get Products %v", err)
		http.Error(w, "Error to get Products", http.StatusBadRequest)
		return
	}
	responseWithJSON(w, http.StatusOK, products)
}

func (apiConfig *ApiConfig) CreateProduct(w http.ResponseWriter, r *http.Request) {
	type Paramater struct {
		Name string `json:"name"`
		Content string `json:"content"`
	}
	params := Paramater{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Failed to decode Product %v", err)
		http.Error(w, "Error to decode Product", http.StatusBadRequest)
		return
	}
	product, err := apiConfig.DB.CreateProduct(r.Context(), database.CreateProductParams{
		Name: params.Name,
		Content: params.Content,
	})
	if err != nil {
		log.Printf("Failed to create Product %v", err)
		http.Error(w, "Error to create Product", http.StatusBadRequest)
		return
	}
	responseWithJSON(w, http.StatusOK, product)
}

func (apiConfig *ApiConfig) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("Failed to parse ID %v", err)
		http.Error(w, "Error to parse ID", http.StatusBadRequest)
		return
	}
	type Paramater struct {
		Name string `json:"name"`
		Content string `json:"content"`
	}
	params := Paramater{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		log.Printf("Failed to decode Product %v", err)
		http.Error(w, "Error to decode Product", http.StatusBadRequest)
		return
	}
	_, err = apiConfig.DB.UpdateProduct(r.Context(), database.UpdateProductParams{
		ID: id,
		Name: params.Name,
		Content: params.Content,
	})
	if err != nil {
		log.Printf("Failed to update Product %v", err)
		http.Error(w, "Error to update Product", http.StatusBadRequest)
		return
	}
	responseWithJSON(w, http.StatusOK, map[string]string{"message": "product updated successfully"})
}

func (apiConfig *ApiConfig) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	parseId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Failed to parse ID %v", err)
		http.Error(w, "Error to parse ID", http.StatusBadRequest)
		return
	}
	err = apiConfig.DB.DeleteProduct(r.Context(), parseId)
	if err != nil {
		log.Printf("Failed to delete Product %v", err)
		http.Error(w, "Error to delete Product", http.StatusBadRequest)
		return
	}
	responseWithJSON(w, http.StatusOK, map[string]string{"message": "product deleted successfully"})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal Product %v", err)
		http.Error(w, "Error to marshal Product", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}