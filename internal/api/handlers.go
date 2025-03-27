package api

import (
	"bytes"
	"net/http"

	"CapitalParser/internal/config"
)

const realApiHost = "https://api-capital.backend-capital.com/api/v1"      // host with real account
const testApiHost = "https://demo-api-capital.backend-capital.com/api/v1" // host with demo account

var conf, _ = config.LoadConfig()
var cst string
var xSecurityToken string

func postCapitalAuth(w http.ResponseWriter) (cst string, xSecurityToken string) {
	path := realApiHost + "/session"
	jsonData := []byte(`{"identifier":"` + conf.CAPITAL_USERNAME + `","password":"` + conf.CAPITAL_PASSWORD + `"}`)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CAP-API-KEY", conf.CAPITAL_API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cst = resp.Header.Get("cst")
	xSecurityToken = resp.Header.Get("x-security-token")

	defer resp.Body.Close()

	return cst, xSecurityToken
}

func getEpic(epic string) {

}
