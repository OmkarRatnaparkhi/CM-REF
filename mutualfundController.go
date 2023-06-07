package controller

import (
	"contractmaster/helper"
	"contractmaster/models"
	sr "contractmaster/service"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/TecXLab/libhttp"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"github.com/rs/zerolog"
)

type MfController interface {
	MutualfundUpload(c *gin.Context)
	MutualfundDrop(c *gin.Context)
	MutualfundSearch(c *gin.Context)
	MutualfundUploadData(c *gin.Context)

	MfSearch(c *gin.Context)
	MfSearchGroup(c *gin.Context)
	MfSearchFacet(c *gin.Context)
	MfSearchfacetfilter(c *gin.Context)
	MasterMfDrop(c *gin.Context)
	MasterMfUploadData(c *gin.Context)
	Searchschcode(c *gin.Context)
	PaymentGateWay(c *gin.Context)
}

type mfController struct {
	mfservice sr.MfService
}

func NewMfController(service sr.MfService) MfController {
	return &mfController{
		mfservice: service,
	}
}

func (service *mfController) PaymentGateWay(c *gin.Context) {
	var request interface{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(request)
	c.JSON(200, "")
}
func (service *mfController) MutualfundDrop(c *gin.Context) {
	service.mfservice.MutualfundDrop()

	c.JSON(http.StatusOK, "Collection Dropped Successfully")
}

func (service *mfController) MutualfundUpload(c *gin.Context) {

	service.mfservice.MutualfundDrop()

	//service.mfservice.GetMfData()
	service.mfservice.GetMfDataOptimized()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }
	c.JSON(http.StatusOK, "")
	// elapsed := time.Since(start)
	// str := fmt.Sprintf("%v", elapsed)
	// c.JSON(http.StatusOK, "data inserted in "+str)

	// zerologs.Info().Msg("Exiting GetContractFromDB Controller")
}

func (service *mfController) MutualfundUploadData(c *gin.Context) {

	service.mfservice.MutualfundDrop()

	service.mfservice.GetMutualFund()

	c.JSON(http.StatusOK, "Successfully Upload Data")

}

func (service *mfController) MutualfundSearch(c *gin.Context) {
	var zerologs zerolog.Logger

	err := container.NamedResolve(&zerologs, "zerologs")
	if err != nil {
		panic("Log Lib Not Initialize" + err.Error())
	}

	zerologs.Info().Msg("In MutualfundSearch Controller")
	var reqModel models.TypeSenseModel
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		zerologs.Error().Err(err).Msg("Error binding json" + err.Error())
		fmt.Println(err)
	}

	q := reqModel.Q
	//https: //www.nuinvstr.com/collections/greenwaremf/documents/search?query_by=SchemeName,AMCName&per_page=250&q=NIPPON INDIA NIVESH LAKSHYA FUND - IDCW REINVESTMENT&exhaustive_search=true
	URL := sr.BaseURL + "/collections/" + helper.MfCollectionName + "/documents/search?query_by=" + sr.MfQuery_By + "&per_page=" + sr.PerPage + "&q=" + url.QueryEscape(q) + "&exhaustive_search=" + sr.Exhaustive_Search + "&use_cache=" + sr.Use_cache
	//fmt.Println(URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add(sr.TypessenseKey, sr.TypessenseValue)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	/*
		jsonString, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var responseModel models.MutualfundResponse

		err = json.Unmarshal(jsonString, &responseModel)
		if err != nil {
			zerologs.Error().Err(err).Msg("Error while unmarshaling response model" + err.Error())
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, responseModel) //Direct Response
	*/

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	c.String(200, string(jsonString))

	zerologs.Info().Msg("Ending MutualfundSearch Controller")
}

func (service *mfController) MasterMfDrop(c *gin.Context) {
	service.mfservice.MasterDataMFDrop()

	c.JSON(http.StatusOK, "Collection Dropped Successfully")
}

func (service *mfController) MasterMfUploadData(c *gin.Context) {

	service.mfservice.MasterDataMFDrop()

	service.mfservice.GetMasterDataMF()

	c.JSON(http.StatusOK, "Data uploaded Successfully")

}

/*
func (service *mfController) MfSearch(c *gin.Context) {
	var zerologs zerolog.Logger

	err := container.NamedResolve(&zerologs, "zerologs")
	if err != nil {
		panic("Log Lib Not Initialize" + err.Error())
	}

	zerologs.Info().Msg("In MutualfundSearch Controller")
	var reqModel models.TypeSenseModel
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		zerologs.Error().Err(err).Msg("Error binding json" + err.Error())
		fmt.Println(err)
	}

	q := reqModel.Q
	//https://www.nuuu.com/collections/greenwaremf2/documents/search?query_by=SchemeName,AMCName&per_page=250&q=IDFC Nifty100 Low Volatility 30 Index Fund - Regular Plan - IDCW&exhaustive_search=true
	URL := helper.BaseURL + "/collections/" + helper.MfCollectionName2 + "/documents/search?query_by=" + helper.MfQuery_By + "&per_page=" + helper.PerPage + "&q=" + url.QueryEscape(q) + "&exhaustive_search=" + helper.Exhaustive_Search
	fmt.Println(URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add(helper.TypessenseKey, helper.TypessenseValue)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Body)
	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var responseModel models.MasterMfResponse

	err = json.Unmarshal(jsonString, &responseModel)
	if err != nil {
		zerologs.Error().Err(err).Msg("Error while unmarshaling response model" + err.Error())
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, responseModel) //Direct Response

	zerologs.Info().Msg("Ending MutualfundSearch Controller")
}
*/

func (service *mfController) MfSearch(c *gin.Context) {
	libhttp.CoreHeader(c)
	var zerologs zerolog.Logger

	err := container.NamedResolve(&zerologs, "zerologs")
	if err != nil {
		panic("Log Lib Not Initialize" + err.Error())
	}

	zerologs.Info().Msg("In MutualfundSearch Controller")
	var reqModel models.MftypeSenseModel
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		zerologs.Error().Err(err).Msg("Error binding json" + err.Error())
		fmt.Println(err)
	}

	q := reqModel.Q
	isnrmlsrch := reqModel.Isnrmlsrch
	isnrmlflag, err := strconv.Atoi(isnrmlsrch)
	if err != nil {
		fmt.Println(err)
		c.String(400, string(err.Error()))
	}

	page := reqModel.Page

	perpage, _ := strconv.Atoi(reqModel.Per_page)
	var per_page string
	if perpage > 250 {
		per_page = "250"
	} else if reqModel.Per_page == "" {
		per_page = "250"
	} else {
		per_page = reqModel.Per_page
	}

	var Filter_by string
	//	Filter_by := "SIPAllowed:=" + sipallowed + "&&LumpSumAllowed:=" + lumpsumallowed + "&&AMCID:[1015,1041]&&CategoryCode:[75,65]&&Risk:=[Low,High]"

	if reqModel.SIPAllowed != "" {
		Filter_by += "SIPAllowed:=" + reqModel.SIPAllowed
	}
	if reqModel.Lumpsumallowed != "" {
		Filter_by += "&&LumpSumAllowed:=" + reqModel.Lumpsumallowed
	}
	if reqModel.AMCID != "" {
		Filter_by += "&&AMCID:[" + reqModel.AMCID + "]"
	}
	if reqModel.CategoryCode != "" {
		Filter_by += "&&CategoryCode:[" + reqModel.CategoryCode + "]"
	}
	if reqModel.Category != "" {
		Filter_by += "&&Category:[" + reqModel.Category + "]"
	}
	if reqModel.Risk != "" {
		Filter_by += "&&Risk:=[" + reqModel.Risk + "]"
	}
	if reqModel.SchemeAUM != "" {
		Filter_by += "&&SchemeAUM:<=" + reqModel.SchemeAUM
	}

	//	fmt.Println(Filter_by)

	var URL string
	if isnrmlflag == 1 {
		URL = sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By + "&per_page=" + per_page + "&q=" + url.QueryEscape(q) + "&exhaustive_search=" + sr.Exhaustive_Search + "&use_cache=" + sr.Use_cache
	} else if reqModel.Include_fields != "" {
		includefields := reqModel.Include_fields
		URL = sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By + "&per_page=" + per_page + "&q=" + url.QueryEscape(q) + "&page=" + page + "&filter_by=" + url.QueryEscape(Filter_by) + "&sort_by=" + sr.Mfsort_by + "&include_fields=" + includefields + "&use_cache=" + sr.Use_cache
	} else {
		URL = sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By + "&per_page=" + per_page + "&q=" + url.QueryEscape(q) + "&page=" + page + "&filter_by=" + url.QueryEscape(Filter_by) + "&sort_by=" + sr.Mfsort_by + "&use_cache=" + sr.Use_cache
	}

	//fmt.Println(URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
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
		fmt.Println(err)
		c.String(400, string(err.Error()))
	}

	/*
		jsonString, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var responseModel models.MasterMfResponse

		err = json.Unmarshal(jsonString, &responseModel)
		if err != nil {
			zerologs.Error().Err(err).Msg("Error while unmarshaling response model" + err.Error())
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, responseModel) //Direct Response
	*/

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		c.String(400, string(err.Error()))
	}

	if resp.StatusCode == 503 {
		c.String(503, "SERVICE UNAVAILABE")
	} else {
		c.String(200, string(jsonString))
	}

	zerologs.Info().Msg("Ending MutualfundSearch Controller")

}

func (service *mfController) MfSearchGroup(c *gin.Context) {
	libhttp.CoreHeader(c)
	var zerologs zerolog.Logger

	err := container.NamedResolve(&zerologs, "zerologs")
	if err != nil {
		panic("Log Lib Not Initialize" + err.Error())
	}

	zerologs.Info().Msg("In MutualfundSearch Controller")
	var reqModel models.MfGroupModel
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		zerologs.Error().Err(err).Msg("Error binding json" + err.Error())
		fmt.Println(err)
	}

	q := reqModel.Q
	isnrmlsrch := reqModel.Isnrmlsrch
	isnrmlflag, err := strconv.Atoi(isnrmlsrch)
	if err != nil {
		fmt.Println(err)
		c.String(400, string(err.Error()))
	}

	page := reqModel.Page

	perpage, _ := strconv.Atoi(reqModel.Per_page)
	var per_page string
	if perpage > 250 {
		per_page = "250"
	} else if reqModel.Per_page == "" {
		per_page = "250"
	} else {
		per_page = reqModel.Per_page
	}

	group_by := reqModel.Group_by
	grouplimit := reqModel.Group_limit

	var Filter_by string
	//	Filter_by := "SIPAllowed:=" + sipallowed + "&&LumpSumAllowed:=" + lumpsumallowed + "&&AMCID:[1015,1041]&&CategoryCode:[75,65]&&Risk:=[Low,High]"

	if reqModel.SIPAllowed != "" {
		Filter_by += "SIPAllowed:=" + reqModel.SIPAllowed
	}
	if reqModel.Lumpsumallowed != "" {
		Filter_by += "&&LumpSumAllowed:=" + reqModel.Lumpsumallowed
	}
	if reqModel.AMCID != "" {
		Filter_by += "&&AMCID:[" + reqModel.AMCID + "]"
	}
	if reqModel.CategoryCode != "" {
		Filter_by += "&&CategoryCode:[" + reqModel.CategoryCode + "]"
	}
	if reqModel.Category != "" {
		Filter_by += "&&Category:[" + reqModel.Category + "]"
	}
	if reqModel.Risk != "" {
		Filter_by += "&&Risk:=[" + reqModel.Risk + "]"
	}
	if reqModel.SchemeAUM != "" {
		Filter_by += "&&SchemeAUM:<=" + reqModel.SchemeAUM
	}

	//	fmt.Println(Filter_by)

	var URL string
	if isnrmlflag == 1 {
		//URL = sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By + "&per_page=" + per_page + "&q=" + url.QueryEscape(q) + "&exhaustive_search=" + sr.Exhaustive_Search + "&use_cache=" + sr.Use_cache
		URL = sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By + "&per_page=" + per_page + "&q=" + url.QueryEscape(q) + "&exhaustive_search=" + sr.Exhaustive_Search + "&group_by=" + group_by + "&group_limit=" + grouplimit + "&use_cache=" + sr.Use_cache
	} else if reqModel.Include_fields != "" {
		includefields := reqModel.Include_fields
		URL = sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By + "&per_page=" + per_page + "&q=" + url.QueryEscape(q) + "&page=" + page + "&filter_by=" + url.QueryEscape(Filter_by) + "&sort_by=" + sr.Mfsort_by + "&include_fields=" + includefields + "&group_by=" + group_by + "&group_limit=" + grouplimit + "&use_cache=" + sr.Use_cache
	} else {
		URL = sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By + "&per_page=" + per_page + "&q=" + url.QueryEscape(q) + "&page=" + page + "&filter_by=" + url.QueryEscape(Filter_by) + "&sort_by=" + sr.Mfsort_by + "&group_by=" + group_by + "&group_limit=" + grouplimit + "&use_cache=" + sr.Use_cache
	}

	fmt.Println(URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add(sr.TypessenseKey, sr.TypessenseValue)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		c.String(400, string(err.Error()))
	}

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		c.String(400, string(err.Error()))
	}

	if resp.StatusCode == 503 {
		c.String(503, "SERVICE UNAVAILABE")
	} else {
		c.String(200, string(jsonString))
	}

	zerologs.Info().Msg("Ending MutualfundSearch Controller")

}

func (service *mfController) MfSearchFacet(c *gin.Context) {
	libhttp.CoreHeader(c)
	var zerologs zerolog.Logger
	err := container.NamedResolve(&zerologs, "zerologs")
	if err != nil {
		panic("Log Lib Not Initialize" + err.Error())
	}
	zerologs.Info().Msg("In MfSearchFacet Controller")
	var reqModel models.MftypeSenseModel
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		zerologs.Error().Err(err).Msg("Error binding json" + err.Error())
		fmt.Println(err)
	}
	q := reqModel.Q
	perpage, _ := strconv.Atoi(reqModel.Per_page)
	var per_page string
	if perpage > 250 {
		per_page = "250"
	} else if reqModel.Per_page == "" {
		per_page = "250"
	} else {
		per_page = reqModel.Per_page
	}
	//https://www.nuuu.com/collections/greenwaremf2/documents/search?query_by=SchemeName,AMCName,SchCode&per_page=250&q=*&page=1&sort_by=Return3Year:desc&facet_by=AMCName,Category,Risk,SchemeAUM,SIPAllowed,LumpSumAllowed&exhaustive_search=true
	var URL string
	URL = sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfsearchfacetQuery_By + "&per_page=" + per_page + "&q=" + url.QueryEscape(q) + "&sort_by=" + sr.Mfsort_by + "&facet_by=" + sr.MfFacet_by + "&exhaustive_search=" + sr.Exhaustive_Search + "&use_cache=" + sr.Use_cache
	//fmt.Println(URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		zerologs.Error().Err(err).Msg("Error in http.NewRequest" + err.Error())
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
		fmt.Println(err)
	}
	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode == 503 {
		c.String(503, "SERVICE UNAVAILABE")
	} else {
		c.String(200, string(jsonString))
	}
	zerologs.Info().Msg("Ending MfSearchFacet Controller")
}

func (service *mfController) Searchschcode(c *gin.Context) {
	libhttp.CoreHeader(c)
	var zerologs zerolog.Logger

	err := container.NamedResolve(&zerologs, "zerologs")
	if err != nil {
		panic("Log Lib Not Initialize" + err.Error())
	}

	zerologs.Info().Msg("In MutualfundSearch Controller")
	var reqModel models.SchcodeModel
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		zerologs.Error().Err(err).Msg("Error binding json" + err.Error())
		fmt.Println(err)
	}

	q := reqModel.Q

	var Filter_by string
	var URL string

	if reqModel.Mulschcode != "" {
		Filter_by += "SchCode:=[" + reqModel.Mulschcode + "]"
		URL = sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By2 + "&q=" + url.QueryEscape(q) + "&filter_by=" + url.QueryEscape(Filter_by) + "&use_cache=" + sr.Use_cache
	} else {
		URL = sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By2 + "&q=" + url.QueryEscape(q) + "&use_cache=" + sr.Use_cache
	}

	//fmt.Println(URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
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
		fmt.Println(err)
	}

	/*
		jsonString, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		var responseModel models.MasterMfResponse
		err = json.Unmarshal(jsonString, &responseModel)
		if err != nil {
			zerologs.Error().Err(err).Msg("Error while unmarshaling response model" + err.Error())
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, responseModel) //Direct Response
	*/

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode == 503 {
		c.String(503, "SERVICE UNAVAILABE")
	} else {
		c.String(200, string(jsonString))
	}

	zerologs.Info().Msg("Ending MutualfundSearch Controller")

}

func (service *mfController) MfSearchfacetfilter(c *gin.Context) {
	libhttp.CoreHeader(c)
	var zerologs zerolog.Logger

	err := container.NamedResolve(&zerologs, "zerologs")
	if err != nil {
		panic("Log Lib Not Initialize" + err.Error())
	}

	zerologs.Info().Msg("In MutualfundSearch Controller")
	var reqModel models.MffacetfilterMod
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		zerologs.Error().Err(err).Msg("Error binding json" + err.Error())
		fmt.Println(err)
	}

	q := reqModel.Q
	page := reqModel.Page
	var Filter_by string

	if reqModel.SIPAllowed != "" {
		Filter_by += "SIPAllowed:=" + reqModel.SIPAllowed
	}
	if reqModel.Lumpsumallowed != "" {
		Filter_by += "&&LumpSumAllowed:=" + reqModel.Lumpsumallowed
	}
	if reqModel.AMCName != "" {
		Filter_by += "&&AMCName:=[" + reqModel.AMCName + "]"
	}
	if reqModel.Category != "" {
		Filter_by += "&&Category:[" + reqModel.Category + "]"
	}
	if reqModel.Risk != "" {
		Filter_by += "&&Risk:=[" + reqModel.Risk + "]"
	}
	if reqModel.SchemeAUM != "" {
		Filter_by += "&&SchemeAUM:<=" + reqModel.SchemeAUM
	}

	perpage, _ := strconv.Atoi(reqModel.Per_page)
	var per_page string
	if perpage > 250 {
		per_page = "250"
	} else if reqModel.Per_page == "" {
		per_page = "250"
	} else {
		per_page = reqModel.Per_page
	}

	//fmt.Println(Filter_by)
	URL := sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By + "&per_page=" + per_page + "&q=" + url.QueryEscape(q) + "&page=" + page + "&sort_by=" + sr.Mfsort_by + "&exhaustive_search=" + sr.Exhaustive_Search + "&filter_by=" + url.QueryEscape(Filter_by) + "&use_cache=" + sr.Use_cache
	//URL := sr.BaseURL + "/collections/" + sr.MfCollectionName2 + "/documents/search?query_by=" + sr.MfQuery_By + "&per_page=" + per_page + "&q=" + url.QueryEscape(q) + "&page=" + page + "&sort_by=" + sr.Mfsort_by + "&facet_by=" + sr.MfFacet_by + "&exhaustive_search=" + sr.Exhaustive_Search + "&filter_by=" + url.QueryEscape(Filter_by)
	//fmt.Println(URL)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Print(err.Error())
		c.String(400, string(err.Error()))
	}
	req.Header.Add(sr.TypessenseKey, sr.TypessenseValue)
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 503 {
		c.String(503, "SERVICE UNAVAILABE")
	} else {
		c.String(200, string(jsonString))
	}

	zerologs.Info().Msg("Ending MutualfundSearch Controller")

}
