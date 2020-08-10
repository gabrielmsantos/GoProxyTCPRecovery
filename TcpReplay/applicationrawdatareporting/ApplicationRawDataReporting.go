package applicationrawdatareporting

import (
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

const(
	TagNoApplSeqNums 		quickfix.Tag = 10054
	TagRawDataOffset		quickfix.Tag = 10055
)

// BEGIN: RawDataOffsetField is a B3 specific field
type RawDataOffsetField struct{ quickfix.FIXInt }

//Tag returns tag.RawDataOffset (10055)
func (f RawDataOffsetField) Tag() quickfix.Tag { return TagRawDataOffset }

//NewRawDataOffsetField returns a new RawDataOffsetField initialized with val
func NewRawDataOffsetField(val int) RawDataOffsetField {
	return RawDataOffsetField{quickfix.FIXInt(val)}
}

func (f RawDataOffsetField) Value() int { return f.Int() }
// END: RawDataOffsetField

type ApplicationRawDataReporting struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

//FromMessage creates a ApplicationRawDataReporting from a quickfix.Message instance
func FromMessage(m *quickfix.Message) ApplicationRawDataReporting {
	return ApplicationRawDataReporting{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ApplicationRawDataReporting) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a ApplicationRawDataReporting initialized with the required fields for ApplicationRawDataReporting
func New(reqId field.ApplReqIDField, respId field.ApplResponseIDField,
	applID field.ApplIDField, applResendFlag field.ApplResendFlagField, rawDataLength field.RawDataLengthField,
	rawData field.RawDataField, TotNumReports field.TotNumReportsField) (m ApplicationRawDataReporting) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("URDR"))
	m.Set(reqId)
	m.Set(respId)
	m.Set(applID)
	m.Set(applResendFlag)
	m.Set(rawDataLength)
	m.Set(rawData)
	m.Set(TotNumReports)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ApplicationRawDataReporting, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "URDR", r
}

//SetApplReqId sets ApplReqId, Tag 1346
func (m ApplicationRawDataReporting) SetApplReqId(v string) {
	m.Set(field.NewApplReqID(v))
}

// SetApplRespId sets ApplRespId, Tag 1353
func (m ApplicationRawDataReporting) SetApplRespId(v string) {
	m.Set(field.NewApplResponseID(v))
}

// Tag 1180
func (m ApplicationRawDataReporting) SetApplId(v string) {
	m.Set(field.NewApplID(v))
}

// Tag 1352
func (m ApplicationRawDataReporting) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

// Tag 95
func (m ApplicationRawDataReporting) SetRawDataLength(v int) {
	m.Set(field.NewRawDataLength(v))
}

// Tag 96
func (m ApplicationRawDataReporting) SetRawData(v string) {
	m.Set(field.NewRawData(v))
}

// Tag 911
func (m ApplicationRawDataReporting) SetTotNumReports(v int) {
	m.Set(field.NewTotNumReports(v))
}

//GetApplReqId gets ApplReqId, Tag 1346
func (m ApplicationRawDataReporting) GetApplReqId() (v string, err quickfix.MessageRejectError) {
	var f field.ApplReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplRespId gets ApplRespId, Tag 1353
func (m ApplicationRawDataReporting) GetApplRespId() (v string) {
	var f field.ApplResponseIDField
	if err := m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//  Tag 1180
func (m ApplicationRawDataReporting) GetApplId() (v string) {
	var f field.ApplIDField
	if err := m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//  Tag 1352
func (m ApplicationRawDataReporting) GetApplResendFlag() (v bool) {
	var f field.ApplResendFlagField
	if err := m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//  Tag 95
func (m ApplicationRawDataReporting) GetRawDataLength() (v int) {
	var f field.RawDataLengthField
	if err := m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//  Tag 96
func (m ApplicationRawDataReporting) GetRawData() (v string) {
	var f field.RawDataField
	if err := m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//  Tag 911
func (m ApplicationRawDataReporting) GetTotNumReports() (v int) {
	var f field.TotNumReportsField
	if err := m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

type NoApplSeqNums struct {
	*quickfix.Group
}

//SetRefApplID sets RefApplID, Tag 1181
func (m NoApplSeqNums) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

// tag 1350
func (m NoApplSeqNums) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

// tag 10055
func (m NoApplSeqNums) SetRawDataOffset(v int) {
	m.Set(NewRawDataOffsetField(v))
}

// tag 95
func (m NoApplSeqNums) SetRawDataLength(v int) {
	m.Set(field.NewRawDataLength(v))
}

// Tag 1350
func (m NoApplSeqNums) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// Tag 1181
func (m NoApplSeqNums) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// Tag 10055
func (m NoApplSeqNums) GetRawDataOffset() (v int, err quickfix.MessageRejectError) {
	var f RawDataOffsetField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// Tag 95
func (m NoApplSeqNums) GetRawDataLength() (v int, err quickfix.MessageRejectError) {
	var f field.RawDataLengthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

func (m ApplicationRawDataReporting) SetNoApplSeqNumsRepeatingGroup(f NoApplSeqNumsRepeatingGroup) {
	m.SetGroup(f)
}

// a repeating group, Tag 10054
type NoApplSeqNumsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMDEntryTypesRepeatingGroup returns an initialized, NoMDEntryTypesRepeatingGroup
func NewNoApplSeqNumsRepeatingGroup() NoApplSeqNumsRepeatingGroup {
	return NoApplSeqNumsRepeatingGroup{
		quickfix.NewRepeatingGroup(TagNoApplSeqNums,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ApplSeqNum), quickfix.GroupElement(tag.ApplLastSeqNum),
				quickfix.GroupElement(TagRawDataOffset), quickfix.GroupElement(tag.RawDataLength)})}
}

//Add create and append a new NoApplIDs to this group
func (m NoApplSeqNumsRepeatingGroup) Add() NoApplSeqNums {
	g := m.RepeatingGroup.Add()
	return NoApplSeqNums{g}
}

//Get returns the ith NoApplIDs in the NewNoApplIDsRepeatingGroup
func (m NoApplSeqNumsRepeatingGroup) Get(i int) NoApplSeqNums {
	return NoApplSeqNums{m.RepeatingGroup.Get(i)}
}

