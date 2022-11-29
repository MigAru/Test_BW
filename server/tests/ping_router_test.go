package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"srv/routers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type PingRouterTest struct {
	RequestMethod string //GET, POST, DELETE etc...
	StatusCode    int
	AssertMsg     string
	Response      interface{}
}

var PingRouterTests = []PingRouterTest{
	{
		RequestMethod: "GET",
		StatusCode:    200,
		AssertMsg:     "correct request GET",
		Response: map[string]interface{}{
			"message": "pong",
		},
	},
	{
		RequestMethod: "POST",
		StatusCode:    404,
		AssertMsg:     "incorrect request POST",
		Response:      "404 page not found",
	},
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	rGroup := r.Group("/api")
	routers.RegisterRouterPing(rGroup)
	return r
}

func TestRouterPingV1(t *testing.T) {
	assert := assert.New(t)
	r := SetUpRouter()

	for _, test := range PingRouterTests {
		var ResponseAssert interface{}
		switch v := test.Response.(type) {
		case string:
			ResponseAssert = []byte(v)

		default:
			byteJsonResp, err := json.Marshal(&v)
			if err != nil {
				t.Error(err)
            }
				ResponseAssert = byteJsonResp
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(test.RequestMethod, "/api/v1/ping", nil)
		r.ServeHTTP(w, req)
		assert.Equal(test.StatusCode, w.Code)
		assert.Equal(ResponseAssert, w.Body.Bytes())
	}
}

