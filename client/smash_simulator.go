package client

import (
	"fmt"
	"strings"
)

// SmashClient represents a SSH connection to the SMASH service
type smashSimulator struct {
	// SSH Endpoint
	Endpoint       endpoint
	PowerOn        bool
	ImageURL       string
	BootStateValue BootState
	MediaConnected bool
}

var CdSystem1 = Response{
	Status:    0,
	StatusTag: "COMMAND COMPLETED",
	Details: `Tue Sep 17 19:29:23 2019



/system1
`,
}

var StartWhenPowerOn = Response{
	Status:    2,
	StatusTag: "COMMAND PROCESSING FAILED",
	ErrorTag:  "COMMAND ERROR-UNSPECIFIED",
	Details: `Tue Sep 17 20:41:10 2019

Server power already on.
`,
}

var StartWhenPowerOff = Response{
	Status:    0,
	StatusTag: "COMMAND COMPLETED",
	Details: `Tue Sep 17 20:40:01 2019



Server powering on .......
`,
}

var StopWhenPowerOn = Response{
	Status:    0,
	StatusTag: "COMMAND COMPLETED",
	Details: `Tue Sep 17 19:39:12 2019



Server powering off .......
`,
}

var StopWhenPowerOff = Response{
	Status:    2,
	StatusTag: "COMMAND PROCESSING FAILED",
	ErrorTag:  "COMMAND ERROR-UNSPECIFIED",
	Details: `Tue Sep 17 19:44:46 2019

Server power already off.
`,
}

var ResetWhenPowerOff = Response{
	Status:    2,
	StatusTag: "COMMAND PROCESSING FAILED",
	ErrorTag:  "COMMAND ERROR-UNSPECIFIED",
	Details: `Tue Sep 17 20:58:33 2019

Server power off.
`,
}

var ResetWhenPowerOn = Response{
	Status:    0,
	StatusTag: "COMMAND COMPLETED",
	Details: `Tue Sep 17 21:01:45 2019

Resetting server.
`,
}

var CdFloppyDr1 = Response{
	Status:    0,
	StatusTag: "COMMAND COMPLETED",
	Details: `Tue Sep 17 22:24:44 2019



/map1/oemhp_vm1/floppydr1
`,
}

var CdCdDr1 = Response{
	Status:    0,
	StatusTag: "COMMAND COMPLETED",
	Details: `Tue Sep 17 22:27:10 2019



/map1/oemhp_vm1/cddr1
`,
}

var SetImageURL = Response{
	Status:    0,
	StatusTag: "COMMAND COMPLETED",
	Details: `Tue Sep 17 22:50:20 2019
`,
}

var SetBootConnect = Response{
	Status:    0,
	StatusTag: "COMMAND COMPLETED",
	Details: `Tue Sep 17 22:53:04 2019
`,
}

var SetBootConnectNoImage = Response{
	Status:    2,
	StatusTag: "COMMAND PROCESSING FAILED",
	ErrorTag:  "COMMAND ERROR-UNSPECIFIED",
	Details: `Tue Sep 17 23:21:06 2019

No image present in the Virtual Media drive.
`,
}

var SetBootOnce = Response{
	Status:    0,
	StatusTag: "COMMAND COMPLETED",
	Details: `Tue Sep 17 22:53:04 2019
`,
}

var SetBootDisconnect = Response{
	Status:    0,
	StatusTag: "COMMAND COMPLETED",
	Details: `Tue Sep 17 22:55:21 2019
`,
}

var SetBootDisconnectNoImage = Response{
	Status:    2,
	StatusTag: "COMMAND PROCESSING FAILED",
	ErrorTag:  "COMMAND ERROR-UNSPECIFIED",
	Details: `Tue Sep 17 22:55:28 2019

No image present in the Virtual Media drive.
`,
}

// creates a new client to Smash simulator.
func NewSimulator(e endpoint) *smashSimulator {
	return &smashSimulator{
		Endpoint:       e,
		PowerOn:        false,
		BootStateValue: NoBootState,
		MediaConnected: false,
	}
}

func copy(r *Response) *Response {
	c := *r
	return &c
}

// Return simulated command output
func (s *smashSimulator) Command(cmd Request) (output *Response, err error) {
	switch cmd.Command {
	case "start":
		if s.PowerOn == true {
			return copy(&StartWhenPowerOn), nil
		} else {
			s.PowerOn = true
			return copy(&StartWhenPowerOff), nil
		}
	case "stop":
		if s.PowerOn == true {
			s.PowerOn = false
			return copy(&StopWhenPowerOn), nil
		} else {
			return copy(&StopWhenPowerOff), nil
		}
	case "reset":
		if s.PowerOn == true {
			return copy(&ResetWhenPowerOn), nil
		} else {
			return copy(&ResetWhenPowerOff), nil
		}
	case "cd":
		if cmd.Args[0] == "/map1/oemhp_vm1/floppydr1" {
			return copy(&CdFloppyDr1), nil
		} else if cmd.Args[0] == "/map1/oemhp_vm1/cddr1" {
			return copy(&CdCdDr1), nil
		} else if cmd.Args[0] == "/system1" {
			return copy(&CdSystem1), nil
		} else {
			return new(Response), fmt.Errorf("Unsupported cd command target: %s", cmd.Args[0])
		}
	case "set":
		switch cmd.Args[0] {
		case "oemhp_boot=connect":
			if s.ImageURL == "" {
				return copy(&SetBootConnectNoImage), nil
			} else {
				s.MediaConnected = true
				s.BootStateValue = BootAlwaysState
				return copy(&SetBootConnect), nil
			}
		case "oemhp_boot=disconnect":
			if s.ImageURL == "" || s.MediaConnected == false {
				return copy(&SetBootDisconnectNoImage), nil
			} else {
				s.ImageURL = ""
				s.MediaConnected = false
				s.BootStateValue = NoBootState
				return copy(&SetBootDisconnect), nil
			}
		case "oemhp_boot=once":
			s.BootStateValue = BootOnceState
			return copy(&SetBootOnce), nil
		default:
			if strings.HasPrefix(cmd.Args[0], "oemhp_image") {
				s.ImageURL = strings.Split(cmd.Args[0], "=")[1]
				return copy(&SetImageURL), nil
			} else {
				return new(Response), fmt.Errorf("Unsupported set command argument: %s", cmd.Args[0])
			}
		}
	default:
		return new(Response), fmt.Errorf("Unrecognized command %s", cmd.Command)
	}
}

// Executes an ordered list of commands to SMASH service. Stop at the first
// execution error.
func (c *smashSimulator) Commands(cmds []Request) ([]Response, error) {
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
