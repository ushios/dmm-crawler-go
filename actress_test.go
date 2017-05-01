package clawrer

import (
	"os"
	"testing"
	"time"

	dmm "github.com/dmmlabo/dmm-go-sdk"
	"github.com/dmmlabo/dmm-go-sdk/api"
)

const (
	TestAPIIDEnvKey       = "TEST_DMM_API_ID"
	TestAffiliateIDEnvKey = "TEST_DMM_AFFILIATE_ID"
)

func ApiID() string {
	return os.Getenv(TestAPIIDEnvKey)
}

func AffiliateID() string {
	return os.Getenv(TestAffiliateIDEnvKey)
}

func TestAllActresses(t *testing.T) {
	table := []struct {
		apiID       string
		affiliateID string
		interval    time.Duration
		maxRepeat   int
	}{
		{ApiID(), AffiliateID(), 10 * time.Millisecond, 2},
	}

	for _, d := range table {
		c := dmm.New(d.affiliateID, d.apiID)
		o := Option{
			Interval:  d.interval,
			MaxRepeat: d.maxRepeat,
		}

		list := []api.Actress{}
		actressChan, doneChan, errChan := AllActresses(c, o)
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

func TestActressList(t *testing.T) {
	table := []struct {
		apiID       string
		affiliateID string
		page        int64
	}{
		{ApiID(), AffiliateID(), 0},
	}

	for _, d := range table {
		c := dmm.New(d.affiliateID, d.apiID)
		res, err := ActressList(c, d.page)
		if err != nil {
			t.Errorf("ActressList got error: %s", err)
		}

		if len(res.Actresses) != APILengthMax {
			t.Errorf("res.Actresses length expected (%d) but (%d)", APILengthMax, len(res.Actresses))
		}
	}
}
