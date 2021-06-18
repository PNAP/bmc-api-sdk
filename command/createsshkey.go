package command

import (
	"net/http"

	"github.com/PNAP/bmc-api-sdk/client"
	"github.com/PNAP/bmc-api-sdk/dto"
)

// CreateSshKeyCommand represents command for sshkey creating. Use NewCreateSshKeyCommand to initilize command properly.
type CreateSshKeyCommand struct {
	requester client.Requester
	sshKey    dto.SshKeyCreate
}

// Execute create new ssh key
func (command *CreateSshKeyCommand) Execute() (*http.Response, error) {
	var req = command.requester
	return req.Post("ssh-keys", command.sshKey.ToBytes())
}

// SetRequester - sets requester to the command
func (command *CreateSshKeyCommand) SetRequester(requester client.Requester) {
	command.requester = requester
}

// SetSshKey - sets sshKey details to the command
func (command *CreateSshKeyCommand) SetSshKey(sshKey dto.SshKeyCreate) {
	command.sshKey = sshKey
}

//NewCreateSshKeyCommand - constructs new command of this type
func NewCreateSshKeyCommand(requester client.Requester, sshKey dto.SshKeyCreate) *CreateSshKeyCommand {
	return &CreateSshKeyCommand{requester, sshKey}
}
