package crawler

import (
	"fmt"
	"time"

	dmm "github.com/dmmlabo/dmm-go-sdk"
	"github.com/dmmlabo/dmm-go-sdk/api"
)

// AllActresses _　TODO: using context.Context
func AllActresses(
	c *dmm.Client, o Option,
	fn func(*dmm.Client, int64) (*api.ActressResponse, error),
) (chan api.Actress, chan struct{}, chan error) {
	actressChan := make(chan api.Actress, 100)
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

			l := res.Actresses

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

// ActressResponse get api response.
func ActressResponse(c *dmm.Client, page int64) (*api.ActressResponse, error) {
	api := c.Actress
	api.SetSort("id")
	api.SetLength(APILengthMax)
	api.SetOffset(page*APILengthMax + 1) // TODO: what is +1

	return api.Execute()
}
