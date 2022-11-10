package initialize

import (
	appAuths "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/application"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	appDeployment "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application"
	appEmail "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/application"
	appModelPackage "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application"
	appMonitor "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application"
	appNoti "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application"
	appProject "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// OpenAllCacheConnection open all cache connection
func InitAllApplication(h *handler.Handler) error {
	// cfg, err := h.GetConfig()
	// if err != nil {
	// 	return err
	// }
	publisher := common.NewEventPublisher()
	err := InintApplication(h, publisher)

	return err
}

// OpenCacheConnection open CacheConnection
func InintApplication(h *handler.Handler, publisher common.EventPublisher) error {
	if h != nil {

		//Application init
		appProjectService, err := appProject.NewProjectApp(h)
		if err != nil {
			return err
		}
		h.SetApp("project", appProjectService)

		appModelPackageService, err := appModelPackage.NewModelPackageApp(h, appProjectService.ProjectSvc)
		if err != nil {
			return err
		}
		h.SetApp("modelpackage", appModelPackageService)

		appMonitorService, err := appMonitor.NewMonitorApp(h, appModelPackageService.ModelPackageSvc, publisher)
		if err != nil {
			return err
		}
		h.SetApp("monitor", appMonitorService)

		appAuthsService, err := appAuths.NewAuthsApp(h)
		if err != nil {
			return err
		}
		h.SetApp("auths", appAuthsService)

		appEmailService, err := appEmail.NewEmailApp(h)
		if err != nil {
			return err
		}
		h.SetApp("email", appEmailService)

		// appResourceService, err := appResource.NewResourceApp(h, ClusterInfoSvc, PredictionEnvSvc)
		// if err != nil {
		// 	return err
		// }
		// h.SetApp("resource", appResourceService)

		appDeploymentService, err := appDeployment.NewDeploymentApp(h, appProjectService.ProjectSvc, appModelPackageService.ModelPackageSvc, appMonitorService.MonitorSvc, publisher)
		if err != nil {
			return err
		}
		h.SetApp("deployment", appDeploymentService)

		appNotiInstance, err := appNoti.NewNotiApp(h, appEmailService.EmailSvc, appDeploymentService.DeploymentSvc, appProjectService.ProjectSvc, appAuthsService.AuthenticationSvc)
		if err != nil {
			return err
		}
		h.SetApp("noti", appNotiInstance)

		publisher.Subscribe(appModelPackageService.ModelPackageSvc, common.DeploymentCreated{}, common.DeploymentModelReplaced{})
		publisher.Subscribe(appMonitorService.MonitorSvc, common.DeploymentActived{}, common.DeploymentInActived{})
		publisher.Subscribe(appDeploymentService.DeploymentSvc, common.MonitoringAccuracyMonitorDisabled{}, common.MonitoringAccuracyMonitorEnabled{}, common.MonitoringDataDriftMonitorDisabled{}, common.MonitoringDataDriftMonitorEnabled{}, common.MonitoringAccuracyStatusChangedToFailing{}, common.MonitoringAccuracyStatusChangedToAtrisk{}, common.MonitoringDataDriftStatusChangedToFailing{}, common.MonitoringDataDriftStatusChangedToAtrisk{}, common.MonitoringServiceHealthStatusChangedToFailing{}, common.MonitoringServiceHealthStatusChangedToAtrisk{})
		publisher.Subscribe(appNotiInstance.NotiSvc, common.MonitoringAccuracyStatusChangedToFailing{}, common.MonitoringAccuracyStatusChangedToAtrisk{}, common.MonitoringDataDriftStatusChangedToFailing{}, common.MonitoringDataDriftStatusChangedToAtrisk{}, common.MonitoringServiceHealthStatusChangedToFailing{}, common.MonitoringServiceHealthStatusChangedToAtrisk{})
	}

	return nil
}
