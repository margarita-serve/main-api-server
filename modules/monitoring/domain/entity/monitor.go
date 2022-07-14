package entity

import (
	"fmt"
	domSvcMonitor "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service"
	domAccuracySvcMonitorDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/accuracy/dto"
	domDriftSvcMonitorDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/data_drift/dto"
	"strings"
)

// Monitor type
type Monitor struct {
	ID                   string `gorm:"size:256"`
	ModelPackageID       string `gorm:"size:256"`
	FeatureDriftTracking bool
	AccuracyMonitoring   bool
	AssociationID        string `gorm:"size:256"`
	DriftStatus          string `gorm:"size:256"`
	AccuracyStatus       string `gorm:"size:256"`
	DataDriftSetting
	AccuracySetting
	ServiceHealthSetting
	BaseEntity
}

//// Validate
//func (m *Monitor) Validate() error {
//	return validation.ValidateStruct(m,
//		validation.Field(&m.ID, validation.Required, validation.NotNil, validation.Length(20, 20)),
//		validation.Field(&m.ModelPackageID, validation.Required, validation.NotNil, validation.Length(20, 20)),
//		validation.Field(&m.AccuracySetting.Measurement, validation.In("Percent", "Value")),
//		validation.Field(&m.AccuracySetting.MetricType, validation.In("rmse", "rmsle", "mae", "mad", "mape", "mean_tweedie_deviance", "gamma_deviance", "tpr", "accuracy",
//			"f1", "ppv", "fnr", "fpr")),
//		validation.Field(&m.DataDriftSetting.DriftThreshold, validation.Min(0.0), validation.Max(1.0)),
//		validation.Field(&m.DataDriftSetting.ImportanceThreshold, validation.Min(0.0), validation.Max(1.0)),
//	)
//}

func NewMonitor(id string, modelPackageID string) (*Monitor, error) {

	monitor := &Monitor{
		id,
		modelPackageID,
		false,
		false,
		"None",
		"unknown",
		"unknown",
		DataDriftSetting{},
		AccuracySetting{},
		ServiceHealthSetting{},
		BaseEntity{},
	}

	return monitor, nil
}

func (m *Monitor) SetDataDriftSetting(monitorRange string, driftMetricType string, driftThreshold float32, importanceThreshold float32, lowImportanceAtRiskCount int, lowImportanceFailingCount int, highImportanceAtRiskCount int, highImportanceFailingCount int) {
	var ds DataDriftSetting

	ds.MonitorRange = monitorRange
	ds.DriftMetricType = driftMetricType
	ds.DriftThreshold = driftThreshold
	ds.ImportanceThreshold = importanceThreshold
	ds.LowImportanceAtRiskCount = lowImportanceAtRiskCount
	ds.LowImportanceFailingCount = lowImportanceFailingCount
	ds.HighImportanceAtRiskCount = highImportanceAtRiskCount
	ds.HighImportanceFailingCount = highImportanceFailingCount

	m.DataDriftSetting = ds
}

func (m *Monitor) SetAccuracySetting(metricType string, measurement string, atRiskValue float32, failingValue float32) {
	var as AccuracySetting

	as.MetricType = metricType
	as.Measurement = strings.ToLower(measurement)
	as.AtRiskValue = atRiskValue
	as.FailingValue = failingValue

	m.AccuracySetting = as

}

func (m *Monitor) SetServiceHealthSetting() {
	var hs ServiceHealthSetting

	m.ServiceHealthSetting = hs
}

// DataDrift Func

func (m *Monitor) SetFeatureDriftTrackingOn(domSvc domSvcMonitor.IExternalDriftMonitorAdapter, reqDom domDriftSvcMonitorDTO.DataDriftCreateRequest) error {
	if m.FeatureDriftTracking == true {
		return fmt.Errorf("drift is already true")
	}

	res, err := domSvc.MonitorCreate(&reqDom)
	fmt.Printf("res : %v\n", res)

	if err != nil {
		return err
	}

	m.FeatureDriftTracking = true
	m.DriftStatus = "unknown"
	return nil
}

func (m *Monitor) SetFeatureDriftTrackingOff(domSvc domSvcMonitor.IExternalDriftMonitorAdapter, reqDom domDriftSvcMonitorDTO.DataDriftDeleteRequest) error {
	if m.FeatureDriftTracking == false {
		return fmt.Errorf("drift is already false")
	}

	err := domSvc.MonitorDelete(&reqDom)
	if err != nil {
		return err
	}

	m.FeatureDriftTracking = false
	m.DriftStatus = "unavailable"

	return nil
}

func (m *Monitor) PatchDataDriftSetting(domSvc domSvcMonitor.IExternalDriftMonitorAdapter, reqDom domDriftSvcMonitorDTO.DataDriftPatchRequest) error {
	//if m.FeatureDriftTracking == false {
	//	return fmt.Errorf("drift tracking is not ready")
	//}
	m.SetDataDriftSetting(
		reqDom.MonitorRange,
		"PSI",
		reqDom.DriftThreshold,
		reqDom.ImportanceThreshold,
		reqDom.LowImportanceAtRiskCount,
		reqDom.LowImportanceFailingCount,
		reqDom.HighImportanceAtRiskCount,
		reqDom.HighImportanceFailingCount,
	)
	res, err := domSvc.MonitorPatch(&reqDom)
	fmt.Printf("res: %v\n", res)

	if err != nil {
		return fmt.Errorf("drift setting change failed")
	}
	m.DriftStatus = "unknown"

	return err
}

func (m *Monitor) GetFeatureDetail(domSvc domSvcMonitor.IExternalDriftMonitorAdapter, reqDom domDriftSvcMonitorDTO.DataDriftGetRequest) (*domDriftSvcMonitorDTO.DataDriftGetResponse, error) {
	res, err := domSvc.MonitorGetDetail(&reqDom)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Monitor) GetFeatureDrift(domSvc domSvcMonitor.IExternalDriftMonitorAdapter, reqDom domDriftSvcMonitorDTO.DataDriftGetRequest) (*domDriftSvcMonitorDTO.DataDriftGetResponse, error) {
	res, err := domSvc.MonitorGetDrift(&reqDom)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Accuracy Func

func (m *Monitor) SetAccuracyMonitoringOn(domSvc domSvcMonitor.IExternalAccuracyMonitorAdapter, reqDom domAccuracySvcMonitorDTO.AccuracyCreateRequest) error {
	if m.AccuracyMonitoring == true {
		return fmt.Errorf("accuracy is already true")
	}
	if m.AssociationID != "None" {
		reqDom.AssociationID = m.AssociationID
	}

	res, err := domSvc.MonitorCreate(&reqDom)
	fmt.Printf("res: %v\n", res)

	if err != nil {
		return err
	}

	m.AccuracyMonitoring = true
	m.AccuracyStatus = "unknown"
	return nil
}

func (m *Monitor) SetAccuracyMonitoringOff(domSvc domSvcMonitor.IExternalAccuracyMonitorAdapter, reqDom domAccuracySvcMonitorDTO.AccuracyDeleteRequest) error {
	if m.AccuracyMonitoring == false {
		return fmt.Errorf("accuracy is already false")
	}

	err := domSvc.MonitorDelete(&reqDom)
	if err != nil {
		return err
	}

	m.AccuracyMonitoring = false
	m.AccuracyStatus = "unavailable"

	return nil
}

func (m *Monitor) PatchAccuracySetting(domSvc domSvcMonitor.IExternalAccuracyMonitorAdapter, reqDom domAccuracySvcMonitorDTO.AccuracyPatchRequest) error {
	//if m.AccuracyMonitoring == false {
	//	return fmt.Errorf("accuracy monitoring is not ready")
	//}
	m.SetAccuracySetting(
		reqDom.DriftMetrics,
		reqDom.DriftMeasurement,
		reqDom.AtriskValue,
		reqDom.FailingValue,
	)

	res, err := domSvc.MonitorPatch(&reqDom)
	fmt.Printf("res: %v\n", res)

	if err != nil {
		return fmt.Errorf("accuracy setting change failed")
	}
	m.AccuracyStatus = "unknown"

	return err

}

func (m *Monitor) GetAccuracy(domSvc domSvcMonitor.IExternalAccuracyMonitorAdapter, reqDom domAccuracySvcMonitorDTO.AccuracyGetRequest) (*domAccuracySvcMonitorDTO.AccuracyGetResponse, error) {
	res, err := domSvc.MonitorGetAccuracy(&reqDom)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Monitor) PostActual(domSvc domSvcMonitor.IExternalAccuracyMonitorAdapter, reqDom domAccuracySvcMonitorDTO.AccuracyPostActualRequest) (*domAccuracySvcMonitorDTO.AccuracyPostActualResponse, error) {
	res, err := domSvc.MonitorPostActual(&reqDom)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Monitor) SetDriftStatusPass() {
	m.DriftStatus = "pass"
}

func (m *Monitor) SetDriftStatusWarning() {
	m.DriftStatus = "warning"
}

func (m *Monitor) SetDriftStatusFailing() {
	m.DriftStatus = "failing"
}

func (m *Monitor) SetDriftStatusUnknown() {
	m.DriftStatus = "unknown"
}

func (m *Monitor) SetAccuracyStatusPass() {
	m.DriftStatus = "pass"
}

func (m *Monitor) SetAccuracyStatusWarning() {
	m.DriftStatus = "warning"
}

func (m *Monitor) SetAccuracyStatusFailing() {
	m.DriftStatus = "failing"
}

func (m *Monitor) SetAccuracyStatusUnknown() {
	m.DriftStatus = "unknown"
}
