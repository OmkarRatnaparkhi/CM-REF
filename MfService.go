package service

import (
	"bytes"
	"contractmaster/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

type MfService interface {
	GetMfData()
	GetMfDataOptimized()
	MutualfundDrop()
	GetMutualFund()

	MasterDataMFDrop()
	GetMasterDataMF()
}

type mfService struct {
}

func NewMfService() MfService {
	return &mfService{}
}

var mfcollectionname = "greenwaremf"
var mfcollectionname2 = "greenwaremf2"

func (repo *mfService) MutualfundDrop() {

	client := typesense.NewClient(
		typesense.WithServer(BaseURL),
		typesense.WithAPIKey(TypessenseValue))

	client.Collection(mfcollectionname).Delete()

	Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Sucessfully collection dropped from Typesense. ")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Sucessfully collection dropped from Typesense. ")
}

func uploadMfDatatypesense(objData interface{}) error {

	Zerologs.Info().Msg("uploadMfDatatypesense Start")
	client := typesense.NewClient(
		typesense.WithServer(BaseURL),
		typesense.WithAPIKey(TypessenseValue))

	_, err := client.Collection(mfcollectionname).Retrieve()
	if err != nil {
		bfacet := true
		schema := &api.CollectionSchema{
			Name: mfcollectionname,
			Fields: []api.Field{
				{
					Name:  "AMCID",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "AMCName",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "ISIN",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchemeCode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchemeName",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "BSEToken",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "LatestNAV",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "NAVDate",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "NAVChange",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Risk",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Category",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchCode",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "Objective",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "TranMode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SIPFrequency",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SIPDates",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMinGAP",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMaxGAP",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMinInstallmentAmount",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMinInstallmentNo",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMultiplierAmt",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMaxInstallmentAmount",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMaxInstallmentNo",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "ISSIP",
					Type:  "bool",
					Facet: &bfacet,
				},
				{
					Name:  "ISLUMPSUM",
					Type:  "bool",
					Facet: &bfacet,
				},
			},
			//DefaultSortingField: &sort1,
		}
		_, err := client.Collections().Create(schema)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(resp)
	}

	objDocModel := objData.(models.MfDocumentModel)
	document := objDocModel

	_, err = client.Collection(mfcollectionname).Documents().Upsert(document)
	if err != nil {
		Zerologs.Error().Err(err).Msgf("Error in uploadtypesenseData")
		fmt.Println(err)
		return err
	}
	return nil
}

func uploadMfDatatypesense1(objData []interface{}) error {
	Zerologs.Info().Msg("uploadMfDatatypesense Start")
	client := typesense.NewClient(
		typesense.WithServer(BaseURL),
		typesense.WithAPIKey(TypessenseValue))

	_, err := client.Collection(mfcollectionname).Retrieve()
	if err != nil {
		bfacet := true
		schema := &api.CollectionSchema{
			Name: mfcollectionname,
			Fields: []api.Field{
				{
					Name:  "AMCID",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "AMCName",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "ISIN",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchemeCode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchemeName",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "BSEToken",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "LatestNAV",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "NAVDate",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "NAVChange",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Risk",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Category",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchCode",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "Objective",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "TranMode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Type",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SIPFrequency",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SIPDates",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMinGAP",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMaxGAP",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMinInstallmentAmount",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMinInstallmentNo",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMultiplierAmt",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMaxInstallmentAmount",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMaxInstallmentNo",
					Type:  "int64",
					Facet: &bfacet,
				},
			},
			//DefaultSortingField: &sort1,
		}
		_, err := client.Collections().Create(schema)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(resp)
	}

	document := objData
	var action = "upsert"
	var batch = 40
	params := &api.ImportDocumentsParams{
		Action:    &action,
		BatchSize: &batch,
	}

	_, err = client.Collection(mfcollectionname).Documents().Import(document, params)
	if err != nil {
		Zerologs.Error().Err(err).Msgf("Error in uploadMfDatatypesense")
		fmt.Println(err)
		return err
	}
	return nil
}

func uploadMfDatatypesense2(objData []interface{}) error {
	Zerologs.Info().Msg("uploadMfDatatypesense Start")
	client := typesense.NewClient(
		typesense.WithServer(BaseURL),
		typesense.WithAPIKey(TypessenseValue))

	_, err := client.Collection(mfcollectionname).Retrieve()
	if err != nil {
		bfacet := true
		schema := &api.CollectionSchema{
			Name: mfcollectionname,
			Fields: []api.Field{
				{
					Name:  "AMCID",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "AMCName",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "ISIN",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchemeCode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchemeName",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "BSEToken",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "LatestNAV",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "NAVDate",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "NAVChange",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Risk",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Category",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchCode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Objective",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "TranMode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SIPFrequency",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SIPDates",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMinGAP",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMaxGAP",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMinInstallmentAmount",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMinInstallmentNo",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMultiplierAmt",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMaxInstallmentAmount",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "SIPMaxInstallmentNo",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "ISSIP",
					Type:  "bool",
					Facet: &bfacet,
				},
				{
					Name:  "ISLUMPSUM",
					Type:  "bool",
					Facet: &bfacet,
				},
			},
			//DefaultSortingField: &sort1,
		}
		_, err := client.Collections().Create(schema)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(resp)
	}

	document := objData
	var action = "upsert"
	var batch = 40
	params := &api.ImportDocumentsParams{
		Action:    &action,
		BatchSize: &batch,
	}

	_, err = client.Collection(mfcollectionname).Documents().Import(document, params)
	if err != nil {
		Zerologs.Error().Err(err).Msgf("Error in uploadMfDatatypesense")
		fmt.Println(err)
		return err
	}
	return nil
}

func Mflogin() string {

	const url = "https://fundcore-nu.azurewebsites.net/fcapi/account/login"

	var login models.Loginmodel
	login.LoginName = "admin"
	login.Password = "India@123"
	login.LoginUserType = 3

	jsonStr, err := json.Marshal(&login)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
		fmt.Println(err)
		//return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	if err != nil {
		//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
		fmt.Println(err)
		// return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "NI_ExternalApiUser:D6&(k-8rp#")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
		fmt.Println(err)
	}

	//fmt.Println(resp.Body)

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
		fmt.Println(err)
	}

	var respModel models.Loginresponse
	err = json.Unmarshal(jsonString, &respModel)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error while unmarshaling response model" + err.Error())
		fmt.Println(err)
	}

	//fmt.Println(respModel.Data.AccessToken)
	return respModel.Data.AccessToken
}

// Lumpsum order
// func (service *mfController) GetPurchaseorder(c *gin.Context) {
func GetPurchaseorder() interface{} {
	//Accesstoken := controller.Mflogin()
	Accesstoken := Mflogin()

	const url = "https://fundcore-nu.azurewebsites.net/fcapi/mf/scheme/list/for-purchase-order"

	var ReqBody models.PurchaseSiporderBody
	ReqBody.TranMode = ""
	//var PurchaseModel Purchse
	// err := c.BindJSON(&ReqBody)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	jsonStr, err := json.Marshal(&ReqBody)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
		fmt.Println(err)
		//return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	if err != nil {
		//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
		fmt.Println(err)
		// return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+Accesstoken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
		fmt.Println(err)
	}

	//fmt.Println(resp.Body)

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
		fmt.Println(err)
	}

	var respModel models.PurchaseSiporderResponse
	err = json.Unmarshal(jsonString, &respModel)
	if err != nil {
		fmt.Println(err)
	}

	return respModel
}

// SIP Order
// func (service *mfController) GetSiporder(c *gin.Context) {
// func GetSiporder() (Sipmodel models.PurchaseSiporderResponse) {
func GetSiporder() interface{} {
	//Accesstoken := controller.Mflogin()
	Accesstoken := Mflogin()
	const url = "https://fundcore-nu.azurewebsites.net/fcapi/mf/scheme/list/for-sip-order"

	var ReqBody models.PurchaseSiporderBody
	ReqBody.TranMode = ""
	//var PurchaseModel Purchse
	// err := c.BindJSON(&ReqBody)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	jsonStr, err := json.Marshal(&ReqBody)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
		fmt.Println(err)
		//return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
		fmt.Println(err)
		// return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+Accesstoken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
		fmt.Println(err)
	}

	//fmt.Println(resp.Body)

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
		fmt.Println(err)
	}

	var respModel models.PurchaseSiporderResponse
	err = json.Unmarshal(jsonString, &respModel)
	if err != nil {
		fmt.Println(err)
	}

	return respModel

}

// func (repo *mfService) GetMfData() []interface{} {
func (repo *mfService) GetMfData() {

	Purchasemodel := GetPurchaseorder()
	//fmt.Println(Purchasemodel)
	objpurchase := Purchasemodel.(models.PurchaseSiporderResponse)

	Sipmodel := GetSiporder()
	// fmt.Println(Sipmodel)
	objsip := Sipmodel.(models.PurchaseSiporderResponse)

	mapPurchasemodel := make(map[string]models.PurchaseSipDataModel)
	mapSipmodel := make(map[string]models.PurchaseSipDataModel)

	for i := 0; i < len(objpurchase.Data); i++ {
		mapPurchasemodel[objpurchase.Data[i].SchemeCode] = objpurchase.Data[i]
	}
	//fmt.Println(mapPurchasemodel)

	for i := 0; i < len(objsip.Data); i++ {
		mapSipmodel[objsip.Data[i].SchemeCode] = objsip.Data[i]
	}
	//fmt.Println(mapSipmodel)

	mapMfPurchaseInfoData := make(map[string]models.Scheme_Info)

	Accesstoken := Mflogin()

	count := 0

	for i, v := range mapPurchasemodel {
		//fmt.Println(v)
		//fmt.Println(v.SchemeCode)

		const url = "https://fundcore-nu.azurewebsites.net/fcapi/mf/scheme/info/for-purchase-order"

		var ReqBody models.PurchaseInfoBody
		ReqBody.BseSchemeCode = v.SchemeCode

		jsonStr, err := json.Marshal(&ReqBody)
		if err != nil {
			//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
			fmt.Println(err)
			//return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

		if err != nil {
			//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
			fmt.Println(err)
			// return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+Accesstoken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
			fmt.Println(err)
		}

		//fmt.Println(resp.Body)

		jsonString, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
			fmt.Println(err)
		}

		var respModel models.InfoPurchaseModel
		err = json.Unmarshal(jsonString, &respModel)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(respModel)

		mapMfPurchaseInfoData[v.SchemeCode] = respModel.Data.SchemeInfo[0]

		objDocModel := models.MfDocumentModel{}
		objDocModel.AMCID = mapMfPurchaseInfoData[v.SchemeCode].AMCID
		objDocModel.AMCName = mapMfPurchaseInfoData[v.SchemeCode].AMCName
		objDocModel.ISIN = mapMfPurchaseInfoData[v.SchemeCode].ISIN
		objDocModel.SchemeCode = mapMfPurchaseInfoData[v.SchemeCode].SchemeCode
		objDocModel.SchemeName = mapMfPurchaseInfoData[v.SchemeCode].SchemeName
		objDocModel.BSEToken = mapMfPurchaseInfoData[v.SchemeCode].BSEToken
		objDocModel.DivReFlag = mapMfPurchaseInfoData[v.SchemeCode].DivReFlag
		objDocModel.LatestNAV = mapMfPurchaseInfoData[v.SchemeCode].LatestNAV
		objDocModel.NAVDate = mapMfPurchaseInfoData[v.SchemeCode].NAVDate
		objDocModel.NAVChange = mapMfPurchaseInfoData[v.SchemeCode].NAVChange
		objDocModel.Risk = mapMfPurchaseInfoData[v.SchemeCode].Risk
		objDocModel.Category = mapMfPurchaseInfoData[v.SchemeCode].Category
		//objDocModel.SchCode = mapMfPurchaseInfoData[v.SchemeCode].SchCode
		schcode := mapMfPurchaseInfoData[v.SchemeCode].SchCode
		strschcode := strconv.FormatInt(schcode, 10)
		objDocModel.SchCode = strschcode
		objDocModel.SchParentCode = mapMfPurchaseInfoData[v.SchemeCode].SchParentCode
		objDocModel.MinPurAmt = mapMfPurchaseInfoData[v.SchemeCode].MinPurAmt
		objDocModel.PurAmountMult = mapMfPurchaseInfoData[v.SchemeCode].PurAmountMult
		objDocModel.MinAddPurAmt = mapMfPurchaseInfoData[v.SchemeCode].MinAddPurAmt
		objDocModel.Objective = mapMfPurchaseInfoData[v.SchemeCode].Objective
		objDocModel.TranMode = mapMfPurchaseInfoData[v.SchemeCode].TranMode
		//objDocModel.Type = "Lumpsum"

		uploadMfDatatypesense(objDocModel)

		count++
		fmt.Println(count, i)
	}

	count = 0
	mapMfSipInfoData := make(map[string]models.InfoSipModel)

	for i, v := range mapSipmodel {
		//Accesstoken := Mflogin()
		const url = "https://fundcore-nu.azurewebsites.net/fcapi/mf/scheme/info/for-sip-order"

		var ReqBody models.PurchaseInfoBody
		ReqBody.BseSchemeCode = v.SchemeCode

		jsonStr, err := json.Marshal(&ReqBody)
		if err != nil {
			//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
			fmt.Println(err)
			//return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

		if err != nil {
			//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
			fmt.Println(err)
			// return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+Accesstoken)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
			fmt.Println(err)
		}

		jsonString, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
			fmt.Println(err)
		}

		var respModel models.InfoSipModel
		err = json.Unmarshal(jsonString, &respModel)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(respModel)

		mapMfSipInfoData[v.SchemeCode] = respModel

		objDocModel := models.MfDocumentModel{}
		objDocModel.AMCID = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].AMCID

		objDocModel.AMCName = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].AMCName
		objDocModel.ISIN = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].ISIN
		objDocModel.SchemeCode = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].SchemeCode
		objDocModel.SchemeName = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].SchemeName
		objDocModel.BSEToken = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].BSEToken
		objDocModel.DivReFlag = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].DivReFlag
		objDocModel.LatestNAV = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].LatestNAV
		objDocModel.NAVDate = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].NAVDate
		objDocModel.NAVChange = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].NAVChange
		objDocModel.Risk = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].Risk
		objDocModel.Category = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].Category
		//objDocModel.SchCode = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].SchCode
		schcode := mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].SchCode
		strschcode := strconv.FormatInt(schcode, 10)
		objDocModel.SchCode = strschcode
		objDocModel.SchParentCode = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].SchParentCode
		objDocModel.MinPurAmt = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].MinPurAmt
		objDocModel.PurAmountMult = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].PurAmountMult
		objDocModel.MinAddPurAmt = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].MinAddPurAmt
		objDocModel.Objective = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].Objective
		objDocModel.TranMode = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].TranMode

		objDocModel.AMCID = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].AMCID
		objDocModel.AMCName = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].AMCName
		objDocModel.ISIN = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].ISIN
		objDocModel.SchemeCode = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SchemeCode
		objDocModel.SchemeName = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SchemeName
		objDocModel.SIPFrequency = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPFrequency
		objDocModel.SIPDates = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPDates
		objDocModel.SIPMinGAP = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMinGAP
		objDocModel.SIPMaxGAP = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMaxGAP
		objDocModel.SIPMinInstallmentAmount = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMinInstallmentAmount
		objDocModel.SIPMinInstallmentNo = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMinInstallmentNo
		objDocModel.SIPMultiplierAmt = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMultiplierAmt
		objDocModel.SIPMaxInstallmentAmount = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMaxInstallmentAmount
		objDocModel.SIPMaxInstallmentNo = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMaxInstallmentNo
		//objDocModel.Type = "sip"

		uploadMfDatatypesense(objDocModel)

		count++
		fmt.Println(count, i)
	}

	fmt.Println("Successfully uploaded documents")

}

// func (repo *mfService) GetMfData() []interface{} {
func (repo *mfService) GetMfDataOptimized() {

	Purchasemodel := GetPurchaseorder()
	//fmt.Println(Purchasemodel)
	objpurchase := Purchasemodel.(models.PurchaseSiporderResponse)

	var objPurchasemodel models.PurchaseSipDataModel
	var ArrPurchasemodel []models.PurchaseSipDataModel
	mapPurchasemodel := make(map[string]models.PurchaseSipDataModel)

	for i := 0; i < len(objpurchase.Data); i++ {
		objPurchasemodel.SchemeCode = objpurchase.Data[i].SchemeCode
		objPurchasemodel.SchemeName = objpurchase.Data[i].SchemeName
		ArrPurchasemodel = append(ArrPurchasemodel, objPurchasemodel)

		mapPurchasemodel[objpurchase.Data[i].SchemeCode] = objpurchase.Data[i]
	}

	// Sipmodel := GetSiporder()
	// // fmt.Println(Sipmodel)
	// objsip := Sipmodel.(models.PurchaseSiporderResponse)

	//mapPurchasemodel := make(map[string]models.PurchaseSipDataModel)
	//mapSipmodel := make(map[string]models.PurchaseSipDataModel)

	// for i := 0; i < len(objpurchase.Data); i++ {
	// 	mapPurchasemodel[objpurchase.Data[i].SchemeCode] = objpurchase.Data[i]
	// }
	//fmt.Println(mapPurchasemodel)

	// for i := 0; i < len(objsip.Data); i++ {
	// 	mapSipmodel[objsip.Data[i].SchemeCode] = objsip.Data[i]
	// }
	//fmt.Println(mapSipmodel)

	mapMfPurchaseInfoData := make(map[string]models.Scheme_Info)
	// var ArrPurchaseInfo []models.Scheme_InfoResponse
	// var PurchaseInfo models.Scheme_InfoResponse

	Accesstoken := Mflogin()

	slcDocument := make([]interface{}, len(ArrPurchasemodel))

	//count := 0
	for i, v := range ArrPurchasemodel {

		const url = "https://fundcore-nu.azurewebsites.net/fcapi/mf/scheme/info/for-purchase-order"

		var ReqBody models.PurchaseInfoBody
		ReqBody.BseSchemeCode = v.SchemeCode

		jsonStr, err := json.Marshal(&ReqBody)
		if err != nil {
			//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
			fmt.Println(err)
			//return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

		if err != nil {
			//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
			fmt.Println(err)
			// return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+Accesstoken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
			fmt.Println(err)
		}

		//fmt.Println(resp.Body)

		jsonString, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
			fmt.Println(err)
		}

		var respModel models.InfoPurchaseModel
		err = json.Unmarshal(jsonString, &respModel)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("schcode:" + v.SchemeCode + " length: " + strconv.Itoa(len(mapMfPurchaseInfoData)) + " lenghtResp: " + strconv.Itoa(len(respModel.Data.SchemeInfo)))

		mapMfPurchaseInfoData[v.SchemeCode] = respModel.Data.SchemeInfo[0]
		//if len(mapMfPurchaseInfoData) > 0 {
		if len(respModel.Data.SchemeInfo) == 0 {
			continue
		}
		objDocModel := models.MfDocumentModel{}
		objDocModel.AMCID = mapMfPurchaseInfoData[v.SchemeCode].AMCID
		objDocModel.AMCName = mapMfPurchaseInfoData[v.SchemeCode].AMCName
		objDocModel.ISIN = mapMfPurchaseInfoData[v.SchemeCode].ISIN
		objDocModel.SchemeCode = mapMfPurchaseInfoData[v.SchemeCode].SchemeCode
		objDocModel.SchemeName = mapMfPurchaseInfoData[v.SchemeCode].SchemeName
		objDocModel.BSEToken = mapMfPurchaseInfoData[v.SchemeCode].BSEToken
		objDocModel.DivReFlag = mapMfPurchaseInfoData[v.SchemeCode].DivReFlag
		objDocModel.LatestNAV = mapMfPurchaseInfoData[v.SchemeCode].LatestNAV
		objDocModel.NAVDate = mapMfPurchaseInfoData[v.SchemeCode].NAVDate
		objDocModel.NAVChange = mapMfPurchaseInfoData[v.SchemeCode].NAVChange
		objDocModel.Risk = mapMfPurchaseInfoData[v.SchemeCode].Risk
		objDocModel.Category = mapMfPurchaseInfoData[v.SchemeCode].Category
		// objDocModel.SchCode = mapMfPurchaseInfoData[v.SchemeCode].SchCode
		schcode := mapMfPurchaseInfoData[v.SchemeCode].SchCode
		strschcode := strconv.FormatInt(schcode, 10)
		objDocModel.SchCode = strschcode
		objDocModel.SchParentCode = mapMfPurchaseInfoData[v.SchemeCode].SchParentCode
		objDocModel.MinPurAmt = mapMfPurchaseInfoData[v.SchemeCode].MinPurAmt
		objDocModel.PurAmountMult = mapMfPurchaseInfoData[v.SchemeCode].PurAmountMult
		objDocModel.MinAddPurAmt = mapMfPurchaseInfoData[v.SchemeCode].MinAddPurAmt
		objDocModel.Objective = mapMfPurchaseInfoData[v.SchemeCode].Objective
		objDocModel.TranMode = mapMfPurchaseInfoData[v.SchemeCode].TranMode
		//objDocModel.Type = "Lumpsum"

		slcDocument[i] = objDocModel
		// count++
		fmt.Println(i, v)

	}

	//fmt.Println(mapMfPurchaseInfoData)

	uploadMfDatatypesense1(slcDocument)

	//return slcDocument

	//Accesstoken := Mflogin()
	Sipmodel := GetSiporder()
	// // fmt.Println(Sipmodel)
	objsip := Sipmodel.(models.PurchaseSiporderResponse)

	var objSipmodel models.PurchaseSipDataModel
	var ArrSipmodel []models.PurchaseSipDataModel
	mapSipmodel := make(map[string]models.PurchaseSipDataModel)

	for i := 0; i < len(objsip.Data); i++ {
		objSipmodel.SchemeCode = objsip.Data[i].SchemeCode
		objSipmodel.SchemeName = objsip.Data[i].SchemeName
		ArrSipmodel = append(ArrSipmodel, objSipmodel)

		mapSipmodel[objsip.Data[i].SchemeCode] = objsip.Data[i]
	}

	mapMfSipInfoData := make(map[string]models.InfoSipModel)

	//Accesstoken := Mflogin()

	slcDocument1 := make([]interface{}, len(ArrSipmodel))

	for i, v := range ArrSipmodel {
		const url = "https://fundcore-nu.azurewebsites.net/fcapi/mf/scheme/info/for-sip-order"

		var ReqBody models.PurchaseInfoBody
		ReqBody.BseSchemeCode = v.SchemeCode

		jsonStr, err := json.Marshal(&ReqBody)
		if err != nil {
			//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
			fmt.Println(err)
			//return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

		if err != nil {
			//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
			fmt.Println(err)
			// return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+Accesstoken)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
			fmt.Println(err)
		}

		jsonString, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
			fmt.Println(err)
		}

		var respModel models.InfoSipModel
		err = json.Unmarshal(jsonString, &respModel)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("schcode:" + v.SchemeCode + " length: " + strconv.Itoa(len(mapMfSipInfoData)) + " lenghtResp: " + strconv.Itoa(len(respModel.Data.SchemeInfo1)))
		mapMfSipInfoData[v.SchemeCode] = respModel

		objDocModel := models.MfDocumentModel{}
		objDocModel.AMCID = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].AMCID
		if len(respModel.Data.SchemeInfo1) == 0 {
			continue
		}
		objDocModel.AMCName = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].AMCName
		objDocModel.ISIN = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].ISIN
		objDocModel.SchemeCode = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].SchemeCode
		objDocModel.SchemeName = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].SchemeName
		objDocModel.BSEToken = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].BSEToken
		objDocModel.DivReFlag = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].DivReFlag
		objDocModel.LatestNAV = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].LatestNAV
		objDocModel.NAVDate = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].NAVDate
		objDocModel.NAVChange = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].NAVChange
		objDocModel.Risk = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].Risk
		objDocModel.Category = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].Category
		// objDocModel.SchCode = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].SchCode
		schcode := mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].SchCode
		strschcode := strconv.FormatInt(schcode, 10)
		objDocModel.SchCode = strschcode
		objDocModel.SchParentCode = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].SchParentCode
		objDocModel.MinPurAmt = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].MinPurAmt
		objDocModel.PurAmountMult = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].PurAmountMult
		objDocModel.MinAddPurAmt = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].MinAddPurAmt
		objDocModel.Objective = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].Objective
		objDocModel.TranMode = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo1[0].TranMode

		objDocModel.AMCID = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].AMCID
		objDocModel.AMCName = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].AMCName
		objDocModel.ISIN = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].ISIN
		objDocModel.SchemeCode = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SchemeCode
		objDocModel.SchemeName = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SchemeName
		objDocModel.SIPFrequency = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPFrequency
		objDocModel.SIPDates = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPDates
		objDocModel.SIPMinGAP = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMinGAP
		objDocModel.SIPMaxGAP = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMaxGAP
		objDocModel.SIPMinInstallmentAmount = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMinInstallmentAmount
		objDocModel.SIPMinInstallmentNo = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMinInstallmentNo
		objDocModel.SIPMultiplierAmt = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMultiplierAmt
		objDocModel.SIPMaxInstallmentAmount = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMaxInstallmentAmount
		objDocModel.SIPMaxInstallmentNo = mapMfSipInfoData[v.SchemeCode].Data.SchemeInfo2[0].SIPMaxInstallmentNo
		//objDocModel.Type = "sip"

		slcDocument1[i] = objDocModel
		fmt.Println(i, v)

	}

	uploadMfDatatypesense1(slcDocument1)

}

func (repo *mfService) GetMutualFund() {

	// t := time.Now()
	// Today := t.Format("02-01-2006 15:04:05")

	Sipmodel := GetSiporder()
	objsip := Sipmodel.(models.PurchaseSiporderResponse)

	Purchasemodel := GetPurchaseorder()
	objpurchase := Purchasemodel.(models.PurchaseSiporderResponse)

	mapSipLumpsum := make(map[string]models.SipLumpsumModel)

	// if len(objsip.Data) > len(objpurchase.Data) {
	// 	for i := 0; i < len(objsip.Data); i++ {
	// 		for j := 0; j < len(objpurchase.Data); j++ {
	// 			if objsip.Data[i].SchemeCode == objpurchase.Data[j].SchemeCode {
	// 				mapSipLumpsum[objsip.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objsip.Data[i].SchemeCode, SchemeName: objsip.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: true}
	// 			} else if objsip.Data[i].SchemeCode != objpurchase.Data[j].SchemeCode {
	// 				mapSipLumpsum[objsip.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objsip.Data[i].SchemeCode, SchemeName: objsip.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: false}
	// 			} else {
	// 				mapSipLumpsum[objsip.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objsip.Data[i].SchemeCode, SchemeName: objsip.Data[i].SchemeName, ISSIP: false, ISLUMPSUM: true}
	// 			}
	// 		}
	// 	}
	// } else if len(objsip.Data) < len(objpurchase.Data) {
	// 	for i := 0; i < len(objpurchase.Data); i++{
	// 		for j := 0; j < len(objsip.Data); j++{
	// 			if objpurchase.Data[i].SchemeCode == objsip.Data[j].SchemeCode {
	// 				mapSipLumpsum[objpurchase.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objpurchase.Data[i].SchemeCode, SchemeName: objpurchase.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: true}
	// 				break
	// 			} else if objpurchase.Data[i].SchemeCode != objsip.Data[j].SchemeCode {
	// 				mapSipLumpsum[objpurchase.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objpurchase.Data[i].SchemeCode, SchemeName: objpurchase.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: false}
	// 			} else {
	// 				mapSipLumpsum[objpurchase.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objpurchase.Data[i].SchemeCode, SchemeName: objpurchase.Data[i].SchemeName, ISSIP: false, ISLUMPSUM: true}
	// 			}
	// 		}
	// 	}
	// }

	if len(objsip.Data) > len(objpurchase.Data) {
		for i := range objsip.Data {
			for j := range objpurchase.Data {
				if objsip.Data[i].SchemeCode == objpurchase.Data[j].SchemeCode {
					mapSipLumpsum[objsip.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objsip.Data[i].SchemeCode, SchemeName: objsip.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: true}
				} else if objsip.Data[i].SchemeCode != objpurchase.Data[j].SchemeCode {
					mapSipLumpsum[objsip.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objsip.Data[i].SchemeCode, SchemeName: objsip.Data[i].SchemeName, ISSIP: false, ISLUMPSUM: true}
				} else {
					mapSipLumpsum[objsip.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objsip.Data[i].SchemeCode, SchemeName: objsip.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: false}
				}
			}
		}
	} else if len(objsip.Data) < len(objpurchase.Data) {
		for i := range objpurchase.Data {
			for j := range objsip.Data {
				if objpurchase.Data[i].SchemeCode == objsip.Data[j].SchemeCode {
					mapSipLumpsum[objpurchase.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objpurchase.Data[i].SchemeCode, SchemeName: objpurchase.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: true}
					break
				} else if objpurchase.Data[i].SchemeCode != objsip.Data[j].SchemeCode {
					mapSipLumpsum[objpurchase.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objpurchase.Data[i].SchemeCode, SchemeName: objpurchase.Data[i].SchemeName, ISSIP: false, ISLUMPSUM: true}
				} else {
					mapSipLumpsum[objpurchase.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objpurchase.Data[i].SchemeCode, SchemeName: objpurchase.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: false}
				}
			}
		}

	}

	fmt.Println("---------------------------------------------------------------------------------------------------------")

	var TestModel []models.MfDocumentModel
	slcDocument := make([]interface{}, (len(mapSipLumpsum) * 2)) //slcDocument := make([]interface{}, len(TestModel))
	count := 1
	for i := range mapSipLumpsum {
		// fmt.Println("SchemeCode:	" + mapSipLumpsum[i].SchemeCode)
		// fmt.Println("SchemeName:	" + mapSipLumpsum[i].SchemeName)
		// fmt.Println("ISSIP:	", mapSipLumpsum[i].ISSIP)
		// fmt.Println("ISLUMPSUM:	", mapSipLumpsum[i].ISLUMPSUM)
		// fmt.Println("---------------------------------------------------------------------------------------------------------")

		fmt.Println(count)

		Accesstoken := Mflogin()

		if mapSipLumpsum[i].ISSIP == true {
			const url = "https://fundcore-nu.azurewebsites.net/fcapi/mf/scheme/info/for-sip-order"

			var ReqBody models.PurchaseInfoBody
			ReqBody.BseSchemeCode = mapSipLumpsum[i].SchemeCode

			jsonStr, err := json.Marshal(&ReqBody)
			if err != nil {
				//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
				fmt.Println(err)
				//return
			}
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

			if err != nil {
				//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
				fmt.Println(err)
				// return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+Accesstoken)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
				fmt.Println(err)
			}

			jsonString, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
				fmt.Println(err)
			}

			var respModel models.InfoSipModel
			err = json.Unmarshal(jsonString, &respModel)
			if err != nil {
				fmt.Println(err)
			}

			//mapSipLumpsum[i].Sip = respModel.Data.SchemeInfo2
			//mapSipLumpsum[i] = models.SipLumpsumModel{Lumpsum: respModel.Data.SchemeInfo1[0]}

			//objDocModel := models.SipLumpsumtypesenseModel{}
			objDocModel := models.MfDocumentModel{}
			objDocModel.SchemeCode = mapSipLumpsum[i].SchemeCode
			objDocModel.SchemeName = mapSipLumpsum[i].SchemeName
			objDocModel.ISSIP = mapSipLumpsum[i].ISSIP
			objDocModel.ISLUMPSUM = mapSipLumpsum[i].ISLUMPSUM
			if len(respModel.Data.SchemeInfo1) == 0 {
				continue
			}
			objDocModel.AMCID = respModel.Data.SchemeInfo1[0].AMCID
			objDocModel.AMCName = respModel.Data.SchemeInfo1[0].AMCName
			objDocModel.ISIN = respModel.Data.SchemeInfo1[0].ISIN
			objDocModel.SchemeCode = respModel.Data.SchemeInfo1[0].SchemeCode
			objDocModel.SchemeName = respModel.Data.SchemeInfo1[0].SchemeName
			objDocModel.BSEToken = respModel.Data.SchemeInfo1[0].BSEToken
			objDocModel.DivReFlag = respModel.Data.SchemeInfo1[0].DivReFlag
			objDocModel.LatestNAV = respModel.Data.SchemeInfo1[0].LatestNAV
			objDocModel.NAVDate = respModel.Data.SchemeInfo1[0].NAVDate
			objDocModel.NAVChange = respModel.Data.SchemeInfo1[0].NAVChange
			objDocModel.Risk = respModel.Data.SchemeInfo1[0].Risk
			objDocModel.Category = respModel.Data.SchemeInfo1[0].Category
			// objDocModel.SchCode = respModel.Data.SchemeInfo1[0].SchCode
			schcode := respModel.Data.SchemeInfo1[0].SchCode
			strschcode := strconv.FormatInt(schcode, 10)
			objDocModel.SchCode = strschcode
			objDocModel.SchParentCode = respModel.Data.SchemeInfo1[0].SchParentCode
			objDocModel.MinPurAmt = respModel.Data.SchemeInfo1[0].MinPurAmt
			objDocModel.PurAmountMult = respModel.Data.SchemeInfo1[0].PurAmountMult
			objDocModel.MinAddPurAmt = respModel.Data.SchemeInfo1[0].MinAddPurAmt
			objDocModel.Objective = respModel.Data.SchemeInfo1[0].Objective
			objDocModel.TranMode = respModel.Data.SchemeInfo1[0].TranMode

			objDocModel.AMCID = respModel.Data.SchemeInfo2[0].AMCID
			objDocModel.AMCName = respModel.Data.SchemeInfo2[0].AMCName
			objDocModel.ISIN = respModel.Data.SchemeInfo2[0].ISIN
			objDocModel.SchemeCode = respModel.Data.SchemeInfo2[0].SchemeCode
			objDocModel.SchemeName = respModel.Data.SchemeInfo2[0].SchemeName
			objDocModel.SIPFrequency = respModel.Data.SchemeInfo2[0].SIPFrequency
			objDocModel.SIPDates = respModel.Data.SchemeInfo2[0].SIPDates
			objDocModel.SIPMinGAP = respModel.Data.SchemeInfo2[0].SIPMinGAP
			objDocModel.SIPMaxGAP = respModel.Data.SchemeInfo2[0].SIPMaxGAP
			objDocModel.SIPMinInstallmentAmount = respModel.Data.SchemeInfo2[0].SIPMinInstallmentAmount
			objDocModel.SIPMinInstallmentNo = respModel.Data.SchemeInfo2[0].SIPMinInstallmentNo
			objDocModel.SIPMultiplierAmt = respModel.Data.SchemeInfo2[0].SIPMultiplierAmt
			objDocModel.SIPMaxInstallmentAmount = respModel.Data.SchemeInfo2[0].SIPMaxInstallmentAmount
			objDocModel.SIPMaxInstallmentNo = respModel.Data.SchemeInfo2[0].SIPMaxInstallmentNo

			TestModel = append(TestModel, objDocModel)
			count++
			// if count >= 100 {
			// 	break
			// }

			//		uploadMfDatatypesense(objDocModel)

		}

		if mapSipLumpsum[i].ISLUMPSUM == true {
			const url = "https://fundcore-nu.azurewebsites.net/fcapi/mf/scheme/info/for-purchase-order"

			var ReqBody models.PurchaseInfoBody
			ReqBody.BseSchemeCode = mapSipLumpsum[i].SchemeCode

			jsonStr, err := json.Marshal(&ReqBody)
			if err != nil {
				//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
				fmt.Println(err)
				//return
			}
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

			if err != nil {
				//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
				fmt.Println(err)
				// return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+Accesstoken)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
				fmt.Println(err)
			}

			//fmt.Println(resp.Body)

			jsonString, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
				fmt.Println(err)
			}

			var respModel models.InfoPurchaseModel
			err = json.Unmarshal(jsonString, &respModel)
			if err != nil {
				fmt.Println(err)
			}

			objDocModel := models.MfDocumentModel{}
			objDocModel.SchemeCode = mapSipLumpsum[i].SchemeCode
			objDocModel.SchemeName = mapSipLumpsum[i].SchemeName
			objDocModel.ISSIP = mapSipLumpsum[i].ISSIP
			objDocModel.ISLUMPSUM = mapSipLumpsum[i].ISLUMPSUM
			if len(respModel.Data.SchemeInfo) == 0 {
				continue
			}
			objDocModel.AMCID = respModel.Data.SchemeInfo[0].AMCID
			objDocModel.AMCName = respModel.Data.SchemeInfo[0].AMCName
			objDocModel.ISIN = respModel.Data.SchemeInfo[0].ISIN
			objDocModel.SchemeCode = respModel.Data.SchemeInfo[0].SchemeCode
			objDocModel.SchemeName = respModel.Data.SchemeInfo[0].SchemeName
			objDocModel.BSEToken = respModel.Data.SchemeInfo[0].BSEToken
			objDocModel.DivReFlag = respModel.Data.SchemeInfo[0].DivReFlag
			objDocModel.LatestNAV = respModel.Data.SchemeInfo[0].LatestNAV
			objDocModel.NAVDate = respModel.Data.SchemeInfo[0].NAVDate
			objDocModel.NAVChange = respModel.Data.SchemeInfo[0].NAVChange
			objDocModel.Risk = respModel.Data.SchemeInfo[0].Risk
			objDocModel.Category = respModel.Data.SchemeInfo[0].Category
			//objDocModel.SchCode = respModel.Data.SchemeInfo[0].SchCode
			schcode := respModel.Data.SchemeInfo[0].SchCode
			strschcode := strconv.FormatInt(schcode, 10)
			objDocModel.SchCode = strschcode
			objDocModel.SchParentCode = respModel.Data.SchemeInfo[0].SchParentCode
			objDocModel.MinPurAmt = respModel.Data.SchemeInfo[0].MinPurAmt
			objDocModel.PurAmountMult = respModel.Data.SchemeInfo[0].PurAmountMult
			objDocModel.MinAddPurAmt = respModel.Data.SchemeInfo[0].MinAddPurAmt
			objDocModel.Objective = respModel.Data.SchemeInfo[0].Objective
			objDocModel.TranMode = respModel.Data.SchemeInfo[0].TranMode

			TestModel = append(TestModel, objDocModel)
			count++

			// if count >= 100 {
			// 	break
			// }
			//uploadMfDatatypesense(objDocModel)

		}

	}

	for i := range TestModel {
		//objDocModel := models.MfDocumentModel{}
		//objScripDetails := TestModel[i]
		// objDocModel.AMCID = objScripDetails.AMCID
		// objDocModel.ISLUMPSUM = objScripDetails.ISLUMPSUM
		//slcDocument[i] = objDocModel
		slcDocument[i] = TestModel[i]

	}

	//uploadMfDatatypesense2(slcDocument)

	if len(slcDocument) > 0 {
		batch := 50
		for i := 0; i < len(slcDocument); i += batch {
			j := i + batch
			if j > len(slcDocument) {
				j = len(slcDocument)
			}
			uploadMfDatatypesense2(slcDocument[i:j]) // Process the batch.
		}
	}

	// elapsed := time.Since(t)
	// ReqTime := fmt.Sprintf("%v", elapsed)
	// str := Today + " Mutual fund data(greenwaremf) inserted in " + ReqTime
	// err := SentToContractCronJob(str)
	// if err != nil {
	// 	fmt.Println("Error while sending to slack")
	// }
	Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Uploading Mutual fund to Typesense End.")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Uploading Mutual fund to Typesense End.")

}

func (repo *mfService) MasterDataMFDrop() {

	t := time.Now()
	Today := t.Format("02-01-2006 15:04:05")

	client := typesense.NewClient(
		typesense.WithServer(BaseURL),
		typesense.WithAPIKey(TypessenseValue))

	client.Collection(mfcollectionname2).Delete()

	elapsed := time.Since(t)
	ReqTime := fmt.Sprintf("%v", elapsed)
	str := Today + "\n Mutual fund data(greenwaremf2-foocut) collection dropped sucessfully in " + ReqTime
	err := SentToContractCronJob(str)
	if err != nil {
		fmt.Println("Error while sending to slack")
	}

	Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Sucessfully collection dropped from Typesense. ")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Sucessfully collection dropped from Typesense. ")
}

// func uploadMasterMFtypesense(objData interface{}) error {
func uploadMasterMFtypesense(objData []interface{}) error {
	Zerologs.Info().Msg("uploadMasterDatatypesense Start")
	client := typesense.NewClient(
		typesense.WithServer(BaseURL),
		typesense.WithAPIKey(TypessenseValue))

	_, err := client.Collection(mfcollectionname2).Retrieve()
	if err != nil {
		bfacet := true
		schema := &api.CollectionSchema{
			Name: mfcollectionname2,
			Fields: []api.Field{
				{
					Name:  "AMCID",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "AMCName",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchCode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchemeName",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Status",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "CategoryCode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Category",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Type",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchemePlan",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Risk",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "LumpSumAllowed",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SIPAllowed",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "MinLumpsumInvestment",
					Type:  "auto",
					Facet: &bfacet,
				},
				{
					Name:  "MinSIPInvestment",
					Type:  "auto",
					Facet: &bfacet,
				},
				{
					Name:  "Return1Day",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Return1Week",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Return1Month",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Return3Month",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Return6Month",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Return1Year",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Return3Year",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Return5Year",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "ReturnSinceInception",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "NAV",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "NAVDate",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "FMgr1Code",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "FMgr1Name",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "FMgr2Code",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "FMgr2Name",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SchemeAUM",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "NatureDes",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "SubNatureDes",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Options",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "MinInvestment",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "Objective",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "AMFICode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "EntryLoad",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "ExitLoad",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "ExitLoadDes",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "ExpenseRatio",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "AumAsOn",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "ISIN_Growth",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "ISIN_DivPayout",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "ISIN_DivReInvest",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "FiftyTwoWeekHigh",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "FiftyTwoWeekHighDate",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "FiftyTwoWeekLow",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "FiftyTwoWeekLowDate",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "OpenDate",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Benchmark",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Address1",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Address2",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "Address3",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "City",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "State",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "StdDev",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "BSESchemecode",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "BSEToken",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "MinPurAmt",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "PurAmountMult",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "MinAddPurAmt",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "SIPSchemeInfo2",
					Type:  "string",
					Facet: &bfacet,
				},
			},
			//DefaultSortingField: &sort1,
		}
		_, err := client.Collections().Create(schema)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(resp)
	}

	document := objData
	var action = "upsert"
	var batch = 40
	params := &api.ImportDocumentsParams{
		Action:    &action,
		BatchSize: &batch,
	}

	_, err = client.Collection(mfcollectionname2).Documents().Import(document, params)
	if err != nil {
		Zerologs.Error().Err(err).Msgf("Error in uploadMasterDatatypesense")
		fmt.Println(err)
		return err
	}
	return nil

}

func GetMasterData() interface{} {

	MasterAccesstoken := "TklfRGF0YUFwaVVzZXI6JHhwfjUreVJXNA=="

	const url = "http://masterdata.fundcore.in/api/v2/scheme/screener"

	var ReqBody models.MasterDataBody
	ReqBody.AmcId = "ALL"
	ReqBody.TypeCode = "ALL"
	ReqBody.CategoryCode = "ALL"
	ReqBody.Option = "ALL"
	ReqBody.Plan = "ALL"
	ReqBody.Risk = "ALL"
	ReqBody.NfoFlag = "ALL"
	ReqBody.RefKey = "ALL"

	//var PurchaseModel Purchse
	// err := c.BindJSON(&ReqBody)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	jsonStr, err := json.Marshal(&ReqBody)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
		fmt.Println(err)
		//return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	if err != nil {
		//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
		fmt.Println(err)
		// return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+MasterAccesstoken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
		fmt.Println(err)
	}

	//fmt.Println(resp.Body)

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
		fmt.Println(err)
	}

	var respModel models.MasterdataMfResponse
	err = json.Unmarshal(jsonString, &respModel)
	if err != nil {
		fmt.Println(err)
	}

	return respModel
}

func GetMasterDataInfo(SchemeCode int64) interface{} {

	MasterAccesstoken := "TklfRGF0YUFwaVVzZXI6JHhwfjUreVJXNA=="

	const url = "http://masterdata.fundcore.in/api/v2/scheme/info"

	var ReqBody models.MasterDataInfoBody
	ReqBody.Code = SchemeCode
	ReqBody.Fields = "string"
	ReqBody.String = "string"
	ReqBody.RefKey = "RefKey"

	jsonStr, err := json.Marshal(&ReqBody)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
		fmt.Println(err)
		//return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	if err != nil {
		//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
		fmt.Println(err)
		// return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+MasterAccesstoken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
		fmt.Println(err)
	}

	//fmt.Println(resp.Body)

	jsonString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
		fmt.Println(err)
	}

	var respModel models.MasterDataInfoResponse
	err = json.Unmarshal(jsonString, &respModel)
	if err != nil {
		fmt.Println(err)
	}

	return respModel
}

func GetMfmap() map[string]models.MfDocumentModel {

	Sipmodel := GetSiporder()
	objsip := Sipmodel.(models.PurchaseSiporderResponse)

	Purchasemodel := GetPurchaseorder()
	objpurchase := Purchasemodel.(models.PurchaseSiporderResponse)

	mapSipLumpsum := make(map[string]models.SipLumpsumModel)

	if len(objsip.Data) > len(objpurchase.Data) {
		for i := range objsip.Data {
			for j := range objpurchase.Data {
				if objsip.Data[i].SchemeCode == objpurchase.Data[j].SchemeCode {
					mapSipLumpsum[objsip.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objsip.Data[i].SchemeCode, SchemeName: objsip.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: true}
				} else if objsip.Data[i].SchemeCode != objpurchase.Data[j].SchemeCode {
					mapSipLumpsum[objsip.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objsip.Data[i].SchemeCode, SchemeName: objsip.Data[i].SchemeName, ISSIP: false, ISLUMPSUM: true}
				} else {
					mapSipLumpsum[objsip.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objsip.Data[i].SchemeCode, SchemeName: objsip.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: false}
				}
			}
		}
	} else if len(objsip.Data) < len(objpurchase.Data) {
		for i := range objpurchase.Data {
			for j := range objsip.Data {
				if objpurchase.Data[i].SchemeCode == objsip.Data[j].SchemeCode {
					mapSipLumpsum[objpurchase.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objpurchase.Data[i].SchemeCode, SchemeName: objpurchase.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: true}
					break
				} else if objpurchase.Data[i].SchemeCode != objsip.Data[j].SchemeCode {
					mapSipLumpsum[objpurchase.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objpurchase.Data[i].SchemeCode, SchemeName: objpurchase.Data[i].SchemeName, ISSIP: false, ISLUMPSUM: true}
				} else {
					mapSipLumpsum[objpurchase.Data[i].SchemeCode] = models.SipLumpsumModel{SchemeCode: objpurchase.Data[i].SchemeCode, SchemeName: objpurchase.Data[i].SchemeName, ISSIP: true, ISLUMPSUM: false}
				}
			}
		}

	}

	var TestModel []models.MfDocumentModel
	count := 1
	for i := range mapSipLumpsum {
		fmt.Println(count)

		Accesstoken := Mflogin()

		objDocModel := models.MfDocumentModel{}

		if mapSipLumpsum[i].ISSIP == true {
			const url = "https://fundcore-nu.azurewebsites.net/fcapi/mf/scheme/info/for-sip-order"

			var ReqBody models.PurchaseInfoBody
			ReqBody.BseSchemeCode = mapSipLumpsum[i].SchemeCode

			jsonStr, err := json.Marshal(&ReqBody)
			if err != nil {
				//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
				fmt.Println(err)
				//return
			}
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

			if err != nil {
				//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
				fmt.Println(err)
				// return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+Accesstoken)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
				fmt.Println(err)
			}

			jsonString, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
				fmt.Println(err)
			}

			var respModel models.InfoSipModel
			err = json.Unmarshal(jsonString, &respModel)
			if err != nil {
				fmt.Println(err)
			}

			objDocModel.SchemeCode = mapSipLumpsum[i].SchemeCode
			objDocModel.SchemeName = mapSipLumpsum[i].SchemeName
			objDocModel.ISSIP = mapSipLumpsum[i].ISSIP
			objDocModel.ISLUMPSUM = mapSipLumpsum[i].ISLUMPSUM
			if len(respModel.Data.SchemeInfo1) == 0 {
				continue
			}
			objDocModel.AMCID = respModel.Data.SchemeInfo1[0].AMCID
			objDocModel.AMCName = respModel.Data.SchemeInfo1[0].AMCName
			objDocModel.ISIN = respModel.Data.SchemeInfo1[0].ISIN
			objDocModel.SchemeCode = respModel.Data.SchemeInfo1[0].SchemeCode
			objDocModel.SchemeName = respModel.Data.SchemeInfo1[0].SchemeName
			objDocModel.BSEToken = respModel.Data.SchemeInfo1[0].BSEToken
			objDocModel.DivReFlag = respModel.Data.SchemeInfo1[0].DivReFlag
			objDocModel.LatestNAV = respModel.Data.SchemeInfo1[0].LatestNAV
			objDocModel.NAVDate = respModel.Data.SchemeInfo1[0].NAVDate
			objDocModel.NAVChange = respModel.Data.SchemeInfo1[0].NAVChange
			objDocModel.Risk = respModel.Data.SchemeInfo1[0].Risk
			objDocModel.Category = respModel.Data.SchemeInfo1[0].Category
			schcode := respModel.Data.SchemeInfo1[0].SchCode
			strschcode := strconv.FormatInt(schcode, 10)
			objDocModel.SchCode = strschcode

			objDocModel.SchParentCode = respModel.Data.SchemeInfo1[0].SchParentCode
			objDocModel.MinPurAmt = respModel.Data.SchemeInfo1[0].MinPurAmt
			objDocModel.PurAmountMult = respModel.Data.SchemeInfo1[0].PurAmountMult
			objDocModel.MinAddPurAmt = respModel.Data.SchemeInfo1[0].MinAddPurAmt
			objDocModel.Objective = respModel.Data.SchemeInfo1[0].Objective
			objDocModel.TranMode = respModel.Data.SchemeInfo1[0].TranMode

			//objDocModel.SIPSchemeInfo = respModel.Data.SchemeInfo2
			SchemeInfo2data, _ := json.Marshal(respModel.Data.SchemeInfo2)
			fmt.Println(string(SchemeInfo2data))
			objDocModel.SIPSchemeInfo2 = string(SchemeInfo2data)

		}

		if mapSipLumpsum[i].ISLUMPSUM == true {
			const url = "https://fundcore-nu.azurewebsites.net/fcapi/mf/scheme/info/for-purchase-order"

			var ReqBody models.PurchaseInfoBody
			ReqBody.BseSchemeCode = mapSipLumpsum[i].SchemeCode

			jsonStr, err := json.Marshal(&ReqBody)
			if err != nil {
				//Zerologs.Error().Err(err).Msg("Error While marshalling watchlist")
				fmt.Println(err)
				//return
			}
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

			if err != nil {
				//Zerologs.Error().Err(err).Msgf("Error While Sending POST Request %s", url)
				fmt.Println(err)
				// return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", "Bearer "+Accesstoken)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				//	Zerologs.Error().Err(err).Msg("Client DO error" + err.Error())
				fmt.Println(err)
			}

			//fmt.Println(resp.Body)

			jsonString, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				//Zerologs.Error().Err(err).Msg("Error reading response body from " + URL + err.Error())
				fmt.Println(err)
			}

			var respModel models.InfoPurchaseModel
			err = json.Unmarshal(jsonString, &respModel)
			if err != nil {
				fmt.Println(err)
			}

			objDocModel.SchemeCode = mapSipLumpsum[i].SchemeCode
			objDocModel.SchemeName = mapSipLumpsum[i].SchemeName
			objDocModel.ISSIP = mapSipLumpsum[i].ISSIP
			objDocModel.ISLUMPSUM = mapSipLumpsum[i].ISLUMPSUM
			if len(respModel.Data.SchemeInfo) == 0 {
				continue
			}
			objDocModel.AMCID = respModel.Data.SchemeInfo[0].AMCID
			objDocModel.AMCName = respModel.Data.SchemeInfo[0].AMCName
			objDocModel.ISIN = respModel.Data.SchemeInfo[0].ISIN
			objDocModel.SchemeCode = respModel.Data.SchemeInfo[0].SchemeCode
			objDocModel.SchemeName = respModel.Data.SchemeInfo[0].SchemeName
			objDocModel.BSEToken = respModel.Data.SchemeInfo[0].BSEToken
			objDocModel.DivReFlag = respModel.Data.SchemeInfo[0].DivReFlag
			objDocModel.LatestNAV = respModel.Data.SchemeInfo[0].LatestNAV
			objDocModel.NAVDate = respModel.Data.SchemeInfo[0].NAVDate
			objDocModel.NAVChange = respModel.Data.SchemeInfo[0].NAVChange
			objDocModel.Risk = respModel.Data.SchemeInfo[0].Risk
			objDocModel.Category = respModel.Data.SchemeInfo[0].Category
			schcode := respModel.Data.SchemeInfo[0].SchCode
			strschcode := strconv.FormatInt(schcode, 10)
			objDocModel.SchCode = strschcode
			objDocModel.SchParentCode = respModel.Data.SchemeInfo[0].SchParentCode
			objDocModel.MinPurAmt = respModel.Data.SchemeInfo[0].MinPurAmt
			objDocModel.PurAmountMult = respModel.Data.SchemeInfo[0].PurAmountMult
			objDocModel.MinAddPurAmt = respModel.Data.SchemeInfo[0].MinAddPurAmt
			objDocModel.Objective = respModel.Data.SchemeInfo[0].Objective
			objDocModel.TranMode = respModel.Data.SchemeInfo[0].TranMode

		}

		TestModel = append(TestModel, objDocModel)
		count++

	}

	mfdatamap := make(map[string]models.MfDocumentModel)

	for i := range TestModel {
		mfdatamap[TestModel[i].SchCode] = TestModel[i]
	}

	return mfdatamap
}

func (repo *mfService) GetMasterDataMF() {

	t := time.Now()
	Today := t.Format("02-01-2006 15:04:05")

	MasterData := GetMasterData()
	objmasterdata := MasterData.(models.MasterdataMfResponse)

	var objmastertypesense models.MasterdataMfTypesenseModel
	var arrmastertypesense []models.MasterdataMfTypesenseModel

	mfmap := GetMfmap()

	checkSchcodemap := make(map[int64]int64)
	//duplicatedatamap := make(map[int64]int64)

	for k := range mfmap {
		for z := range objmasterdata.Data {

			if strconv.FormatInt(objmasterdata.Data[z].SchCode, 10) == mfmap[k].SchCode {
				if checkSchcodemap[objmasterdata.Data[z].SchCode] != 0 {
					//duplicatedatamap[objmasterdata.Data[z].SchCode] = objmasterdata.Data[z].SchCode
					continue
				}
				checkSchcodemap[objmasterdata.Data[z].SchCode] = objmasterdata.Data[z].SchCode

				amcid := mfmap[k].AMCID
				stramcid := strconv.FormatInt(int64(amcid), 10)
				objmastertypesense.AMCID = stramcid
				objmastertypesense.AMCName = mfmap[k].AMCName
				// schcode := mfmap[k].SchCode
				// strschcode := strconv.FormatInt(schcode, 10)
				objmastertypesense.SchCode = mfmap[k].SchCode //strschcode
				//objmastertypesense.SchCode = objmasterdata.Data[z].SchCode
				objmastertypesense.SchemeName = mfmap[k].SchemeName
				objmastertypesense.Status = objmasterdata.Data[z].Status
				categorycode := objmasterdata.Data[z].CategoryCode
				strcategorycode := strconv.FormatInt(categorycode, 10)
				objmastertypesense.CategoryCode = strcategorycode
				objmastertypesense.Category = mfmap[k].Category
				objmastertypesense.Type = objmasterdata.Data[z].Type
				objmastertypesense.SchemePlan = objmasterdata.Data[z].SchemePlan
				objmastertypesense.Risk = mfmap[k].Risk
				objmastertypesense.LumpSumAllowed = objmasterdata.Data[z].LumpSumAllowed
				objmastertypesense.SIPAllowed = objmasterdata.Data[z].SIPAllowed

				//var s string
				//var f float64
				//fmt.Println("var1 = ", reflect.TypeOf(objmasterdata.Data[z].MinLumpsumInvestment))
				// if reflect.TypeOf(objmasterdata.Data[z].MinLumpsumInvestment) == reflect.TypeOf(objmasterdata.Data[z].MinLumpsumInvestment.(string)) {
				// 	objmastertypesense.MinLumpsumInvestment = objmasterdata.Data[z].MinLumpsumInvestment
				// }
				objmastertypesense.MinLumpsumInvestment = objmasterdata.Data[z].MinLumpsumInvestment

				objmastertypesense.MinSIPInvestment = objmasterdata.Data[z].MinSIPInvestment
				objmastertypesense.Return1Day = objmasterdata.Data[z].Return1Day
				objmastertypesense.Return1Week = objmasterdata.Data[z].Return1Week
				objmastertypesense.Return1Month = objmasterdata.Data[z].Return1Month
				objmastertypesense.Return3Month = objmasterdata.Data[z].Return3Month
				objmastertypesense.Return6Month = objmasterdata.Data[z].Return6Month
				objmastertypesense.Return1Year = objmasterdata.Data[z].Return1Year
				objmastertypesense.Return3Year = objmasterdata.Data[z].Return3Year
				objmastertypesense.Return5Year = objmasterdata.Data[z].Return5Year
				objmastertypesense.ReturnSinceInception = objmasterdata.Data[z].ReturnSinceInception
				//objmastertypesense.NAV = mfmap[k].LatestNAV
				//objmastertypesense.NAVDate = mfmap[k].NAVDate
				objmastertypesense.FMgr1Code = objmasterdata.Data[z].FMgr1Code
				objmastertypesense.FMgr1Name = objmasterdata.Data[z].FMgr1Name
				objmastertypesense.FMgr2Code = objmasterdata.Data[z].FMgr2Code
				objmastertypesense.FMgr2Name = objmasterdata.Data[z].FMgr2Name

				MasterDataInfo := GetMasterDataInfo(objmasterdata.Data[z].SchCode)
				objmasterdatainfo := MasterDataInfo.(models.MasterDataInfoResponse)
				if objmasterdatainfo.Data == nil {
					continue
				}
				fmt.Printf("schcode:%d\n", objmasterdata.Data[z].SchCode)
				fmt.Printf("count:%d\n", z)
				objmastertypesense.NatureDes = objmasterdatainfo.Data[0].NatureDes
				objmastertypesense.SubNatureDes = objmasterdatainfo.Data[0].SubNatureDes
				objmastertypesense.Options = objmasterdatainfo.Data[0].Options
				objmastertypesense.MinInvestment = objmasterdatainfo.Data[0].MinInvestment
				objmastertypesense.Objective = objmasterdatainfo.Data[0].Objective
				objmastertypesense.AMFICode = objmasterdatainfo.Data[0].AMFICode
				objmastertypesense.EntryLoad = objmasterdatainfo.Data[0].EntryLoad
				objmastertypesense.ExitLoad = objmasterdatainfo.Data[0].ExitLoad
				objmastertypesense.ExitLoadDes = objmasterdatainfo.Data[0].ExitLoadDes
				objmastertypesense.ExpenseRatio = objmasterdatainfo.Data[0].ExpenseRatio
				objmastertypesense.AumAsOn = objmasterdatainfo.Data[0].AumAsOn
				objmastertypesense.SchemeAUM = objmasterdatainfo.Data[0].AUM
				objmastertypesense.ISIN_Growth = objmasterdatainfo.Data[0].ISIN_Growth
				objmastertypesense.ISIN_DivPayout = objmasterdatainfo.Data[0].ISIN_DivPayout
				objmastertypesense.ISIN_DivReInvest = objmasterdatainfo.Data[0].ISIN_DivReInvest
				objmastertypesense.FiftyTwoWeekHigh = objmasterdatainfo.Data[0].FiftyTwoWeekHigh
				objmastertypesense.FiftyTwoWeekHighDate = objmasterdatainfo.Data[0].FiftyTwoWeekHighDate
				objmastertypesense.FiftyTwoWeekLow = objmasterdatainfo.Data[0].FiftyTwoWeekLow
				objmastertypesense.FiftyTwoWeekLowDate = objmasterdatainfo.Data[0].FiftyTwoWeekLowDate
				objmastertypesense.OpenDate = objmasterdatainfo.Data[0].OpenDate
				objmastertypesense.Benchmark = objmasterdatainfo.Data[0].Benchmark
				objmastertypesense.Address1 = objmasterdatainfo.Data[0].Address1
				objmastertypesense.Address2 = objmasterdatainfo.Data[0].Address2
				objmastertypesense.Address3 = objmasterdatainfo.Data[0].Address3
				objmastertypesense.City = objmasterdatainfo.Data[0].City
				objmastertypesense.State = objmasterdatainfo.Data[0].State
				objmastertypesense.StdDev = objmasterdatainfo.Data[0].StdDev
				objmastertypesense.NAV = objmasterdatainfo.Data[0].NAV
				objmastertypesense.NAVDate = objmasterdatainfo.Data[0].NAVDate

				objmastertypesense.MinPurAmt = mfmap[k].MinPurAmt
				objmastertypesense.PurAmountMult = mfmap[k].PurAmountMult
				objmastertypesense.MinAddPurAmt = mfmap[k].MinAddPurAmt
				objmastertypesense.BSESchemecode = mfmap[k].SchemeCode
				objmastertypesense.BSEToken = mfmap[k].BSEToken
				objmastertypesense.SIPSchemeInfo2 = mfmap[k].SIPSchemeInfo2

				arrmastertypesense = append(arrmastertypesense, objmastertypesense)
			}
		}

	}

	slcDocument := make([]interface{}, len(arrmastertypesense))

	for i := range arrmastertypesense {

		slcDocument[i] = arrmastertypesense[i]

	}

	if len(slcDocument) > 0 {
		batch := 50
		for i := 0; i < len(slcDocument); i += batch {
			j := i + batch
			if j > len(slcDocument) {
				j = len(slcDocument)
			}
			uploadMasterMFtypesense(slcDocument[i:j])
		}
	}

	elapsed := time.Since(t)
	ReqTime := fmt.Sprintf("%v", elapsed)
	str := Today + "\n Mutual fund data(greenwaremf2) inserted in " + ReqTime + "\n Length of greenwaremf2 data = " + strconv.Itoa(len(arrmastertypesense))
	err := SentToContractCronJob(str)
	if err != nil {
		fmt.Println("Error while sending to slack")
	}
	Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Uploading contract to Typesense End.")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Uploading contract to Typesense End.")
}
