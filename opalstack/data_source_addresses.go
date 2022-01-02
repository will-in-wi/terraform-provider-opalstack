package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAddresses() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAddressesRead,
		Schema: map[string]*schema.Schema{
			"addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: addressesSchema(false)},
			},
		},
	}
}

func dataSourceAddressesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	addressesResponse, _, err := r.client.AddressApi.AddressList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(addressesResponse))
	addresses := make([]map[string]interface{}, 0, len(addressesResponse))
	for _, elem := range addressesResponse {
		uuids = append(uuids, elem.Id)
		addresses = append(addresses, map[string]interface{}{
			"id":           elem.Id,
			"state":        elem.State,
			"ready":        elem.Ready,
			"source":       elem.Source,
			"destinations": elem.Destinations,
			"forwards":     elem.Forwards,
		})
	}
	if err := d.Set("addresses", addresses); err != nil {
		return diag.Errorf("Unable to assign addresses: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diag.Diagnostics{}
}
