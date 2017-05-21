package crawler

import "os"

const (
	TestAPIIDEnvKey       = "DMM_TEST_API_ID"
	TestAffiliateIDEnvKey = "DMM_TEST_AFFILIATE_ID"
)

func APIID() string {
	return os.Getenv(TestAPIIDEnvKey)
}

func AffiliateID() string {
	return os.Getenv(TestAffiliateIDEnvKey)
}
