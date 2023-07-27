package fusionauth

import (
	"context"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceTheme() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceThemelRead,
		Schema: map[string]*schema.Schema{
			"default_messages": {
				Type:             schema.TypeString,
				Required:         true,
				Computed:         true,
				Description:      "A properties file formatted String containing at least all of the message keys defined in the FusionAuth shipped messages file. Required if not copying an existing Theme.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique Id of the Email Template",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the Email Template.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func dataSourceThemeRead(_ context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(Client)

	resp, err := client.FAClient.RetrieveThemes()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := checkResponse(resp.StatusCode, nil); err != nil {
		return diag.FromErr(err)
	}
	name := data.Get("name").(string)
	var t *fusionauth.EmailTemplate

	if len(resp.Themes) > 0 {
		for i := range resp.Themes {
			if resp.Themes[i].Name == name {
				t = &resp.Themes[i]
				break
			}
		}
	}
	if t == nil {
		return diag.Errorf("couldn't find theme %s", name)
	}
	data.SetId(t.Id)
	if err := data.Set("default_messages", t.DefaultMessages); err != nil {
		return diag.Errorf("theme.default_messages: %s", err.Error())
	}
	// if err := data.Set("default_html_template", t.DefaultHtmlTemplate); err != nil {
	// 	return diag.Errorf("email.default_html_template: %s", err.Error())
	// }
	// if err := data.Set("default_subject", t.DefaultSubject); err != nil {
	// 	return diag.Errorf("email.default_subject: %s", err.Error())
	// }
	// if err := data.Set("default_text_template", t.DefaultTextTemplate); err != nil {
	// 	return diag.Errorf("email.default_text_template: %s", err.Error())
	// }
	// if err := data.Set("from_email", t.FromEmail); err != nil {
	// 	return diag.Errorf("email.from_email: %s", err.Error())
	// }
	// if err := data.Set("localized_from_names", t.LocalizedFromNames); err != nil {
	// 	return diag.Errorf("email.localized_from_names: %s", err.Error())
	// }
	// if err := data.Set("localized_html_templates", t.LocalizedHtmlTemplates); err != nil {
	// 	return diag.Errorf("email.localized_html_templates: %s", err.Error())
	// }
	// if err := data.Set("localized_subjects", t.LocalizedSubjects); err != nil {
	// 	return diag.Errorf("email.localized_subjects: %s", err.Error())
	// }
	// if err := data.Set("localized_text_templates", t.LocalizedTextTemplates); err != nil {
	// 	return diag.Errorf("email.localized_text_templates: %s", err.Error())
	// }
	// if err := data.Set("name", t.Name); err != nil {
	// 	return diag.Errorf("email.name: %s", err.Error())
	// }
	return nil
}
