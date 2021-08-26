package bitbucket

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

const testRepo string = "test-repo"

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"bitbucket": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("BITBUCKET_USERNAME"); v == "" {
		t.Fatal("BITBUCKET_USERNAME must be set for acceptence tests")
	}

	if v := os.Getenv("BITBUCKET_PASSWORD"); v == "" {
		t.Fatal("BITBUCKET_PASSWORD must be set for acceptence tests")
	}

	if v := os.Getenv("BITBUCKET_TEAM"); v == "" {
		t.Fatal("BITBUCKET_TEAM must be set for acceptence tests")
	}
}
