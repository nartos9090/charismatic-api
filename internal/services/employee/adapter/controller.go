package adapter

import (
	response "go-api-echo/internal/pkg/helpers/response"
	"net/http"
)

func HandleListEmployee(req ListEmployeeReq, repo EmployeeRepoInterface) (resp response.HttpRes) {
	employees, err := repo.ListEmployeeTx(&req)
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Data = employees
	resp.Status = http.StatusOK
	resp.Message = `data found`
	resp.Pagination = &req.Pagination

	return
}

func HandleDetailemployee(id int, repo EmployeeRepoInterface) (resp response.HttpRes) {
	employee, err := repo.DetailEmployee(id)
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Data = employee
	resp.Status = http.StatusOK
	resp.Message = `data found`

	return
}

func HandleInsertEmployee(req UpsertEmployeeReq, repo EmployeeRepoInterface) (resp response.HttpRes) {
	lastInsertedId, err := repo.InsertEmployee(req)
	if err != nil {
		return err.ToHttpRes()
	}

	if lastInsertedId <= 0 {
		resp.Data = 0
		resp.Status = http.StatusInternalServerError
		resp.Message = `no data added`
	}

	resp.Data = lastInsertedId
	resp.Status = http.StatusOK
	resp.Message = `data added`

	return
}

func HandleUpdateEmployee(req UpsertEmployeeReq, repo EmployeeRepoInterface) (resp response.HttpRes) {
	rowsAffected, err := repo.UpdateEmployee(req)
	if err != nil {
		return err.ToHttpRes()
	}

	if rowsAffected <= 0 {
		resp.Data = 0
		resp.Status = http.StatusInternalServerError
		resp.Message = `no data updated`
	}

	resp.Data = rowsAffected
	resp.Status = http.StatusOK
	resp.Message = `data updated`

	return
}

func HandleDeleteEmployee(id int, repo EmployeeRepoInterface) (resp response.HttpRes) {
	rowsAffected, err := repo.DeleteEmployeeTx(id)
	if err != nil {
		return err.ToHttpRes()
	}

	if rowsAffected <= 0 {
		resp.Data = 0
		resp.Status = http.StatusInternalServerError
		resp.Message = `no data deleted`
	}

	resp.Data = rowsAffected
	resp.Status = http.StatusOK
	resp.Message = `data deleted`

	return
}
