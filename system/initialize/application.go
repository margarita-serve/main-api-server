package initialize

import (
	appAuths "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/application"
	appAuthsSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/application/service"
	appDeployment "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application"
	appDeploymentSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/service"
	appEmail "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/application"
	appEmailSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/application/service"
	appModelPackage "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application"
	appModelPackageSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"
	appMonitor "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application"
	appMonitorSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/service"
	appNoti "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application"
	appNotiSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/service"
	appProject "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application"
	appProjectSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application/service"
	appResource "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/application"
	appResourceSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// OpenAllCacheConnection open all cache connection
func InitAllApplication(h *handler.Handler) error {
	// cfg, err := h.GetConfig()
	// if err != nil {
	// 	return err
	// }

	err := InintApplication(h)

	return err
}

// OpenCacheConnection open CacheConnection
func InintApplication(h *handler.Handler) error {
	if h != nil {
		//Service init
		AuthenticationSvc, err := appAuthsSvc.NewAuthenticationSvc(h)
		if err != nil {
			return err
		}

		EmailSvc, err := appEmailSvc.NewEmailService(h)
		if err != nil {
			return err
		}

		EmailTemplateSvc, err := appEmailSvc.NewEmailTemplateService(h)
		if err != nil {
			return err
		}

		ClusterInfoSvc, err := appResourceSvc.NewClusterInfoService(h)
		if err != nil {
			return err
		}

		PredictionEnvSvc, err := appResourceSvc.NewPredictionEnvService(h, ClusterInfoSvc)
		if err != nil {
			return err
		}

		ProjectSvc, err := appProjectSvc.NewProjectService(h)
		if err != nil {
			return err
		}

		ModelPackageSvc, err := appModelPackageSvc.NewModelPackageService(h, ProjectSvc)
		if err != nil {
			return err
		}

		DeploymentGetByIDInternalSvc, err := appDeploymentSvc.NewDeploymentGetByIDInternalService(h)
		if err != nil {
			return err
		}
		DeploymentGovernanceHistorySvc, err := appDeploymentSvc.NewDeploymentGovernanceHistoryService(h)
		if err != nil {
			return err
		}

		WebHookEventSvc, err := appNotiSvc.NewWebHookEventService(h)
		if err != nil {
			return err
		}

		WebHookSvc, err := appNotiSvc.NewWebHookService(h, WebHookEventSvc)
		if err != nil {
			return err
		}

		NotiSvc, err := appNotiSvc.NewNotiService(h, EmailSvc, DeploymentGetByIDInternalSvc, ProjectSvc, AuthenticationSvc, DeploymentGovernanceHistorySvc, WebHookSvc)
		if err != nil {
			return err
		}

		MonitorSvc, err := appMonitorSvc.NewMonitorService(h, ModelPackageSvc, NotiSvc)
		if err != nil {
			return err
		}

		MonitorMessagingSvc, err := appMonitorSvc.NewMessagingService(h, MonitorSvc)
		if err != nil {
			return err
		}

		DeploymentSvc, err := appDeploymentSvc.NewDeploymentService(h, PredictionEnvSvc, ProjectSvc, ModelPackageSvc, MonitorSvc)
		if err != nil {
			return err
		}

		//Application init
		appAuthsInstance, err := appAuths.NewAuthsApp(h, AuthenticationSvc)
		if err != nil {
			return err
		}
		h.SetApp("auths", appAuthsInstance)

		appEmailInstance, err := appEmail.NewEmailApp(h, EmailSvc, EmailTemplateSvc)
		if err != nil {
			return err
		}
		h.SetApp("email", appEmailInstance)

		appResourceInstance, err := appResource.NewResourceApp(h, ClusterInfoSvc, PredictionEnvSvc)
		if err != nil {
			return err
		}
		h.SetApp("resource", appResourceInstance)

		appProjectInstance, err := appProject.NewProjectApp(h, ProjectSvc)
		if err != nil {
			return err
		}
		h.SetApp("project", appProjectInstance)

		appModelPackageInstance, err := appModelPackage.NewModelPackageApp(h, ModelPackageSvc)
		if err != nil {
			return err
		}
		h.SetApp("modelpackage", appModelPackageInstance)

		appNotiInstance, err := appNoti.NewNotiApp(h, NotiSvc, WebHookSvc)
		if err != nil {
			return err
		}
		h.SetApp("noti", appNotiInstance)

		appMonitorInstance, err := appMonitor.NewMonitorApp(h, MonitorSvc, MonitorMessagingSvc)
		if err != nil {
			return err
		}
		h.SetApp("monitor", appMonitorInstance)

		appDeploymentInstance, err := appDeployment.NewDeploymentApp(h, DeploymentSvc)
		if err != nil {
			return err
		}
		h.SetApp("deployment", appDeploymentInstance)

	}

	return nil
}
