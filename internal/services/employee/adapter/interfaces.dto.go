package adapter

import errors "go-api-echo/internal/pkg/helpers/errors"

type EmployeeRepoInterface interface {
	BeginTransaction() *errors.Error
	CommitTransaction() *errors.Error
	RollbackTransaction()

	ListEmployeeTx(req *ListEmployeeReq) (*[]Employee, *errors.Error)
	DetailEmployee(id int) (*Employee, *errors.Error)
	InsertEmployee(req UpsertEmployeeReq) (int, *errors.Error)
	UpdateEmployee(req UpsertEmployeeReq) (int, *errors.Error)
	DeleteEmployeeTx(id int) (int, *errors.Error)
}
