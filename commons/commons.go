package commons

import (
	"alerts/commons/constants"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/configs"

	"strconv"
	"strings"
)

func GetAlertId(input string) (uint64, error) {
	// Define a regular expression to match the alert ID
	regex := configs.GetRegexPattern(genericConstants.NonDigitSequenceKey)
	result := regex.ReplaceAllString(input, "")

	// Extract the numeric part and convert it to an integer
	alertID, err := strconv.ParseUint(result, 10, 32)
	if err != nil {
		return 0, fmt.Errorf(constants.AlertIdConvertError, err)
	}
	return alertID, nil
}

type SendAlertsOn string

const (
	Email SendAlertsOn = "EMAIL"
	SMS   SendAlertsOn = "SMS"
	Both  SendAlertsOn = "EMAILSMS"
)

func (alerts SendAlertsOn) IsValid() bool {
	upperAlerts := strings.ToUpper(string(alerts))
	switch upperAlerts {
	case string(Email), string(SMS), string(Both):
		return true
	default:
		return false
	}
}

type Target string

const (
	Cancelled Target = "CANCELLED"
	Complete  Target = "COMPLETE"
	Modified  Target = "MODIFIED"
)

func (alerts Target) IsValid() bool {
	upperAlerts := strings.ToUpper(string(alerts))
	switch upperAlerts {
	case string(Cancelled), string(Complete), string(Modified):
		return true
	default:
		return false
	}
}

func ExtractValue(input, key string) string {
	// Constructing a regex pattern to find key and its value
	regex := configs.GetRegexPattern(key)
	match := regex.FindStringSubmatch(strings.ToUpper(input))

	if len(match) == 2 {
		return match[1]
	}

	return ""
}
