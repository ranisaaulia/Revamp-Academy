package hrService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EmployeeService struct {
	employeeRepository *hrRepository.EmployeeRepository
}

func NewEmployeeService(employeeRepository *hrRepository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepository: employeeRepository,
	}
}

func (es EmployeeService) GetListEmployee(ctx *gin.Context) ([]*models.HrEmployee, *models.ResponseError) {
	return es.employeeRepository.GetListEmployee(ctx)
}

func (es EmployeeService) GetEmployee(ctx *gin.Context, id int64) (*models.HrEmployee, *models.ResponseError) {
	return es.employeeRepository.GetEmployee(ctx, id)
}

func (es EmployeeService) CreateEmployee(ctx *gin.Context, employeeParams *dbContext.CreateEmployeeParams) (*models.HrEmployee, *models.ResponseError) {
	responseErr := validateEmployee(employeeParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return es.employeeRepository.CreateEmployee(ctx, employeeParams)
}

func (es EmployeeService) UpdateEmployee(ctx *gin.Context, employeeParams *dbContext.CreateEmployeeParams, id int64) *models.ResponseError {
	responseErr := validateEmployee(employeeParams)
	if responseErr != nil {
		return responseErr
	}

	return es.employeeRepository.UpdateEmployee(ctx, employeeParams)
}

func (es EmployeeService) DeleteEmployee(ctx *gin.Context, id int64) *models.ResponseError {
	return es.employeeRepository.DeleteEmployee(ctx, id)
}

func validateEmployee(employeeParams *dbContext.CreateEmployeeParams) *models.ResponseError {
	if employeeParams.EmpEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid EmpEntityID",
			Status:  http.StatusBadRequest,
		}
	}

	if employeeParams.EmpEmpNumber == "" {
		return &models.ResponseError{
			Message: "Emp Number Required",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
