package handlers

import (
	"log"
  "fmt"
  "net/http"
  "github.com/goweb4/models"
  "github.com/gorilla/mux"
  "strconv"
  "github.com/goweb4/utils"
  "html/template"
)

/**
  * Index Product
  */
func IndexProduct(w http.ResponseWriter) {
	fmt.Fprintln(w, "Need to be implement")
}

/**
  * Show product
  */
func ShowProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 32);
  product, err := models.GetProduct(uint(id)); if err != nil {
    fmt.Fprintln(w, err);
  }
  fmt.Fprintln(w, product)
}

/**
  * Show form create new product
  */
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	templates, err := template.ParseFiles(
    "templates/admin/master.html",
    fmt.Sprintf("templates/admin/add_product.html"),		
  )

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := templates.ExecuteTemplate(w, "master", ""); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

/**
  * Admin create new product
  */
func StoreProduct(w http.ResponseWriter, r *http.Request) {
  product := new(models.Product)
  
  errMap := utils.MapFormValues(product, r); if errMap != nil {
    log.Fatal(errMap)
    fmt.Fprintln(w, errMap);
    return
  }
  
  errCreate := models.CreateProduct(product); if errCreate != nil {
    fmt.Fprintln(w, errCreate);
    return
  } else {
    fmt.Fprintln(w, "Create Product Success")
  }
}

/**
  * Show form product's edit 
  */
func EditProduct(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 32);
  product, err := models.GetProduct(uint(id)); if err != nil {
    fmt.Fprintln(w, err);
    return
  }
  fmt.Fprintln(w, product)
}

/**
  * User update product's infor
  */
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 32);
  product, err := models.GetProduct(uint(id)); if err != nil {
    fmt.Fprintln(w, err);
    return
  }
  
  errMap := utils.MapFormValues(&product, r); if errMap != nil {
    log.Fatal(errMap)
    fmt.Fprintln(w, errMap);
    return
  }
  errUpdate := models.UpdateProduct(&product); if errUpdate != nil {
    fmt.Fprintln(w, errUpdate)
    return
  }
  fmt.Fprintln(w, "Update this product success")
}

/**
  * Delete product
  */
func DestroyProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 32);
  err := models.DeleteProduct(uint(id)); if err != nil {
    fmt.Fprintln(w, err);
    return
  } else {
    fmt.Fprintln(w, "Delete this product success")
  }
}
