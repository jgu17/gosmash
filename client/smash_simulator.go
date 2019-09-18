package client

import (
    "fmt"
    "strings"
)

// SmashClient represents a SSH connection to the SMASH service
type smashSimulator struct {
    // SSH Endpoint
    Endpoint endpoint
    PowerOn bool
    ImageURL string
    BootStateValue BootState
    MediaConnected bool
}

const ( 
    CdSystem1 = `status=0
status_tag=COMMAND COMPLETED
Tue Sep 17 19:29:23 2019



/system1
`

    StartWhenPowerOn = `status=2
status_tag=COMMAND PROCESSING FAILED
error_tag=COMMAND ERROR-UNSPECIFIED
Tue Sep 17 20:41:10 2019

Server power already on.
`

    StartWhenPowerOff = `status=0
status_tag=COMMAND COMPLETED
Tue Sep 17 20:40:01 2019



Server powering on .......
`

    StopWhenPowerOn = `status=0
status_tag=COMMAND COMPLETED
Tue Sep 17 19:39:12 2019



Server powering off .......
`

    StopWhenPowerOff = `status=2
status_tag=COMMAND PROCESSING FAILED
error_tag=COMMAND ERROR-UNSPECIFIED
Tue Sep 17 19:44:46 2019

Server power already off.
`

    ResetWhenPowerOff = `status=2
status_tag=COMMAND PROCESSING FAILED
error_tag=COMMAND ERROR-UNSPECIFIED
Tue Sep 17 20:58:33 2019

Server power off.
`

    ResetWhenPowerOn = `status=0
status_tag=COMMAND COMPLETED
Tue Sep 17 21:01:45 2019

Resetting server.
`

    CdFloppyDr1 = `status=0
status_tag=COMMAND COMPLETED
Tue Sep 17 22:24:44 2019



/map1/oemhp_vm1/floppydr1
`

   CdCdDr1 = `status=0
status_tag=COMMAND COMPLETED
Tue Sep 17 22:27:10 2019



/map1/oemhp_vm1/cddr1
`

    SetImageURL = `status=0
status_tag=COMMAND COMPLETED
Tue Sep 17 22:50:20 2019
`

    SetBootConnect = `status=0
status_tag=COMMAND COMPLETED
Tue Sep 17 22:53:04 2019
`

    SetBootConnectNoImage = `status=2
status_tag=COMMAND PROCESSING FAILED
error_tag=COMMAND ERROR-UNSPECIFIED
Tue Sep 17 23:21:06 2019

No image present in the Virtual Media drive.
`

    SetBootOnce = `status=0
status_tag=COMMAND COMPLETED
Tue Sep 17 22:53:04 2019
`

    SetBootDisconnect = `status=0
status_tag=COMMAND COMPLETED
Tue Sep 17 22:55:21 2019
`

    SetBootDisconnectNoImage = `status=2
status_tag=COMMAND PROCESSING FAILED
error_tag=COMMAND ERROR-UNSPECIFIED
Tue Sep 17 22:55:28 2019

No image present in the Virtual Media drive.
`
)

// creates a new client to Smash simulator.
func NewSimulator(e endpoint) *smashSimulator {
    se := &smashSimulator{
           Endpoint: e,
           PowerOn: false,
           BootStateValue: NoBootState,
           MediaConnected: false,
    }
    return se
}

// Return simulated command output
func (s *smashSimulator) Command(cmd Request) (output string, err error) {
    switch cmd.Command {
    case "cd /system1":
        output = CdSystem1
    case "start":
        if s.PowerOn == true {
            output = StartWhenPowerOn
        } else {
            output = StartWhenPowerOff
            s.PowerOn = true
        }
    case "stop":
        if s.PowerOn == true {
            output = StopWhenPowerOn
            s.PowerOn = false
        } else {
            output = StopWhenPowerOff
        }
    case "reset":
        if s.PowerOn == true {
            output = ResetWhenPowerOn
        } else {
            output = ResetWhenPowerOff
        }
    case "cd":
        if cmd.Args[0] ==  "/map1/oemhp_vm1/floppydr1" {
            output = CdFloppyDr1
        } else if cmd.Args[0] == "/map1/oemhp_vm1/cddr1" {
            output = CdCdDr1
        } else {       
            return "", fmt.Errorf("Unsupported cd command target: %s", cmd.Args[0])
        }
    case "set":
        switch cmd.Args[0] {
        case "oemhp_boot=connect":
            if s.ImageURL == "" {
                output = SetBootConnectNoImage
            } else {
                output = SetBootConnect
                s.MediaConnected = true
                s.BootStateValue = BootAlwaysState
            }
        case "oemhp_boot=disconnect":
            if (s.ImageURL == "" || s.MediaConnected == false) {
                output = SetBootDisconnectNoImage
            } else {
                output = SetBootDisconnect
                s.ImageURL = ""
                s.MediaConnected = false
                s.BootStateValue = NoBootState
            }
        case "oemhp_boot=once":
            output = SetBootOnce
            s.BootStateValue = BootOnceState
        default:
            if strings.HasPrefix(cmd.Args[0], "oemhp_image") {
                output = SetImageURL
                s.ImageURL = strings.Split(cmd.Args[0], "=")[1]
            } else {
                return "", fmt.Errorf("Unsupported set command argument: %s", cmd.Args[0])
            }
        }
    default:
       return "", fmt.Errorf("Unrecognized command %s", cmd.Command)
    }
    return
}

// Executes an ordered list of commands to SMASH service. Stop at the first
// execution error.
func (c *smashSimulator) Commands(cmds []Request) (string, error) {
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
