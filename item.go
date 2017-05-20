package crawler

import (
	"fmt"
	"time"

	dmm "github.com/dmmlabo/dmm-go-sdk"
	"github.com/dmmlabo/dmm-go-sdk/api"
)

// AllItems _　TODO: using context.Context
func AllItems(
	c *dmm.Client, o Option,
	fn func(*dmm.Client, int64) (*api.ProductResponse, error),
) (chan api.Item, chan struct{}, chan error) {
	actressChan := make(chan api.Item, 100)
	errChan := make(chan error)
	doneChan := make(chan struct{}, 1)

	go func() {
		defer func() {
			doneChan <- struct{}{}
		}()

		var page int64
		for {
			res, err := fn(c, page)
			if err != nil {
				errChan <- fmt.Errorf("crawler ActressResponse gor error: %s", err)
				return
			}

			l := res.Items

			if len(l) == 0 {
				return
			}

			if o.MaxRepeat > 0 {
				if page >= int64(o.MaxRepeat) {
					return
				}
			}

			for _, a := range l {
				actressChan <- a
			}

			page = page + 1

			time.Sleep(o.Interval)
		}
	}()

	return actressChan, doneChan, errChan
}

// AdultVideoItemResponse get adult video items.
func AdultVideoItemResponse(c *dmm.Client, page int64) (*api.ProductResponse, error) {
	api := c.Product
	api.SetSort("date")
	api.SetLength(APILengthMax)
	api.SetOffset(page)
	api.SetSite("DMM.R18")
	api.SetFloor("videoa")
	api.SetService("digital")

	return api.Execute()
}
