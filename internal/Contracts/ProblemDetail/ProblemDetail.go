package ProblemDetail

type problemDetail struct {
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
	Status   uint   `json:"status"`
	Type     string `json:"type"`
}

func New(
	detail string,
	instance string,
	status uint,
	errtype string) problemDetail {

	return problemDetail{
		Detail:   detail,
		Instance: instance,
		Status:   status,
		Type:     errtype,
	}
}
