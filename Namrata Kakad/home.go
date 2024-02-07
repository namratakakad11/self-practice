package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Company struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    CEO string  `json:"ceo"`
    Revenue  string `json:"revenue"`
}

var companies = []Company{
    {ID: "1", Name: "Dell", CEO: "Michael Dell", Revenue: "92.2 billion"},
    {ID: "2", Name: "Netflix", CEO: "Reed Hastings", Revenue: "20.2 billion"},
    {ID: "3", Name: "Microsoft", CEO: "Satya Nadella", Revenue: "320 million"},
}

func NewCompanyHandler(c *gin.Context) {
    var newCompany Company
    if err := c.ShouldBindJSON(&newCompany); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    newCompany.ID = xid.New().String()
    companies = append(companies, newCompany)
    c.JSON(http.StatusCreated,  newCompany)
}

func GetCompaniesHandler(c *gin.Context) {
    c.JSON(http.StatusOK, companies)
}

func main() {
    router := gin.Default()
	router.POST("/company", NewCompanyHandler)
	router.GET("/companies", GetCompaniesHandler)
    router.Run()
}