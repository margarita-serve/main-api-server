package data

import (
	domEmailEtt "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/entity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/utils"
)

// EmailTemplate07 data (TEXT)
func EmailTemplate07() domEmailEtt.EmailTemplate {
	return domEmailEtt.EmailTemplate{
		UUID:        utils.GenerateUUID(),
		Code:        "service-alert-html",
		Name:        "Service Alert Email (HTML)",
		IsActive:    true,
		EmailFormat: "HTML",
	}
}

// EmailTemplate07Version data
func EmailTemplate07Version() domEmailEtt.EmailTemplateVersion {
	return domEmailEtt.EmailTemplateVersion{
		Version:    utils.GenSemVersion(""),
		SubjectTpl: "Koreserve Deployment Service Alert",
		BodyTpl: `{{define "T"}}<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
	</head>
	<body>
		<p>
			Service Alert
		</p>
		<p>
		    {{index . "Body.deploymentName"}} Deployment's Service status changed to {{index . "Body.additionalData"}}
		</p>
		<p>
			<!-- <button title="button title" class="" onclick=" window.open('{{index . "Body.deploymentURL"}}', '_blank'); return false;">Manage Deployment</button> -->
		</p>
	</body>
</html>{{end}}`,
	}
}
