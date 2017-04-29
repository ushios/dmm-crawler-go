package clawrer

import dmm "github.com/dmmlabo/dmm-go-sdk"

// Actress _
type Actress struct {
	client *dmm.Client
}

// NewActress _
func NewActress(c *dmm.Client) *Actress {
	return &Actress{
		client: c,
	}
}
