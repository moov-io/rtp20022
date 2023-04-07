package msg_status

import (
	"strings"
)

type ErrorCode struct {
	Code        string
	Description string
	Level       ErrorLevel
}

type ErrorLevel string

func (e ErrorLevel) Error() string {
	return string(e)
}

const (
	// ErrAccountFatal indicates a problem with the (Creditor) account.
	// It likely should be reverified.
	ErrAccountFatal ErrorLevel = "Fatal error for the Account"

	// ErrFatal indicates a problem that likely requires a human be notified.
	ErrFatal ErrorLevel = "Fatal error"

	// ErrTemporary indicates a problem which should go away with time.
	ErrTemporary ErrorLevel = "Temporary"

	// ErrNetwork indicates a problem with the network / connection.
	ErrNetwork ErrorLevel = "Network issue"

	// ErrLogic indicates a bug to fix or work around.
	ErrLogic ErrorLevel = "Logic Bug"
)

func IsError(code string) *ErrorCode {
	code = strings.ToUpper(code)

	ec, exists := errorCodes[code]
	if exists {
		ec.Code = code
		return &ec
	}
	return nil
}

var errorCodes = map[string]ErrorCode{
	"1100": {Description: "Failed For Another Reason", Level: ErrFatal},
	"9909": {Description: "TCH Central Switch Malfunction", Level: ErrFatal},
	"9910": {Description: "TCH Instructed Agent Signed Off", Level: ErrTemporary}, // Creditor
	"9912": {Description: "TCH Recipient Is Not Available", Level: ErrTemporary},
	"9914": {Description: "Element Mandatory For Zelle", Level: ErrLogic},
	"9934": {Description: "TCH Instructing Agent Signed Off", Level: ErrLogic}, // Debtor
	"9941": {Description: "Missing Proxy Identification", Level: ErrFatal},
	"9946": {Description: "TCH Instructing Agent Suspended", Level: ErrFatal},       // Debtor
	"9947": {Description: "TCH Instructed Agent Suspended", Level: ErrAccountFatal}, // Creditor
	"9948": {Description: "TCH Central Switch Suspended", Level: ErrFatal},
	"9952": {Description: "Mapping Incompaibility Between Versions", Level: ErrLogic},
	"9953": {Description: "Missing Code FULL", Level: ErrLogic},
	"9954": {Description: "Missing Instructions For Creditor Agent Code In Zelle RFP", Level: ErrLogic},
	"9956": {Description: "Instructing Agent Funding Account Suspended", Level: ErrFatal},       // Debtor
	"9957": {Description: "Instructed Agent Funding Account Suspended", Level: ErrAccountFatal}, // Creditor
	"9964": {Description: "Invalid Participant Identification", Level: ErrAccountFatal},
	"AC02": {Description: "Debtor Account Invalid", Level: ErrFatal},
	"AC03": {Description: "Creditor Account Invalid", Level: ErrAccountFatal},
	"AC04": {Description: "Account Closed", Level: ErrAccountFatal},
	"AC06": {Description: "Account Blocked", Level: ErrAccountFatal},
	"AC07": {Description: "Creditor Account Closed", Level: ErrAccountFatal},
	"AC10": {Description: "Debtor Currency Invalid", Level: ErrFatal},
	"AC11": {Description: "Creditor Currency Invalid", Level: ErrAccountFatal},
	"AC13": {Description: "Debtor Account Type Invalid", Level: ErrFatal},
	"AC14": {Description: "Creditor Account Type Invalid", Level: ErrAccountFatal},
	"AG01": {Description: "Transactions Forbidden On Account", Level: ErrFatal},
	"AG03": {Description: "Transaction Type Not Supported", Level: ErrFatal},
	"AGNT": {Description: "Incorrect Agent", Level: ErrFatal},
	"AM02": {Description: "Transaction Amount Exceeds Limit", Level: ErrTemporary},
	"AM04": {Description: "Transaction Amount Exceeds Funding", Level: ErrTemporary},
	"AM09": {Description: "Unexpected Amount Received", Level: ErrLogic},
	"AM11": {Description: "Transaction Currency Invalid", Level: ErrLogic},
	"AM12": {Description: "Transaction Amount Invalid", Level: ErrLogic},
	"AM13": {Description: "Transaction Amount Exceeds Clearing Limit", Level: ErrTemporary},
	"AM14": {Description: "Transaction Amount Exceeds Agreement Limit", Level: ErrTemporary},
	"BE04": {Description: "Incorrect Creditor Address", Level: ErrAccountFatal},
	"BE06": {Description: "Customer Not Known To FI", Level: ErrAccountFatal},
	"BE07": {Description: "Incorrect Debtor Address", Level: ErrFatal},
	"BE10": {Description: "Debtor Country Code Invalid", Level: ErrFatal},
	"BE11": {Description: "Creditor Country Code Invalid", Level: ErrAccountFatal},
	"BE13": {Description: "Debtor Residence Country Code Invalid", Level: ErrFatal},
	"BE14": {Description: "Creditor Residence Country Code Invalid", Level: ErrAccountFatal},
	"BE16": {Description: "Debtor Identification Code Invalid", Level: ErrFatal},
	"BE17": {Description: "Creditor Identification Code Invalid", Level: ErrAccountFatal},
	"CH11": {Description: "Creditor Identifier Incorrect", Level: ErrAccountFatal},
	"CUST": {Description: "Requested By Customer", Level: ErrAccountFatal},
	"DS04": {Description: "Order Rejected", Level: ErrAccountFatal},
	"DS0H": {Description: "Signer Not Allowed", Level: ErrFatal},
	"DS24": {Description: "Waiting Time Expired", Level: ErrTemporary},
	"DT04": {Description: "Future Date Not Supported", Level: ErrLogic},
	"DUPL": {Description: "Duplicate Payment", Level: ErrTemporary},
	"FF02": {Description: "Syntax Erorr", Level: ErrFatal},
	"FF03": {Description: "Invalid Payment Type", Level: ErrFatal},
	"FF08": {Description: "End To End Id Missing", Level: ErrFatal},
	"MD07": {Description: "Customer Deceased", Level: ErrAccountFatal},
	"NARR": {Description: "Check Narrative Information", Level: ErrLogic},
	"NOAT": {Description: "Message Type Not Supported", Level: ErrLogic},
	"RC01": {Description: "Incorrect Format For RoutingCode", Level: ErrFatal},
	"RC02": {Description: "Invalid Bank", Level: ErrFatal},
	"RC03": {Description: "Debtor FI Invalid", Level: ErrFatal},
	"RC04": {Description: "Creditor FI Invalid", Level: ErrAccountFatal},
	"SL03": {Description: "Token Service Did Not Respond", Level: ErrNetwork},
	"TK01": {Description: "Invalid Token", Level: ErrAccountFatal},
	"TK02": {Description: "Sender Token Not Found", Level: ErrAccountFatal},
	"TK03": {Description: "Receiver Token Not Found", Level: ErrAccountFatal},
	"TK04": {Description: "Token Expired", Level: ErrTemporary},
	"TK05": {Description: "Token Found With Counterparty Mismatch", Level: ErrLogic},
	"TK06": {Description: "Token Found With Value Limit Rule Violation", Level: ErrLogic},
	"TK07": {Description: "Single Use Token Already Used", Level: ErrLogic},
	"TK08": {Description: "Token Suspended", Level: ErrAccountFatal},
	"TM01": {Description: "Invalid CutOff Time", Level: ErrLogic},
}
