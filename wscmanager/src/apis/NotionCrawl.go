package wsc_apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func NotionCrawl(c *gin.Context) {
	AuthorizaionHeader := c.GetHeader("Authorizaion")

	if AuthorizaionHeader == "" || !strings.Contains(AuthorizaionHeader, "Bearer secret_") {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get Authorization header.",
		})
		return
	}

	var Crawljson NotionCrawljson

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

	result := make([]interface{}, 0)
	length := len(data["results"].([]interface{}))
	startTime, _ := time.Parse("2006-01-02", Crawljson.StartDate)
	EndTime, _ := time.Parse("2006-01-02", Crawljson.EndDate)

	for i := 0; i < length; i++ {
		targetdata, targetkeycheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Target"].(map[string]interface{})["select"].(map[string]interface{})["name"]
		titledata, titlekeycheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Title"].(map[string]interface{})["title"].([]interface{})[0].(map[string]interface{})["plain_text"]
		datedata, datekeycheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["FindDate"].(map[string]interface{})["date"].(map[string]interface{})["start"]
		humandata, humankeycheck := data["results"].([]interface{})[i].(map[string]interface{})["properties"].(map[string]interface{})["Human"].(map[string]interface{})["people"].([]interface{})[0].(map[string]interface{})["name"]

		if !(targetkeycheck && titlekeycheck && datekeycheck && humankeycheck) {
			continue
		}

		if Crawljson.Target != targetdata.(string) {
			continue
		}

		var resultdata = gin.H{
			"title": titledata.(string),
			"date":  datedata.(string),
			"human": humandata.(string),
		}

		dataDate, _ := time.Parse("2006-01-02", resultdata["date"].(string))

		if startTime.After(dataDate) || EndTime.Before(dataDate) {
			continue
		}

		result = append(result, resultdata)
	}

	c.JSON(http.StatusOK, result)
}
