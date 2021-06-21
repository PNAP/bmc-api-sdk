package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//SshKey represents sshkey
type SshKey struct {
	ID            string `json:"id"`
	Default       bool   `json:"default"`
	Name          string `json:"name"`
	Key           string `json:"key"`
	Fingerprint   string `json:"fingerprint"`
	CreatedOn     string `json:"createdOn"`
	LastUpdatedOn string `jason:"lastUpdatedOn"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *SshKey) FromBytes(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	} else {
		fmt.Println("API response body can not be converted to SshKey dto:", err)
		os.Exit(1)
	}
}
