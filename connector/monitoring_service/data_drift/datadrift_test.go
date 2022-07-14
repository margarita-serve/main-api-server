package data_drift

import (
	monType "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/data_drift/types"
	"testing"
)

func newDataDriftMonitor(t *testing.T) *DriftMonitor {
	config := Config{Server: "http://192.168.88.151:30071"}
	return NewDriftMonitor(config, nil)
}

func TestDriftMonitor_CreateDriftMonitor(t *testing.T) {
	c := newDataDriftMonitor(t)
	resp, err := c.CreateDriftMonitor(
		&monType.CreateDataDriftRequest{
			TrainDatasetPath:    "dataset/train_data.csv",
			ModelPath:           "testmodel/mpg2/1",
			InferenceName:       "mpg-sample",
			ModelID:             "000001",
			TargetLabel:         "MPG",
			ModelType:           "Regression",
			Framework:           "tensorflow",
			DriftThreshold:      0.15,
			ImportanceThreshold: 0.5,
			MonitorRange:        "7d",
			LowImpAtRiskCount:   1,
			LowImpFailingCount:  0,
			HighImpAtRiskCount:  0,
			HighImpFailingCount: 1,
		},
	)
	if err != nil {
		t.Logf("TestERROR: %s", err)
	}

	if resp != nil {
		t.Logf("RESPONSE.Message: %#v", resp.Message)
	}
}
