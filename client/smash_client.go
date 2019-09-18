package client

import (
        "io/ioutil"
        "strings"
        "golang.org/x/crypto/ssh"
)

// SmashClient represents a SSH connection to the SMASH service
type smashClient struct {
	// SSH Endpoint
	Endpoint endpoint
        SSHClient *ssh.Client
}


// creates a new client to a Smash service.
func NewClient(e endpoint, auth ssh.AuthMethod) (client *smashClient, err error) {
    client = &smashClient{Endpoint: e}
    config := &ssh.ClientConfig {
        User: e.User,
        Auth: []ssh.AuthMethod {
                auth,
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client.SSHClient, err = ssh.Dial("tcp", *e.HostString(), config)
    return
}


func KeyAuth(keyFile *string) ssh.AuthMethod {
    key, err := ioutil.ReadFile(*keyFile)
    if err != nil {
        panic(err)
    }

    signer, err := ssh.ParsePrivateKey(key)
    if err != nil {
        panic(err)
    }
    return ssh.PublicKeys(signer)
}


func PasswordAuth(pass *string) ssh.AuthMethod {
    return ssh.Password(*pass)
}


// Executes a command to SMASH service
func (c *smashClient) Command(cmd Request) (string, error) {
    s, err := c.SSHClient.NewSession()
    defer s.Close()

    if err != nil {
        return "", err
    }
    req := cmd.Command
    if (cmd.Args != nil && len(cmd.Args) > 0) {
        req = req + strings.Join(cmd.Args, " ")
    }
    output, e := s.CombinedOutput(req)
    return string(output), e
}

// Executes an ordered list of commands to SMASH service. Stop at the first
// execution error.
func (c *smashClient) Commands(cmds []Request) (string, error) {
    var output []string
    for _, cmd := range cmds {     
        s, err := c.Command(cmd)
        if s != "" {
            output = append(output, s)
        }
        if err != nil {
            return strings.Join(output, "\n*****\n"), err
        }
    }

    return strings.Join(output, "\n*****\n"), nil
}
