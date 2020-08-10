package applicationmessagerequestack

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

type ApplicationMessageRequestAck struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

//FromMessage creates a ApplicationMessageRequestAck from a quickfix.Message instance
func FromMessage(m *quickfix.Message) ApplicationMessageRequestAck {
	return ApplicationMessageRequestAck{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ApplicationMessageRequestAck) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a ApplicationMessageRequestAck initialized with the required fields for ApplicationMessageRequestAck
func New(reqId field.ApplReqIDField, reqType field.ApplReqTypeField, respId field.ApplResponseIDField,
	respType field.ApplResponseTypeField) (m ApplicationMessageRequestAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BX"))
	m.Set(reqId)
	m.Set(reqType)
	m.Set(respId)
	m.Set(respType)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ApplicationMessageRequestAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "BX", r
}

// SetApplReqId sets ApplReqId, Tag 1346
func (m ApplicationMessageRequestAck) SetApplReqId(v string) {
	m.Set(field.NewApplReqID(v))
}

// SetApplRespId sets ApplRespId, Tag 1353
func (m ApplicationMessageRequestAck) SetApplRespId(v string) {
	m.Set(field.NewApplResponseID(v))
}

// SetApplReqType sets ApplReqType, Tag 1347
func (m ApplicationMessageRequestAck) SetApplReqType(v string) {
	m.Set(field.NewApplReqType(enum.ApplReqType(v)))
}

// SetApplRespType sets ApplRespType, Tag 1348
func (m ApplicationMessageRequestAck) SetApplRespType(v string) {
	m.Set(field.NewApplResponseType(enum.ApplResponseType(v)))
}

// GetApplReqId gets ApplReqId, Tag 1346
func (m ApplicationMessageRequestAck) GetApplReqId() (v string) {
	var f field.ApplReqIDField
	if err := m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplRespId gets ApplRespId, Tag 1353
func (m ApplicationMessageRequestAck) GetApplRespId() (v string) {
	var f field.ApplResponseIDField
	if err := m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplRespType gets ApplRespType, Tag 1348
func (m ApplicationMessageRequestAck) GetApplRespType() (v string, err quickfix.MessageRejectError) {
	var f field.ApplReqTypeField
	if err = m.Get(&f); err == nil {
		v = string(f.Value())
	}
	return
}

//GetAccount gets ApplReqId, Tag 1347
func (m ApplicationMessageRequestAck) GetApplReqType() (v string, err quickfix.MessageRejectError) {
	var f field.ApplReqTypeField
	if err = m.Get(&f); err == nil {
		v = string(f.Value())
	}
	return
}

//NoMDEntryTypes is a repeating group element, Tag 1351
type NoApplIDs struct {
	*quickfix.Group
}

//SetMDEntryType sets MDEntryType, Tag 1355
func (m NoApplIDs) SetRefApplID(v string) {
	m.Set(field.NewRefApplID(v))
}

// Tag 1182
func (m NoApplIDs) SetApplBegSeqNum(v int) {
	m.Set(field.NewApplBegSeqNum(v))
}

// Tag 1183
func (m NoApplIDs) SetApplEndSeqNum(v int) {
	m.Set(field.NewApplEndSeqNum(v))
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

// Tag 1182
func (m NoApplIDs) GetApplBegSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplBegSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// Tag 1183
func (m NoApplIDs) GetApplEndSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplEndSeqNumField
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

func (m NoApplIDs) HasApplBegSeqNum() bool {
	return m.Has(tag.ApplBegSeqNum)
}

func (m NoApplIDs) HasApplEndSeqNum() bool {
	return m.Has(tag.ApplEndSeqNum)
}

func (m NoApplIDs) HasApplRespError() bool {
	return m.Has(tag.ApplResponseError)
}

func (m ApplicationMessageRequestAck) SetNoApplIDs(f NoApplIDsRepeatingGroup) {
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RefApplID), quickfix.GroupElement(tag.ApplBegSeqNum),
				quickfix.GroupElement(tag.ApplEndSeqNum), quickfix.GroupElement(tag.ApplResponseError)})}
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
