package commands

import (
	"fmt"

	"github.com/jgu17/gosmash/client"
)

// Smash CLP virtual media commands
type VirtualMedia struct {
	Client client.Client
}

var SetUSBTargetCmd = client.Request{
	Command: "cd",
	Args:    []string{"/map1/oemhp_vm1/floppydr1"},
}
var SetCDROMTargetCmd = client.Request{
	Command: "cd",
	Args:    []string{"/map1/oemhp_vm1/cddr1"},
}
var SetOEMHPImageCmd = client.Request{
	Command: "set",
	Args:    []string{"oemhp_image="},
}
var SetOEMHPConnectCmd = client.Request{
	Command: "set",
	Args:    []string{"oemhp_boot=connect"},
}
var SetOEMHPDisconnectCmd = client.Request{
	Command: "set",
	Args:    []string{"oemhp_boot=disconnect"},
}
var SetOEMHPBootCmd = client.Request{
	Command: "set",
	Args:    []string{"oemhp_boot="},
}

func (v *VirtualMedia) InsertUSBImage(url string) ([]client.Response, error) {
	cmds := []client.Request{
		SetUSBTargetCmd,
		client.Request{
			Command: "set",
			Args:    []string{"oemhp_image=" + url},
		},
		SetOEMHPConnectCmd,
	}
	return v.Client.Commands(cmds)
}

func (v *VirtualMedia) InsertUSBImageSingleBoot(url string) ([]client.Response, error) {
	cmds := []client.Request{
		SetUSBTargetCmd,
		client.Request{
			Command: "set",
			Args:    []string{"oemhp_image=" + url},
		},
		SetOEMHPConnectCmd,
		client.Request{
			Command: "set",
			Args:    []string{fmt.Sprintf("oemhp_boot=%s", client.BootOnce)},
		},
	}
	return v.Client.Commands(cmds)
}

func (v *VirtualMedia) EjectUSBImage() ([]client.Response, error) {
	cmds := []client.Request{
		SetUSBTargetCmd,
		SetOEMHPDisconnectCmd,
	}
	return v.Client.Commands(cmds)
}

func (v *VirtualMedia) InsertCDRomImage(url string) ([]client.Response, error) {
	cmds := []client.Request{
		SetCDROMTargetCmd,
		client.Request{
			Command: "set",
			Args:    []string{"oemhp_image=" + url},
		},
		SetOEMHPConnectCmd,
	}
	return v.Client.Commands(cmds)
}

func (v *VirtualMedia) InsertCDRomImageSingleBoot(url string) ([]client.Response, error) {
	cmds := []client.Request{
		SetCDROMTargetCmd,
		client.Request{
			Command: "set",
			Args:    []string{"oemhp_image=" + url},
		},
		SetOEMHPConnectCmd,
		client.Request{
			Command: "set",
			Args:    []string{fmt.Sprintf("oemhp_boot=%s", client.BootOnce)},
		},
	}
	return v.Client.Commands(cmds)
}

func (v *VirtualMedia) EjectCDRomImage() ([]client.Response, error) {
	cmds := []client.Request{
		SetCDROMTargetCmd,
		SetOEMHPDisconnectCmd,
	}
	return v.Client.Commands(cmds)
}
