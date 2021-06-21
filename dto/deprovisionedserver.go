package dto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//DeprovisionedServer represents server delete response type
type DeprovisionedServer struct {
	Result   string `json:"result"`
	ServerID string `json:"serverId"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *DeprovisionedServer) FromBytes(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	} else {
		fmt.Println("API response body can not be converted to DeprovisionedServer dto:", err)
		os.Exit(1)
	}
}
