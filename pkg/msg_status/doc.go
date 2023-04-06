// Package msg_status implements the RTP Message Status Report (pacs.002)
//
// The Message Status Report (pacs.002) is used to provide a real time response to all Payment
// Related Messages. The pacs.002 will confirm that the message has been received and whether
// or not it passed business validation. As such, the Message Status Report is effectively a
// "technical" response. Where a Payment Related Message has been accepted by the receiving
// party, a subsequent response may also be required in order to fulfil the request contained
// within the message, or to advise that the request is being declined.
//
//	Payment Acknowledgement camt.035.001.03
package msg_status
