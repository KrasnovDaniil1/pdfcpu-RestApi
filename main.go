package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Use(cors.Default())
	router.POST("/pagenum", collectRest)
	// router.GET("/cards", handlers.GetCards)
	// router.DELETE("/cards/:id", handlers.DeleteCard)
	// router.GET("/tags", handlers.GetTags)

	router.Run("localhost:8080")
}

func collectRest(c *gin.Context) {

	filePDF, err := c.FormFile("file")

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить файл."})
		return
	}
	err = c.SaveUploadedFile(filePDF, "./file/"+filePDF.Filename)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить файл."})
		return
	}
	// cmd.Collect()
	num, _ := pdfcpu.PageCountFile("./file/" + filePDF.Filename)
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"Количество страниц": num})

}

// func saveImage(c *gin.Context, image *multipart.FileHeader) error {
// 	return err
// }
