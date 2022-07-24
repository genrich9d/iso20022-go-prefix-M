package iso20022

// Key elements used to refer the original transaction.
type OriginalTransactionReference21 struct {

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType4Choice `xml:"Amt,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod4Code `xml:"PmtMtd,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation6 `xml:"RmtInf,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`
}

func (o *OriginalTransactionReference21) AddAmount() *AmountType4Choice {
	o.Amount = new(AmountType4Choice)
	return o.Amount
}

func (o *OriginalTransactionReference21) SetRequestedExecutionDate(value string) {
	o.RequestedExecutionDate = (*ISODate)(&value)
}

func (o *OriginalTransactionReference21) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	o.PaymentTypeInformation = new(PaymentTypeInformation19)
	return o.PaymentTypeInformation
}

func (o *OriginalTransactionReference21) SetPaymentMethod(value string) {
	o.PaymentMethod = (*PaymentMethod4Code)(&value)
}

func (o *OriginalTransactionReference21) AddRemittanceInformation() *RemittanceInformation6 {
	o.RemittanceInformation = new(RemittanceInformation6)
	return o.RemittanceInformation
}

func (o *OriginalTransactionReference21) AddUltimateDebtor() *PartyIdentification43 {
	o.UltimateDebtor = new(PartyIdentification43)
	return o.UltimateDebtor
}

func (o *OriginalTransactionReference21) AddDebtor() *PartyIdentification43 {
	o.Debtor = new(PartyIdentification43)
	return o.Debtor
}

func (o *OriginalTransactionReference21) AddDebtorAccount() *CashAccount24 {
	o.DebtorAccount = new(CashAccount24)
	return o.DebtorAccount
}

func (o *OriginalTransactionReference21) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalTransactionReference21) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.CreditorAgent
}

func (o *OriginalTransactionReference21) AddCreditor() *PartyIdentification43 {
	o.Creditor = new(PartyIdentification43)
	return o.Creditor
}

func (o *OriginalTransactionReference21) AddCreditorAccount() *CashAccount24 {
	o.CreditorAccount = new(CashAccount24)
	return o.CreditorAccount
}

func (o *OriginalTransactionReference21) AddUltimateCreditor() *PartyIdentification43 {
	o.UltimateCreditor = new(PartyIdentification43)
	return o.UltimateCreditor
}
