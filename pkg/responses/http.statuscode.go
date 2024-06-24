package response

const (
	StatusSuccess      = 200
	StatusParamInvalid = 400
)

var msg = map[int]string{
	StatusSuccess:      "Success",
	StatusParamInvalid: "Fail",
}
