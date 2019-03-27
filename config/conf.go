package config

import (
	"html/template"
)

type viewData struct {
	OpenIETag  template.HTML
	CloseIETag template.HTML
	Version    int
}

func Templates() []string {

	return []string{
		"templates/index.html",
		"templates/show_in_browser.html",
		"templates/logo.html",
		"templates/header_menu.html",
		"templates/social.html",
		"templates/footer_menu.html",
		"templates/footer_phone.html",
		"templates/footer_unsubscribe.html",
		"templates/not_an_offer.html",
		"templates/copy_right.html",
		"templates/content/262.html",
		"templates/content/264.html",
		"templates/content/266.html",
		"templates/content/271.html",
		"templates/content/274.html"}

}

func GetData() viewData {
	return viewData{
		OpenIETag: `<!--[if (gte mso 9)|(IE)]>
						  <table align="center" cellspacing="0" cellpadding="0" border="0" width="600" style="width: 600px;">
							  <tr>
								  <td style="padding-top: 0px; padding-bottom: 0px; padding-right: 0px; padding-left: 0px; margin-top: 0px; margin-right: 0px; margin-bottom: 0px; margin-left: 0px;">
					<![endif]-->`,
		CloseIETag: "<!--[if (gte mso 9)|(IE)]></td></tr></table><![endif]-->",
		Version:    262,
	}

}
