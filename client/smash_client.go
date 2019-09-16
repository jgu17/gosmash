package client

import (
        "fmt"
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


var reset_hard_cmds = []string {
                         "cd /system1",
                         "reset hard",
}


var reset_soft_cmds = []string {
                         "cd /system1",
                         "reset soft",
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
func (c *smashClient) Command(cmd string) (string, error) {
    s, err := c.SSHClient.NewSession()
    defer s.Close()

    if err != nil {
        return "", err
    }

    output, e := s.CombinedOutput(cmd)
    return string(output), e
}

// Executes an ordered list of commands to SMASH service
func (c *smashClient) Commands(cmds []string) (string, error) {
    var output []string
    for _, cmd := range cmds {
        fmt.Println("Executing cmd:", cmd)
        s, err := c.SSHClient.NewSession()

        defer s.Close()

        if err != nil {
            return strings.Join(output, "\n*****\n"), err
        }

        o, e := s.CombinedOutput(cmd)
        if len(o) > 0 {
            output = append(output, string(o))
        }

        if e != nil {
            return strings.Join(output, "\n*****\n"), err
        }
        s.Close()
    }

    return strings.Join(output, "\n*****\n"), nil
}

func (c *smashClient) InsertUSBImage(url string) (string, error) {
   cmds := []string {
             "cd /map1/oemhp_vm1/floppydr1",
             "set oemhp_image=" + url,
             "set oemhp_boot=connect",
   }
   return c.Commands(cmds)
}


func (c *smashClient) InsertUSBImageSingleBoot(url string) (string, error) {
   cmds := []string {
             "cd /map1/oemhp_vm1/floppydr1",
             "set oemhp_image=" + url,
             "set oemhp_boot=connect",
             "set oemhp_boot=once",
   }
   return c.Commands(cmds)
}

func (c *smashClient) EjectUSBImage() (string, error) {
   cmds := []string {
             "cd /map1/oemhp_vm1/floppydr1",
             "set oemhp_boot=disconnect",
   }
   return c.Commands(cmds)
}

func (c *smashClient) InsertCDRomImage(url string) (string, error) {
   cmds := []string {
             "cd /map1/oemhp_vm1/cddr1",
             "set oemhp_image=" + url,
             "set oemhp_boot=connect",
   }
   return c.Commands(cmds)
}


func (c *smashClient) InsertCDRomImageSingleBoot(url string) (string, error) {
   cmds := []string {
             "cd /map1/oemhp_vm1/cddr1",
             "set oemhp_image=" + url,
             "set oemhp_boot=connect",
             "set oemhp_boot=once",
   }
   return c.Commands(cmds)
}


func (c *smashClient) EjectCDRomImage() (string, error) {
   cmds := []string {
             "cd /map1/oemhp_vm1/cddr1",
             "set oemhp_boot=disconnect",
   }
   return c.Commands(cmds)
}


func (c *smashClient) StartServer() (string, error) {
   cmds := []string {
             "cd /system1",
             "start",
   }
   return c.Commands(cmds)
}

func (c *smashClient) StopServer() (string, error) {
   cmds := []string {
             "cd /system1",
             "stop",
   }
   return c.Commands(cmds)
}

func (c *smashClient) ResetServerHard() (string, error) {
   cmds := []string {
             "cd /system1",
             "reset hard",
   }
   return c.Commands(cmds)
}


func (c *smashClient) ResetServerSoft() (string, error) {
   return c.Commands(reset_soft_cmds)
}

