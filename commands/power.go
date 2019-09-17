package commands

//SMASH CLP power commands
type Power struct {
    Client *gosmash.Client
}

const (
    SetSystemTargetCmd = "cd /system1"
    StartCmd = "start"
    StopCmd = "stop"
    ResetHardCmd = "reset hard"
    ResetSoftCmd = "reset soft"
)

func (c *smashClient) StartServer() (string, error) {
   cmds := []string {
             SetSystemTargetCmd,
             StartCmd,
   }
   return c.Commands(cmds)
}

func (c *smashClient) StopServer() (string, error) {
   cmds := []string {
             SetSystemTargetCmd,
             StopCmd,
   }
   return c.Commands(cmds)
}

func (c *smashClient) ResetServerHard() (string, error) {
   cmds := []string {
             SetSystemTargetCmd,
             ResetHardCmd,
   }
   return c.Commands(cmds)
}


func (c *smashClient) ResetServerSoft() (string, error) {
   cmds := []string {
             SetSystemTargetCmd,
             ResetSoftCmd,
   }
   return c.Commands(cmds)
}
