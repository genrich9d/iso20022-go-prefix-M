package iso20022

// Provides details information on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction23 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *ExternalPaymentGroupStatus1Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction82 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction23) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction23) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction23) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction23) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*ExternalPaymentGroupStatus1Code)(&value)
}

func (o *OriginalPaymentInstruction23) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction23) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus5 {
	newValue := new(NumberOfTransactionsPerStatus5)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction23) AddTransactionInformationAndStatus() *PaymentTransaction82 {
	newValue := new(PaymentTransaction82)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}
