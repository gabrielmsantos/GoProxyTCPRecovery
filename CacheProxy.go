package main

import (
	"fmt"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fix44"
	"github.com/quickfixgo/fix44/applicationmessagereport"
	"github.com/quickfixgo/fix44/applicationmessagerequestacknowledgment"
	"github.com/quickfixgo/fix44/applicationrawdatareporting"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
	"os"
	"strconv"
	"time"
)

type CacheProxy struct {
	*quickfix.MessageRouter
	*quickfix.Initiator
	inputChannel  chan quickfix.Messagable
	outputChannel chan quickfix.Messagable
	targetCompID  string
	senderCompID  string
}

func newCacheProxy(in chan quickfix.Messagable, out chan quickfix.Messagable) *CacheProxy {
	c := &CacheProxy{MessageRouter: quickfix.NewMessageRouter()}

	c.inputChannel = in
	c.outputChannel = out

	//Message cracker
	c.AddRoute(applicationmessagerequestacknowledgment.Route(c.onApplicationMessageRequestAck))
	c.AddRoute(applicationmessagereport.Route(c.onApplicationMessageReport))
	c.AddRoute(applicationrawdatareporting.Route(c.onApplicationRawDataReporting))
	return c
}

//quickfix.Application interface
func (c *CacheProxy) OnCreate(sessionID quickfix.SessionID) {
	c.targetCompID = sessionID.TargetCompID
	c.senderCompID = sessionID.SenderCompID
	return
}

func (c CacheProxy) OnLogon(sessionID quickfix.SessionID) {
	return
}

func (c CacheProxy) OnLogout(sessionID quickfix.SessionID)                           { return }
func (c CacheProxy) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID)     { return }
func (c CacheProxy) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) error { return nil }
func (c CacheProxy) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

//Use Message Cracker on Incoming Application Messages
func (c *CacheProxy) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	return c.Route(msg, sessionID)
}

func (c *CacheProxy) SendToClient(msg quickfix.Messagable) {
	c.outputChannel <- msg
}

// Handling 35=BX here
func (m *CacheProxy) onApplicationMessageRequestAck(msg applicationmessagerequestacknowledgment.ApplicationMessageRequestAcknowledgment, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	fmt.Println("Proxy: Received 35=BX: " + msg.Message.String())

	msCopied := applicationmessagerequestacknowledgment.ApplicationMessageRequestAcknowledgment{}
	msCopied.Message = new(quickfix.Message)
	msg.Message.CopyInto(msCopied.Message)

	msCopied.Header.Header = &msCopied.Message.Header
	msCopied.Body = &msCopied.Message.Body
	msCopied.Trailer.Trailer = &msCopied.Message.Trailer

	msCopied.Message.Unsorted = true
	m.SendToClient(msCopied)

	return
}

// Handling 35=BY here
func (m *CacheProxy) onApplicationMessageReport(msg applicationmessagereport.ApplicationMessageReport, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	fmt.Println("Proxy: Received 35=BY: " + msg.Message.String())

	msCopied := applicationmessagereport.ApplicationMessageReport{}
	msCopied.Message = new(quickfix.Message)
	msg.Message.CopyInto(msCopied.Message)

	msCopied.Header.Header = &msCopied.Message.Header
	msCopied.Body = &msCopied.Message.Body
	msCopied.Trailer.Trailer = &msCopied.Message.Trailer

	msCopied.Message.Unsorted = true
	m.SendToClient(msCopied)

	return
}

// Handling 35=BY here
func (m *CacheProxy) onApplicationRawDataReporting(msg applicationrawdatareporting.ApplicationRawDataReporting, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	fmt.Println("Proxy: Received 35=URDR: " + msg.Message.String())

	/*msCopied := applicationrawdatareporting.ApplicationRawDataReporting{}
	msCopied.Message = new(quickfix.Message)
	msg.Message.CopyInto(msCopied.Message)

	msCopied.Header.Header = &msCopied.Message.Header
	msCopied.Body = &msCopied.Message.Body
	msCopied.Trailer.Trailer= &msCopied.Message.Trailer

	//msCopied.Message.Unsorted = true
	m.SendToClient(msCopied.Message)
	*/

	msNew := applicationrawdatareporting.ApplicationRawDataReporting{}
	msNew.Message = quickfix.NewMessage()
	msNew.Header = fix44.NewHeader(&msNew.Message.Header)
	msNew.Body = &msNew.Message.Body
	msNew.Trailer.Trailer = &msNew.Message.Trailer
	msNew.Header.Set(field.NewMsgType("URDR"))

	mgroup := applicationrawdatareporting.NewNoApplSeqNumsRepeatingGroup()
	var g applicationrawdatareporting.NoApplSeqNums
	total := 0
	vt := quickfix.TagValue{}
	df := msg.Message.GetFields()
	for _, vt = range df {
		switch vt.GetTag() {
		case tag.ApplReqID:
			msNew.SetApplReqID(string(vt.GetValue()))
		case tag.ApplRespID:
			msNew.SetApplRespID(string(vt.GetValue()))
		case tag.ApplID:
			msNew.SetApplID(string(vt.GetValue()))
		case tag.ApplResendFlag:
			flag, _ := strconv.ParseBool(string(vt.GetValue()))
			msNew.SetApplResendFlag(flag)
		case tag.RawDataLength:
			mInt32, _ := strconv.ParseInt(string(vt.GetValue()), 10, 32)
			if total > 0 {
				g.SetRawDataLength(int(mInt32))
				total--
			} else {
				msNew.SetRawDataLength(int(mInt32))
			}
		case tag.RawData:
			msNew.SetRawData(string(vt.GetValue()))
		case tag.TotNumReports:
			mInt32, _ := strconv.ParseInt(string(vt.GetValue()), 10, 32)
			msNew.SetTotNumReports(int(mInt32))
		case tag.NoApplSeqNums:
			fmt.Println("URDR: Processing repeating group!")
			mInt32, _ := strconv.ParseInt(string(vt.GetValue()), 10, 32)
			total = int(mInt32)
		case tag.ApplSeqNum:
			g = mgroup.Add()
			g.SetApplSeqNum(string(vt.GetValue()))
		case tag.ApplLastSeqNum:
			mInt32, _ := strconv.ParseInt(string(vt.GetValue()), 10, 32)
			g.SetApplLastSeqNum(int(mInt32))
		case tag.RawDataOffset:
			mInt32, _ := strconv.ParseInt(string(vt.GetValue()), 10, 32)
			g.SetRawDataOffset(int(mInt32))
		}
	}
	msNew.SetNoApplSeqNums(mgroup)
	m.SendToClient(msNew)
	return
}

func (m *CacheProxy) Start(cfgFileName string) {
	cfg, err := os.Open(cfgFileName)

	if err != nil {
		fmt.Printf("Proxy: Error opening %v, %v\n", cfgFileName, err)
		return
	}
	defer cfg.Close()

	appSettings, err := quickfix.ParseSettings(cfg)
	if err != nil {
		fmt.Println("Proxy: Error reading cfg,", err)
		return
	}

	logFactory := quickfix.NewScreenLogFactory()

	m.Initiator, err = quickfix.NewInitiator(m, quickfix.NewMemoryStoreFactory(), appSettings, logFactory)
	if err != nil {
		fmt.Printf("Proxy: Unable to create Initiator: %s\n", err)
		return
	}

	err = m.Initiator.Start()
	if err != nil {
		fmt.Printf("Proxy: Unable to start Initiator: %s\n", err)
		return
	}
	fmt.Println("Proxy: Initiator started: ")

	// Wait for client messages to forward to B3
	m.waitForClientMessages()
}

func (m *CacheProxy) waitForClientMessages() {
	for msg := range m.inputChannel {

		mp := msg.ToMessage()

		mp.Header.SetString(tag.SenderCompID, m.senderCompID)
		mp.Header.SetString(tag.TargetCompID, m.targetCompID)
		msg.ToMessage().Unsorted = true
		time.Sleep(40 * time.Millisecond)

		err2 := quickfix.Send(msg)
		if err2 != nil {
			fmt.Printf("Proxy: Unable to send message to B3: %s\n", err2)
		}
	}
}

func (m *CacheProxy) Stop() {
	m.Initiator.Stop()
}
