package applicationmessagerequest

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
) //ApplicationMessageRequest is the fix44 ApplicationMessageRequest type, MsgType = BW

type ApplicationMessageRequest struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

//FromMessage creates a ApplicationMessageRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) ApplicationMessageRequest {
	return ApplicationMessageRequest{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ApplicationMessageRequest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a MarketDataRequest initialized with the required fields for MarketDataRequest
func New(reqId field.ApplReqIDField, reqType field.ApplReqTypeField) (m ApplicationMessageRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BW"))
	m.Set(reqId)
	m.Set(reqType)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ApplicationMessageRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "BW", r
}

//SetAccount sets ApplReqId, Tag 1346
func (m ApplicationMessageRequest) SetApplReqId(v string) {
	m.Set(field.NewApplReqID(v))
}

//SetAccount sets ApplReqId, Tag 1347
func (m ApplicationMessageRequest) SetApplReqType(v string) {
	m.Set(field.NewApplReqType(enum.ApplReqType(v)))
}


func (m ApplicationMessageRequest) SetNoApplIDs(f NoApplIDsRepeatingGroup) {
	m.SetGroup(f)
}

func (m ApplicationMessageRequest) GetNoApplIDs() (NoApplIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoApplIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAccount gets ApplReqId, Tag 1346
func (m ApplicationMessageRequest) GetApplReqId() (v string, err quickfix.MessageRejectError) {
	var f field.ApplReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccount gets ApplReqId, Tag 1346
func (m ApplicationMessageRequest) GetApplReqType() (v string, err quickfix.MessageRejectError) {
	var f field.ApplReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
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

func (m NoApplIDs) SetApplBegSeqNum(v int) {
	m.Set(field.NewApplBegSeqNum(v))
}

func (m NoApplIDs) SetApplEndSeqNum(v int) {
	m.Set(field.NewApplEndSeqNum(v))
}

func (m NoApplIDs) GetRefApplID() (v string, err quickfix.MessageRejectError) {
	var f field.RefApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

func (m NoApplIDs) GetApplBegSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplBegSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

func (m NoApplIDs) GetApplEndSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplEndSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
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

//NoApplIDsRepeatingGroup is a repeating group, Tag 1351
type NoApplIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMDEntryTypesRepeatingGroup returns an initialized, NoMDEntryTypesRepeatingGroup
func NewNoApplIDsRepeatingGroup() NoApplIDsRepeatingGroup {
	return NoApplIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoApplIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RefApplID), quickfix.GroupElement(tag.ApplBegSeqNum), quickfix.GroupElement(tag.ApplEndSeqNum)})}
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

