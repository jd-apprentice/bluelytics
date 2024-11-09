package constants

const (
	EnvNotFound      = "Env file not found"
	FailedFetch      = "Failed to fetch data"
	FailtedSecrets   = "Failed to get secrets"
	BaseErrorMessage = "error: "
)

const (
	SendCause      = "sending request"
	CreateCause    = "creating request"
	ReadCause      = "reading response body"
	UnmarshalCause = "unmarshalling response"
)

const (
	SendError      = BaseErrorMessage + SendCause
	CreateError    = BaseErrorMessage + CreateCause
	ReadError      = BaseErrorMessage + ReadCause
	UnmarshalError = BaseErrorMessage + UnmarshalCause
)
