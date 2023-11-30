package adapter

import (
	"go-api-echo/internal/pkg/helpers/errors"
	response "go-api-echo/internal/pkg/helpers/response"
	"go-api-echo/internal/services/employee/usecase"
	"net/http"
)

func HandleListEmployee(req ListEmployeeReq, repo EmployeeRepoInterface) (resp response.HttpRes) {
	employees, err := repo.ListEmployeeTx(&req)
	if err != nil {
		return err.ToHttpRes()
	}
	if len(*employees) == 0 {
		resp.Status = http.StatusNotFound
		resp.Message = "data not found"

		goto set_data
	}

	resp.Status = http.StatusOK
	resp.Message = "data found"
set_data:
	resp.Data = employees
	resp.Pagination = &req.Pagination

	return
}

func HandleDetailEmployee(id int, repo EmployeeRepoInterface) (resp response.HttpRes) {
	employee, err := repo.DetailEmployee(id)
	if err != nil {
		return err.ToHttpRes()
	}
	if employee == nil {
		resp.Status = http.StatusNotFound
		resp.Message = "data not found"

		goto set_data
	}

	resp.Status = http.StatusOK
	resp.Message = "data found"
set_data:
	resp.Data = employee

	return
}

func HandleInsertEmployee(req UpsertEmployeeReq, repo EmployeeRepoInterface) (resp response.HttpRes) {
	lastInsertedId, err := repo.InsertEmployee(req)
	if err != nil {
		return err.ToHttpRes()
	}

	if lastInsertedId <= 0 {
		resp.Status = http.StatusInternalServerError
		resp.Message = "no data added"
		resp.Data = 0

		return
	}

	resp.Status = http.StatusOK
	resp.Message = "data added"
	resp.Data = lastInsertedId

	return
}

func HandleUpdateEmployee(req UpsertEmployeeReq, repo EmployeeRepoInterface) (resp response.HttpRes) {
	rowsAffected, err := repo.UpdateEmployee(req)
	if err != nil {
		return err.ToHttpRes()
	}

	if rowsAffected <= 0 {
		resp.Status = http.StatusInternalServerError
		resp.Message = "no data updated"
		resp.Data = 0

		return
	}

	resp.Status = http.StatusOK
	resp.Message = "data updated"
	resp.Data = rowsAffected

	return
}

func HandleDeleteEmployee(id int, repo EmployeeRepoInterface) (resp response.HttpRes) {
	rowsAffected, err := repo.DeleteEmployeeTx(id)
	if err != nil {
		return err.ToHttpRes()
	}

	if rowsAffected <= 0 {
		resp.Status = http.StatusInternalServerError
		resp.Message = "no data deleted"
		resp.Data = 0

		return
	}

	resp.Status = http.StatusOK
	resp.Message = "data deleted"
	resp.Data = rowsAffected

	return
}

func HandleListEmployeeLeave(id int, repo LeaveSubmissionRepoInterface) (resp response.HttpRes) {
	leaveDates, err := repo.ListEmployeeLeave(id)
	if err != nil {
		return err.ToHttpRes()
	}
	if len(*leaveDates) == 0 {
		resp.Status = http.StatusNotFound
		resp.Message = "data not found"

		goto set_data
	}

	resp.Status = http.StatusOK
	resp.Message = "data found"
set_data:
	resp.Data = leaveDates

	return
}

func HandleSubmitLeave(rawReq LeaveSubmissionReq, repo LeaveSubmissionRepoInterface) (resp response.HttpRes) {
	if len(rawReq.Dates) == 0 {
		resp.Status = http.StatusBadRequest
		resp.Message = "select leave dates first"
		resp.Data = 0

		return
	}

	req := TransformLeaveSubmissionRequest(rawReq)
	res, rawErr := usecase.ProcessLeaveSubmission(req, repo.SubmitEmployeeLeave)
	if rawErr != nil {
		err := errors.InternalServerError
		err.AddError("error while processing leave submission")
		err.AddError(rawErr.Error())

		return err.ToHttpRes()
	}

	resp.Status = http.StatusOK
	resp.Message = "data submitted"
	resp.Data = res

	return
}
