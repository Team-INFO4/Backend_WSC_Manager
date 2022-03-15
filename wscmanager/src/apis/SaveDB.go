package wsc_apis

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	payload := strings.NewReader("{\"page_size\":100,\"filter\":{\"or\":[{\"property\":\"FindDate\",\"date\":{\"equals\":\"" + Crawljson.StartDate + "\"}},{\"property\":\"FindDate\",\"date\":{\"equals\":\"" + Crawljson.EndDate + "\"}},{\"and\":[{\"property\":\"FindDate\",\"date\":{\"after\":\"" + Crawljson.StartDate + "\"}},{\"property\":\"FindDate\",\"date\":{\"before\":\"" + Crawljson.EndDate + "\"}}]}]}}")

	req, reqerr := http.NewRequest("POST", url, payload)
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
	if length == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get pages.",
		})
		return
	}
	enverr := godotenv.Load(".env")

	if enverr != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	dbip := os.Getenv("DBIP")
	dbport := os.Getenv("DBPORT")
	dbid := os.Getenv("DBID")
	dbpassword := os.Getenv("DBPASSWORD")
	dbschema := os.Getenv("DBSCHEMA")

	db, dberr := sql.Open("mysql", dbid+":"+dbpassword+"@tcp("+dbip+":"+dbport+")/"+dbschema)
	if dberr != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer db.Close()

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

		sqldatedata = datedata["start"].(string)

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
			c.Status(http.StatusInternalServerError)
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

		pagemap := make(map[string]interface{})

		pageapimarshalerr := json.Unmarshal(pageapibody, &pagemap)
		if pageapimarshalerr != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		pagelength := len(pagemap["results"].([]interface{}))
		var referencebool bool = false
		for j := 0; j < pagelength; j++ {
			blocktype, blocktypecheck := pagemap["results"].([]interface{})[j].(map[string]interface{})["type"]
			if !blocktypecheck {
				c.Status(http.StatusInternalServerError)
				return
			}

			pagedata, pagedatacheck := pagemap["results"].([]interface{})[j].(map[string]interface{})[blocktype.(string)].(map[string]interface{})["rich_text"].([]interface{})
			if !pagedatacheck {
				c.Status(http.StatusInternalServerError)
				return
			}

			if len(pagedata) == 0 {
				continue
			}

			if pagedata[0].(map[string]interface{})["plain_text"].(string) == "참고 자료" {
				referencebool = true
				continue
			}

			if !referencebool {
				sqldescription += pagedata[0].(map[string]interface{})["plain_text"].(string) + "\n"
			} else {
				sqlreference += pagedata[0].(map[string]interface{})["plain_text"].(string) + "\n"
			}
		}

		var execerr error
		if string(sqltypedata[0]) != "V" {
			_, execerr = db.Exec("INSERT INTO report (id, title, type, vulnerability, date, target, human, description, reference) values (0, ?, ?, ?, ?, ?, ?, ?, ?)", sqltitledata, string(sqltypedata[0]), nil, sqldatedata, sqltargetdata, sqlhumandata, sqldescription, sqlreference)
		} else {
			_, execerr = db.Exec("INSERT INTO report (id, title, type, vulnerability, date, target, human, description, reference) values (0, ?, ?, ?, ?, ?, ?, ?, ?)", sqltitledata, string(sqltypedata[0]), sqlvulnerabilitydata, sqldatedata, sqltargetdata, sqlhumandata, sqldescription, sqlreference)
		}

		if execerr != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.Status(http.StatusOK)
}
