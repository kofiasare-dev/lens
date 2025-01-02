package constants

const (
	PERMISSION_FACIAL_RECOGNITION = "FACIAL_RECOGNITION"
	APIKEY_REVOKED                = "revoked"
	APIKEY_ENABLED                = "enabled"
	CLIENT_INACTIVE               = "inactive"
	CLIENT_ACTIVE                 = "active"

	ERROR_INSUFFICIENT_PRIVELEGES = "INSUFFICIENT PRIVILEGES"

	// TASKS
	VERIFICATION_REQUEST_PROCESSING_TASK = "verification_request:process"
	QUEUE_DEFAULT                        = "default"
	QUEUE_PRIORITY                       = "priority"
)

var (
	SupportedApiKeyPermissions = []string{
		PERMISSION_FACIAL_RECOGNITION,
	}
)
