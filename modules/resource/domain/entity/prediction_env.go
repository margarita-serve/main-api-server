package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PredictionEnv struct {
	ID            string   `gorm:"size:256"`
	ClusterInfoID string   `gorm:"size:256"`
	Name          string   `gorm:"size:256"`
	Description   string   `gorm:"size:256"`
	Namespace     string   `gorm:"size:256"`
	Projects      []string `gorm:"size:256"`
	BaseEntity
}

// Validate
func (d *PredictionEnv) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.ID, validation.Required),
		validation.Field(&d.Name, validation.Length(0, 255)),
		validation.Field(&d.Description, validation.Length(0, 255)),
	)
}

func NewPredictionEnv(id string, name string, description string, ownerID string) (*PredictionEnv, error) {

	var baseEntity BaseEntity
	baseEntity.CreatedBy = ownerID

	if name == "" {
		name = "default-name"
	}

	project := &PredictionEnv{
		ID:          id,
		Name:        name,
		Description: description,
		BaseEntity:  baseEntity,
	}

	// Validate
	err := project.Validate()
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (d *PredictionEnv) SetName(req string) error {
	d.Name = req

	err := d.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (d *PredictionEnv) SetDescription(req string) error {
	d.Description = req

	err := d.Validate()
	if err != nil {
		return err
	}
	return nil
}
