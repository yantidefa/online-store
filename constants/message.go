package constants

var (
	DateFromStd = "2006-01-02 15:04:05"
)

var (
	ErrMissingLoginValues   = "missing Username or Password"
	ErrFailedAuthentication = "incorrect Username or Password"
	ErrExpiredToken         = "token is expired"
	ErrExistEmail           = "Email Already Registered"
	ErrConfirmPassword      = "Password is not the same"
	ErrEmptyAuthHeader      = "auth header is empty"
	ErrInvalidToken         = "Invalid Token"
	ErrLogout               = "Account Has Been Logged Out"
	ErrActivation           = "Account Has Been Activation"
	ErrEmail                = "Email Doesn't Exist"
	ErrPassword             = "Incorrect Password"
	ErrLogin                = "The Credentials Already Login"
)

var (
	DataFound            = "Data Found"
	SuccessAddData       = "Successfully Added Data"
	SuccessRegister      = "Account Successfully To Register"
	SuccessUpdateData    = "Successfully Updated Data"
	SuccessDeleteData    = "Successfully Deleted Data"
	SuccessDownload      = "Successfully Download"
	SuccessDisplayedData = "Successfully Displayed The Data"
	SuccessLogin         = "Successfully Login"
)

var (
	FailedRegister             = "Account Failed To Register"
	ErrDataNotFound            = "Data Not Found"
	ErrInternalServer          = "Internal Server Error"
	DataIsAvailable            = "Data Is Available"
	ContentNotFound            = "Content Not Found"
	ModuleNotFound             = "Module Not Found"
	FailedAddData              = "Failed To Added Data"
	FailedUpdateData           = "Failed To Change Data"
	FailedDeleteData           = "Failed deleted Data"
	FailedDownload             = "Failed Download"
	FailedDisplayedData        = "Failed Displayed The Data"
	FailedDisplayedDataDeleted = "Data Has Been Deleted"
	InvalidJsonRequest         = "Invalid Json Request"
)
