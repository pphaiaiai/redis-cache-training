package errors

import "github.com/gofiber/fiber/v2"

const (
	ERROR_CODE_INVALID_USERNAME_PASSWORD        = "E001"
	ERROR_CODE_INVALID_JWT                      = "E002"
	ERROR_CODE_INVALID_TOKEN                    = "E003"
	ERROR_CODE_TOKEN_EXPIRED                    = "E004"
	ERROR_CODE_INVALID_PASSWORD                 = "E005"
	ERROR_CODE_ACCESS_DENIED                    = "E006"
	ERROR_CODE_VERIFY_EMAIL_DUPLICATE           = "E007"
	ERROR_CODE_DELETE_ACCOUNT                   = "E008"
	ERROR_CODE_USER_BLOCK                       = "E009"
	ERROR_CODE_EMAIL_NOTVERIFY                  = "E010"
	ERROR_CODE_PROCESS_USER_APPROVE             = "E011"
	ERROR_CODE_EMAIL_INUSED                     = "E012"
	ERROR_CODE_USERNAME_INUSED                  = "E013"
	ERROR_CODE_WRONG_CURRENT_PASSWORD           = "E014"
	ERROR_CODE_NEW_PASSWORD_NOT_SAME            = "E015"
	ERROR_CODE_OLD_AND_CURRENT_PASSWORD_SAME    = "E016"
	ERROR_CODE_VERIFY_EMAIL_FAIL                = "E017"
	ERROR_CODE_VERIFY_EMAIL_EXPIRE              = "E018"
	ERROR_CODE_INSERT_FAIL                      = "E019"
	ERROR_CODE_UPDATE_FAIL                      = "E020"
	ERROR_CODE_DELETE_FAIL                      = "E021"
	ERROR_CODE_EMAIL_NOT_FOUND                  = "E022"
	ERROR_CODE_DATA_NOT_FOUND                   = "E023"
	ERROR_CODE_EMAIL_IN_USE                     = "E024"
	ERROR_CODE_USER_EXPIRED                     = "E025"
	ERROR_CODE_DUPLICATE_USERNAME               = "E026"
	ERROR_CODE_PERMISSION_INCORRECT             = "E027"
	ERROR_CODE_INTERNAL_SERVER_ERROR            = "E028"
	ERROR_CODE_ACCECPT_PRIVACY_POLICY           = "E029"
	ERROR_CODE_YOUR_ACCOUNT_IS_EXPIRED          = "E030"
	ERROR_CODE_TERM_AND_CONDITION_NOT_ACCEPTED  = "E031"
	ERROR_CODE_PRIVACY_POLICY_VERSION           = "E032"
	ERROR_CODE_PRIVACY_NOTICE_VERSION           = "E033"
	ERROR_CODE_ERROR_FILE_EXCEPTION             = "E034"
	ERROR_CODE_FILE_MUST_BE_IMAGE               = "E035"
	ERROR_CODE_USER_NOT_FOUND                   = "E036"
	ERROR_CODE_URL_INCORRECT                    = "E037"
	ERROR_CODE_REGISTRATION_TOKEN_FAILED        = "E038"
	ERROR_CODE_NOTIFICATION_FAILED              = "E039"
	ERROR_CODE_FORBIDDEN                        = "E040"
	ERROR_CODE_FAIL_PASE_MULTIPART_FORM         = "E041"
	ERROR_CODE_INVALID_ID                       = "E042"
	ERROR_CODE_MISSING_REQUIRED_FORM_DATA       = "E043"
	ERROR_CODE_FAIL_PASE_REQUEST_BODY           = "E044"
	ERROR_CODE_VALIDATE_FAIL                    = "E045"
	ERROR_CODE_FAIL_KAFKA_PRODUCER              = "E046"
	ERROR_CODE_FAIL_KAFKA_CONSUMER              = "E047"
	ERROR_CODE_PAGINATION_FAIL                  = "E048"
	ERROR_CODE_QUERY_FAIL                       = "E049"
	ERROR_CODE_FAIL_COMPRESS_JSON               = "E050"
	ERROR_CODE_FAIL_KAFKA_READER_INITIALIZATION = "E051"
	ERROR_CODE_FAIL_KAFKA_READER_FETCH          = "E052"
	ERROR_CODE_FAIL_KAFKA_READER_COMMIT         = "E053"
	ERROR_CODE_CONNECTION_NOT_SUPPORTED         = "E054"
	ERROR_CODE_FAIL_FETCH_REDIS                 = "E055"
	ERROR_CODE_FAIL_SET_DATA_REDIS              = "E056"
	ERROR_CODE_FAIL_UNCOMPRESS_JSON             = "E057"
	ERROR_CODE_MISMATCHED_DOCUMENT_DATA         = "E058"
)

type Error struct {
	Code       string
	Message    string
	HTTPStatus int
}

var errorData = map[string]Error{
	"E001": {Code: ERROR_CODE_INVALID_USERNAME_PASSWORD, Message: "Invalid username or password", HTTPStatus: fiber.StatusBadRequest},
	"E002": {Code: ERROR_CODE_INVALID_JWT, Message: "Invalid JWT", HTTPStatus: fiber.StatusUnauthorized},
	"E003": {Code: ERROR_CODE_INVALID_TOKEN, Message: "Invalid token", HTTPStatus: fiber.StatusBadRequest},
	"E004": {Code: ERROR_CODE_TOKEN_EXPIRED, Message: "Token expired", HTTPStatus: fiber.StatusUnauthorized},
	"E005": {Code: ERROR_CODE_INVALID_PASSWORD, Message: "Invalid password", HTTPStatus: fiber.StatusBadRequest},
	"E006": {Code: ERROR_CODE_ACCESS_DENIED, Message: "Access denied", HTTPStatus: fiber.StatusForbidden},
	"E007": {Code: ERROR_CODE_VERIFY_EMAIL_DUPLICATE, Message: "Email already verified", HTTPStatus: fiber.StatusBadRequest},
	"E008": {Code: ERROR_CODE_DELETE_ACCOUNT, Message: "Failed to delete account", HTTPStatus: fiber.StatusInternalServerError},
	"E009": {Code: ERROR_CODE_USER_BLOCK, Message: "User is blocked", HTTPStatus: fiber.StatusForbidden},
	"E010": {Code: ERROR_CODE_EMAIL_NOTVERIFY, Message: "Email not verified", HTTPStatus: fiber.StatusForbidden},
	"E011": {Code: ERROR_CODE_PROCESS_USER_APPROVE, Message: "Failed to process user approval", HTTPStatus: fiber.StatusInternalServerError},
	"E012": {Code: ERROR_CODE_EMAIL_INUSED, Message: "Email already in use", HTTPStatus: fiber.StatusBadRequest},
	"E013": {Code: ERROR_CODE_USERNAME_INUSED, Message: "Username already in use", HTTPStatus: fiber.StatusBadRequest},
	"E014": {Code: ERROR_CODE_WRONG_CURRENT_PASSWORD, Message: "Wrong current password", HTTPStatus: fiber.StatusBadRequest},
	"E015": {Code: ERROR_CODE_NEW_PASSWORD_NOT_SAME, Message: "New password must be different from the old password", HTTPStatus: fiber.StatusBadRequest},
	"E016": {Code: ERROR_CODE_OLD_AND_CURRENT_PASSWORD_SAME, Message: "Old and new password must be different", HTTPStatus: fiber.StatusBadRequest},
	"E017": {Code: ERROR_CODE_VERIFY_EMAIL_FAIL, Message: "Failed to verify email", HTTPStatus: fiber.StatusInternalServerError},
	"E018": {Code: ERROR_CODE_VERIFY_EMAIL_EXPIRE, Message: "Verification email expired", HTTPStatus: fiber.StatusBadRequest},
	"E019": {Code: ERROR_CODE_INSERT_FAIL, Message: "Failed to insert data", HTTPStatus: fiber.StatusInternalServerError},
	"E020": {Code: ERROR_CODE_UPDATE_FAIL, Message: "Failed to update data", HTTPStatus: fiber.StatusInternalServerError},
	"E021": {Code: ERROR_CODE_DELETE_FAIL, Message: "Failed to delete data", HTTPStatus: fiber.StatusInternalServerError},
	"E022": {Code: ERROR_CODE_EMAIL_NOT_FOUND, Message: "Email not found", HTTPStatus: fiber.StatusBadRequest},
	"E023": {Code: ERROR_CODE_DATA_NOT_FOUND, Message: "Data not found", HTTPStatus: fiber.StatusNotFound},
	"E024": {Code: ERROR_CODE_EMAIL_IN_USE, Message: "Email in use", HTTPStatus: fiber.StatusBadRequest},
	"E025": {Code: ERROR_CODE_USER_EXPIRED, Message: "User expired", HTTPStatus: fiber.StatusForbidden},
	"E026": {Code: ERROR_CODE_DUPLICATE_USERNAME, Message: "Duplicate username", HTTPStatus: fiber.StatusBadRequest},
	"E027": {Code: ERROR_CODE_PERMISSION_INCORRECT, Message: "Permission incorrect", HTTPStatus: fiber.StatusForbidden},
	"E028": {Code: ERROR_CODE_INTERNAL_SERVER_ERROR, Message: "Internal server error", HTTPStatus: fiber.StatusInternalServerError},
	"E029": {Code: ERROR_CODE_ACCECPT_PRIVACY_POLICY, Message: "Please accept privacy policy", HTTPStatus: fiber.StatusBadRequest},
	"E030": {Code: ERROR_CODE_YOUR_ACCOUNT_IS_EXPIRED, Message: "Your account is expired", HTTPStatus: fiber.StatusForbidden},
	"E031": {Code: ERROR_CODE_TERM_AND_CONDITION_NOT_ACCEPTED, Message: "Please accept term and condition", HTTPStatus: fiber.StatusBadRequest},
	"E032": {Code: ERROR_CODE_PRIVACY_POLICY_VERSION, Message: "Privacy policy version is not updated", HTTPStatus: fiber.StatusBadRequest},
	"E033": {Code: ERROR_CODE_PRIVACY_NOTICE_VERSION, Message: "Privacy notice version is not updated", HTTPStatus: fiber.StatusBadRequest},
	"E034": {Code: ERROR_CODE_ERROR_FILE_EXCEPTION, Message: "Error file exception", HTTPStatus: fiber.StatusBadRequest},
	"E035": {Code: ERROR_CODE_FILE_MUST_BE_IMAGE, Message: "File must be an image", HTTPStatus: fiber.StatusBadRequest},
	"E036": {Code: ERROR_CODE_USER_NOT_FOUND, Message: "User not found", HTTPStatus: fiber.StatusBadRequest},
	"E037": {Code: ERROR_CODE_URL_INCORRECT, Message: "URL is incorrect", HTTPStatus: fiber.StatusBadRequest},
	"E038": {Code: ERROR_CODE_REGISTRATION_TOKEN_FAILED, Message: "Failed to generate registration token", HTTPStatus: fiber.StatusInternalServerError},
	"E039": {Code: ERROR_CODE_NOTIFICATION_FAILED, Message: "Failed to send notification", HTTPStatus: fiber.StatusInternalServerError},
	"E040": {Code: ERROR_CODE_FORBIDDEN, Message: "Forbidden", HTTPStatus: fiber.StatusForbidden},
	"E041": {Code: ERROR_CODE_FAIL_PASE_MULTIPART_FORM, Message: "Failed to parse multipart form", HTTPStatus: fiber.StatusBadRequest},
	"E042": {Code: ERROR_CODE_INVALID_ID, Message: "Invalid ID", HTTPStatus: fiber.StatusBadRequest},
	"E043": {Code: ERROR_CODE_MISSING_REQUIRED_FORM_DATA, Message: "Missing required form data", HTTPStatus: fiber.StatusBadRequest},
	"E044": {Code: ERROR_CODE_FAIL_PASE_REQUEST_BODY, Message: "Failed to parse request body", HTTPStatus: fiber.StatusBadRequest},
	"E045": {Code: ERROR_CODE_VALIDATE_FAIL, Message: "Failed to validate request body", HTTPStatus: fiber.StatusBadRequest},
	"E046": {Code: ERROR_CODE_FAIL_KAFKA_PRODUCER, Message: "Failed to produce message", HTTPStatus: fiber.StatusInternalServerError},
	"E047": {Code: ERROR_CODE_FAIL_KAFKA_CONSUMER, Message: "Failed to consume message", HTTPStatus: fiber.StatusInternalServerError},
	"E048": {Code: ERROR_CODE_PAGINATION_FAIL, Message: "Failed to paginate data", HTTPStatus: fiber.StatusInternalServerError},
	"E049": {Code: ERROR_CODE_QUERY_FAIL, Message: "Failed to query data", HTTPStatus: fiber.StatusInternalServerError},
	"E050": {Code: ERROR_CODE_FAIL_COMPRESS_JSON, Message: "Failed to compress JSON", HTTPStatus: fiber.StatusInternalServerError},
	"E051": {Code: ERROR_CODE_FAIL_KAFKA_READER_INITIALIZATION, Message: "Failed to initialize Kafka reader", HTTPStatus: fiber.StatusInternalServerError},
	"E052": {Code: ERROR_CODE_FAIL_KAFKA_READER_FETCH, Message: "Failed to fetch message from Kafka", HTTPStatus: fiber.StatusInternalServerError},
	"E053": {Code: ERROR_CODE_FAIL_KAFKA_READER_COMMIT, Message: "Failed to commit message", HTTPStatus: fiber.StatusInternalServerError},
	"E054": {Code: ERROR_CODE_CONNECTION_NOT_SUPPORTED, Message: "Connection not supported", HTTPStatus: fiber.StatusBadRequest},
	"E055": {Code: ERROR_CODE_FAIL_FETCH_REDIS, Message: "Failed to fetch cache data from Redis", HTTPStatus: fiber.StatusInternalServerError},
	"E056": {Code: ERROR_CODE_FAIL_SET_DATA_REDIS, Message: "Failed to set data in Redis", HTTPStatus: fiber.StatusInternalServerError},
	"E057": {Code: ERROR_CODE_FAIL_UNCOMPRESS_JSON, Message: "Failed to uncompress JSON", HTTPStatus: fiber.StatusInternalServerError},
	"E058": {Code: ERROR_CODE_MISMATCHED_DOCUMENT_DATA, Message: "Mismatched document data", HTTPStatus: fiber.StatusBadRequest},
}

func GetError(code string) *Error {
	if err, exists := errorData[code]; exists {
		return &err
	}

	return &Error{
		Code:       "E000",
		Message:    code,
		HTTPStatus: fiber.StatusInternalServerError,
	}
}

func GetAllErrors() map[string]Error {
	return errorData
}
