package infra

import (
	"context"
	"go-api-echo/internal/pkg/helpers/errors"
	"go-api-echo/internal/services/employee/adapter"
	"go-api-echo/internal/services/employee/entity"

	"github.com/jmoiron/sqlx"
)

type EmployeeRepo struct {
	ctx context.Context
	db  *sqlx.DB
	tx  *sqlx.Tx
}

func (r *EmployeeRepo) BeginTransaction() *errors.Error {
	tx, err := r.db.Beginx()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error starting transaction`)

		return &sqlErr
	}

	r.tx = tx

	return nil
}

func (r EmployeeRepo) CommitTransaction() *errors.Error {
	err := r.tx.Commit()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error committing transaction`)

		return &sqlErr
	}

	return nil
}

func (r EmployeeRepo) RollbackTransaction() {
	_ = r.tx.Rollback()
}

func (r EmployeeRepo) ListEmployeeTx(req *adapter.ListEmployeeReq) (*[]adapter.Employee, *errors.Error) {
	tx, err := r.db.Beginx()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error starting transaction`)
		return nil, &sqlErr
	}

	employees := make([]adapter.Employee, 0)
	err = tx.SelectContext(
		r.ctx,
		&employees,
		`
		SELECT
			e.id,
			e.fullname,
			e.leaveQuota,
			COUNT(el.id) AS onLeave
		FROM employees e
		LEFT JOIN employee_leave_submissions el ON e.id = el.id AND el.leaveDate = datetime('now','localtime')
		WHERE e.fullname LIKE ?
		GROUP BY e.id
		ORDER BY e.fullname ASC
		LIMIT ? OFFSET ?
		`,
		`%`+req.Search+`%`,
		req.Pagination.Limit,
		req.Pagination.Offset,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error querying employee list`)
		_ = tx.Rollback()

		return &employees, &sqlErr
	}

	var total int
	err = tx.GetContext(
		r.ctx,
		&total,
		`
		SELECT
			COUNT(*)
		FROM employees e
		WHERE e.fullname LIKE ?
		`,
		`%`+req.Search+`%`,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error querying total employee`)
		_ = tx.Rollback()

		return &employees, &sqlErr
	}

	req.Pagination.Total = total

	err = tx.Commit()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error committing transaction`)
		_ = tx.Rollback()

		return &employees, &sqlErr
	}

	return &employees, nil
}

func (r EmployeeRepo) DetailEmployee(id int) (*adapter.Employee, *errors.Error) {
	employee := adapter.Employee{}
	err := r.db.GetContext(
		r.ctx,
		&employee,
		`
		SELECT
			e.id,
			e.fullname,
			e.leaveQuota,
			COUNT(el.id) AS onLeave
		FROM employees e
		LEFT JOIN employee_leave_submissions el ON e.id = el.id AND el.leaveDate = datetime('now','localtime')
		WHERE e.id = ?
		GROUP BY e.id
		`,
		id,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error querying employee detail`)

		return &employee, &sqlErr
	}

	return &employee, nil
}

func (r EmployeeRepo) InsertEmployee(req adapter.UpsertEmployeeReq) (int, *errors.Error) {
	res, err := r.db.NamedExecContext(
		r.ctx,
		`
		INSERT INTO employees
		(
			fullname,
			leaveQuota
		) VALUES (
			:fullname,
			:leavequota
		)
		`,
		req,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error adding new employee`)

		return 0, &sqlErr
	}

	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error getting new employee id`)

		return 0, &sqlErr
	}

	return int(lastInsertedId), nil
}

func (r EmployeeRepo) UpdateEmployee(req adapter.UpsertEmployeeReq) (int, *errors.Error) {
	res, err := r.db.NamedExecContext(
		r.ctx,
		`
		UPDATE employees SET
			fullname = :fullname,
			leaveQuota = :leavequota
		WHERE id = :id
		`,
		req,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error updating employee`)

		return 0, &sqlErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error getting affected rows after updating employee`)

		return 0, &sqlErr
	}

	return int(rowsAffected), nil
}

func (r EmployeeRepo) DeleteEmployeeTx(id int) (int, *errors.Error) {
	res, err := r.db.ExecContext(
		r.ctx,
		`
		DELETE FROM employees
		WHERE id = ?
		`,
		id,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error deleting employee`)

		return 0, &sqlErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error getting affected rows after deleting employee`)

		return 0, &sqlErr
	}

	return int(rowsAffected), nil
}

func (r EmployeeRepo) ListEmployeeLeave(id int) (*[]string, *errors.Error) {
	leaveDates := make([]string, 0)
	err := r.db.SelectContext(
		r.ctx,
		&leaveDates,
		`
		SELECT
			e.leaveDate
		FROM employee_leave_submissions e
		WHERE e.employeeId = ?
		ORDER BY e.leaveDate ASC
		`,
		id,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error querying employee list`)

		return &leaveDates, &sqlErr
	}

	return &leaveDates, nil
}

func (r EmployeeRepo) SubmitEmployeeLeave(req *[]entity.LeaveSubmission) (int, error) {
	res, err := r.db.NamedExecContext(
		r.ctx,
		`
		INSERT INTO employee_leave_submissions
		(
			employeeId,
			leaveDate
		) VALUES (
			:id,
			:date
		)
		`,
		*req,
	)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertedId), nil
}
