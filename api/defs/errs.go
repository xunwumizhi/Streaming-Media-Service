package defs

type ErrResponse struct {
	HttpSC int
	Error MyErr
}

type MyErr struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{
		HttpSC: 400,
		Error: MyErr{Error: "Request body is not correct.", ErrorCode: "001"},
	}
	ErrorNotAuthUser = ErrResponse{
		HttpSC: 401,
		Error: MyErr{Error: "User anthentication failed.", ErrorCode: "002"},
	}
	ErrorDBError = ErrResponse{
		HttpSC: 500,
		Error: MyErr{Error: "DB ops failed.", ErrorCode: "003"},
	}
	ErrorInternalFaults = ErrResponse{
		HttpSC: 500,
		Error: MyErr{Error: "Internal service error.", ErrorCode: "004"},
	}
)