package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"github.com/fronbasal/vertretungsplan/structs"
	"fmt"
)

func Request(url string) (*http.Response, error) {
	creds := LoadCredentials()
	req, err := http.NewRequest("GET", creds.Host+url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(creds.Username, creds.Password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func LoadCredentials() structs.Credentials {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatal("Failed to read config file!")
	}
	var c structs.Credentials
	json.Unmarshal(b, &c)
	return c
}

func IServLogin(username, password string) (error, bool) {
	s := fmt.Sprintf("login_act=%s&login_pwd=%s", username, password)
	body := strings.NewReader(s)
	req, err := http.NewRequest("POST", "https://steinbart-gym.eu/idesk/index.php", body)
	if err != nil {
		fmt.Println(err)
		return err, false
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, false
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	return nil, resp.StatusCode == 302
}
