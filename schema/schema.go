package schema

import (
	"github.com/labstack/echo"
	gs "github.com/gorilla/schema"
	"go-crawler/utils"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
)

type CustomContext struct {
	echo.Context
	request_body []byte
	request_form url.Values
	method       string
}
type RequestData struct {
	Url         string
	Method_data string
	Cookie      string
	Proxy       string
}

func (c *CustomContext) DoRequire() error {
	var (
		err error
	)
	err = c.Request().ParseForm()
	if err != nil {
		return c.ResponseJson(err, nil)
	}
	c.request_form = c.Request().Form
	c.method = c.Request().Method
	request_body := c.Request().Body
	c.request_body, err = ioutil.ReadAll(request_body)
	if err != nil {
		return c.ResponseJson(err, nil)
	}
	defer request_body.Close()
	return nil
}
func (c *CustomContext) Method() string {
	return c.method
}
func (c *CustomContext) RequestBody() []byte {
	return c.request_body
}
func (c *CustomContext) RequestForm() url.Values {
	return c.request_form
}
func (c *CustomContext) ParseRequest(req ...interface{}) error {
	var (
		err error
	)
	for _, v := range req {
		if utils.IsJSON(string(c.request_body)) {
			if err = json.Unmarshal(c.request_body, v); err != nil {
				log.Println("parseRequest", err)
				return err
			}
		} else {
			decoder := gs.NewDecoder()
			if err = decoder.Decode(v, c.request_form); err != nil {
				log.Println("parseRequest", err)
				return err
			}
		}
	}
	return nil
}
func (c CustomContext) RenderJson(o interface{}) error {
	return c.JSON(http.StatusOK, o)
}
func (c CustomContext) ResponseJson(err error, data interface{}, messages ...string) error {
	app_err := make(map[string]interface{})
	app_err["error"] = err
	if len(messages) == 0 {
		if err != nil {
			app_err["message"] = err.Error()
		}
	} else {
		app_err["message"] = messages[0]
	}
	if data != nil {
		app_err["data"] = data
	}

	return c.JSON(http.StatusOK, app_err)
}
