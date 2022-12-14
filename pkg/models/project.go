package models

import (
	"errors"
	"fmt"
	"profiler/pkg/constants"
	"time"
)

type Project struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Deleted     bool      `json:"deleted"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CompanyID   uint      `json:"company" gorm:"constraint:OnUpdate:CASCADE"`
}

func (p *Project) ToString() string {
	me := fmt.Sprintf("%v with id: %d", p.Name, p.ID)
	return me
}

func (p *Project) Validate() error {
	if p.Name == "" {
		return fmt.Errorf(constants.EmptyFieldErrorTmp, "Name")
	}
	if len(p.Name) <= 1 || len(p.Name) > 150 {
		return errors.New("name size should be between 2 and 150")
	}
	if p.Description == "" {
		return fmt.Errorf(constants.EmptyFieldErrorTmp, "Description")
	}
	minSize := 10
	maxSize := 300
	if len(p.Description) < minSize || len(p.Description) > maxSize {
		return fmt.Errorf(constants.OutOfSizeValueErrorTmp, "Description", minSize, maxSize)
	}

	return nil
}

func (p *Project) Me() *Project {
	return p
}

func NewProject() *Project {
	return &Project{}
}
