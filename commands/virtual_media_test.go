package commands

import (
    "strings"
    "testing"
    "gosmash/client"
    "gosmash/commands"
)

var host = "localhost"
var ep = *client.NewEndpoint(&host)

func TestInsertUSBImage(t *testing.T) {
    c := client.NewSimulator(ep)
    c.ImageURL = ""
    c.MediaConnected = false
    vm := commands.VirtualMedia{c}
    res, err := vm.InsertUSBImage("http://xyz.com/foo.iso")
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if !strings.Contains(res, "status=0") {
        t.Errorf("Unexpected non zero status code in output: %s", res)
    }
    if c.MediaConnected != true {
        t.Errorf("Media is not connected. \nReceived response: \n%s", res)
    }
    if c.ImageURL != "http://xyz.com/foo.iso" {
        t.Errorf("Media is not set to the correct image url. \nReceived response: \n%s", res)
    }
    if (c.BootStateValue != client.BootAlwaysState) {
        t.Errorf("oemhp_boot is not set to always as expected: %s", *c)
    }
}

func TestInsertUSBImageSingleBoot(t *testing.T) {
    c := client.NewSimulator(ep)
    c.ImageURL = ""
    c.MediaConnected = false
    vm := commands.VirtualMedia{c}
    res, err := vm.InsertUSBImageSingleBoot("http://xyz.com/foo.iso")
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if !strings.Contains(res, "status=0") {
        t.Errorf("Unexpected non zero status code in output: %s", res)
    }
    if c.MediaConnected != true {
        t.Errorf("Media is not connected. \nReceived response: \n%s", res)
    }
    if c.ImageURL != "http://xyz.com/foo.iso" {
        t.Errorf("Media is not set to the correct image url. \nReceived response: \n%s", res)
    }
    if (c.BootStateValue != client.BootOnceState) {
        t.Errorf("oemhp_boot is not set to once as expected: %s", *c)
    }
}

func TestInsertCDRomImage(t *testing.T) {
    c := client.NewSimulator(ep)
    c.ImageURL = ""
    c.MediaConnected = false
    vm := commands.VirtualMedia{c}
    res, err := vm.InsertCDRomImage("http://xyz.com/foo.iso")
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if !strings.Contains(res, "status=0") {
        t.Errorf("Unexpected non zero status code in output: %s", res)
    }
    if c.MediaConnected != true {
        t.Errorf("Media is not connected. \nReceived response: \n%s", res)
    }
    if c.ImageURL != "http://xyz.com/foo.iso" {
        t.Errorf("Media is not set to the correct image url. \nReceived response: \n%s", res)
    }
    if (c.BootStateValue != client.BootAlwaysState) {
        t.Errorf("oemhp_boot is not set to always as expected: %s", *c)
    }
}

func TestInsertCDRomImageSingleBoot(t *testing.T) {
    c := client.NewSimulator(ep)
    c.ImageURL = ""
    c.MediaConnected = false
    vm := commands.VirtualMedia{c}
    res, err := vm.InsertCDRomImageSingleBoot("http://xyz.com/foo.iso")
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if !strings.Contains(res, "status=0") {
        t.Errorf("Unexpected non zero status code in output: %s", res)
    }
    if c.MediaConnected != true {
        t.Errorf("Media is not connected. \nReceived response: \n%s", res)
    }
    if c.ImageURL != "http://xyz.com/foo.iso" {
        t.Errorf("Media is not set to the correct image url. \nReceived response: \n%s", res)
    }
    if (c.BootStateValue != client.BootOnceState) {
        t.Errorf("oemhp_boot is not set to once as expected: %s", *c)
    }
}

func TestEjectCDRomImage(t *testing.T) {
    c := client.NewSimulator(ep)
    c.ImageURL = "http://xyz.com/foo.iso"
    c.MediaConnected = false
    vm := commands.VirtualMedia{c}
    res, err := vm.EjectCDRomImage()
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if !strings.Contains(res, "status=0") {
        t.Errorf("Unexpected non zero status code in output: %s", res)
    }
    if c.MediaConnected != false {
        t.Errorf("Media is still connected: \n%s", *c)
    }
    if c.ImageURL != "" {
        t.Errorf("Media is not ejected: %s", *c)
    }
}

func TestEjectUSBImage(t *testing.T) {
    c := client.NewSimulator(ep)
    c.ImageURL = "http://xyz.com/foo.iso"
    c.MediaConnected = false
    vm := commands.VirtualMedia{c}
    res, err := vm.EjectUSBImage()
    if err != nil {
        t.Errorf("Command execution error: %s", err)
    }
    if !strings.Contains(res, "status=0") {
        t.Errorf("Unexpected non zero status code in output: %s", res)
    }
    if c.MediaConnected != false {
        t.Errorf("Media is still connected: \n%s", *c)
    }
    if c.ImageURL != "" {
        t.Errorf("Media is not ejected: %s", *c)
    }
}
