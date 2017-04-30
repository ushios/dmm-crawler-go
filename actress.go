package clawrer

import (
	dmm "github.com/dmmlabo/dmm-go-sdk"
	"github.com/dmmlabo/dmm-go-sdk/api"
)

const (
	// APILengthMax is max size of length
	APILengthMax = 100
)

// AllActresses _
// func AllActresses(ctx context.Context, c *dmm.Client) (chan api.Actress, chan error) {
//
// }

// PartOfActresses _
func PartOfActresses(c *dmm.Client, page int64) (*api.ActressResponse, error) {
	api := c.Actress
	api.SetSort("id")
	api.SetLength(APILengthMax)
	api.SetOffset(page)

	return api.Execute()
}
