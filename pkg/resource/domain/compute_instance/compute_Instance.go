package compute_instance

type ComputeInstanceID string

type ComputeInstance struct {
	ID               ComputeInstanceID
	Name             string
	UseType          UseType
	K8sClientIP      string
	K8sClientPort    string
	KserveClientIP   string
	KserveClientPort string
	Namespace        string
	SystemInfo       SystemInfo
}

type UseType string

const (
	PackageTest UseType = "package_test"
	Default     UseType = "default"
	Custom      UseType = "custom"
)

type SystemInfo struct {
	CPU    float32
	Memory float32
	Disk   float32
	GPU    float32
}
