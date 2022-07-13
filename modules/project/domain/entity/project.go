package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Project struct {
	ID          string `gorm:"size:256"`
	Name        string `gorm:"size:256"`
	Description string `gorm:"size:256"`
	BaseEntity
}

// Validate
func (d *Project) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.ID, validation.Required),
		validation.Field(&d.Name, validation.Length(0, 255)),
		validation.Field(&d.Description, validation.Length(0, 255)),
	)
}

func NewProject(id string, name string, description string, ownerID string) (*Project, error) {

	var baseEntity BaseEntity
	baseEntity.CreatedBy = ownerID

	if name == "" {
		name = "default-name"
	}

	project := &Project{
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

func (d *Project) SetName(req string) error {
	d.Name = req

	err := d.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (d *Project) SetDescription(req string) error {
	d.Description = req

	err := d.Validate()
	if err != nil {
		return err
	}
	return nil
}
