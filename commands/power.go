package commands

import (
    "fmt"
    "strings"
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


func (p *Power) StartServer() ([]client.Response, error) {
   cmds := []client.Request {
             SetSystemTargetCmd,
             StartCmd,
   }
   res, err :=  p.Client.Commands(cmds)
   if err == nil {
       if res[0].Status != 0 {
           err = fmt.Errorf("Failed to set target to system1:", res[0].ErrorTag)
       } else if !(res[1].Status == 0 || (res[1].Status == 2 && strings.Contains(res[1].Details, "Server power already on"))) {
           err = fmt.Errorf("Failed to execute stop:", res[1].ErrorTag)
       }
   }
   return res, err
}

func (p *Power) StopServer() ([]client.Response, error) {
   cmds := []client.Request {
             SetSystemTargetCmd,
             StopCmd,
   }
   res, err :=  p.Client.Commands(cmds)
   if err == nil {
       if res[0].Status != 0 {
           err = fmt.Errorf("Failed to set target to system1:", res[0].ErrorTag)
       } else if !(res[1].Status == 0 || (res[1].Status == 2 && strings.Contains(res[1].Details, "Server power already off"))) {
           err = fmt.Errorf("Failed to execute stop:", res[1].ErrorTag)
       }
   }
   return res, err
}

func (p *Power) ResetServerHard() ([]client.Response, error) {
   cmds := []client.Request {
             SetSystemTargetCmd,
             ResetHardCmd,
   }
   res, err :=  p.Client.Commands(cmds)
   if err == nil {
       if res[0].Status != 0 {
           err = fmt.Errorf("Failed to set target to system1:", res[0].ErrorTag)
       } else if !(res[1].Status == 0 || (res[1].Status == 2 && strings.Contains(res[1].Details, "Server power off"))) {
           err = fmt.Errorf("Failed to execute reset:", res[1].ErrorTag)
       }
   }
   return res, err
}


func (p *Power) ResetServerSoft() ([]client.Response, error) {
   cmds := []client.Request {
             SetSystemTargetCmd,
             ResetSoftCmd,
   }
   res, err :=  p.Client.Commands(cmds)
   if err == nil {
       if res[0].Status != 0 {
           err = fmt.Errorf("Failed to set target to system1:", res[0].ErrorTag)
       } else if !(res[1].Status == 0 || (res[1].Status == 2 && strings.Contains(res[1].Details, "Server power off"))) {
           err = fmt.Errorf("Failed to execute reset:", res[1].ErrorTag)
       }
   }
   return res, err
}
