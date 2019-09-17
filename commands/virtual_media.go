package commands

import (
    "gosmash/client"
)

// Smash CLP virtual media commands
type VirtualMedia struct {
    Client client.Client
}

const (
    SetUSBTargetCmd = "cd /map1/oemhp_vm1/floppydr1"
    SetCDROMTargetCmd = "cd /map1/oemhp_vm1/cddr1"
    SetOEMHPImageCmd = "set oemhp_image="
    SetOEMHPConnectCmd = "set oemhp_boot=connect"
    SetOEMHPDisconnectCmd = "set oemhp_boot=disconnect"
    SetOEMHPBootCmd = "set oemhp_boot="
    BootOnceOption = "once"
    BootAlwaysOption = "always"
)

func (v *VirtualMedia) InsertUSBImage(url string) (string, error) {
   cmds := []string {
             SetUSBTargetCmd,
             SetOEMHPImageCmd + url,
             SetOEMHPConnectCmd,
   }
   return v.Client.Commands(cmds)
}


func (v *VirtualMedia) InsertUSBImageSingleBoot(url string) (string, error) {
   cmds := []string {
             SetUSBTargetCmd,
             SetOEMHPImageCmd + url,
             SetOEMHPConnectCmd,
             SetOEMHPBootCmd + BootOnceOption,
   }
   return v.Client.Commands(cmds)
}

func (v *VirtualMedia) EjectUSBImage() (string, error) {
   cmds := []string {
             SetUSBTargetCmd,
             SetOEMHPDisconnectCmd,
   }
   return v.Client.Commands(cmds)
}

func (v *VirtualMedia) InsertCDRomImage(url string) (string, error) {
   cmds := []string {
             SetCDROMTargetCmd,
             SetOEMHPImageCmd + url,
             SetOEMHPConnectCmd,
   }
   return v.Client.Commands(cmds)
}


func (v *VirtualMedia) InsertCDRomImageSingleBoot(url string) (string, error) {
   cmds := []string {
             SetCDROMTargetCmd,
             SetOEMHPImageCmd + url,
             SetOEMHPConnectCmd,
             SetOEMHPBootCmd + BootOnceOption,
   }
   return v.Client.Commands(cmds)
}

func (v *VirtualMedia) EjectCDRomImage() (string, error) {
   cmds := []string {
             SetCDROMTargetCmd,
             SetOEMHPDisconnectCmd,
   }
   return v.Client.Commands(cmds)
}
