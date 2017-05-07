package clawrer

import (
	"testing"
	"time"

	dmm "github.com/dmmlabo/dmm-go-sdk"
	"github.com/dmmlabo/dmm-go-sdk/api"
)

func TestAllProduct(t *testing.T) {
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

		list := []api.Item{}
		itemChan, doneChan, errChan := AllItems(c, o, AdultVideoItemResponse)
	Product:
		for {
			select {
			case err := <-errChan:
				t.Fatalf("clawrer.AllActresses got error: %s", err)
			case item := <-itemChan:
				list = append(list, item)
			case <-doneChan:
				break Product
			}
		}

		if len(list) != d.maxRepeat*APILengthMax {
			t.Errorf("list length expected (%d) but (%d)", d.maxRepeat*APILengthMax, len(list))
		}

	}
}

func TestAdultVideoItemResponse(t *testing.T) {
	table := []struct {
		apiID       string
		affiliateID string
		page        int64
	}{
		{APIID(), AffiliateID(), 0},
	}

	for _, d := range table {
		c := dmm.New(d.affiliateID, d.apiID)
		res, err := AdultVideoItemResponse(c, d.page)
		if err != nil {
			t.Errorf("ActressList got error: %s", err)
		}

		if len(res.Items) != APILengthMax {
			t.Errorf("res.Items length expected (%d) but (%d)", APILengthMax, len(res.Items))
		}
	}
}
