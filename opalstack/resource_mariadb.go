package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func allowedMariaDbCharsets() []string {
	return []string{
		// See swagger/model_maria_charset.go for list.
		"utf8",
		"armscii8",
		"ascii",
		"big5",
		"binary",
		"cp1250",
		"cp1251",
		"cp1256",
		"cp1257",
		"cp850",
		"cp852",
		"cp866",
		"cp932",
		"dec8",
		"eucjpms",
		"euckr",
		"gb2312",
		"gbk",
		"geostd8",
		"greek",
		"hebrew",
		"hp8",
		"keybcs2",
		"koi8r",
		"koi8u",
		"latin1",
		"latin2",
		"latin5",
		"latin7",
		"macce",
		"macroman",
		"sjis",
		"swe7",
		"tis620",
		"ucs2",
		"ujis",
		"utf16",
		"utf16le",
		"utf32",
		"utf8mb4",
	}
}

func resourceMariadb() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMariadbCreate,
		ReadContext:   resourceMariadbRead,
		UpdateContext: resourceMariadbUpdate,
		DeleteContext: resourceMariadbDelete,
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
				Default:      "utf8mb4",
				ValidateFunc: validation.StringInSlice(allowedMariaDbCharsets(), false),
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

func resourceMariadbCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	charset := swagger.MariaCharset(d.Get("charset").(string))

	create := swagger.MariaDbCreate{
		Name:             d.Get("name").(string),
		Server:           d.Get("server").(string),
		Charset:          &charset,
		DbusersReadwrite: stringSetToStringArray(d.Get("dbusers_readwrite").(*schema.Set)),
		DbusersReadonly:  stringSetToStringArray(d.Get("dbusers_readonly").(*schema.Set)),
	}

	mariadbResponse, _, err := r.client.MariadbApi.MariadbCreate(*r.auth, []swagger.MariaDbCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(mariadbResponse[0].Id)

	retryErr := waitForResourceReady(ctx, d, mariadbChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for mariadb to be created: %s", retryErr)
	}

	resourceMariadbRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceMariadbRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	mariadbResponse, _, err := r.client.MariadbApi.MariadbRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("name", mariadbResponse.Name)
	d.Set("server", mariadbResponse.Server)
	d.Set("charset", mariadbResponse.Charset)
	d.Set("dbusers_readwrite", mariadbResponse.DbusersReadwrite)
	d.Set("dbusers_readonly", mariadbResponse.DbusersReadonly)

	return diag.Diagnostics{}
}

func resourceMariadbUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		update := swagger.MariaDbUpdate{
			Id:               d.Id(),
			DbusersReadwrite: stringSetToStringArray(d.Get("dbusers_readwrite").(*schema.Set)),
			DbusersReadonly:  stringSetToStringArray(d.Get("dbusers_readonly").(*schema.Set)),
		}

		_, _, err := r.client.MariadbApi.MariadbUpdate(*r.auth, []swagger.MariaDbUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForResourceReady(ctx, d, mariadbChecker(r, d))
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for mariadb to be updated: %s", retryErr)
		}
	}

	return resourceMariadbRead(ctx, d, m)
}

func resourceMariadbDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.MariadbApi.MariadbDelete(*r.auth, []swagger.MariaDbRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	retryErr := waitForResourceDestroyed(ctx, d, mariadbChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for mariadb to be destroyed: %s", retryErr)
	}

	d.SetId("")

	return diag.Diagnostics{}
}

func mariadbChecker(r *requester, d *schema.ResourceData) func() (bool, error) {
	return func() (bool, error) {
		mariadbResponse, _, err := r.client.MariadbApi.MariadbRead(*r.auth, d.Id())
		return mariadbResponse.Ready, err
	}
}
