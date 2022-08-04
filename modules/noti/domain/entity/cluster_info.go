package entity

import validation "github.com/go-ozzo/ozzo-validation/v4"

type ClusterInfo struct {
	ID               string           `gorm:"size:256"`
	Name             string           `gorm:"size:256"`
	InferenceSvcInfo InferenceSvcInfo `gorm:"size:256"`
	BaseEntity
}

type InferenceSvcInfo struct {
	InfereceSvcAPISvrEndPoint   string `gorm:"size:256"`
	InfereceSvcHostName         string `gorm:"size:256"`
	InferenceSvcIngressEndPoint string `gorm:"size:256"`
}

// Validate
func (d *ClusterInfo) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Name, validation.Length(0, 255)),
		validation.Field(&d.InferenceSvcInfo.InfereceSvcAPISvrEndPoint, validation.Length(0, 255)),
		validation.Field(&d.InferenceSvcInfo.InfereceSvcHostName, validation.Length(0, 255)),
		validation.Field(&d.InferenceSvcInfo.InferenceSvcIngressEndPoint, validation.Length(0, 255)),
	)
}

func NewClusterInfo(id string, name string, infereceSvcAPISvrEndPoint string, infereceSvcHostName string, inferenceSvcIngressEndPoint string, userName string) (*ClusterInfo, error) {

	var baseEntity BaseEntity
	baseEntity.CreatedBy = userName

	if name == "" {
		name = "default-name"
	}

	ClusterInfo := &ClusterInfo{
		ID:   id,
		Name: name,
		InferenceSvcInfo: InferenceSvcInfo{
			InfereceSvcAPISvrEndPoint:   infereceSvcAPISvrEndPoint,
			InfereceSvcHostName:         infereceSvcHostName,
			InferenceSvcIngressEndPoint: inferenceSvcIngressEndPoint,
		},
		BaseEntity: baseEntity,
	}

	// Validate
	err := ClusterInfo.Validate()
	if err != nil {
		return nil, err
	}

	return ClusterInfo, nil
}

func (d *ClusterInfo) SetName(req string) error {
	d.Name = req

	err := d.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (d *ClusterInfo) SetInfereceSvcAPISvrEndPoint(req string) error {
	d.InferenceSvcInfo.InfereceSvcAPISvrEndPoint = req

	err := d.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (d *ClusterInfo) SetInfereceSvcHostName(req string) error {
	d.InferenceSvcInfo.InfereceSvcHostName = req

	err := d.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (d *ClusterInfo) SetInferenceSvcIngressEndPoint(req string) error {
	d.InferenceSvcInfo.InferenceSvcIngressEndPoint = req

	err := d.Validate()
	if err != nil {
		return err
	}
	return nil
}
