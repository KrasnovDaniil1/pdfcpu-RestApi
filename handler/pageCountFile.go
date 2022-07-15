package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func PageCountFileRestApi(c *gin.Context) {

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
	num, _ := pdfcpu.PageCountFile("./file/" + filePDF.Filename)
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"Количество страниц": num})

}
