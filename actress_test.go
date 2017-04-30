package clawrer

import (
	"testing"

	dmm "github.com/dmmlabo/dmm-go-sdk"
)

const (
	TestApiID       = "PVNKaRNVkvQbatzKa7Q5"
	TestAffiliateID = "FDGdnq4T6rvz-990"
)

func TestPartOfActresses(t *testing.T) {
	table := []struct {
		apiID       string
		affiliateID string
		page        int64
	}{
		{TestApiID, TestAffiliateID, 0},
	}

	for _, d := range table {
		c := dmm.New(d.affiliateID, d.apiID)
		res, err := PartOfActresses(c, d.page)
		if err != nil {
			t.Errorf("PartOfActresses got error: %s", err)
		}

		if len(res.Actresses) != APILengthMax {
			t.Errorf("res.Actresses length expected (%d) but (%d)", APILengthMax, len(res.Actresses))
		}
	}
}
