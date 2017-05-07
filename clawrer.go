package clawrer

import (
	"fmt"
	"log"
	"time"

	dmm "github.com/dmmlabo/dmm-go-sdk"
)

// Clawrer _
type Clawrer struct {
}

type (
	// Option is clawrer options
	Option struct {
		Interval  time.Duration
		MaxRepeat int
	}

	// DmmAPIResponse is interface object of dmm api response
	DmmAPIResponse interface{}
)

// All _
func All(
	c *dmm.Client, o Option, ch chan interface{},
	fnResponse func(*dmm.Client, int64) (DmmAPIResponse, error),
	fnFromResponse func(DmmAPIResponse) []interface{},
) (chan struct{}, chan error) {
	resChan := make(chan DmmAPIResponse, 5)
	doneChan := make(chan struct{}, 1)

	listDoneChan, errChan := AllList(c, o, resChan, fnResponse)

	go func() {
		defer func() {
			doneChan <- struct{}{}
		}()
	LOOP:
		for {
			select {
			case <-listDoneChan:
				break LOOP
			case res := <-resChan:
				l := fnFromResponse(res)

				if len(l) < 1 {
					break LOOP
				}

				for _, obj := range l {
					ch <- obj
				}

			default:
				log.Fatal("unknown chan in dmm-clawrer-go.All select")
			}
		}
	}()

	return doneChan, errChan
}

// AllList get all list
func AllList(
	c *dmm.Client, o Option, ch chan DmmAPIResponse,
	fn func(*dmm.Client, int64) (DmmAPIResponse, error),
) (chan struct{}, chan error) {
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
				errChan <- fmt.Errorf("clawrer.All gor error: %s", err)
				return
			}

			if o.MaxRepeat > 0 {
				if page >= int64(o.MaxRepeat) {
					return
				}
			}

			ch <- res

			page = page + 1

			time.Sleep(o.Interval)
		}
	}()

	return doneChan, errChan
}
