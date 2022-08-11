package entity

import (
	"fmt"
	domSvcMonitor "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service"
	domAccuracySvcMonitorDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/accuracy/dto"
	domDriftSvcMonitorDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/data_drift/dto"
	domGraphSvcMonitorDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/graph/dto"
	"strings"
)

// Monitor type
type Monitor struct {
	// drift, accuracy 생성 상태값 추가
	ID                   string `gorm:"size:256"`
	ModelPackageID       string `gorm:"size:256"`
	FeatureDriftTracking bool
	AccuracyMonitoring   bool
	AssociationID        string `gorm:"size:256"`
	DriftStatus          string `gorm:"size:256"`
	AccuracyStatus       string `gorm:"size:256"`
	DriftCreated         bool
	AccuracyCreated      bool
	DataDriftSetting
	AccuracySetting
	ServiceHealthSetting
	BaseEntity
}

func NewMonitor(id string, modelPackageID string) (*Monitor, error) {
	// drift, accuracy 생성 상태값 추가

	monitor := &Monitor{
		id,
		modelPackageID,
		false,
		false,
		"None",
		"unknown",
		"unknown",
		false,
		false,
		DataDriftSetting{},
		AccuracySetting{},
		ServiceHealthSetting{},
		BaseEntity{},
	}

	return monitor, nil
}

func (m *Monitor) SetDataDriftSetting(monitorRange string, driftMetricType string, driftThreshold float32, importanceThreshold float32, lowImportanceAtRiskCount int, lowImportanceFailingCount int, highImportanceAtRiskCount int, highImportanceFailingCount int) {
	if monitorRange == "" {
		m.MonitorRange = "7d"
	} else {
		m.MonitorRange = monitorRange
	}
	if driftMetricType == "" {
		m.DriftMetricType = "PSI"
	} else {
		m.DriftMetricType = driftMetricType
	}
	if driftThreshold == 0 {
		m.DriftThreshold = 0.15
	} else {
		m.DriftThreshold = driftThreshold
	}
	if importanceThreshold == 0 {
		m.ImportanceThreshold = 0.5
	} else {
		m.ImportanceThreshold = importanceThreshold
	}
	if lowImportanceAtRiskCount == 0 {
		m.LowImportanceAtRiskCount = 0
	} else {
		m.LowImportanceAtRiskCount = lowImportanceAtRiskCount
	}
	if lowImportanceFailingCount == 0 {
		m.LowImportanceFailingCount = 0
	} else {
		m.LowImportanceFailingCount = lowImportanceFailingCount
	}
	if highImportanceAtRiskCount == 0 {
		m.HighImportanceAtRiskCount = 0
	} else {
		m.HighImportanceAtRiskCount = highImportanceAtRiskCount
	}
	if highImportanceFailingCount == 0 {
		m.HighImportanceFailingCount = 0
	} else {
		m.HighImportanceFailingCount = highImportanceFailingCount
	}
}

func (m *Monitor) UpdateDriftSetting(monitorRange string, driftMetricType string, driftThreshold float32, importanceThreshold float32, lowImportanceAtRiskCount int, lowImportanceFailingCount int, highImportanceAtRiskCount int, highImportanceFailingCount int) {
	m.MonitorRange = monitorRange

	m.DriftMetricType = driftMetricType

	m.DriftThreshold = driftThreshold

	m.ImportanceThreshold = importanceThreshold

	m.LowImportanceAtRiskCount = lowImportanceAtRiskCount

	m.LowImportanceFailingCount = lowImportanceFailingCount

	m.HighImportanceAtRiskCount = highImportanceAtRiskCount

	m.HighImportanceFailingCount = highImportanceFailingCount

}

func (m *Monitor) SetAccuracySetting(metricType string, measurement string, atRiskValue float32, failingValue float32, targetType string) {
	if metricType == "" {
		if targetType == "Regression" {
			m.MetricType = "rmse"
		} else if targetType == "Binary" {
			m.MetricType = "f1"
		}
	} else {
		m.MetricType = metricType
	}
	if measurement == "" {
		m.Measurement = "percent"
	} else {
		m.Measurement = strings.ToLower(measurement)
	}
	if atRiskValue == 0 {
		m.AtRiskValue = 0
	} else {
		m.AtRiskValue = atRiskValue
	}
	if failingValue == 0 {
		m.FailingValue = 0
	} else {
		m.FailingValue = failingValue
	}

}

func (m *Monitor) UpdateAccuracySetting(metricType string, measurement string, atRiskValue float32, failingValue float32) {
	m.MetricType = metricType

	m.Measurement = measurement

	m.AtRiskValue = atRiskValue

	m.FailingValue = failingValue
}

func (m *Monitor) SetServiceHealthSetting() {
	var hs ServiceHealthSetting

	m.ServiceHealthSetting = hs
}

// DataDrift Func

func (m *Monitor) SetFeatureDriftTrackingOn(domSvc domSvcMonitor.IExternalDriftMonitorAdapter, reqDom domDriftSvcMonitorDTO.DataDriftCreateRequest) error {

	if m.DriftCreated == true {
		// 만들어져 있을경우엔 단순 on
		reqEnable := new(domDriftSvcMonitorDTO.DataDriftEnableRequest)
		reqEnable.InferenceName = reqDom.InferenceName
		res, err := domSvc.MonitorEnable(reqEnable)
		fmt.Printf("res : %v\n", res)
		if err != nil {
			return err
		}

		m.FeatureDriftTracking = true
		m.DriftStatus = "unknown"
		return nil
	} else {
		// 만들어져있지 않은 경우엔 생성 이후 on
		res, err := domSvc.MonitorCreate(&reqDom)
		fmt.Printf("res : %v\n", res)

		if err != nil {
			return err
		}

		m.DriftCreated = true
		m.FeatureDriftTracking = true
		m.DriftStatus = "unknown"

		return nil
	}

}

func (m *Monitor) SetFeatureDriftTrackingOff(domSvc domSvcMonitor.IExternalDriftMonitorAdapter, reqDom domDriftSvcMonitorDTO.DataDriftDeleteRequest) error {

	err := domSvc.MonitorDisable(&reqDom)
	if err != nil {
		return err
	}

	m.FeatureDriftTracking = false
	m.DriftStatus = "unknown"

	return nil
}

func (m *Monitor) PatchDataDriftSetting(domSvc domSvcMonitor.IExternalDriftMonitorAdapter, reqDom domDriftSvcMonitorDTO.DataDriftPatchRequest) error {
	//if m.FeatureDriftTracking == false {
	//	return fmt.Errorf("drift tracking is not ready")
	//}
	m.UpdateDriftSetting(
		reqDom.MonitorRange,
		"PSI",
		reqDom.DriftThreshold,
		reqDom.ImportanceThreshold,
		reqDom.LowImportanceAtRiskCount,
		reqDom.LowImportanceFailingCount,
		reqDom.HighImportanceAtRiskCount,
		reqDom.HighImportanceFailingCount,
	)
	if m.DriftCreated == true {
		res, err := domSvc.MonitorPatch(&reqDom)
		fmt.Printf("res: %v\n", res)
		if err != nil {
			return fmt.Errorf("drift setting change failed")
		}

		return err
	}

	return nil
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

func (m *Monitor) GetDetailGraph(domSvc domSvcMonitor.IExternalGraphMonitorAdapter, reqDom domGraphSvcMonitorDTO.DetailGraphGetRequest) (*domGraphSvcMonitorDTO.DetailGraphGetResponse, error) {
	res, err := domSvc.MonitorGetDetailGraph(&reqDom)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Monitor) GetDriftGraph(domSvc domSvcMonitor.IExternalGraphMonitorAdapter, reqDom domGraphSvcMonitorDTO.DriftGraphGetRequest) (*domGraphSvcMonitorDTO.DriftGraphGetResponse, error) {
	res, err := domSvc.MonitorGetDriftGraph(&reqDom)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Accuracy Func

func (m *Monitor) SetAccuracyMonitoringOn(domSvc domSvcMonitor.IExternalAccuracyMonitorAdapter, reqDom domAccuracySvcMonitorDTO.AccuracyCreateRequest) error {

	if m.AccuracyCreated == true {
		// 만들어져 있을경우엔 단순 on
		reqEnable := new(domAccuracySvcMonitorDTO.AccuracyEnableRequest)
		reqEnable.InferenceName = reqDom.InferenceName
		res, err := domSvc.MonitorEnable(reqEnable)
		fmt.Printf("res : %v\n", res)
		if err != nil {
			return err
		}

		m.AccuracyMonitoring = true
		m.AccuracyStatus = "unknown"

		return nil
	} else {
		// 만들어지지 않은 경우에는 생성 이후 on
		res, err := domSvc.MonitorCreate(&reqDom)
		fmt.Printf("res: %v\n", res)

		if err != nil {
			return err
		}
		m.AssociationID = reqDom.AssociationID
		m.AccuracyCreated = true
		m.AccuracyMonitoring = true
		m.AccuracyStatus = "unknown"

		return nil
	}

}

func (m *Monitor) SetAccuracyMonitoringOff(domSvc domSvcMonitor.IExternalAccuracyMonitorAdapter, reqDom domAccuracySvcMonitorDTO.AccuracyDeleteRequest) error {

	err := domSvc.MonitorDisable(&reqDom)
	if err != nil {
		return err
	}

	m.AccuracyMonitoring = false
	m.AccuracyStatus = "unknown"

	return nil
}

func (m *Monitor) PatchAccuracySetting(domSvc domSvcMonitor.IExternalAccuracyMonitorAdapter, reqDom domAccuracySvcMonitorDTO.AccuracyPatchRequest) error {
	//if m.AccuracyMonitoring == false {
	//	return fmt.Errorf("accuracy monitoring is not ready")
	//}
	m.UpdateAccuracySetting(
		reqDom.DriftMetrics,
		reqDom.DriftMeasurement,
		reqDom.AtriskValue,
		reqDom.FailingValue,
	)
	if m.AccuracyCreated == true {
		res, err := domSvc.MonitorPatch(&reqDom)
		fmt.Printf("res: %v\n", res)

		if err != nil {
			return fmt.Errorf("accuracy setting change failed")
		}
		m.AccuracyStatus = "unknown"
		return err
	}

	return nil

}

func (m *Monitor) SetAssociationID(domSvc domSvcMonitor.IExternalAccuracyMonitorAdapter, reqDom domAccuracySvcMonitorDTO.AccuracyUpdateAssociationIDRequest) error {
	m.AssociationID = reqDom.AssociationID

	if m.AccuracyCreated == true {
		res, err := domSvc.MonitorAssociationIDPatch(&reqDom)
		fmt.Printf("res: %v\n", res)
		if err != nil {
			return fmt.Errorf("association ID change failed")
		}
		return err
	}
	return nil
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

func (m *Monitor) CheckDriftStatus(status string) bool {
	if m.DriftStatus != status {
		m.DriftStatus = status
		return true
	} else {
		return false
	}
}

func (m *Monitor) SetDriftStatusPass() {
	m.DriftStatus = "pass"
}

func (m *Monitor) SetDriftStatusAtRisk() {
	m.DriftStatus = "atrisk"
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

func (m *Monitor) SetAccuracyStatusAtRisk() {
	m.DriftStatus = "atrisk"
}

func (m *Monitor) SetAccuracyStatusFailing() {
	m.DriftStatus = "failing"
}

func (m *Monitor) SetAccuracyStatusUnknown() {
	m.DriftStatus = "unknown"
}

func (m *Monitor) SetDriftCreatedTrue() {
	m.DriftCreated = true
}

func (m *Monitor) SetDriftCreatedFalse() {
	m.DriftCreated = false
}

func (m *Monitor) SetAccuracyCreatedTrue() {
	m.AccuracyCreated = true
}

func (m *Monitor) SetAccuracyCreatedFalse() {
	m.AccuracyCreated = false
}
