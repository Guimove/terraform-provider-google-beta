// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAppEngineDomainMapping_appEngineDomainMappingBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppEngineDomainMappingDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAppEngineDomainMapping_appEngineDomainMappingBasicExample(context),
			},
			{
				ResourceName:            "google_app_engine_domain_mapping.domain_mapping",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"override_strategy"},
			},
		},
	})
}

func testAccAppEngineDomainMapping_appEngineDomainMappingBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_app_engine_domain_mapping" "domain_mapping" {
  domain_name = "tf-test-domain%{random_suffix}.gcp.tfacc.hashicorptest.com"

  ssl_settings {
    ssl_management_type = "AUTOMATIC"
  }
}
`, context)
}

func testAccCheckAppEngineDomainMappingDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_app_engine_domain_mapping" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{AppEngineBasePath}}apps/{{project}}/domainMappings/{{domain_name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("AppEngineDomainMapping still exists at %s", url)
			}
		}

		return nil
	}
}
