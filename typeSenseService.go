package service

import (
	"contractmaster/helper"
	"contractmaster/models"
	"fmt"
	"strconv"
	"time"

	"github.com/TecXLab/libdb/contracts"
	"github.com/TecXLab/libdb/stockal"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

type TypeSenseService interface {
	GetContractFromDB()
	DropCollection()
}

type typeSenseService struct {
}

func NewTypeSenseService() TypeSenseService {
	return &typeSenseService{}
}

var collectionname = "stocksearch2"

var sort1 = "priorityno" //"cnam"
//var sort2 = "expry"

func (repo *typeSenseService) DropCollection() {

	t := time.Now()
	Today := t.Format("02-01-2006 15:04:05")

	client := typesense.NewClient(
		typesense.WithServer(BaseURL),
		typesense.WithAPIKey(TypessenseValue))

	client.Collection(collectionname).Delete()

	elapsed := time.Since(t)
	ReqTime := fmt.Sprintf("%v", elapsed)
	str := Today + "\n Data of " + collectionname + " collection dropped sucessfully in " + ReqTime
	err := SentToContractCronJob(str)
	if err != nil {
		fmt.Println("Error while sending to slack")
	}

	Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Sucessfully collection dropped from Typesense. ")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Sucessfully collection dropped from Typesense. ")
}

func uploadtypesenseData(objData interface{}) error {
	Zerologs.Info().Msg("uploadtypesenseData Start")
	client := typesense.NewClient(
		typesense.WithServer(BaseURL), typesense.WithAPIKey(TypessenseValue))

	_, err := client.Collection(collectionname).Retrieve()
	if err != nil {
		bfacet := true
		schema := &api.CollectionSchema{
			Name: collectionname,
			Fields: []api.Field{
				{
					Name:  "fullname",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "omexs",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "omtkn",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "cnam",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "tsym",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "exseg",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "uomtkn",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "expry",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "optyp",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "strikprc",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "stktyp",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "seris",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "symdes",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "usym",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "wgt",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "last",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "pchng",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "chng",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "time",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "vol",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "lotSz",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "hprcchg",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "lprcchg",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "opdiff",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "cm_icod",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "cm_inm",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "cm_mc",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "cm_mcty",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "cm_scn",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "cm_scod",
					Type:  "string",
					Facet: &bfacet,
				},
			},
			// DefaultSortingField: &sort1,
		}
		_, err := client.Collections().Create(schema)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(resp)
	}

	objDocModel := models.StockSearchDocumentModel{}

	switch objData.(type) {
	case contracts.Contract_NSEEQ:
		objScripDetails := objData.(contracts.Contract_NSEEQ)
		objDocModel.Fullname = ""
		objDocModel.Chng = 0
		objDocModel.Omtkn = strconv.Itoa(objScripDetails.NToken)
		objDocModel.Cm_icod = ""
		objDocModel.Cm_inm = ""
		objDocModel.Cm_mc = 0
		objDocModel.Cm_mcty = ""
		objDocModel.Cm_scn = ""
		objDocModel.Cm_scod = " "
		objDocModel.Cnam = objScripDetails.SSymbol
		objDocModel.Expry = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
		objDocModel.Last = 0
		objDocModel.Opdiff = 0
		objDocModel.Optyp = objScripDetails.SSeries
		objDocModel.Pchng = 0
		objDocModel.Strikprc = "0"
		objDocModel.Time = "0"
		objDocModel.Vol = 0
		objDocModel.Wgt = "0"
		objDocModel.Exseg = "nse_cm"
		objDocModel.Omexs = objDocModel.Omtkn + "_" + "nse_cm"
		objDocModel.Stktyp = "spot"
		objDocModel.Tsym = objScripDetails.SSymbol + "-" + objScripDetails.SSeries
		objDocModel.Symdes = objScripDetails.SSymbolName
		objDocModel.Uomtkn = "0"
		objDocModel.Seris = objScripDetails.SSeries
		objDocModel.Usym = objScripDetails.SSymbol

	case contracts.Contract_NSEFO:
		objScripDetails := objData.(contracts.Contract_NSEFO)
		objDocModel.Fullname = ""
		objDocModel.Chng = 0
		objDocModel.Omtkn = strconv.Itoa(objScripDetails.NToken)
		objDocModel.Cm_icod = ""
		objDocModel.Cm_inm = ""
		objDocModel.Cm_mc = 0
		objDocModel.Cm_mcty = ""
		objDocModel.Cm_scn = ""
		objDocModel.Cm_scod = ""
		objDocModel.Cnam = objScripDetails.SSymbol
		objDocModel.Expry = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
		objDocModel.Last = 0
		objDocModel.Opdiff = 0
		objDocModel.Optyp = objScripDetails.SSeries
		objDocModel.Pchng = 0
		objDocModel.Strikprc = "0"
		objDocModel.Time = "0"
		objDocModel.Vol = 0
		objDocModel.Wgt = "0"

		if objScripDetails.NStrikePrice > 0 {
			objDocModel.Exseg = "nse_fo"
			objDocModel.Omexs = objDocModel.Omtkn + "_" + "nse_fo"
			objDocModel.Stktyp = "option"
			objDocModel.Tsym = objScripDetails.SSymbol + helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + objScripDetails.SOptionType + strconv.Itoa(objScripDetails.NStrikePrice)
			objDocModel.Uomtkn = strconv.Itoa(objScripDetails.NAssetToken)
			objDocModel.Symdes = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + " " + objScripDetails.SOptionType + " " + strconv.Itoa(objScripDetails.NStrikePrice)
			objDocModel.Seris = objScripDetails.SInstrumentName
			objDocModel.Usym = objScripDetails.SSymbol
		} else {
			objDocModel.Exseg = "nse_fo"
			objDocModel.Omexs = objDocModel.Omtkn + "_" + "nse_fo"
			objDocModel.Stktyp = "future"
			objDocModel.Tsym = objScripDetails.SSymbol + helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
			objDocModel.Uomtkn = strconv.Itoa(objScripDetails.NAssetToken)
			objDocModel.Symdes = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + " FUT"
			objDocModel.Seris = objScripDetails.SInstrumentName
			objDocModel.Usym = objScripDetails.SSymbol
		}
	case contracts.Contract_NSECD:
		fmt.Println("NSCD")
	default:
		fmt.Println("unknown")
	}
	document := objDocModel

	_, err = client.Collection(collectionname).Documents().Upsert(document)
	if err != nil {
		Zerologs.Error().Err(err).Msgf("Error in uploadtypesenseData")
		fmt.Println(err)
		return err
	}
	return nil
}

func uploadtypesenseData1(objData []interface{}) error {
	Zerologs.Info().Msg("uploadtypesenseData Start")
	client := typesense.NewClient(
		typesense.WithServer(BaseURL),
		typesense.WithAPIKey(TypessenseValue))

	_, err := client.Collection(collectionname).Retrieve()
	if err != nil {
		bfacet := true
		schema := &api.CollectionSchema{
			Name: collectionname,
			Fields: []api.Field{
				{
					Name:  "fullname",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "omexs",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "omtkn",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "cnam",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "tsym",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "exseg",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "uomtkn",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "expry",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "nexpry",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "optyp",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "optyp2",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "strikprc",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "teststrikprc",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "stktyp",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "seris",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "symdes",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "usym",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "wgt",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "last",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "pchng",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "chng",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "time",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "vol",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "lotSz",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "hprcchg",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "lprcchg",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "opdiff",
					Type:  "float",
					Facet: &bfacet,
				},
				{
					Name:  "cm_icod",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "cm_inm",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "cm_mc",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "cm_mcty",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "cm_scn",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "cm_scod",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "exmnt",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "exdate",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "exyear",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "exdtemnt",
					Type:  "string",
					Facet: &bfacet,
				},
				{
					Name:  "priorityno",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "nInstrumentType",
					Type:  "int64",
					Facet: &bfacet,
				},
				{
					Name:  "sisin",
					Type:  "string",
					Facet: &bfacet,
				},
			},
			DefaultSortingField: &sort1,
		}
		_, err := client.Collections().Create(schema)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(resp)
	}
	// data111, err := json.Marshal(objData)
	document := objData
	var action = "upsert"
	var batch = 40
	params := &api.ImportDocumentsParams{
		Action:    &action,
		BatchSize: &batch,
	}

	_, err = client.Collection(collectionname).Documents().Import(document, params)
	// _, err = client.Collection(collectionname).Documents().Upsert(document)
	if err != nil {
		Zerologs.Error().Err(err).Msgf("Error in uploadtypesenseData")
		fmt.Println(err)
		return err
	}
	return nil
}

func getContractDay() int64 {
	t1 := time.Date(2018, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Now()
	days := int64(t2.Sub(t1).Hours() / 24)
	return days
}

var ExchId = map[string]int16{
	"nsefo":         1,
	"nseeq":         2,
	"nsecd":         3,
	"nseeq_Indices": 4,
}
var EnumNSECMInstrumentType = map[string]int{
	"Equities":          0,
	"Preference_Shares": 1,
	"Debentures":        2,
	"Warrants":          3,
	"Miscellaneous":     4,
}

var ScripType = map[string]int16{
	"SPOT":    1,
	"FUTURES": 2,
	"OPTIONS": 3,
}

// var EnumExchange = map[string]int16{
// 	"nse_cm": 1,
// 	"nse_fo": 2,
// }

func (repo *typeSenseService) GetContractFromDB() {
	t := time.Now()
	Today := t.Format("02-01-2006 15:04:05")

	Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Contract Reading from DB Start.")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Contract Reading from DB Start.")
	var (
		nsefo            []contracts.Contract_NSEFO
		nseeq            []contracts.Contract_NSEEQ
		nsecd            []contracts.Contract_NSECD
		indexmaster      []contracts.IndexMaster
		instrumentmaster []stockal.InstrumentMaster
	)
	days := getContractDay()

	Db.Client.Where("n_contract_date=? AND (s_sec_normal_mkt_eligibility=? OR s_sec_additional_mkt_eligibility2=?) AND (n_instrument_type = ? OR n_instrument_type = ?) AND s_series != ?", days, "1", "1", EnumNSECMInstrumentType["Equities"], EnumNSECMInstrumentType["Miscellaneous"], "MF").Find(&nseeq)

	Db.Client.Where("n_contract_date=?", days).Find(&nsefo)

	//Db.Client.Where("n_contract_date=?", days).Find(&nsecd)
	Db.Client.Where("n_contract_date=? AND (s_instrument_name=? OR s_instrument_name=? OR s_instrument_name=? OR s_instrument_name=? OR s_instrument_name=?)", days, "OPTCUR", "OPTIRC", "FUTCUR", "FUTIRT", "FUTIRC").Find(&nsecd)

	Db.Client.Find(&indexmaster)
	StockalDB.Stockal.Find(&instrumentmaster)

	fmt.Println("nsefo Count: ", len(nsefo))
	fmt.Println("nseeq Count: ", len(nseeq))
	fmt.Println("nsecd Count: ", len(nsecd))
	fmt.Println("indexmaster Count: ", len(indexmaster))
	fmt.Println("Stockal instrumentmaster Count: ", len(instrumentmaster))
	Zerologs.Info().Msgf("nseeq Count %d", len(nsefo))
	Zerologs.Info().Msgf("nseeq Count %d", len(nseeq))
	Zerologs.Info().Msgf("nsecd Count %d", len(nsecd))
	Zerologs.Info().Msgf("indexmaster Count %d", len(indexmaster))

	Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Contract Reading from DB End.")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Contract Reading from DB End.")

	Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Uploading contract to Typesense Start.")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Uploading contract to Typesense Start.")

	if len(nseeq) > 0 {
		doc := GenerateNSEQ(nseeq)
		if len(doc) > 0 {
			batch := 100
			for i := 0; i < len(doc); i += batch {
				j := i + batch
				if j > len(doc) {
					j = len(doc)
				}
				uploadtypesenseData1(doc[i:j]) // Process the batch.
				fmt.Println("Uploading typesense data...", i, j)
			}
		}

		Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Uploading NSEEQ contract to Typesense Done.")
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Uploading NSEEQ contract to Typesense Done.")
	}

	if len(nsefo) > 0 {
		doc := GenerateNSFO(nsefo)
		if len(doc) > 0 {
			batch := 100
			for i := 0; i < len(doc); i += batch {
				j := i + batch
				if j > len(doc) {
					j = len(doc)
				}
				uploadtypesenseData1(doc[i:j]) // Process the batch.
				fmt.Println("Uploading typesense data...", i, j)
			}
			Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Uploading NSEFO contract to Typesense Done.")
			fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Uploading NSEFO contract to Typesense Done.")
		}
	}

	if len(nsecd) > 0 {
		doc := GenerateNSCD(nsecd)
		if len(doc) > 0 {
			batch := 100
			for i := 0; i < len(doc); i += batch {
				j := i + batch
				if j > len(doc) {
					j = len(doc)
				}
				uploadtypesenseData1(doc[i:j]) // Process the batch.
				fmt.Println("Uploading typesense data...", i, j)
			}
		}
		Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Uploading NSECD contract to Typesense Done.")
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Uploading NSECD  contract to Typesense Done.")
	}

	if len(indexmaster) > 0 {
		doc := GenerateIndexMaster(indexmaster)
		if len(doc) > 0 {
			uploadtypesenseData1(doc)
		}
		Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Uploading Indexmaster data to Typesense Done.")
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Uploading Indexmaster data to Typesense Done.")
	}

	if len(instrumentmaster) > 0 {
		doc := GenerateStockal(instrumentmaster)
		if len(doc) > 0 {
			batch := 100
			for i := 0; i < len(doc); i += batch {
				j := i + batch
				if j > len(doc) {
					j = len(doc)
				}
				uploadtypesenseData1(doc[i:j]) // Process the batch.
				fmt.Println("Uploading typesense data...", i, j)
			}
		}

		Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Uploading Stockal contract to Typesense Done.")
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Uploading Stockal contract to Typesense Done.")
	}

	elapsed := time.Since(t)
	ReqTime := fmt.Sprintf("%v", elapsed)
	str := Today + "\n stocksearch2(Foocut-Dev) data inserted in " + ReqTime + "\n Length of nseeq = " + strconv.Itoa(len(nseeq)) + "\n Length of nsefo = " + strconv.Itoa(len(nsefo)) + "\n Length of nsecd = " + strconv.Itoa(len(nsecd)) + "\n Length of indexmaster = " + strconv.Itoa(len(indexmaster)) + "\n Length of stockal = " + strconv.Itoa(len(instrumentmaster))
	err := SentToContractCronJob(str)
	if err != nil {
		fmt.Println("Error while sending to slack")
	}
	Zerologs.Info().Msg(time.Now().Format("2006-01-02 15:04:05") + " Uploading contract to Typesense End.")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Uploading contract to Typesense End.")
}

func GenerateNSEQ(nseeq []contracts.Contract_NSEEQ) []interface{} {
	if len(nseeq) > 0 {
		// slcDocument := []models.StockSearchDocumentModel{}
		// There is no need to check, we want to panic if it's not slice or array
		slcDocument := make([]interface{}, len(nseeq))
		// for i := 0; i < v.Len(); i++ {
		// 	intf[i] = v.Index(i).Interface()
		// }
		// slcDocument = []interface{}
		for i := range nseeq {
			//fmt.Println("NInstrumentType:- ", nseeq[i].NInstrumentType, "	Fullname:- "+nseeq[i].SSymbolName+"	", i)
			objDocModel := models.StockSearchDocumentModel{}
			objScripDetails := nseeq[i]
			objDocModel.Fullname = objScripDetails.SSymbolName
			objDocModel.Chng = 0
			objDocModel.Omtkn = strconv.Itoa(objScripDetails.NToken)
			objDocModel.Cm_icod = ""
			objDocModel.Cm_inm = ""
			objDocModel.Cm_mc = 0
			objDocModel.Cm_mcty = ""
			objDocModel.Cm_scn = ""
			objDocModel.Cm_scod = " "
			objDocModel.Cnam = objScripDetails.SSymbol
			objDocModel.Expry = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
			objDocModel.NExpry = int64(objScripDetails.NExpiryDate)
			objDocModel.Last = 0
			objDocModel.Opdiff = 0
			objDocModel.Optyp = objScripDetails.SSeries
			objDocModel.Pchng = 0
			objDocModel.Strikprc = "0"
			objDocModel.Time = "0"
			objDocModel.Vol = 0
			objDocModel.Wgt = "0"
			objDocModel.Exseg = "nse_cm"
			objDocModel.Omexs = objDocModel.Omtkn + "_" + "nse_cm"
			objDocModel.Stktyp = "spot"
			objDocModel.Tsym = objScripDetails.SSymbol + "-" + objScripDetails.SSeries
			objDocModel.Symdes = objScripDetails.SSymbolName
			objDocModel.Uomtkn = "0"
			objDocModel.Seris = objScripDetails.SSeries
			objDocModel.Usym = objScripDetails.SSymbol
			//objDocModel.Priorityno = 1
			//objDocModel.NInstrumentType = int64(objScripDetails.NInstrumentType)
			if objScripDetails.NInstrumentType == 0 {
				objDocModel.Priorityno = 1
				objDocModel.NInstrumentType = 0
			} else if objScripDetails.NInstrumentType == 4 {
				objDocModel.Priorityno = 5
				objDocModel.NInstrumentType = 10
				objDocModel.Seris = "Etf"
			}
			objDocModel.SISIN = objScripDetails.SISIN
			//fmt.Println("#Series:- "+objDocModel.Seris+"	Fullname:- "+objDocModel.Fullname+"	", i)
			slcDocument[i] = objDocModel
		}
		return slcDocument
	}
	return nil
}

func GenerateNSFO(nsefo []contracts.Contract_NSEFO) []interface{} {
	if len(nsefo) > 0 {
		// slcDocument := []models.StockSearchDocumentModel{}
		// There is no need to check, we want to panic if it's not slice or array
		slcDocument := make([]interface{}, len(nsefo))
		// for i := 0; i < v.Len(); i++ {
		// 	intf[i] = v.Index(i).Interface()
		// }
		// slcDocument = []interface{}
		days := getContractDay()
		var nseeq []contracts.Contract_NSEEQ
		Db.Client.Where("n_contract_date=?", days).Find(&nseeq)
		IsinMap := make(map[int]string)
		for i := 0; i < len(nseeq); i++ {
			IsinMap[nseeq[i].NToken] = nseeq[i].SSymbolName
		}
		for i := range nsefo {
			objDocModel := models.StockSearchDocumentModel{}
			objScripDetails := nsefo[i]
			objDocModel.Fullname = IsinMap[objScripDetails.NAssetToken] //objScripDetails.SSymbolFullName
			objDocModel.Chng = 0
			objDocModel.Omtkn = strconv.Itoa(objScripDetails.NToken)
			objDocModel.Cm_icod = ""
			objDocModel.Cm_inm = ""
			objDocModel.Cm_mc = 0
			objDocModel.Cm_mcty = ""
			objDocModel.Cm_scn = ""
			objDocModel.Cm_scod = ""
			objDocModel.Cnam = objScripDetails.SSymbol
			objDocModel.Expry = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
			objDocModel.NExpry = int64(objScripDetails.NExpiryDate)
			objDocModel.Last = 0
			objDocModel.Opdiff = 0
			objDocModel.Pchng = 0
			objDocModel.Time = "0"
			objDocModel.Vol = 0
			objDocModel.Wgt = "0"
			//objDocModel.Exmnt = helper.ConvertDateTime1980OnlyMonth(int64(objScripDetails.NExpiryDate))
			ArrDate := helper.ConvertDateTimeTo1980DateSplit(objScripDetails.NExpiryDate)
			objDocModel.Exdate = ArrDate[0]
			objDocModel.Exmnt = ArrDate[1]
			objDocModel.Exyear = ArrDate[2]
			objDocModel.Exdtemnt = ArrDate[0] + " " + ArrDate[1]

			if objScripDetails.NStrikePrice > 0 {
				objDocModel.Exseg = "nse_fo"
				objDocModel.Omexs = objDocModel.Omtkn + "_" + "nse_fo"
				objDocModel.Stktyp = "option"
				objDocModel.Tsym = objScripDetails.SSymbol + helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + objScripDetails.SOptionType + strconv.Itoa(objScripDetails.NStrikePrice/100)
				objDocModel.Uomtkn = strconv.Itoa(objScripDetails.NAssetToken)
				objDocModel.Symdes = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + " " + strconv.Itoa(objScripDetails.NStrikePrice/100) + " " + objScripDetails.SOptionType
				objDocModel.Seris = objScripDetails.SInstrumentName
				objDocModel.Usym = objScripDetails.SSymbol
				objDocModel.Strikprc = strconv.Itoa(objScripDetails.NStrikePrice / 100)
				objDocModel.TestStrikprc = objDocModel.Strikprc[0:1]
				if objScripDetails.SSymbol == "MOTHERSON" {
					var StrikePrice float64
					StrikePrice = float64(objScripDetails.NStrikePrice)
					StrikePrice = StrikePrice / 100
					objDocModel.Strikprc = fmt.Sprintf("%v", StrikePrice)
					objDocModel.Symdes = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + " " + fmt.Sprintf("%v", StrikePrice) + " " + objScripDetails.SOptionType
					objDocModel.Tsym = objScripDetails.SSymbol + helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + objScripDetails.SOptionType + fmt.Sprintf("%v", StrikePrice) //strconv.Itoa(objScripDetails.NStrikePrice/100)
				}
				objDocModel.Optyp = objScripDetails.SOptionType
				if objDocModel.Optyp == "PE" {
					objDocModel.Optyp2 = "PUT"
				} else if objDocModel.Optyp == "CE" {
					objDocModel.Optyp2 = "CALL"
				}
				objDocModel.Priorityno = 3
				objDocModel.NInstrumentType = 8
				// if objScripDetails.SSymbol == "MIDCPNIFTY" {
				// 	objDocModel.Priorityno = 5
				// 	objDocModel.NInstrumentType = 7
				// }
				if objScripDetails.SSymbol == "NIFTY" {
					objDocModel.Priorityno = 2
					objDocModel.NInstrumentType = 3
					objDocModel.Seris = "indicesnifty"
				}
				if objScripDetails.SSymbol == "BANKNIFTY" {
					objDocModel.Priorityno = 2
					objDocModel.NInstrumentType = 6
					objDocModel.Seris = "indicesnifty"
				}
				objDocModel.SISIN = objScripDetails.SISIN

			} else {
				objDocModel.Exseg = "nse_fo"
				objDocModel.Omexs = objDocModel.Omtkn + "_" + "nse_fo"
				objDocModel.Stktyp = "future"
				objDocModel.Tsym = objScripDetails.SSymbol + helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
				objDocModel.Uomtkn = strconv.Itoa(objScripDetails.NAssetToken)
				objDocModel.Symdes = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + " FUT"
				objDocModel.Seris = objScripDetails.SInstrumentName
				objDocModel.Usym = objScripDetails.SSymbol
				objDocModel.Strikprc = "0"
				objDocModel.Optyp = objScripDetails.SSeries
				objDocModel.Priorityno = 3
				objDocModel.NInstrumentType = 7
				if objScripDetails.SSymbol == "NIFTY" {
					objDocModel.Priorityno = 2
					objDocModel.NInstrumentType = 2
					objDocModel.Seris = "indicesnifty"
				}
				if objScripDetails.SSymbol == "BANKNIFTY" {
					objDocModel.Priorityno = 2
					objDocModel.NInstrumentType = 5
					objDocModel.Seris = "indicesnifty"
				}
				objDocModel.SISIN = objScripDetails.SISIN
			}
			slcDocument[i] = objDocModel
		}
		return slcDocument
	}
	return nil
}

func GenerateNSCD(nsecd []contracts.Contract_NSECD) []interface{} {
	if len(nsecd) > 0 {
		slcDocument := make([]interface{}, len(nsecd))
		var nstrikprice float64
		for i := range nsecd {
			objDocModel := models.StockSearchDocumentModel{}
			objScripDetails := nsecd[i]

			if objScripDetails.NStrikePrice > 0 {
				if objScripDetails.SInstrumentName == "OPTCUR" || objScripDetails.SInstrumentName == "OPTIRC" {
					nstrikprice = float64(objScripDetails.NStrikePrice)
					nstrikprice = nstrikprice / 10000000
					objDocModel.Fullname = objScripDetails.SSymbolName
					objDocModel.Chng = 0
					objDocModel.Omtkn = strconv.Itoa(objScripDetails.NToken)
					objDocModel.Cm_icod = ""
					objDocModel.Cm_inm = ""
					objDocModel.Cm_mc = 0
					objDocModel.Cm_mcty = ""
					objDocModel.Cm_scn = ""
					objDocModel.Cm_scod = ""
					objDocModel.Cnam = objScripDetails.SSymbol
					objDocModel.Expry = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
					objDocModel.NExpry = int64(objScripDetails.NExpiryDate)
					objDocModel.Last = 0
					objDocModel.Opdiff = 0
					objDocModel.Pchng = 0
					objDocModel.Time = "0"
					objDocModel.Vol = 0
					objDocModel.Wgt = "0"
					//objDocModel.Exmnt = helper.ConvertDateTime1980OnlyMonth(int64(objScripDetails.NExpiryDate))
					ArrDate := helper.ConvertDateTimeTo1980DateSplit(objScripDetails.NExpiryDate)
					objDocModel.Exdate = ArrDate[0]
					objDocModel.Exmnt = ArrDate[1]
					objDocModel.Exyear = ArrDate[2]
					objDocModel.Exdtemnt = ArrDate[0] + " " + ArrDate[1]
					objDocModel.Exseg = "cde_fo"
					objDocModel.Omexs = objDocModel.Omtkn + "_" + "cde_fo"
					objDocModel.Stktyp = "opt"
					objDocModel.Tsym = objScripDetails.SSymbol + helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + objScripDetails.SOptionType + fmt.Sprintf("%v", nstrikprice)
					objDocModel.Uomtkn = strconv.Itoa(objScripDetails.NAssetToken)
					objDocModel.Symdes = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + " " + fmt.Sprintf("%v", nstrikprice) + " " + objScripDetails.SOptionType
					objDocModel.Seris = objScripDetails.SInstrumentName
					objDocModel.Usym = objScripDetails.SSymbol
					objDocModel.Strikprc = fmt.Sprintf("%v", nstrikprice)
					objDocModel.TestStrikprc = objDocModel.Strikprc[0:1]
					//fmt.Println(objDocModel.Strikprc)
					objDocModel.Optyp = objScripDetails.SOptionType
					if objDocModel.Optyp == "PE" {
						objDocModel.Optyp2 = "PUT"
					} else if objDocModel.Optyp == "CE" {
						objDocModel.Optyp2 = "CALL"
					}
					if objScripDetails.SInstrumentName == "OPTCUR" {
						objDocModel.Priorityno = 6
						objDocModel.NInstrumentType = 14
					} else if objScripDetails.SInstrumentName == "OPTIRC" {
						objDocModel.Priorityno = 6
						objDocModel.NInstrumentType = 15
					}
				}
			} else {
				if objScripDetails.SInstrumentName == "FUTCUR" || objScripDetails.SInstrumentName == "FUTIRT" || objScripDetails.SInstrumentName == "FUTIRC" {
					objDocModel.Fullname = objScripDetails.SSymbolName
					objDocModel.Chng = 0
					objDocModel.Omtkn = strconv.Itoa(objScripDetails.NToken)
					objDocModel.Cm_icod = ""
					objDocModel.Cm_inm = ""
					objDocModel.Cm_mc = 0
					objDocModel.Cm_mcty = ""
					objDocModel.Cm_scn = ""
					objDocModel.Cm_scod = ""
					objDocModel.Cnam = objScripDetails.SSymbol
					objDocModel.Expry = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
					objDocModel.NExpry = int64(objScripDetails.NExpiryDate)
					objDocModel.Last = 0
					objDocModel.Opdiff = 0
					objDocModel.Pchng = 0
					objDocModel.Time = "0"
					objDocModel.Vol = 0
					objDocModel.Wgt = "0"
					//objDocModel.Exmnt = helper.ConvertDateTime1980OnlyMonth(int64(objScripDetails.NExpiryDate))
					ArrDate := helper.ConvertDateTimeTo1980DateSplit(objScripDetails.NExpiryDate)
					objDocModel.Exdate = ArrDate[0]
					objDocModel.Exmnt = ArrDate[1]
					objDocModel.Exyear = ArrDate[2]
					objDocModel.Exdtemnt = ArrDate[0] + " " + ArrDate[1]
					objDocModel.Exseg = "cde_fo"
					objDocModel.Omexs = objDocModel.Omtkn + "_" + "cde_fo"
					objDocModel.Stktyp = "fut"
					objDocModel.Tsym = objScripDetails.SSymbol + helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
					objDocModel.Uomtkn = strconv.Itoa(objScripDetails.NAssetToken)
					objDocModel.Symdes = helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate)) + " FUT"
					objDocModel.Seris = objScripDetails.SInstrumentName
					objDocModel.Usym = objScripDetails.SSymbol
					objDocModel.Strikprc = "0"
					objDocModel.Optyp = objScripDetails.SSeries
					if objScripDetails.SInstrumentName == "FUTCUR" {
						objDocModel.Priorityno = 6
						objDocModel.NInstrumentType = 11
					} else if objScripDetails.SInstrumentName == "FUTIRT" {
						objDocModel.Priorityno = 6
						objDocModel.NInstrumentType = 12
					} else if objScripDetails.SInstrumentName == "FUTIRC" {
						objDocModel.Priorityno = 6
						objDocModel.NInstrumentType = 13
					}
				}
			}
			slcDocument[i] = objDocModel
		}
		return slcDocument
	}
	return nil
}

func GenerateIndexMaster(indexmaster []contracts.IndexMaster) []interface{} {
	if len(indexmaster) > 0 {

		slcDocument := make([]interface{}, len(indexmaster))
		for i := range indexmaster {
			objDocModel := models.StockSearchDocumentModel{}
			objScripDetails := indexmaster[i]

			if indexmaster[i].SIndexName == "Nifty Bank" {
				objDocModel.Fullname = "Bank Nifty" //objScripDetails.SIndexName //objScripDetails.SSymbolName
				objDocModel.Chng = 0
				objDocModel.Omtkn = strconv.Itoa(objScripDetails.NtokenNo)
				objDocModel.Cm_icod = ""
				objDocModel.Cm_inm = ""
				objDocModel.Cm_mc = 0
				objDocModel.Cm_mcty = ""
				objDocModel.Cm_scn = ""
				objDocModel.Cm_scod = " "
				objDocModel.Cnam = "Bank Nifty" //objScripDetails.SIndexName
				objDocModel.Expry = ""          //helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
				objDocModel.Last = 0
				objDocModel.Opdiff = 0
				objDocModel.Optyp = "EQ" //objScripDetails.SSeries
				objDocModel.Pchng = 0
				objDocModel.Strikprc = "0"
				objDocModel.Time = "0"
				objDocModel.Vol = 0
				objDocModel.Wgt = "0"
				objDocModel.Exseg = "nse_indices"                           //"nse_cm"
				objDocModel.Omexs = objDocModel.Omtkn + "_" + "nse_indices" //"nse_cm"
				objDocModel.Stktyp = "spot"
				objDocModel.Tsym = "Bank Nifty" + "-" + "EQ" //objScripDetails.SIndexName + "-" + "EQ"
				objDocModel.Symdes = "Bank Nifty"            //objScripDetails.SIndexName
				objDocModel.Uomtkn = "0"
				objDocModel.Usym = "BANKNIFTY" //objScripDetails.S_Symbol
				objDocModel.Priorityno = 2
				objDocModel.NInstrumentType = 4
				objDocModel.Seris = "indicesnifty"
				slcDocument[i] = objDocModel
			} else if indexmaster[i].SIndexName == "Nifty 50" {
				objDocModel.Fullname = objScripDetails.SIndexName //objScripDetails.SSymbolName
				objDocModel.Chng = 0
				objDocModel.Omtkn = strconv.Itoa(objScripDetails.NtokenNo)
				objDocModel.Cm_icod = ""
				objDocModel.Cm_inm = ""
				objDocModel.Cm_mc = 0
				objDocModel.Cm_mcty = ""
				objDocModel.Cm_scn = ""
				objDocModel.Cm_scod = " "
				objDocModel.Cnam = objScripDetails.SIndexName
				objDocModel.Expry = "" //helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
				objDocModel.Last = 0
				objDocModel.Opdiff = 0
				objDocModel.Optyp = "EQ" //objScripDetails.SSeries
				objDocModel.Pchng = 0
				objDocModel.Strikprc = "0"
				objDocModel.Time = "0"
				objDocModel.Vol = 0
				objDocModel.Wgt = "0"
				objDocModel.Exseg = "nse_indices"                           //"nse_cm"
				objDocModel.Omexs = objDocModel.Omtkn + "_" + "nse_indices" //"nse_cm"
				objDocModel.Stktyp = "spot"
				objDocModel.Tsym = objScripDetails.SIndexName + "-" + "EQ"
				objDocModel.Symdes = objScripDetails.SIndexName
				objDocModel.Uomtkn = "0"
				objDocModel.Usym = objScripDetails.S_Symbol
				objDocModel.Priorityno = 2
				objDocModel.NInstrumentType = 1
				objDocModel.Seris = "indicesnifty"
				slcDocument[i] = objDocModel
			} else {
				objDocModel.Fullname = objScripDetails.SIndexName //objScripDetails.SSymbolName
				objDocModel.Chng = 0
				objDocModel.Omtkn = strconv.Itoa(objScripDetails.NtokenNo)
				objDocModel.Cm_icod = ""
				objDocModel.Cm_inm = ""
				objDocModel.Cm_mc = 0
				objDocModel.Cm_mcty = ""
				objDocModel.Cm_scn = ""
				objDocModel.Cm_scod = " "
				objDocModel.Cnam = objScripDetails.SIndexName
				objDocModel.Expry = "" //helper.ConvertDateTime1980(int64(objScripDetails.NExpiryDate))
				objDocModel.Last = 0
				objDocModel.Opdiff = 0
				objDocModel.Optyp = "EQ" //objScripDetails.SSeries
				objDocModel.Pchng = 0
				objDocModel.Strikprc = "0"
				objDocModel.Time = "0"
				objDocModel.Vol = 0
				objDocModel.Wgt = "0"
				objDocModel.Exseg = "nse_indices"                           //"nse_cm"
				objDocModel.Omexs = objDocModel.Omtkn + "_" + "nse_indices" //"nse_cm"
				objDocModel.Stktyp = "spot"
				objDocModel.Tsym = objScripDetails.SIndexName + "-" + "EQ"
				objDocModel.Symdes = objScripDetails.SIndexName
				objDocModel.Uomtkn = "0"
				//objDocModel.Seris = "EQ"
				objDocModel.Seris = "nse_indices"
				objDocModel.Usym = objScripDetails.S_Symbol
				objDocModel.Priorityno = 4
				objDocModel.NInstrumentType = 9
				slcDocument[i] = objDocModel
			}
		}
		return slcDocument
	}
	return nil
}

func GenerateStockal(Instrumentmaster []stockal.InstrumentMaster) []interface{} {
	if len(Instrumentmaster) > 0 {
		// slcDocument := []models.StockSearchDocumentModel{}
		// There is no need to check, we want to panic if it's not slice or array
		slcDocument := make([]interface{}, len(Instrumentmaster))
		// for i := 0; i < v.Len(); i++ {
		// 	intf[i] = v.Index(i).Interface()
		// }
		// slcDocument = []interface{}
		for i := range Instrumentmaster {
			//fmt.Println("NInstrumentType:- ", nseeq[i].NInstrumentType, "	Fullname:- "+nseeq[i].SSymbolName+"	", i)
			objDocModel := models.StockSearchDocumentModel{}
			objScripDetails := Instrumentmaster[i]
			objDocModel.Fullname = objScripDetails.Company
			objDocModel.Chng = 0
			objDocModel.Omtkn = strconv.Itoa(objScripDetails.Token)
			objDocModel.Cm_icod = ""
			objDocModel.Cm_inm = ""
			objDocModel.Cm_mc = 0
			objDocModel.Cm_mcty = ""
			objDocModel.Cm_scn = ""
			objDocModel.Cm_scod = " "
			objDocModel.Cnam = objScripDetails.Symbol
			objDocModel.Expry = "0"
			objDocModel.NExpry = 0
			objDocModel.Last = 0
			objDocModel.Opdiff = 0
			objDocModel.Optyp = objScripDetails.Series
			objDocModel.Pchng = 0
			objDocModel.Strikprc = "0"
			objDocModel.Time = "0"
			objDocModel.Vol = 0
			objDocModel.Wgt = "0"
			objDocModel.Exseg = "USSTOCK"
			objDocModel.Omexs = objDocModel.Omtkn + "_" + "stockal"
			objDocModel.Stktyp = "spot"
			objDocModel.Tsym = objScripDetails.Symbol + "-" + objScripDetails.Series
			objDocModel.Symdes = objScripDetails.Company
			objDocModel.Uomtkn = "0"
			objDocModel.Seris = objScripDetails.Series
			objDocModel.Usym = objScripDetails.Symbol
			//objDocModel.Priorityno = 1
			//objDocModel.NInstrumentType = int64(objScripDetails.NInstrumentType)
			// if objScripDetails.NInstrumentType == 0 {
			// 	objDocModel.Priorityno = 1
			// 	objDocModel.NInstrumentType = 0
			// } else if objScripDetails.NInstrumentType == 4 {
			// 	objDocModel.Priorityno = 5
			// 	objDocModel.NInstrumentType = 10
			// 	objDocModel.Seris = "Etf"
			// }
			objDocModel.Priorityno = 1
			objDocModel.NInstrumentType = 0
			objDocModel.SISIN = objScripDetails.Isin
			//fmt.Println("#Series:- "+objDocModel.Seris+"	Fullname:- "+objDocModel.Fullname+"	", i)
			slcDocument[i] = objDocModel
		}
		return slcDocument
	}
	return nil
}
