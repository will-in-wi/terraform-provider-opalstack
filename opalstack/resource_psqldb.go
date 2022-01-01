package opalstack

import (
	"context"
	"fmt"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func allowedPgDbCharsets() []string {
	return []string{
		// See swagger/model_psql_charset.go for list.
		"utf8",
		"euc_cn",
		"euc_jis_2004",
		"euc_jp",
		"euc_kr",
		"euc_tw",
		"iso_8859_5",
		"iso_8859_6",
		"iso_8859_7",
		"iso_8859_8",
		"koi8r",
		"koi8u",
		"latin1",
		"latin2",
		"latin3",
		"latin4",
		"latin5",
		"latin6",
		"latin7",
		"latin8",
		"latin9",
		"latin10",
		"sql_ascii",
		"win866",
		"win874",
		"win1250",
		"win1251",
		"win1252",
		"win1253",
		"win1254",
		"win1255",
		"win1256",
		"win1257",
		"win1258",
	}
}

func resourcePsqldb() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePsqldbCreate,
		ReadContext:   resourcePsqldbRead,
		UpdateContext: resourcePsqldbUpdate,
		DeleteContext: resourcePsqldbDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"server": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"charset": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "utf8",
				ValidateFunc: validation.StringInSlice(allowedPgDbCharsets(), false),
			},
			"dbusers_readwrite": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dbusers_readonly": {
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

func resourcePsqldbCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	charset := swagger.PsqlCharset(d.Get("charset").(string))

	create := swagger.PsqlDbCreate{
		Name:             d.Get("name").(string),
		Server:           d.Get("server").(string),
		Charset:          &charset,
		DbusersReadwrite: stringSetToStringArray(d.Get("dbusers_readwrite").(*schema.Set)),
		DbusersReadonly:  stringSetToStringArray(d.Get("dbusers_readonly").(*schema.Set)),
	}

	psqldbResponse, _, err := r.client.PsqldbApi.PsqldbCreate(*r.auth, []swagger.PsqlDbCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(psqldbResponse[0].Id)

	retryErr := waitForPsqldbReady(ctx, d, r)
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for psqldb to be created: %s", retryErr)
	}

	resourcePsqldbRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourcePsqldbRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	psqldbResponse, _, err := r.client.PsqldbApi.PsqldbRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("name", psqldbResponse.Name)
	d.Set("server", psqldbResponse.Server)
	d.Set("charset", psqldbResponse.Charset)
	d.Set("dbusers_readwrite", psqldbResponse.DbusersReadwrite)
	d.Set("dbusers_readonly", psqldbResponse.DbusersReadonly)

	return diags
}

func resourcePsqldbUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		update := swagger.PsqlDbUpdate{
			Id:               d.Id(),
			DbusersReadwrite: stringSetToStringArray(d.Get("dbusers_readwrite").(*schema.Set)),
			DbusersReadonly:  stringSetToStringArray(d.Get("dbusers_readonly").(*schema.Set)),
		}

		_, _, err := r.client.PsqldbApi.PsqldbUpdate(*r.auth, []swagger.PsqlDbUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForPsqldbReady(ctx, d, r)
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for psqldb to be updated: %s", retryErr)
		}
	}

	return resourcePsqldbRead(ctx, d, m)
}

func resourcePsqldbDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := r.client.PsqldbApi.PsqldbDelete(*r.auth, []swagger.PsqlDbRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId("")

	return diags
}

func waitForPsqldbReady(ctx context.Context, d *schema.ResourceData, r *requester) error {
	return resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		var err error
		psqldbResponse, _, err := r.client.PsqldbApi.PsqldbRead(*r.auth, d.Id())
		if err != nil {
			return resource.NonRetryableError(err)
		}

		if !psqldbResponse.Ready {
			return resource.RetryableError(fmt.Errorf("not ready yet"))
		}

		return nil
	})
}
