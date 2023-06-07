package controller

import (
	"contractmaster/models"
	sr "contractmaster/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/TecXLab/libhttp"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"github.com/rs/zerolog"
)

type TypeSenseController interface {
	TypeSenseUpload(c *gin.Context)
	TypeSenseSearch(c *gin.Context)
	Search(c *gin.Context)
	TypeSenseDrop(c *gin.Context)
}

type typeSenseController struct {
	typesenseservice sr.TypeSenseService
}

func NewTypeSenseController(service sr.TypeSenseService) TypeSenseController {
	return &typeSenseController{
		typesenseservice: service,
	}
}

func (service *typeSenseController) TypeSenseUpload(c *gin.Context) {

	service.typesenseservice.DropCollection()

	service.typesenseservice.GetContractFromDB()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }
	c.JSON(http.StatusOK, "Uploaded successfully")
	// elapsed := time.Since(start)
	// str := fmt.Sprintf("%v", elapsed)
	// c.JSON(http.StatusOK, "data inserted in "+str)

	// zerologs.Info().Msg("Exiting GetContractFromDB Controller")
}

func (service *typeSenseController) TypeSenseDrop(c *gin.Context) {
	//Droping collection from typesense
	service.typesenseservice.DropCollection()

	c.JSON(http.StatusOK, "Collection Dropped Successfully")

}

func (service *typeSenseController) TypeSenseSearch(c *gin.Context) {
	libhttp.CoreHeader(c)
	var zerologs zerolog.Logger

	err := container.NamedResolve(&zerologs, "zerologs")
	if err != nil {
		panic("Log Lib Not Initialize" + err.Error())
	}

	zerologs.Info().Msg("In TypeSenseSearch Controller")
	var reqModel models.TypeSenseModel

	if err := c.ShouldBindJSON(&reqModel); err != nil {
		zerologs.Error().Err(err).Msg("Error binding json" + err.Error())
		fmt.Println(err)
	}

	q := reqModel.Q
	var URL string
	var length = len([]rune(q))
	if length <= 2 {
		URL = sr.BaseURL + "/collections/" + sr.CollectionName + "/documents/search?query_by=" + sr.Query_By1 + "&q=" + url.QueryEscape(q) + "&group_by=" + sr.Group_by + "&group_limit=" + sr.Group_Limit + "&per_page=" + sr.PerPage + "&sort_by=" + sr.Sort_by + "&use_cache=" + sr.Use_cache
	} else {
		URL = sr.BaseURL + "/collections/" + sr.CollectionName + "/documents/search?query_by=" + sr.Query_By + "&q=" + url.QueryEscape(q) + "&group_by=" + sr.Group_by + "&group_limit=" + sr.Group_Limit + "&per_page=" + sr.PerPage + "&sort_by=" + sr.Sort_by + "&use_cache=" + sr.Use_cache
	}
	fmt.Println(URL)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		zerologs.Error().Err(err).Msg("Http NewRequest Error" + err.Error())
		fmt.Print(err.Error())
	}
	req.Header.Add(sr.TypessenseKey, sr.TypessenseValue)
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
		fmt.Println(err)
	}

	// fmt.Println(resp.Body)

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
		fmt.Println(err)
	}
	// fmt.Println(jsonString)

	var responseModel models.TypeSenseResponse

	err = json.Unmarshal(jsonString, &responseModel)
	if err != nil {
		zerologs.Error().Err(err).Msg("Error while unmarshaling response model" + err.Error())
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, responseModel) //Direct Response

	zerologs.Info().Msg("Ending TypeSenseSearch Controller")

}

// func (service *typeSenseController) TypeSenseSearch(c *gin.Context) {
// 	libhttp.CoreHeader(c)
// 	var zerologs zerolog.Logger

// 	err := container.NamedResolve(&zerologs, "zerologs")
// 	if err != nil {
// 		panic("Log Lib Not Initialize" + err.Error())
// 	}

// 	zerologs.Info().Msg("In TypeSenseSearch Controller")
// 	var reqModel models.TypeSenseModel

// 	if err := c.ShouldBindJSON(&reqModel); err != nil {
// 		zerologs.Error().Err(err).Msg("Error binding json" + err.Error())
// 		fmt.Println(err)
// 	}

// 	q := reqModel.Q
// 	page := "1"
// 	//q := strings.ReplaceAll(reqModel.Q, " ", "")

// 	//req, err := http.NewRequest("GET", "https://www.nuuu.com/collections/companymasters/documents/search?query_by=compname,s_name,isin&q="+cname, nil)

// 	//URL := sr.BaseURL + "/collections/" + sr.CollectionName + "/documents/search?query_by=" + sr.Query_By + "&q=" + url.QueryEscape(q) + "&group_by=" + sr.Group_by + "&group_limit=" + sr.Group_Limit + "&per_page=" + sr.PerPage + "&sort_by=" + sr.Sort_by
// 	//URL := sr.BaseURL + "/collections/" + sr.CollectionName + "/documents/search?query_by=" + sr.Query_By + "&q=" + url.QueryEscape(q) + "&per_page=" + sr.PerPage + "&sort_by=" + sr.Sort_by

// 	URL := sr.BaseURL + "/collections/" + sr.CollectionName + "/documents/search?exhaustive_search=" + sr.Exhaustive_Search + "&page=" + page + "&query_by=" + sr.Query_By + "&q=" + url.QueryEscape(q) + "&per_page=" + sr.PerPage + "&sort_by=" + sr.Sort_by + "&use_cache=true"
// 	//https://www.nuuu.com/collections/{{collection_name}}/documents/search?exhaustive_search=true&page=1&query_by=cnam,fullname,exmnt,strikprc,optyp,optyp2,seris,expry,usym&q=nifty 50&per_page=250&sort_by=nInstrumentType:asc,nexpry:asc&infix=always
// 	fmt.Println(URL)

// 	req, err := http.NewRequest("GET", URL, nil)
// 	if err != nil {
// 		zerologs.Error().Err(err).Msg("Http NewRequest Error" + err.Error())
// 		fmt.Print(err.Error())
// 	}
// 	req.Header.Add(sr.TypessenseKey, sr.TypessenseValue)
// 	// tr := &http.Transport{
// 	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	// }
// 	// client := &http.Client{Transport: tr}
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
// 		fmt.Println(err)
// 	}

// 	// fmt.Println(resp.Body)

// 	jsonString, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
// 		fmt.Println(err)
// 	}
// 	// fmt.Println(jsonString)

// 	// var responseModel models.TypeSenseResponse

// 	// err = json.Unmarshal(jsonString, &responseModel)
// 	// if err != nil {
// 	// 	zerologs.Error().Err(err).Msg("Error while unmarshaling response model" + err.Error())
// 	// 	fmt.Println(err)
// 	// }

// 	// c.JSON(http.StatusOK, responseModel) //Direct Response

// 	if resp.StatusCode == 503 {
// 		c.String(503, "SERVICE UNAVAILABE")
// 	} else {
// 		c.String(200, string(jsonString))
// 	}

// 	zerologs.Info().Msg("Ending TypeSenseSearch Controller")

// }

func (service *typeSenseController) Search(c *gin.Context) {
	libhttp.CoreHeader(c)
	var zerologs zerolog.Logger

	err := container.NamedResolve(&zerologs, "zerologs")
	if err != nil {
		panic("Log Lib Not Initialize" + err.Error())
	}

	zerologs.Info().Msg("In TypeSenseSearch Controller")
	var reqModel models.TypeSenseModel

	if err := c.ShouldBindJSON(&reqModel); err != nil {
		zerologs.Error().Err(err).Msg("Error binding json" + err.Error())
		fmt.Println(err)
	}

	q := reqModel.Q
	page := "1"
	//q := strings.ReplaceAll(reqModel.Q, " ", "")

	//req, err := http.NewRequest("GET", "https://www.nuuu.com/collections/companymasters/documents/search?query_by=compname,s_name,isin&q="+cname, nil)

	//URL := sr.BaseURL + "/collections/" + sr.CollectionName + "/documents/search?query_by=" + sr.Query_By + "&q=" + url.QueryEscape(q) + "&group_by=" + sr.Group_by + "&group_limit=" + sr.Group_Limit + "&per_page=" + sr.PerPage + "&sort_by=" + sr.Sort_by
	//URL := sr.BaseURL + "/collections/" + sr.CollectionName + "/documents/search?query_by=" + sr.Query_By + "&q=" + url.QueryEscape(q) + "&per_page=" + sr.PerPage + "&sort_by=" + sr.Sort_by

	URL := sr.BaseURL + "/collections/" + sr.CollectionName + "/documents/search?exhaustive_search=" + sr.Exhaustive_Search + "&page=" + page + "&query_by=" + sr.Query_By + "&q=" + url.QueryEscape(q) + "&per_page=" + sr.PerPage + "&sort_by=" + sr.Sort_by + "&use_cache=" + sr.Use_cache
	//https://www.nuuu.com/collections/{{collection_name}}/documents/search?exhaustive_search=true&page=1&query_by=cnam,fullname,exmnt,strikprc,optyp,optyp2,seris,expry,usym&q=nifty 50&per_page=250&sort_by=nInstrumentType:asc,nexpry:asc&infix=always
	fmt.Println(URL)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		zerologs.Error().Err(err).Msg("Http NewRequest Error" + err.Error())
		fmt.Print(err.Error())
	}
	req.Header.Add(sr.TypessenseKey, sr.TypessenseValue)
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
		fmt.Println(err)
	}

	// fmt.Println(resp.Body)

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
		fmt.Println(err)
	}
	// fmt.Println(jsonString)

	// var responseModel models.TypeSenseResponse

	// err = json.Unmarshal(jsonString, &responseModel)
	// if err != nil {
	// 	zerologs.Error().Err(err).Msg("Error while unmarshaling response model" + err.Error())
	// 	fmt.Println(err)
	// }

	// c.JSON(http.StatusOK, responseModel) //Direct Response

	if resp.StatusCode == 503 {
		c.String(503, "SERVICE UNAVAILABE")
	} else {
		c.String(200, string(jsonString))
	}

	zerologs.Info().Msg("Ending TypeSenseSearch Controller")

}
