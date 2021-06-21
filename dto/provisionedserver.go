package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//ProvisionedServer represents server that should be provisioned (created) for the first time
type ProvisionedServer struct {
	Name                  string          `json:"hostname"`
	Description           string          `json:"description"`
	Os                    string          `json:"os"`
	Type                  string          `json:"type"`
	Location              string          `json:"location"`
	Status                string          `json:"status"`
	InstallDefaultSshKeys bool            `json:"installDefaultSshKeys"`
	SshKeys               []string        `json:"sshKeys"`
	SshKeyIds             []string        `json:"sshKeyIds"`
	ReservationId         string          `json:"reservationId"`
	PricingModel          string          `json:"pricingModel"`
	NetworkType           string          `json:"networkType"`
	OsConfiguration       OsConfiguration `json:"osConfiguration"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto ProvisionedServer) ToBytes() (reader io.Reader) {
	requestByte, err := json.Marshal(dto)
	if err != nil {
		fmt.Println("Server details entered in invalid format:", err)
		os.Exit(1)
	}
	return bytes.NewBuffer(requestByte)
}
