package applicationmessagereport

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

type ApplicationMessageReport struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

//FromMessage creates a ApplicationMessageRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) ApplicationMessageReport {
	return ApplicationMessageReport{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ApplicationMessageReport) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a ApplicationMessageRequestAck initialized with the required fields for MarketDataReApplicationMessageRequestAckquest
func New(reqId field.ApplReqIDField, reqType field.ApplReqTypeField, respId field.ApplResponseIDField,
	reportId field.ApplReportIDField, reportType field.ApplReportTypeField) (m ApplicationMessageReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BY"))
	m.Set(reqId)
	m.Set(reqType)
	m.Set(respId)
	m.Set(reportId)
	m.Set(reportType)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ApplicationMessageReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "BY", r
}

//SetAccount sets ApplReqId, Tag 1346
func (m ApplicationMessageReport) SetApplReqId(v string) {
	m.Set(field.NewApplReqID(v))
}

// SetApplReqType sets ApplReqType, Tag 1347
func (m ApplicationMessageReport) SetApplReqType(v string) {
	m.Set(field.NewApplReqType(enum.ApplReqType(v)))
}

// SetApplRespId sets ApplRespId, Tag 1353
func (m ApplicationMessageReport) SetApplRespId(v string) {
	m.Set(field.NewApplResponseID(v))
}

// SetApplRespId sets ApplRespId, Tag 1356
func (m ApplicationMessageReport) SetApplReportId(v string) {
	m.Set(field.NewApplReportID(v))
}

// SetApplRespId sets ApplRespId, Tag 1426
func (m ApplicationMessageReport) SetApplReportType(v string) {
	m.Set(field.NewApplReportType(enum.ApplReportType(v)))
}

//GetApplReqId gets ApplReqId, Tag 1346
func (m ApplicationMessageReport) GetApplReqId() (v string, err quickfix.MessageRejectError) {
	var f field.ApplReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplReqType gets ApplReqType, Tag 1347
func (m ApplicationMessageReport) GetApplReqType() (v string, err quickfix.MessageRejectError) {
	var f field.ApplReqTypeField
	if err = m.Get(&f); err == nil {
		v = string(f.Value())
	}
	return
}

// GetApplRespId gets ApplRespId, Tag 1353
func (m ApplicationMessageReport) GetApplRespId() (v string) {
	var f field.ApplResponseIDField
	if err := m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplReportId gets ApplReportId, Tag 1356
func (m ApplicationMessageReport) GetApplReportId() (v string) {
	var f field.ApplReportIDField
	if err := m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplReportType gets ApplReportType, Tag 1426
func (m ApplicationMessageReport) GetApplReportType() (v string, err quickfix.MessageRejectError) {
	var f field.ApplReportTypeField
	if err = m.Get(&f); err == nil {
		v = string(f.Value())
	}
	return
}

//NoApplIDs is a repeating group element, Tag 1351
type NoApplIDs struct {
	*quickfix.Group
}

//SetRefApplID sets RefApplID, Tag 1355
func (m NoApplIDs) SetRefApplID(v string) {
	m.Set(field.NewRefApplID(v))
}

//SetRefApplID sets RefApplID, Tag 1357
func (m NoApplIDs) SetRefApplLastSeqNum(v int) {
	m.Set(field.NewRefApplLastSeqNum(v))
}

// Tag 1354
func (m NoApplIDs) SetApplRespError(v string) {
	m.Set(field.NewApplResponseError(enum.ApplResponseError(v)))
}

// Tag 1355
func (m NoApplIDs) GetRefApplID() (v string, err quickfix.MessageRejectError) {
	var f field.RefApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// Tag 1357
func (m NoApplIDs) GetRefApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.RefApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}
// Tag 1354
func (m NoApplIDs) GetApplRespError() (v string, err quickfix.MessageRejectError) {
	var f field.ApplEndSeqNumField
	if err = m.Get(&f); err == nil {
		v = string(f.Value())
	}
	return
}

func (m NoApplIDs) HasRefApplID() bool {
	return m.Has(tag.RefApplID)
}

func (m NoApplIDs) HasRefApplLastSeqNum() bool {
	return m.Has(tag.RefApplLastSeqNum)
}

func (m NoApplIDs) HasApplRespError() bool {
	return m.Has(tag.ApplResponseError)
}

func (m ApplicationMessageReport) SetNoApplIDs(f NoApplIDsRepeatingGroup) {
	m.SetGroup(f)
}

//NoApplIDsRepeatingGroup is a repeating group, Tag 1351
type NoApplIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMDEntryTypesRepeatingGroup returns an initialized, NoMDEntryTypesRepeatingGroup
func NewNoApplIDsRepeatingGroup() NoApplIDsRepeatingGroup {
	return NoApplIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoApplIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RefApplID), quickfix.GroupElement(tag.RefApplLastSeqNum),
				quickfix.GroupElement(tag.ApplResponseError)})}
}

//Add create and append a new NoApplIDs to this group
func (m NoApplIDsRepeatingGroup) Add() NoApplIDs {
	g := m.RepeatingGroup.Add()
	return NoApplIDs{g}
}

//Get returns the ith NoApplIDs in the NewNoApplIDsRepeatingGroup
func (m NoApplIDsRepeatingGroup) Get(i int) NoApplIDs {
	return NoApplIDs{m.RepeatingGroup.Get(i)}
}

