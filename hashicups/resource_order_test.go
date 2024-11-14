// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hashicup

import (
	"fmt"
	"testing"

	hc "github.com/hashicorp-demoapp/hashicup-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAcchashicupOrderBasic(t *testing.T) {
	coffeeID := "1"
	quantity := "2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckhashicupOrderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckhashicupOrderConfigBasic(coffeeID, quantity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckhashicupOrderExists("hashicup_order.new"),
				),
			},
		},
	})
}

func testAccCheckhashicupOrderDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*hc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "hashicup_order" {
			continue
		}

		orderID := rs.Primary.ID

		err := c.DeleteOrder(orderID, &c.Token)
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckhashicupOrderConfigBasic(coffeeID, quantity string) string {
	return fmt.Sprintf(`
	resource "hashicup_order" "new" {
		items {
			coffee {
				id = %s
			}
    		quantity = %s
  		}
	}
	`, coffeeID, quantity)
}

func testAccCheckhashicupOrderExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No OrderID set")
		}

		return nil
	}
}
