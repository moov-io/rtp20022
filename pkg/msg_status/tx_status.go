package msg_status

type TransactionStatus string

const (
	TxStatusAccepted TransactionStatus = "ACTC" // Payment has been Accepted

	// TxStatusAcceptedWithoutPosting is the "Accepted Without Posting" transaction status
	//
	// Payment instruction included in the Credit Transfer is Accepted
	// without being Posted to the Creditor Customers account.
	// This is done so further checks and status update can be performed
	// (eg. OFAC checks). See System Message Flows for more information.
	TxStatusAcceptedWithoutPosting TransactionStatus = "ACWP"

	TxStatusReceived TransactionStatus = "RCVD"
	TxStatusRejected TransactionStatus = "RJCT"
)
