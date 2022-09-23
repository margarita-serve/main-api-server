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
		//Service init
		// AuthenticationSvc, err := appAuthsSvc.NewAuthenticationSvc(h)
		// if err != nil {
		// 	return err
		// }

		// ClusterInfoSvc, err := appResourceSvc.NewClusterInfoService(h)
		// if err != nil {
		// 	return err
		// }

		// PredictionEnvSvc, err := appResourceSvc.NewPredictionEnvService(h, ClusterInfoSvc)
		// if err != nil {
		// 	return err
		// }

		// ProjectSvc, err := appProjectSvc.NewProjectService(h)
		// if err != nil {
		// 	return err
		// }

		// ModelPackageSvc, err := appModelPackageSvc.NewModelPackageService(h, ProjectSvc)
		// if err != nil {
		// 	return err
		// }

		// DeploymentGetByIDInternalSvc, err := appDeploymentSvc.NewDeploymentGetByIDInternalService(h)
		// if err != nil {
		// 	return err
		// }

		// MonitorSvc, err := appMonitorSvc.NewMonitorService(h, appModelPackageInstance.ModelPackageSvc, publisher)
		// if err != nil {
		// 	return err
		// }

		// MonitorMessagingSvc, err := appMonitorSvc.NewMessagingService(h, MonitorSvc)
		// if err != nil {
		// 	return err
		// }

		// DeploymentSvc, err := appDeploymentSvc.NewDeploymentService(h, PredictionEnvSvc, ProjectSvc, appModelPackageInstance.ModelPackageSvc, appMonitorInstance.MonitorSvc, publisher)
		// if err != nil {
		// 	return err
		// }

		// DeploymentGovernanceHistorySvc, err := appDeploymentSvc.NewDeploymentGovernanceHistoryService(h)
		// if err != nil {
		// 	return err
		// }

		// NotiSvc, err := appNotiSvc.NewNotiService(h, EmailSvc, DeploymentSvc, ProjectSvc, AuthenticationSvc, DeploymentGovernanceHistorySvc, WebHookSvc)
		// if err != nil {
		// 	return err
		// }

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
		publisher.Subscribe(appMonitorService.MonitorSvc, common.DeploymentInferenceServiceCreated{}, common.DeploymentDeleted{}, common.DeploymentActived{}, common.DeploymentInActived{}, common.DeploymentModelReplaced{}, common.DeploymentFeatureDriftTrackingEnabled{}, common.DeploymentFeatureDriftTrackingDisabled{}, common.DeploymentAccuracyAnalyzeEnabled{}, common.DeploymentAccuracyAnalyzeDisabled{}, common.DeploymentAssociationIDUpdated{})
		publisher.Subscribe(appDeploymentService.DeploymentSvc, common.MonitoringCreated{}, common.MonitoringCreateFailed{}, common.MonitoringAccuracyStatusChangedToFailing{}, common.MonitoringAccuracyStatusChangedToAtrisk{}, common.MonitoringDataDriftStatusChangedToFailing{}, common.MonitoringDataDriftStatusChangedToAtrisk{}, common.MonitoringServiceHealthStatusChangedToFailing{}, common.MonitoringServiceHealthStatusChangedToAtrisk{})
		publisher.Subscribe(appNotiInstance.NotiSvc, common.MonitoringAccuracyStatusChangedToFailing{}, common.MonitoringAccuracyStatusChangedToAtrisk{}, common.MonitoringDataDriftStatusChangedToFailing{}, common.MonitoringDataDriftStatusChangedToAtrisk{}, common.MonitoringServiceHealthStatusChangedToFailing{}, common.MonitoringServiceHealthStatusChangedToAtrisk{})
	}

	return nil
}
