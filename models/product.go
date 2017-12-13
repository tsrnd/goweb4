package models

import (
	"fmt"
	"time"

	"github.com/goweb4/database"
)

type Product struct {
	Model
	Size          string         `schema:"size"`
	Color         string         `schema:"color"`
	Price         float64        `schema:"price"`
	Name          string         `schema:"name"`
	InStock       uint           `schema:"in_stock"`
	GroupID       uint           `schema:"group_id"`
	ProductGroup  *ProductGroup  //belong To Product Group
	OrderProducts []OrderProduct //has many order products
	Images        []Image        //has many image
}

func (product *Product) GetSchema() map[string]interface{} {
	return map[string]interface{}{
		"id":         &product.ID,
		"size":       &product.Size,
		"color":      &product.Color,
		"price":      &product.Price,
		"in_stock":   &product.InStock,
		"group_id":   &product.GroupID,
		"created_at": &product.CreatedAt,
		"updated_at": &product.UpdatedAt,
		"deleted_at": &product.DeletedAt,
		"name":       &product.Name,
	}
}

func (product *Product) TableName() string {
	return "products"
}

// func GetProducts() (products []Product) {
// 	err := database.DBCon.Find(&products).Error
// 	if err != nil {
// 		return products
// 	}
// 	return relationship
// }

// func GetProducts() (products []Product) {
// err := database.DBCon.Find(&products).Error
// if err != nil {
// 	return products
// }
// return products
// }

func GetProduct(id uint) (product Product, err error) {
	rows, err := database.DBCon.Db.
		Query("SELECT * FROM images WHERE product_id = $1", product.ID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		image := Image{Product: &product}
		err = rows.Scan(
			&image.ID, &image.Name, &image.URL, &image.CreatedAt, &image.UpdatedAt, &image.DeletedAt, &image.ProductId,
		)
		if err != nil {
			return
		}
		product.Images = append(product.Images, image)
	}
	return
}

// func UpdateProduct(product *Product) (errUpdate error) {
// 	errUpdate = database.DBCon.Save(&product).Error
// 	return errUpdate
// }

// func DeleteProduct(id uint) (errGet error) {
// 	product := Product{}
// 	product, errGet = GetProduct(id)
// 	if errGet == nil {
// 		errGet = database.DBCon.Delete(&product).Error
// 	}
// 	return errGet
// }

func CreateProduct(product *Product) (proId uint, err error) {
	fmt.Println(product)
	err = database.DBCon.Db.
		QueryRow(
			"INSERT INTO products (size, color, price, in_stock, group_id, created_at, name) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;",
			product.Size, product.Color, product.Price, product.InStock,
			product.GroupID, time.Now(), product.Name).
		Scan(&product.ID)
	return product.ID, err
}

func GetTrendProducts(limit int) (listProduct []Product) {
	rows, err := database.DBCon.Db.Query("SELECT product_id, quantity from order_products ORDER BY quantity DESC limit $1", limit)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id, quantity uint
		product := Product{}
		if err := rows.Scan(&id, &quantity); err != nil {
			fmt.Println(err)
			return
		}
		database.DBCon.Db.QueryRow("SELECT id, size, name, price from products where id = $1", id).
			Scan(&product.ID, &product.Size, &product.Name, &product.Price)
		listProduct = append(listProduct, product)
	}
	return listProduct
}

func GetLatestProduct(limit int) (products []Product) {
	rows, err := database.DBCon.Db.Query("SELECT id, size, name, price from products ORDER BY id DESC limit $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.ID, &product.Size, &product.Name, &product.Price)
		if err != nil {
			return
		}
		products = append(products, product)
	}
	return products
}
