package clawrer

import (
	"testing"
	"time"

	dmm "github.com/dmmlabo/dmm-go-sdk"
	"github.com/dmmlabo/dmm-go-sdk/api"
)

func TestAllActresses(t *testing.T) {
	table := []struct {
		apiID       string
		affiliateID string
		interval    time.Duration
		maxRepeat   int
	}{
		{APIID(), AffiliateID(), 10 * time.Millisecond, 2},
	}

	for _, d := range table {
		c := dmm.New(d.affiliateID, d.apiID)
		o := Option{
			Interval:  d.interval,
			MaxRepeat: d.maxRepeat,
		}

		list := []api.Actress{}
		actressChan, doneChan, errChan := AllActresses(c, o, ActressResponse)
	ACTRESS:
		for {
			select {
			case err := <-errChan:
				t.Fatalf("clawrer.AllActresses got error: %s", err)
			case actress := <-actressChan:
				list = append(list, actress)
			case <-doneChan:
				break ACTRESS
			}
		}

		if len(list) != d.maxRepeat*APILengthMax {
			t.Errorf("list length expected (%d) but (%d)", d.maxRepeat*APILengthMax, len(list))
		}

	}
}

func TestActressResponse(t *testing.T) {
	table := []struct {
		apiID       string
		affiliateID string
		page        int64
	}{
		{APIID(), AffiliateID(), 0},
	}

	for _, d := range table {
		c := dmm.New(d.affiliateID, d.apiID)
		res, err := ActressResponse(c, d.page)
		if err != nil {
			t.Errorf("ActressList got error: %s", err)
		}

		if len(res.Actresses) != APILengthMax {
			t.Errorf("res.Actresses length expected (%d) but (%d)", APILengthMax, len(res.Actresses))
		}
	}
}
