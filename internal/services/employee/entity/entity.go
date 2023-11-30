package entity

type LeaveSubmission struct {
	ID   int
	Date string
}

const MAX_SUBMISSION_DAY_BEFORE_LEAVE_DATE = 3
const MAX_LEAVE_PER_SUBMISSION = 5
