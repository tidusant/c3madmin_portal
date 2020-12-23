package main

import (
	"bytes"
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/c3m-common/mycrypto"

	"github.com/tidusant/chadmin-repo/models"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var testsession string

func decodeResponse(requeststring string, data string) (rs models.RequestResult, err error) {
	//encode data
	requeststring = mycrypto.EncDat2(requeststring)
	data = "data=" + mycrypto.EncDat2(data)

	//add body into request
	body := bytes.NewReader([]byte(data))
	req, err := http.NewRequest(http.MethodPost, "/"+requeststring, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {

		return
	}
	// Create a response recorder so you can inspect the response test
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		err = errors.New(fmt.Sprintf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code))
		return

	}

	//check data
	//get response body
	bodyresp, err := ioutil.ReadAll(w.Body)
	rtstr := string(bodyresp)
	//decode data

	rtstr = mycrypto.DecodeOld(rtstr, 8)
	json.Unmarshal([]byte(rtstr), &rs)
	return
}

func doCall(testname, requesturl, queryData string, t *testing.T) models.RequestResult {
	fmt.Println("\n\n==== " + testname + " ====")
	fmt.Printf("Data: url: %s - data:%s\n", requesturl, queryData)
	rs, err := decodeResponse(requesturl, queryData)
	if err != nil {
		t.Fatalf("Test fail: request error: %s", err.Error())
	}
	fmt.Printf("Request return: %+v\n", rs)
	return rs
}

func setup() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and
	// register your routes
	r = gin.Default()
	r.POST("/*name", postHandler)
}
func TestMain(m *testing.M) {
	setup()
	exitVal := m.Run()
	os.Exit(exitVal)
}

//test special char
func TestSpecialChar(t *testing.T) {
	rs := doCall("TestSpecialChar", c3mcommon.GetSpecialChar(), "", t)
	//check test data
	if rs.Status == 1 {
		t.Fatalf("Test fail")
	}
}
