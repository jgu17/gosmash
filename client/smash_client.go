package client

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"strings"
)

// SmashClient represents a SSH connection to the SMASH service
type smashClient struct {
	// SSH Endpoint
	Endpoint  endpoint
	sshClient *ssh.Client
}

// creates a new client to a Smash service.
func NewClient(e endpoint) *smashClient {
	return &smashClient{Endpoint: e}
}

func KeyAuth(keyFile string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		panic(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}
	return ssh.PublicKeys(signer)
}

func PasswordAuth(pass string) ssh.AuthMethod {
	return ssh.Password(pass)
}

func (c *smashClient) Connect(auth ssh.AuthMethod) error {
	config := &ssh.ClientConfig{
		User: c.Endpoint.User,
		Auth: []ssh.AuthMethod{
			auth,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	var err error
	c.sshClient, err = ssh.Dial("tcp", *(c.Endpoint).HostString(), config)
	return err
}

// Executes a command to SMASH service
func (c *smashClient) Command(cmd Request) (*Response, error) {
	s, err := c.sshClient.NewSession()
	defer s.Close()

	if err != nil {
		return new(Response), err
	}
	req := cmd.Command
	if cmd.Args != nil && len(cmd.Args) > 0 {
		req = req + " " + strings.Join(cmd.Args, " ")
	}
	output, e := s.CombinedOutput(req)
	return NewResponse(string(output)), e
}

// Executes an ordered list of commands to SMASH service. Stop at the first
// execution error.
func (c *smashClient) Commands(cmds []Request) ([]Response, error) {
	var resList []Response
	for _, cmd := range cmds {
		r, err := c.Command(cmd)
		if err != nil {
			return resList, err
		}
		resList = append(resList, *r)
		if r.Status != 0 {
			return resList, nil
		}
	}

	return resList, nil
}
