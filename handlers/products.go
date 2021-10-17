package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Ethical-Ralph/simple-crud-go/database"
	"github.com/Ethical-Ralph/simple-crud-go/models"
	"github.com/Ethical-Ralph/simple-crud-go/response"
	"github.com/Ethical-Ralph/simple-crud-go/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ProductHandler struct {
	l *log.Logger
	db *database.Database
}

func Product(l *log.Logger, db *database.Database) *ProductHandler {
	return &ProductHandler{l, db}
} 

func (p *ProductHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	helper.ContentJSON(rw)
	if r.Method == http.MethodGet {

		split := strings.Split(r.URL.Path, "/")
		if split[1] != "" {
			id := split[1]
			p.getProduct(id, rw, r)
			return
		}

		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw,r)
		return
	}

	if r.Method == http.MethodPut {
		split := strings.Split(r.URL.Path, "/")
		id := split[1]

		p.updateProduct(id, rw, r)
		return
	}

	if r.Method == http.MethodDelete {
		split := strings.Split(r.URL.Path, "/")
		id := split[1]

		p.deleteProduct(id, rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *ProductHandler) getProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle get all data")


	products, err := models.GetProducts(p.db)
	
	if err != nil {
		helper.HandleError(err, rw)
		return
	}

	response.Success(rw, products, "")
}

func (p *ProductHandler) addProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle create product")

	prod := models.NewProduct()

	json.NewDecoder(r.Body).Decode(&prod)
	result, err := prod.Save(p.db)

	if err != nil {
		helper.HandleError(err, rw)
		return
	}

	prod.ID =  result.InsertedID.(primitive.ObjectID)

	response.Success(rw, prod, "")
}

func (p *ProductHandler) updateProduct(id string, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle update product")

	data := models.NewProduct()
	json.NewDecoder(r.Body).Decode(&data)
	
	result, err := models.UpdateProduct(id, data, p.db)

	if err != nil {
		helper.HandleError(err, rw)
		return
	}

	response.Success(rw, result, "")
}

func (p *ProductHandler) getProduct(id string, rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle get product with id %s", id)


	product, err := models.GetProduct(id, p.db)
	
	if err != nil {
		helper.HandleError(err, rw)
		return
	}

	response.Success(rw, product, "")
}

func (p *ProductHandler) deleteProduct(id string, rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle delete product with the id %s", id)

	product, err := models.DeleteProduct(id, p.db)
	
	if err != nil {
		helper.HandleError(err, rw)
		return
	}

	response.Success(rw, product, "document deleted successfully")
}

