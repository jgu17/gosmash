package client

type BootOption string

const (
	BootOnce       BootOption = "once"
	BootAlways     BootOption = "always"
	BootNever      BootOption = "never"
	BootConnect    BootOption = "connect"
	BootDisconnect BootOption = "disconnect"
)

type BootState string

const (
	NoBootState     = "No_Boot"
	BootOnceState   = "Once"
	BootAlwaysState = "Always"
)
