package iso20022

// Provides details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction15 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	PaymentCancellationIdentification *Max35Text `xml:"PmtCxlId,omitempty"`

	// Identifies the case.
	Case *Case3 `xml:"Case,omitempty"`

	// Unique and unambiguous identifier of the original payment information block, as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the cancellation payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the cancellation payment information group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether or not the cancellation applies to a whole group of transactions or to individual transactions within the original group.
	PaymentInformationCancellation *GroupCancellationIndicator `xml:"PmtInfCxl,omitempty"`

	// Detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`

	// Information concerning the original transactions, to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction61 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction15) SetPaymentCancellationIdentification(value string) {
	o.PaymentCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction15) AddCase() *Case3 {
	o.Case = new(Case3)
	return o.Case
}

func (o *OriginalPaymentInstruction15) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction15) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction15) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction15) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction15) SetPaymentInformationCancellation(value string) {
	o.PaymentInformationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalPaymentInstruction15) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction15) AddTransactionInformation() *PaymentTransaction61 {
	newValue := new(PaymentTransaction61)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}
