package crawler

import "os"

const (
	TestAPIIDEnvKey       = "TEST_DMM_API_ID"
	TestAffiliateIDEnvKey = "TEST_DMM_AFFILIATE_ID"
)

func APIID() string {
	return os.Getenv(TestAPIIDEnvKey)
}

func AffiliateID() string {
	return os.Getenv(TestAffiliateIDEnvKey)
}
