// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingPppKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_PPP_NOAUTH:
		t = ktypeBoolean
	case NM_SETTING_PPP_REFUSE_EAP:
		t = ktypeBoolean
	case NM_SETTING_PPP_REFUSE_PAP:
		t = ktypeBoolean
	case NM_SETTING_PPP_REFUSE_CHAP:
		t = ktypeBoolean
	case NM_SETTING_PPP_REFUSE_MSCHAP:
		t = ktypeBoolean
	case NM_SETTING_PPP_REFUSE_MSCHAPV2:
		t = ktypeBoolean
	case NM_SETTING_PPP_REQUIRE_MPPE:
		t = ktypeBoolean
	case NM_SETTING_PPP_REQUIRE_MPPE_128:
		t = ktypeBoolean
	case NM_SETTING_PPP_MPPE_STATEFUL:
		t = ktypeBoolean
	case NM_SETTING_PPP_NOBSDCOMP:
		t = ktypeBoolean
	case NM_SETTING_PPP_NODEFLATE:
		t = ktypeBoolean
	case NM_SETTING_PPP_NO_VJ_COMP:
		t = ktypeBoolean
	case NM_SETTING_PPP_CRTSCTS:
		t = ktypeBoolean
	case NM_SETTING_PPP_BAUD:
		t = ktypeUint32
	case NM_SETTING_PPP_MRU:
		t = ktypeUint32
	case NM_SETTING_PPP_MTU:
		t = ktypeUint32
	case NM_SETTING_PPP_LCP_ECHO_FAILURE:
		t = ktypeUint32
	case NM_SETTING_PPP_LCP_ECHO_INTERVAL:
		t = ktypeUint32
	}
	return
}

// Check is key in current setting field
func isKeyInSettingPpp(key string) bool {
	switch key {
	case NM_SETTING_PPP_NOAUTH:
		return true
	case NM_SETTING_PPP_REFUSE_EAP:
		return true
	case NM_SETTING_PPP_REFUSE_PAP:
		return true
	case NM_SETTING_PPP_REFUSE_CHAP:
		return true
	case NM_SETTING_PPP_REFUSE_MSCHAP:
		return true
	case NM_SETTING_PPP_REFUSE_MSCHAPV2:
		return true
	case NM_SETTING_PPP_REQUIRE_MPPE:
		return true
	case NM_SETTING_PPP_REQUIRE_MPPE_128:
		return true
	case NM_SETTING_PPP_MPPE_STATEFUL:
		return true
	case NM_SETTING_PPP_NOBSDCOMP:
		return true
	case NM_SETTING_PPP_NODEFLATE:
		return true
	case NM_SETTING_PPP_NO_VJ_COMP:
		return true
	case NM_SETTING_PPP_CRTSCTS:
		return true
	case NM_SETTING_PPP_BAUD:
		return true
	case NM_SETTING_PPP_MRU:
		return true
	case NM_SETTING_PPP_MTU:
		return true
	case NM_SETTING_PPP_LCP_ECHO_FAILURE:
		return true
	case NM_SETTING_PPP_LCP_ECHO_INTERVAL:
		return true
	}
	return false
}

// Get key's default value
func getSettingPppKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_PPP_NOAUTH:
		valueJSON = `true`
	case NM_SETTING_PPP_REFUSE_EAP:
		valueJSON = `false`
	case NM_SETTING_PPP_REFUSE_PAP:
		valueJSON = `false`
	case NM_SETTING_PPP_REFUSE_CHAP:
		valueJSON = `false`
	case NM_SETTING_PPP_REFUSE_MSCHAP:
		valueJSON = `false`
	case NM_SETTING_PPP_REFUSE_MSCHAPV2:
		valueJSON = `false`
	case NM_SETTING_PPP_REQUIRE_MPPE:
		valueJSON = `false`
	case NM_SETTING_PPP_REQUIRE_MPPE_128:
		valueJSON = `false`
	case NM_SETTING_PPP_MPPE_STATEFUL:
		valueJSON = `false`
	case NM_SETTING_PPP_NOBSDCOMP:
		valueJSON = `false`
	case NM_SETTING_PPP_NODEFLATE:
		valueJSON = `false`
	case NM_SETTING_PPP_NO_VJ_COMP:
		valueJSON = `false`
	case NM_SETTING_PPP_CRTSCTS:
		valueJSON = `false`
	case NM_SETTING_PPP_BAUD:
		valueJSON = `0`
	case NM_SETTING_PPP_MRU:
		valueJSON = `0`
	case NM_SETTING_PPP_MTU:
		valueJSON = `0`
	case NM_SETTING_PPP_LCP_ECHO_FAILURE:
		valueJSON = `0`
	case NM_SETTING_PPP_LCP_ECHO_INTERVAL:
		valueJSON = `0`
	}
	return
}

// Get JSON value generally
func generalGetSettingPppKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingPppKeyJSON: invalide key", key)
	case NM_SETTING_PPP_NOAUTH:
		value = getSettingPppNoauthJSON(data)
	case NM_SETTING_PPP_REFUSE_EAP:
		value = getSettingPppRefuseEapJSON(data)
	case NM_SETTING_PPP_REFUSE_PAP:
		value = getSettingPppRefusePapJSON(data)
	case NM_SETTING_PPP_REFUSE_CHAP:
		value = getSettingPppRefuseChapJSON(data)
	case NM_SETTING_PPP_REFUSE_MSCHAP:
		value = getSettingPppRefuseMschapJSON(data)
	case NM_SETTING_PPP_REFUSE_MSCHAPV2:
		value = getSettingPppRefuseMschapv2JSON(data)
	case NM_SETTING_PPP_REQUIRE_MPPE:
		value = getSettingPppRequireMppeJSON(data)
	case NM_SETTING_PPP_REQUIRE_MPPE_128:
		value = getSettingPppRequireMppe128JSON(data)
	case NM_SETTING_PPP_MPPE_STATEFUL:
		value = getSettingPppMppeStatefulJSON(data)
	case NM_SETTING_PPP_NOBSDCOMP:
		value = getSettingPppNobsdcompJSON(data)
	case NM_SETTING_PPP_NODEFLATE:
		value = getSettingPppNodeflateJSON(data)
	case NM_SETTING_PPP_NO_VJ_COMP:
		value = getSettingPppNoVjCompJSON(data)
	case NM_SETTING_PPP_CRTSCTS:
		value = getSettingPppCrtsctsJSON(data)
	case NM_SETTING_PPP_BAUD:
		value = getSettingPppBaudJSON(data)
	case NM_SETTING_PPP_MRU:
		value = getSettingPppMruJSON(data)
	case NM_SETTING_PPP_MTU:
		value = getSettingPppMtuJSON(data)
	case NM_SETTING_PPP_LCP_ECHO_FAILURE:
		value = getSettingPppLcpEchoFailureJSON(data)
	case NM_SETTING_PPP_LCP_ECHO_INTERVAL:
		value = getSettingPppLcpEchoIntervalJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingPppKeyJSON(data connectionData, key, valueJSON string) {
	switch key {
	default:
		logger.Error("generalSetSettingPppKeyJSON: invalide key", key)
	case NM_SETTING_PPP_NOAUTH:
		setSettingPppNoauthJSON(data, valueJSON)
	case NM_SETTING_PPP_REFUSE_EAP:
		setSettingPppRefuseEapJSON(data, valueJSON)
	case NM_SETTING_PPP_REFUSE_PAP:
		setSettingPppRefusePapJSON(data, valueJSON)
	case NM_SETTING_PPP_REFUSE_CHAP:
		setSettingPppRefuseChapJSON(data, valueJSON)
	case NM_SETTING_PPP_REFUSE_MSCHAP:
		setSettingPppRefuseMschapJSON(data, valueJSON)
	case NM_SETTING_PPP_REFUSE_MSCHAPV2:
		setSettingPppRefuseMschapv2JSON(data, valueJSON)
	case NM_SETTING_PPP_REQUIRE_MPPE:
		logicSetSettingPppRequireMppeJSON(data, valueJSON)
	case NM_SETTING_PPP_REQUIRE_MPPE_128:
		setSettingPppRequireMppe128JSON(data, valueJSON)
	case NM_SETTING_PPP_MPPE_STATEFUL:
		setSettingPppMppeStatefulJSON(data, valueJSON)
	case NM_SETTING_PPP_NOBSDCOMP:
		setSettingPppNobsdcompJSON(data, valueJSON)
	case NM_SETTING_PPP_NODEFLATE:
		setSettingPppNodeflateJSON(data, valueJSON)
	case NM_SETTING_PPP_NO_VJ_COMP:
		setSettingPppNoVjCompJSON(data, valueJSON)
	case NM_SETTING_PPP_CRTSCTS:
		setSettingPppCrtsctsJSON(data, valueJSON)
	case NM_SETTING_PPP_BAUD:
		setSettingPppBaudJSON(data, valueJSON)
	case NM_SETTING_PPP_MRU:
		setSettingPppMruJSON(data, valueJSON)
	case NM_SETTING_PPP_MTU:
		setSettingPppMtuJSON(data, valueJSON)
	case NM_SETTING_PPP_LCP_ECHO_FAILURE:
		setSettingPppLcpEchoFailureJSON(data, valueJSON)
	case NM_SETTING_PPP_LCP_ECHO_INTERVAL:
		setSettingPppLcpEchoIntervalJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingPppNoauthExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOAUTH)
}
func isSettingPppRefuseEapExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_EAP)
}
func isSettingPppRefusePapExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_PAP)
}
func isSettingPppRefuseChapExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_CHAP)
}
func isSettingPppRefuseMschapExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAP)
}
func isSettingPppRefuseMschapv2Exists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAPV2)
}
func isSettingPppRequireMppeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE)
}
func isSettingPppRequireMppe128Exists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE_128)
}
func isSettingPppMppeStatefulExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MPPE_STATEFUL)
}
func isSettingPppNobsdcompExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOBSDCOMP)
}
func isSettingPppNodeflateExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NODEFLATE)
}
func isSettingPppNoVjCompExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NO_VJ_COMP)
}
func isSettingPppCrtsctsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_CRTSCTS)
}
func isSettingPppBaudExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_BAUD)
}
func isSettingPppMruExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MRU)
}
func isSettingPppMtuExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MTU)
}
func isSettingPppLcpEchoFailureExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_FAILURE)
}
func isSettingPppLcpEchoIntervalExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_INTERVAL)
}

// Ensure field and key exists and not empty
func ensureFieldSettingPppExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_PPP_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_PPP_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_PPP_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_PPP_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_PPP_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_PPP_SETTING_NAME))
	}
}
func ensureSettingPppNoauthNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppNoauthExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOAUTH, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppRefuseEapNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppRefuseEapExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_EAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppRefusePapNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppRefusePapExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_PAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppRefuseChapNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppRefuseChapExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_CHAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppRefuseMschapNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppRefuseMschapExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppRefuseMschapv2NoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppRefuseMschapv2Exists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAPV2, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppRequireMppeNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppRequireMppeExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppRequireMppe128NoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppRequireMppe128Exists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE_128, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppMppeStatefulNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppMppeStatefulExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MPPE_STATEFUL, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppNobsdcompNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppNobsdcompExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOBSDCOMP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppNodeflateNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppNodeflateExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NODEFLATE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppNoVjCompNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppNoVjCompExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NO_VJ_COMP, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppCrtsctsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppCrtsctsExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_CRTSCTS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppBaudNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppBaudExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_BAUD, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppMruNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppMruExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MRU, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppMtuNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppMtuExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MTU, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppLcpEchoFailureNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppLcpEchoFailureExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_FAILURE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingPppLcpEchoIntervalNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingPppLcpEchoIntervalExists(data) {
		rememberError(errs, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_INTERVAL, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingPppNoauth(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOAUTH).(bool)
	return
}
func getSettingPppRefuseEap(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_EAP).(bool)
	return
}
func getSettingPppRefusePap(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_PAP).(bool)
	return
}
func getSettingPppRefuseChap(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_CHAP).(bool)
	return
}
func getSettingPppRefuseMschap(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAP).(bool)
	return
}
func getSettingPppRefuseMschapv2(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAPV2).(bool)
	return
}
func getSettingPppRequireMppe(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE).(bool)
	return
}
func getSettingPppRequireMppe128(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE_128).(bool)
	return
}
func getSettingPppMppeStateful(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MPPE_STATEFUL).(bool)
	return
}
func getSettingPppNobsdcomp(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOBSDCOMP).(bool)
	return
}
func getSettingPppNodeflate(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NODEFLATE).(bool)
	return
}
func getSettingPppNoVjComp(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NO_VJ_COMP).(bool)
	return
}
func getSettingPppCrtscts(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_CRTSCTS).(bool)
	return
}
func getSettingPppBaud(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_BAUD).(uint32)
	return
}
func getSettingPppMru(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MRU).(uint32)
	return
}
func getSettingPppMtu(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MTU).(uint32)
	return
}
func getSettingPppLcpEchoFailure(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_FAILURE).(uint32)
	return
}
func getSettingPppLcpEchoInterval(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_INTERVAL).(uint32)
	return
}

// Setter
func setSettingPppNoauth(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOAUTH, value)
}
func setSettingPppRefuseEap(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_EAP, value)
}
func setSettingPppRefusePap(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_PAP, value)
}
func setSettingPppRefuseChap(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_CHAP, value)
}
func setSettingPppRefuseMschap(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAP, value)
}
func setSettingPppRefuseMschapv2(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAPV2, value)
}
func setSettingPppRequireMppe(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE, value)
}
func setSettingPppRequireMppe128(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE_128, value)
}
func setSettingPppMppeStateful(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MPPE_STATEFUL, value)
}
func setSettingPppNobsdcomp(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOBSDCOMP, value)
}
func setSettingPppNodeflate(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NODEFLATE, value)
}
func setSettingPppNoVjComp(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NO_VJ_COMP, value)
}
func setSettingPppCrtscts(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_CRTSCTS, value)
}
func setSettingPppBaud(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_BAUD, value)
}
func setSettingPppMru(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MRU, value)
}
func setSettingPppMtu(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MTU, value)
}
func setSettingPppLcpEchoFailure(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_FAILURE, value)
}
func setSettingPppLcpEchoInterval(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_INTERVAL, value)
}

// JSON Getter
func getSettingPppNoauthJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOAUTH, getSettingPppKeyType(NM_SETTING_PPP_NOAUTH))
	return
}
func getSettingPppRefuseEapJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_EAP, getSettingPppKeyType(NM_SETTING_PPP_REFUSE_EAP))
	return
}
func getSettingPppRefusePapJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_PAP, getSettingPppKeyType(NM_SETTING_PPP_REFUSE_PAP))
	return
}
func getSettingPppRefuseChapJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_CHAP, getSettingPppKeyType(NM_SETTING_PPP_REFUSE_CHAP))
	return
}
func getSettingPppRefuseMschapJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAP, getSettingPppKeyType(NM_SETTING_PPP_REFUSE_MSCHAP))
	return
}
func getSettingPppRefuseMschapv2JSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAPV2, getSettingPppKeyType(NM_SETTING_PPP_REFUSE_MSCHAPV2))
	return
}
func getSettingPppRequireMppeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE, getSettingPppKeyType(NM_SETTING_PPP_REQUIRE_MPPE))
	return
}
func getSettingPppRequireMppe128JSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE_128, getSettingPppKeyType(NM_SETTING_PPP_REQUIRE_MPPE_128))
	return
}
func getSettingPppMppeStatefulJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MPPE_STATEFUL, getSettingPppKeyType(NM_SETTING_PPP_MPPE_STATEFUL))
	return
}
func getSettingPppNobsdcompJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOBSDCOMP, getSettingPppKeyType(NM_SETTING_PPP_NOBSDCOMP))
	return
}
func getSettingPppNodeflateJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NODEFLATE, getSettingPppKeyType(NM_SETTING_PPP_NODEFLATE))
	return
}
func getSettingPppNoVjCompJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NO_VJ_COMP, getSettingPppKeyType(NM_SETTING_PPP_NO_VJ_COMP))
	return
}
func getSettingPppCrtsctsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_CRTSCTS, getSettingPppKeyType(NM_SETTING_PPP_CRTSCTS))
	return
}
func getSettingPppBaudJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_BAUD, getSettingPppKeyType(NM_SETTING_PPP_BAUD))
	return
}
func getSettingPppMruJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MRU, getSettingPppKeyType(NM_SETTING_PPP_MRU))
	return
}
func getSettingPppMtuJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MTU, getSettingPppKeyType(NM_SETTING_PPP_MTU))
	return
}
func getSettingPppLcpEchoFailureJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_FAILURE, getSettingPppKeyType(NM_SETTING_PPP_LCP_ECHO_FAILURE))
	return
}
func getSettingPppLcpEchoIntervalJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_INTERVAL, getSettingPppKeyType(NM_SETTING_PPP_LCP_ECHO_INTERVAL))
	return
}

// JSON Setter
func setSettingPppNoauthJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOAUTH, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_NOAUTH))
}
func setSettingPppRefuseEapJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_EAP, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_REFUSE_EAP))
}
func setSettingPppRefusePapJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_PAP, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_REFUSE_PAP))
}
func setSettingPppRefuseChapJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_CHAP, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_REFUSE_CHAP))
}
func setSettingPppRefuseMschapJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAP, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_REFUSE_MSCHAP))
}
func setSettingPppRefuseMschapv2JSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAPV2, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_REFUSE_MSCHAPV2))
}
func setSettingPppRequireMppeJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_REQUIRE_MPPE))
}
func setSettingPppRequireMppe128JSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE_128, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_REQUIRE_MPPE_128))
}
func setSettingPppMppeStatefulJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MPPE_STATEFUL, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_MPPE_STATEFUL))
}
func setSettingPppNobsdcompJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOBSDCOMP, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_NOBSDCOMP))
}
func setSettingPppNodeflateJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NODEFLATE, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_NODEFLATE))
}
func setSettingPppNoVjCompJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NO_VJ_COMP, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_NO_VJ_COMP))
}
func setSettingPppCrtsctsJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_CRTSCTS, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_CRTSCTS))
}
func setSettingPppBaudJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_BAUD, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_BAUD))
}
func setSettingPppMruJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MRU, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_MRU))
}
func setSettingPppMtuJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MTU, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_MTU))
}
func setSettingPppLcpEchoFailureJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_FAILURE, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_LCP_ECHO_FAILURE))
}
func setSettingPppLcpEchoIntervalJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_INTERVAL, valueJSON, getSettingPppKeyType(NM_SETTING_PPP_LCP_ECHO_INTERVAL))
}

// Remover
func removeSettingPppNoauth(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOAUTH)
}
func removeSettingPppRefuseEap(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_EAP)
}
func removeSettingPppRefusePap(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_PAP)
}
func removeSettingPppRefuseChap(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_CHAP)
}
func removeSettingPppRefuseMschap(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAP)
}
func removeSettingPppRefuseMschapv2(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REFUSE_MSCHAPV2)
}
func removeSettingPppRequireMppe(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE)
}
func removeSettingPppRequireMppe128(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_REQUIRE_MPPE_128)
}
func removeSettingPppMppeStateful(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MPPE_STATEFUL)
}
func removeSettingPppNobsdcomp(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NOBSDCOMP)
}
func removeSettingPppNodeflate(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NODEFLATE)
}
func removeSettingPppNoVjComp(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_NO_VJ_COMP)
}
func removeSettingPppCrtscts(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_CRTSCTS)
}
func removeSettingPppBaud(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_BAUD)
}
func removeSettingPppMru(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MRU)
}
func removeSettingPppMtu(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_MTU)
}
func removeSettingPppLcpEchoFailure(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_FAILURE)
}
func removeSettingPppLcpEchoInterval(data connectionData) {
	removeSettingKey(data, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_INTERVAL)
}
