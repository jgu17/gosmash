package commands

import (
    "gosmash/client"
)

//SMASH CLP power commands
type Power struct {
    Client client.Client
}


var SetSystemTargetCmd = client.Request{Command: "cd", Args: []string {"/system1"}}
var StartCmd = client.Request{Command: "start"}
var StopCmd = client.Request{Command: "stop"}
var ResetHardCmd = client.Request{Command: "reset", Args: []string {"hard"}}
var ResetSoftCmd = client.Request{Command: "reset", Args: []string {"soft"}}


func (p *Power) StartServer() (string, error) {
   cmds := []client.Request {
             SetSystemTargetCmd,
             StartCmd,
   }
   return p.Client.Commands(cmds)
}

func (p *Power) StopServer() (string, error) {
   cmds := []client.Request {
             SetSystemTargetCmd,
             StopCmd,
   }
   return p.Client.Commands(cmds)
}

func (p *Power) ResetServerHard() (string, error) {
   cmds := []client.Request {
             SetSystemTargetCmd,
             ResetHardCmd,
   }
   return p.Client.Commands(cmds)
}


func (p *Power) ResetServerSoft() (string, error) {
   cmds := []client.Request {
             SetSystemTargetCmd,
             ResetSoftCmd,
   }
   return p.Client.Commands(cmds)
}
