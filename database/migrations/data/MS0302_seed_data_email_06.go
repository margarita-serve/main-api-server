package data

import (
	domEmailEtt "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/entity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/utils"
)

// EmailTemplate06 data (TEXT)
func EmailTemplate06() domEmailEtt.EmailTemplate {
	return domEmailEtt.EmailTemplate{
		UUID:        utils.GenerateUUID(),
		Code:        "accuracy-alert-html",
		Name:        "Accuracy Alert Email (HTML)",
		IsActive:    true,
		EmailFormat: "HTML",
	}
}

// EmailTemplate06Version data
func EmailTemplate06Version() domEmailEtt.EmailTemplateVersion {
	return domEmailEtt.EmailTemplateVersion{
		Version:    utils.GenSemVersion(""),
		SubjectTpl: "Koreserve Deployment Accuracy Alert",
		BodyTpl: `{{define "T"}}<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
	</head>
	<body>
		<p>
			Accuracy Alert
		</p>
		<p>
		    {{index . "Body.deploymentName"}} Deployment's Accuracy status changed to {{index . "Body.additionalData"}}
		</p>
		<p>
			<!-- <button title="button title" class="" onclick=" window.open('{{index . "Body.deploymentURL"}}', '_blank'); return false;">Manage Deployment</button> -->
		</p>
	</body>
</html>{{end}}`,
	}
}
