package iso20022

// Elements which allow to identify a meeting.
type MeetingReference2 struct {

	// Identification assigned to a general meeting by the party notifying the meeting. It must be unique for the party notifying the meeting.
	MeetingIdentification *Max35Text `xml:"MtgId"`

	// Identification assigned to a meeting by the issuer. It must be unique for the issuer.
	IssuerMeetingIdentification *Max35Text `xml:"IssrMtgId,omitempty"`

	// Date and time at which the meeting will take place.
	MeetingDateAndTime *ISODateTime `xml:"MtgDtAndTm,omitempty"`

	// Specifies the type of meeting for which  instructions are sent.
	Type *MeetingType2Code `xml:"Tp,omitempty"`

	// Classifies the type of meeting.
	Classification *MeetingTypeClassification1Code `xml:"Clssfctn,omitempty"`

	// This code can be used in case another meeting classifications is required.
	ExtendedClassification *Extended350Code `xml:"XtndedClssfctn,omitempty"`

	// Place of the company meeting for the scheduled meeting date.
	Location []*PostalAddress1 `xml:"Lctn,omitempty"`
}

func (m *MeetingReference2) SetMeetingIdentification(value string) {
	m.MeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingReference2) SetIssuerMeetingIdentification(value string) {
	m.IssuerMeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingReference2) SetMeetingDateAndTime(value string) {
	m.MeetingDateAndTime = (*ISODateTime)(&value)
}

func (m *MeetingReference2) SetType(value string) {
	m.Type = (*MeetingType2Code)(&value)
}

func (m *MeetingReference2) SetClassification(value string) {
	m.Classification = (*MeetingTypeClassification1Code)(&value)
}

func (m *MeetingReference2) SetExtendedClassification(value string) {
	m.ExtendedClassification = (*Extended350Code)(&value)
}

func (m *MeetingReference2) AddLocation() *PostalAddress1 {
	newValue := new(PostalAddress1)
	m.Location = append(m.Location, newValue)
	return newValue
}
