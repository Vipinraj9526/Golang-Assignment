package repositories

import (
	"context"
	"data-manage/models"
	"errors"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, db *gorm.DB, employees *models.Employee) error
	UpdateEmployee(ctx context.Context, db *gorm.DB, record interface{}, conditions map[string]interface{}, updates interface{}) error
	ReadEmployeeWithCondition(ctx context.Context, db *gorm.DB, employees models.Employee, condition map[string]interface{}) (*models.Employee, error)
	DeleteEmployee(ctx context.Context, db *gorm.DB, employees *models.Employee, condition map[string]interface{}) error
}

type employeeRepository struct{}

func NewEmployeesRepository() *employeeRepository {
	return &employeeRepository{}
}

func EmployeeRepositoryInstance() EmployeeRepository {
	return NewEmployeesRepository()
}

func (employeeRepository *employeeRepository) CreateEmployee(ctx context.Context, db *gorm.DB, employees *models.Employee) error {
	result := db.WithContext(ctx).Create(employees)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (employeeRepository *employeeRepository) UpdateEmployee(ctx context.Context, db *gorm.DB, record interface{}, conditions map[string]interface{}, updates interface{}) error {
	updateResponse := db.WithContext(ctx).Model(record).Where(conditions).Updates(updates)

	if updateResponse.Error != nil {
		return updateResponse.Error
	}
	if updateResponse.RowsAffected == 0 {
		return errors.New("no rows were affected")
	}
	return nil
}

func (employeeRepository *employeeRepository) ReadEmployeeWithCondition(ctx context.Context, db *gorm.DB, employees models.Employee, condition map[string]interface{}) (*models.Employee, error) {
	result := db.WithContext(ctx).Where(condition).Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows found")
	}
	return &employees, nil
}

func (employeeRepository *employeeRepository) DeleteEmployee(ctx context.Context, db *gorm.DB, employees *models.Employee, condition map[string]interface{}) error {
	result := db.WithContext(ctx).Where(condition).Delete(&employees)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
