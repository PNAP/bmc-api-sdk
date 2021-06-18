package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//SshKeyCreate represents sshKey that should be created
type SshKeyCreate struct {
	Default bool   `json:"default"`
	Name    string `json:"name"`
	Key     string `json:"key"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto SshKeyCreate) ToBytes() (reader io.Reader) {
	requestByte, err := json.Marshal(dto)
	if err != nil {
		fmt.Println("SshKey details entered in invalid format:", err)
		os.Exit(1)
	}
	return bytes.NewBuffer(requestByte)
}
