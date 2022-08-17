package entity

type WebHook struct {
	ID                         string
	DeploymentID               string
	URL                        string //Webhook target URL, including the http:// or https:// prefix along with the url params
	Method                     string //GET, POST - Default is POST
	CustomHeader               string //Specify any custom header lines here
	MessageBody                string //Put the body of your message here
	CustomParametersAndSecrets string //Custom parameters and secrets allow you to add unique parameters and secure elements such as passwords
}
