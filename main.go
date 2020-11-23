package main

import (
  "github.com/gin-gonic/gin"
  "simple_restfull_api/models"
  "simple_restfull_api/controllers"
)

func main() {

  r := gin.Default()
  
  
  // Connect to database
  models.Connection()

  //Routes
   r.GET("/books/add", controllers.AddBooks)
  r.GET("/books", controllers.FindBooks)
  r.GET("/books/find/:id", controllers.FindBook)
  // r.POST("/books/create", controllers.CreateBook) //json_data
  r.POST("/books/create-form", controllers.CreateBookForm) //x-www-form-urlencoded
	r.PATCH("/books/update/:id", controllers.UpdateBook)
	r.DELETE("/books/delete/:id", controllers.DeleteBook)


	r.Run()

  //db.First(&product, 1) // find product with integer primary key
  //db.First(&product, "code = ?", "D42") // find product with code D42

  // // Update - update product's price to 200
  // db.Model(&product).Update("Price", 200)
  // // Update - update multiple fields
  // db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  // db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // // Delete - delete product
  // db.Delete(&product, 1)
}