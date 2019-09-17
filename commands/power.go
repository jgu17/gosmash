package commands

import (
    "gosmash/client"
)

//SMASH CLP power commands
type Power struct {
    Client client.Client
}

const (
    SetSystemTargetCmd = "cd /system1"
    StartCmd = "start"
    StopCmd = "stop"
    ResetHardCmd = "reset hard"
    ResetSoftCmd = "reset soft"
)

func (p *Power) StartServer() (string, error) {
   cmds := []string {
             SetSystemTargetCmd,
             StartCmd,
   }
   return p.Client.Commands(cmds)
}

func (p *Power) StopServer() (string, error) {
   cmds := []string {
             SetSystemTargetCmd,
             StopCmd,
   }
   return p.Client.Commands(cmds)
}

func (p *Power) ResetServerHard() (string, error) {
   cmds := []string {
             SetSystemTargetCmd,
             ResetHardCmd,
   }
   return p.Client.Commands(cmds)
}


func (p *Power) ResetServerSoft() (string, error) {
   cmds := []string {
             SetSystemTargetCmd,
             ResetSoftCmd,
   }
   return p.Client.Commands(cmds)
}
