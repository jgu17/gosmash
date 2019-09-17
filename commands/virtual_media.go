package commands

// Smash CLP virtual media commands
type VirtualMedia struct {
    Client *gosmash.Client
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

func (c *smashClient) InsertUSBImage(url string) (string, error) {
   cmds := []string {
             SetUSBTargetCmd,
             SetOEMHPImageCmd + url,
             SetOEMHPConnectCmd,
   }
   return c.Commands(cmds)
}


func (c *smashClient) InsertUSBImageSingleBoot(url string) (string, error) {
   cmds := []string {
             SetUSBTargetCmd,
             SetOEMHPImageCmd + url,
             SetOEMHPConnectCmd,
             SetOEMHPBootCmd + BootOnceOption,
   }
   return c.Commands(cmds)
}

func (c *smashClient) EjectUSBImage() (string, error) {
   cmds := []string {
             SetUSBTargetCmd,
             SetOEMHPDisconnectCmd,
   }
   return c.Commands(cmds)
}

func (c *smashClient) InsertCDRomImage(url string) (string, error) {
   cmds := []string {
             SetCDROMTargetCmd,
             SetOEMHPImageCmd + url,
             SetOEMHPConnectCmd,
   }
   return c.Commands(cmds)
}


func (c *smashClient) InsertCDRomImageSingleBoot(url string) (string, error) {
   cmds := []string {
             SetCDROMTargetCmd,
             SetOEMHPImageCmd + url,
             SetOEMHPConnectCmd,
             SetOEMHPBootCmd + BootOnceOption,
   }
   return c.Commands(cmds)
}

func (c *smashClient) EjectCDRomImage() (string, error) {
   cmds := []string {
             SetCDROMTargetCmd,
             SetOEMHPDisconnectCmd,
   }
   return c.Commands(cmds)
}
