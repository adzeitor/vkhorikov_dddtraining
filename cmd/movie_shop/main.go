package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"dddtraining/application"
	"dddtraining/domain"
)

func toJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(data)
}

func main() {
	customerController := application.NewCustomerController()

	http.HandleFunc(
		"/api/customers/",
		func(w http.ResponseWriter, r *http.Request) {
			pathId := strings.TrimPrefix(r.RequestURI, "/api/customers/")
			customerId, err := strconv.Atoi(pathId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if r.Method == http.MethodGet {
				customer, err := customerController.Get(customerId)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				toJson(w, customer)
				return
			}
			if r.Method == http.MethodPut {
				var customer domain.Customer
				err := json.NewDecoder(r.Body).Decode(&customer)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				err = customerController.Update(customerId, customer)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		},
	)

	http.HandleFunc(
		"/api/customers",
		func(w http.ResponseWriter, r *http.Request) {
			var customer domain.Customer
			err := json.NewDecoder(r.Body).Decode(&customer)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = customerController.Create(customer)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)

	http.HandleFunc(
		"/api/purchase",
		func(w http.ResponseWriter, r *http.Request) {
			customerId, err := strconv.Atoi(r.FormValue("customerId"))
			if err != nil {
				http.Error(w, "invalid customerId", http.StatusBadRequest)
				return
			}

			movieId, err := strconv.Atoi(r.FormValue("movieId"))
			if err != nil {
				http.Error(w, "invalid movieId", http.StatusBadRequest)
				return
			}

			err = customerController.PurchaseMovie(customerId, movieId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)

	http.HandleFunc(
		"/api/promotion",
		func(w http.ResponseWriter, r *http.Request) {
			customerId, err := strconv.Atoi(r.FormValue("customerId"))
			if err != nil {
				http.Error(w, "invalid customerId", http.StatusBadRequest)
				return
			}

			err = customerController.PromoteCustomer(customerId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		},
	)

	fmt.Println("go to http://localhost:8000/api/customers/1")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
