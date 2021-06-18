package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//SshKeyUpdate represents sshKey that should be updated
type SshKeyUpdate struct {
	Default bool   `json:"default"`
	Name    string `json:"name"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto SshKeyUpdate) ToBytes() (reader io.Reader) {
	requestByte, err := json.Marshal(dto)
	if err != nil {
		fmt.Println("SshKey update details entered in invalid format:", err)
		os.Exit(1)
	}
	return bytes.NewBuffer(requestByte)
}
