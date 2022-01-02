package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAddressRead,
		Schema:      addressesSchema(true),
	}
}

func dataSourceAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	addressResponse, _, err := r.client.AddressApi.AddressRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", addressResponse.State)
	d.Set("ready", addressResponse.Ready)
	d.Set("source", addressResponse.Source)
	d.Set("destinations", addressResponse.Destinations)
	d.Set("forwards", addressResponse.Forwards)
	d.SetId(addressResponse.Id)

	return diag.Diagnostics{}
}
