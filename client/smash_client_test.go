package client

import "testing"

func TestCommand(t *testing.T) {
    e := smashendpoint.
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}
