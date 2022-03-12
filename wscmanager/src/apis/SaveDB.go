package wsc_apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	wsc_jsonstructs "wscmanager.com/jsonstructs"
)

func SaveDB(c *gin.Context) {
	AuthorizaionHeader := c.GetHeader("Authorizaion")

	if AuthorizaionHeader == "" || !strings.Contains(AuthorizaionHeader, "Bearer secret_") {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get Authorization header.",
		})
		return
	}

	var Crawljson wsc_jsonstructs.SaveDBjson

	binderr := c.BindJSON(&Crawljson)
	if binderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get body.",
		})
		return
	}

	if Crawljson.StartDate == "" || Crawljson.EndDate == "" || Crawljson.Target == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Insufficient parameters.",
		})
		return
	}

	url := "https://api.notion.com/v1/databases/12a67850f2c243bba567110741f39ef7/query"

	req, reqerr := http.NewRequest("POST", url, nil)
	if reqerr != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Notion-Version", "2022-02-22")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", AuthorizaionHeader)

	res, doerr := http.DefaultClient.Do(req)
	if doerr != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if res.StatusCode == http.StatusUnauthorized {
		c.Status(http.StatusUnauthorized)
		return
	}
	defer res.Body.Close()

	body, readerr := ioutil.ReadAll(res.Body)
	if readerr != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})

	marshalerr := json.Unmarshal(body, &data)
	if marshalerr != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	length := len(data["results"].([]interface{}))
	startTime, _ := time.Parse("2006-01-02", Crawljson.StartDate)
	EndTime, _ := time.Parse("2006-01-02", Crawljson.EndDate)

	for i := 0; i < length; i++ {
		titledata, titlecheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Title"].(map[string]interface{})["title"].([]interface{})
		typedata, typecheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Type"].(map[string]interface{})["select"].(map[string]interface{})
		vulnerabilitydata, vulnerabilitycheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Vulnerability"].(map[string]interface{})["multi_select"].([]interface{})
		datedata, datecheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["FindDate"].(map[string]interface{})["date"].(map[string]interface{})
		targetdata, targetcheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Target"].(map[string]interface{})["select"].(map[string]interface{})
		humandata, humancheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Human"].(map[string]interface{})["people"].([]interface{})
		pageid, pageidcheck := data["results"].([]interface{})[i].(map[string]interface{})["id"]
		var sqltitledata string = ""
		var sqltypedata string = ""
		var sqlvulnerabilitydata string = ""
		var sqldatedata string = ""
		var sqltargetdata string = ""
		var sqlhumandata string = ""
		var sqldescription string = ""
		var sqlreference string = ""

		if !(titlecheck && typecheck && vulnerabilitycheck && datecheck && targetcheck && humancheck && pageidcheck) {
			continue
		}

		if Crawljson.Target != targetdata["name"].(string) {
			continue
		} else {
			sqltargetdata = targetdata["name"].(string)
		}

		dataDate, _ := time.Parse("2006-01-02", datedata["start"].(string))

		if startTime.After(dataDate) || EndTime.Before(dataDate) {
			continue
		} else {
			sqldatedata = datedata["start"].(string)
		}

		if len(titledata) == 0 {
			continue
		} else {
			sqltitledata = titledata[0].(map[string]interface{})["plain_text"].(string)
		}

		if len(humandata) == 0 {
			continue
		} else {
			sqlhumandata = humandata[0].(map[string]interface{})["name"].(string)
		}

		sqltypedata = typedata["name"].(string)
		sqlvulnerabilitydata = vulnerabilitydata[0].(map[string]interface{})["name"].(string)

		pageapiurl := "https://api.notion.com/v1/blocks/" + pageid.(string) + "/children"

		pageapireq, pageapireqdoerr := http.NewRequest("GET", pageapiurl, nil)
		if pageapireqdoerr != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		pageapireq.Header.Add("Accept", "application/json")
		pageapireq.Header.Add("Notion-Version", "2022-02-22")
		pageapireq.Header.Add("Authorization", AuthorizaionHeader)

		pageapires, pageapiresdoerr := http.DefaultClient.Do(pageapireq)
		if pageapiresdoerr != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		defer pageapires.Body.Close()

		pageapibody, pageapireaderr := ioutil.ReadAll(pageapires.Body)
		if pageapireaderr != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		pagedata := make(map[string]interface{})

		pageapimarshalerr := json.Unmarshal(pageapibody, &pagedata)
		if pageapimarshalerr != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		// [TODO] : 문자열 이을때 개행문자 추가

		// [TODO] : sql 서버 쿼리 작업

	}

	c.Status(http.StatusOK)
}

/*
	if !(targetnamecheck && titlecheck && datecheck && humannamecheck && typecheck) {
			continue
		}

		if Crawljson.Target != targetdata["name"].(string) {
			continue
		} else {
			sqltargetdata = targetdata["name"].(string)
		}

		dataDate, _ := time.Parse("2006-01-02", datedata["start"].(string))

		if startTime.After(dataDate) || EndTime.Before(dataDate) {
			continue
		} else {
			sqldatedata = datedata["start"].(string)
		}

		if typedata["name"].(string) == "Vulnability" {

			if vulnerabilitycheck {
				sqlvulnerabilitydata = vulnerabilitydata
			}
		}

		if sqltargetdata != "" && sqldatedata != "" && sqlvulnerabilitydata != "" {
			continue
		}

		if len(titledata) == 0 {
			titledata = nil
		}

		if len(humandata) == 0 {
			humandata = nil
		}
*/

/*
	targetdata, targetnamecheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Target"].(map[string]interface{})["select"].(map[string]interface{})
	titledata, plaintextcheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Title"].(map[string]interface{})["title"].([]interface{})
	datedata, startcheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["FindDate"].(map[string]interface{})["date"].(map[string]interface{})
	humandata, humannamecheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Human"].(map[string]interface{})["people"].([]interface{})
	typedata, typecheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Type"].(map[string]interface{})["select"].(map[string]interface{})

	if !(targetnamecheck && plaintextcheck && startcheck && humannamecheck) {
		continue
	}

	if Crawljson.Target != targetdata["name"].(string) {
		continue
	}

	dataDate, _ := time.Parse("2006-01-02", datedata["start"].(string))

	if startTime.After(dataDate) || EndTime.Before(dataDate) {
		continue
	} else {
		resultdata["date"] = datedata["start"].(string)
	}

	if len(titledata) == 0 {
		resultdata["title"] = nil
	} else {
		resultdata["title"] = titledata[0].(map[string]interface{})["plain_text"].(string)
	}

	if len(humandata) == 0 {
		resultdata["human"] = nil
	} else {
		resultdata["human"] = humandata[0].(map[string]interface{})["name"].(string)
	}

	result = append(result, resultdata)
*/
