package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAddressCreate,
		ReadContext:   resourceAddressRead,
		UpdateContext: resourceAddressUpdate,
		DeleteContext: resourceAddressDelete,
		Schema: map[string]*schema.Schema{
			"source": {
				Type:     schema.TypeString,
				Required: true,
			},
			"destinations": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"forwards": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAddressCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	create := swagger.VirtualAliasCreate{
		Source:       d.Get("source").(string),
		Destinations: stringSetToStringArray(d.Get("destinations").(*schema.Set)),
		Forwards:     stringSetToStringArray(d.Get("forwards").(*schema.Set)),
	}

	addressResponse, _, err := r.client.AddressApi.AddressCreate(*r.auth, []swagger.VirtualAliasCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(addressResponse[0].Id)

	retryErr := waitForResourceReady(ctx, d, addressChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for address to be created: %s", retryErr)
	}

	resourceAddressRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	addressResponse, _, err := r.client.AddressApi.AddressRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("source", addressResponse.Source)
	d.Set("destinations", addressResponse.Destinations)
	d.Set("forwards", addressResponse.Forwards)

	return diag.Diagnostics{}
}

func resourceAddressUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		update := swagger.VirtualAliasUpdate{
			Id:           d.Id(),
			Source:       d.Get("source").(string),
			Destinations: stringSetToStringArray(d.Get("destinations").(*schema.Set)),
			Forwards:     stringSetToStringArray(d.Get("forwards").(*schema.Set)),
		}

		_, _, err := r.client.AddressApi.AddressUpdate(*r.auth, []swagger.VirtualAliasUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForResourceReady(ctx, d, addressChecker(r, d))
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for address to be updated: %s", retryErr)
		}
	}

	return resourceAddressRead(ctx, d, m)
}

func resourceAddressDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.AddressApi.AddressDelete(*r.auth, []swagger.VirtualAliasRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	retryErr := waitForResourceDestroyed(ctx, d, addressChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for address to be destroyed: %s", retryErr)
	}

	d.SetId("")

	return diag.Diagnostics{}
}

func addressChecker(r *requester, d *schema.ResourceData) func() (bool, error) {
	return func() (bool, error) {
		addressResponse, _, err := r.client.AddressApi.AddressRead(*r.auth, d.Id())
		return addressResponse.Ready, err
	}
}
