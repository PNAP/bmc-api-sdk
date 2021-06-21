package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
	"github.com/PNAP/bmc-api-sdk/dto"
)

// UpdateSshKeyCommand represents command for sshkey updating. Use NewUpdateSshKeyCommand to initilize command properly.
type UpdateSshKeyCommand struct {
	requester client.Requester
	sshKey    dto.SshKeyUpdate
	sshKeyID  string
}

// Execute create new ssh key
func (command *UpdateSshKeyCommand) Execute() (*http.Response, error) {
	var req = command.requester
	return req.Post("ssh-keys/"+command.sshKeyID, command.sshKey.ToBytes()) //implement PUT
}

// SetRequester - sets requester to the command
func (command *UpdateSshKeyCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetSshKey - sets sshKey details to the command
func (command *UpdateSshKeyCommand) SetSshKey(sshKey dto.SshKeyUpdate) {
	command.sshKey = sshKey
}

//NewUpdateSshKeyCommand - constructs new command of this type
func NewUpdateSshKeyCommand(requester client.Requester, sshKey dto.SshKeyUpdate, sshKeyID string) *UpdateSshKeyCommand {
	return &UpdateSshKeyCommand{requester, sshKey, sshKeyID}
}
