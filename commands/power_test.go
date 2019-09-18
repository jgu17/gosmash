package commands

import (
    "testing"
    "gosmash/client"
)


func TestStartServerWhenPowerOff(t *testing.T) {
    c := client.NewSimulator(client.NewEndpoint("localhost"))
    c.PowerOn = false
    power := commands.Power{c}
    res, err = power.StartServer()
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if c.PowerOn != true {
        t.Errorf("Server power is not on. \nReceived response: \n%s", res)
    }
    if !strings.Contains(res, "status=0") {
        t.Errorf("Unexpected non zero status code in output: %s", res)
    }
}

func TestStartServerWhenPowerOn(t *testing.T) {
    c := client.NewSimulator(client.NewEndpoint("localhost"))
    c.PowerOn = true
    power := commands.Power{c}
    res, err = power.StartServer()
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if c.PowerOn != true {
        t.Errorf("Server power is not on. \nReceived response: \n%s", res)
    }
    if strings.Contains(res, "status=0") {
        t.Errorf("Unexpected success status code in output: %s", res)
    }
}

func TestStopServerWhenPowerOn(t *testing.T) {
    c := client.NewSimulator(client.NewEndpoint("localhost"))
    c.PowerOn = true
    power := commands.Power{c}
    res, err = power.StopServer()
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if c.PowerOn != false {
        t.Errorf("Server power is not off. \nReceived response: \n%s", res)
    }
    if !strings.Contains(res, "status=0") {
        t.Errorf("Unexpected non zero status code in output: %s", res)
    }
}

func TestStopServerWhenPowerOff(t *testing.T) {
    c := client.NewSimulator(client.NewEndpoint("localhost"))
    c.PowerOn = false
    power := commands.Power{c}
    res, err = power.StopServer()
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if c.PowerOn != false {
        t.Errorf("Server power is not off. \nReceived response: \n%s", res)
    }
    if strings.Contains(res, "status=0") {
        t.Errorf("Unexpected success status code in output: %s", res)
    }
}

func TestResetServerHardWhenPowerOn(t *testing.T) {
    c := client.NewSimulator(client.NewEndpoint("localhost"))
    c.PowerOn = true
    power := commands.Power{c}
    res, err = power.ResetServerHard()
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if !strings.Contains(res, "status=0") {
        t.Errorf("Unexpected non zero status code in output: %s", res)
    }
    if c.PowerOn == false {
        t.Errorf("Server power is off. \nReceived response: \n%s", res)
    }
}

func TestResetServerHardWhenPowerOff(t *testing.T) {
    c := client.NewSimulator(client.NewEndpoint("localhost"))
    c.PowerOn = off
    power := commands.Power{c}
    res, err = power.ResetServerHard()
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if strings.Contains(res, "status=0") {
        t.Errorf("Unexpected success status code in output: %s", res)
    }
    if c.PowerOn != false {
        t.Errorf("Server power is unexpectedly on. \nReceived response: \n%s", res)
    }
}

func TestResetServerSoftWhenPowerOn(t *testing.T) {
    c := client.NewSimulator(client.NewEndpoint("localhost"))
    c.PowerOn = true
    power := commands.Power{c}
    res, err = power.ResetServerSoft()
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if !strings.Contains(res, "status=0") {
        t.Errorf("Unexpected non zero status code in output: %s", res)
    }
    if c.PowerOn == false {
        t.Errorf("Server power is off. \nReceived response: \n%s", res)
    }
}

func TestResetServerSoftWhenPowerOff(t *testing.T) {
    c := client.NewSimulator(client.NewEndpoint("localhost"))
    c.PowerOn = off
    power := commands.Power{c}
    res, err = power.ResetServerSoft()
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if strings.Contains(res, "status=0") {
        t.Errorf("Unexpected success status code in output: %s", res)
    }
    if c.PowerOn != false {
        t.Errorf("Server power is unexpectedly on. \nReceived response: \n%s", res)
    }
}
