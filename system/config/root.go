package config

// Config represent Config
type Config struct {
	Applications Applications `json:"applications" yaml:"applications"`
	IAM          IAM          `json:"IAM" yaml:"IAM"`
	Securities   Securities   `json:"securities" yaml:"securities"`
	DirLocations DirLocations `json:"dirLocations" yaml:"dirLocations"`
	Databases    Databases    `json:"databases" yaml:"databases"`
	Connectors   Connectors   `json:"connectors" yaml:"connectors"`
	SMTPServers  SMTPServers  `json:"SMTPServers" yaml:"SMTPServers"`
	Caches       Caches       `json:"caches" yaml:"caches"`
	Indexers     Indexers     `json:"indexers" yaml:"indexers"`
}

// DirLocations represent DirLocations
type DirLocations struct {
	Conf                     string `json:"conf" yaml:"conf"`
	WWW                      string `json:"www" yaml:"www"`
	Temp                     string `json:"temp" yaml:"temp"`
	Log                      string `json:"log" yaml:"log"`
	Cache                    string `json:"cache" yaml:"cache"`
	ModelPackageFileRootPath string `json:"modelPackageFileRootPath" yaml:"modelPackageFileRootPath"`
	ModelPath                string `json:"modelPath" yaml:"modelPath"`
	TrainingDatasetPath      string `json:"trainingDatasetPath" yaml:"tainingDatasetPath"`
	HoldoutDatasetPath       string `json:"holdoutDatasetPath" yaml:"holdoutDatasetPath"`
	MonitoringFileRootPath   string `json:"monitoringFileRootPath" yaml:"monitoringFileRootPath"`
	ActualDataPath           string `json:"actualDataPath" yaml:"actualDataPath"`
}
