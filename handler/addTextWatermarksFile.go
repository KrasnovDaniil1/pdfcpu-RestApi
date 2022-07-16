package handler

import (
	"mime/multipart"
	"net/http"
	"strings"

	// "strings"

	"github.com/gin-gonic/gin"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	pdfcpuType "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

var err error
var inFile *multipart.FileHeader
var outFile string
var selectedPages []string = nil
var onTop bool = false
var text string
var desc string
var conf *pdfcpuType.Configuration

func AddTextWatermarksFileRestApi(c *gin.Context) {

	inFile, err = c.FormFile("inFile")
	outFile = c.PostForm("outFile")
	if c.PostForm("selectedPages") != "all" {
		selectedPages = strings.Split(c.PostForm("selectedPages"), ",")
	}

	if c.PostForm("onTop") == "true" {
		onTop = true
	}
	text = c.PostForm("text")

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить файл."})
		return
	}

	err = c.SaveUploadedFile(inFile, "./file/"+inFile.Filename)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить файл."})
		return
	}

	err = pdfcpu.AddTextWatermarksFile("./file/"+inFile.Filename, "./file/"+outFile+".pdf", selectedPages, onTop, text, desc, conf)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Не удалось прочитать файл."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"error": "Нету"})

}
