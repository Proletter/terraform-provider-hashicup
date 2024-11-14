// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hashicup

import (
	"context"

	"github.com/hashicorp-demoapp/hashicup-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("hashicup_HOST", nil),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("hashicup_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("hashicup_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hashicup_order": resourceOrder(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"hashicup_coffees":     dataSourceCoffees(),
			"hashicup_order":       dataSourceOrder(),
			"hashicup_ingredients": dataSourceIngredients(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := hashicup.NewClient(host, &username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create hashicup client",
				Detail:   "Unable to authenticate user for authenticated hashicup client",
			})

			return nil, diags
		}

		return c, diags
	}

	c, err := hashicup.NewClient(host, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create hashicup client",
			Detail:   "Unable to create anonymous hashicup client",
		})
		return nil, diags
	}

	return c, diags
}
