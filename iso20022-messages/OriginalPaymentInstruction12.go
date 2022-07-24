package iso20022

// Provides details information on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction12 struct {

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of the payment information group.
	PaymentInformationStatus *TransactionGroupStatus3Code `xml:"PmtInfSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`

	// Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus []*PaymentTransaction57 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction12) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction12) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction12) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction12) SetPaymentInformationStatus(value string) {
	o.PaymentInformationStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalPaymentInstruction12) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction12) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction12) AddTransactionInformationAndStatus() *PaymentTransaction57 {
	newValue := new(PaymentTransaction57)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}
