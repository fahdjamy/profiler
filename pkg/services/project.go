package services

import (
	"errors"
	"fmt"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/storage/database"
	"gopher-camp/pkg/types"
	"gopher-camp/pkg/types/dto"
	"gorm.io/gorm"
)

type ProjectService struct {
	db        *gorm.DB
	logger    types.Logger
	coService types.Storage[models.Company]
}

func (p ProjectService) FindAll() []models.Project {
	var projects []models.Project
	p.db.Preload("Companies").Find(&projects)

	return projects
}

func (p ProjectService) Delete(id uint) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProjectService) FindById(id uint) (*models.Project, error) {
	project := &models.Project{}
	rec := p.db.Where("id = ?", id).Limit(1).Find(project)

	if rec.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("project with id (%v) does not exist", id))
	}

	return project, nil
}

func (p ProjectService) Create(newProject *models.Project) (*models.Project, error) {
	project := models.NewProject()

	companyResponse, err := p.coService.FindById(newProject.CompanyID)
	if err != nil {
		return nil, err
	}

	project.CompanyID = companyResponse.ID

	p.db.Create(project)
	p.db.Preload("Companies").Find(project)

	return newProject, nil
}

func (p ProjectService) Update(id uint, project *models.Project) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func NewProjectService(db database.Database, logger types.Logger, coService types.Storage[models.Company]) ProjectService {
	return ProjectService{
		db:        db.GetDB(),
		logger:    logger,
		coService: coService,
	}
}

func convertProjectDTOToProject(projectDTO types.DTOMapper[models.Project, dto.ProjectReqDTO], project *models.Project) error {
	projectDTO.MapToDO(project)

	return nil
}
