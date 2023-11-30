package usecase

import (
	"fmt"
	"go-api-echo/internal/pkg/utils"
	"go-api-echo/internal/services/employee/entity"
	"time"
)

func ProcessLeaveSubmission(
	req *[]entity.LeaveSubmission,
	submit func(req *[]entity.LeaveSubmission) (int, error),
) (res int, err error) {
	if err = checkSubmissionDayUntilLeaveDate((*req)[0].Date); err != nil {
		return 0, err
	}
	if err = checkSubmittedLeave(len(*req)); err != nil {
		return 0, err
	}

	res, err = submit(req)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func checkSubmissionDayUntilLeaveDate(d string) error {
	date, parseErr := time.Parse(utils.FULLDATE_FORMAT_LAYOUT, d)
	if parseErr != nil {
		return fmt.Errorf("error parsing date while checking day before leave submission")
	}

	duration := int(time.Until(date).Hours() / 24)

	if duration < entity.MAX_SUBMISSION_DAY_BEFORE_LEAVE_DATE {
		return fmt.Errorf(
			"can't submit leave. submission must be made %d days before leave date",
			entity.MAX_SUBMISSION_DAY_BEFORE_LEAVE_DATE,
		)
	}

	return nil
}

func checkSubmittedLeave(t int) error {
	if t > entity.MAX_LEAVE_PER_SUBMISSION {
		return fmt.Errorf(
			"can't submit leave. maximum %d days per submission",
			entity.MAX_LEAVE_PER_SUBMISSION,
		)
	}

	return nil
}
