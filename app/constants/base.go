package constants

const (
	PERMISSION_FACIAL_RECOGNITION = "FACIAL_RECOGNITION"
	APIKEY_REVOKED                = "revoked"
	APIKEY_ENABLED                = "enabled"
	CLIENT_INACTIVE               = "inactive"
	CLIENT_ACTIVE                 = "active"

	ERROR_INSUFFICIENT_PRIVELEGES = "INSUFFICIENT PRIVILEGES"
)

var SupportedApiKeyPermissions = []string{
	PERMISSION_FACIAL_RECOGNITION,
}
