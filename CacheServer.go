package main

import (
	"fmt"
	"github.com/quickfixgo/fix44/applicationmessagerequest"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
	"os"
	"strings"
)

type CacheServer struct {
	*quickfix.MessageRouter
	*quickfix.Acceptor
	inputChannel  chan quickfix.Messagable
	outputChannel chan quickfix.Messagable
	senderCompId  string
}

func newCacheServer(in chan quickfix.Messagable, out chan quickfix.Messagable) *CacheServer {
	c := &CacheServer{MessageRouter: quickfix.NewMessageRouter()}

	c.inputChannel = in
	c.outputChannel = out

	c.AddRoute(applicationmessagerequest.Route(c.onApplicationMessageRequest))

	return c
}

func (c *CacheServer) OnCreate(sessionID quickfix.SessionID) {
	c.senderCompId = sessionID.SenderCompID
	return
}

func (c CacheServer) OnLogon(sessionID quickfix.SessionID) {
	return
}

func (c CacheServer) OnLogout(sessionID quickfix.SessionID)                           { return }
func (c CacheServer) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID)     { return }
func (c CacheServer) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) error { return nil }
func (c CacheServer) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	return nil
}

//Use Message Cracker on Incoming Application Messages
func (c *CacheServer) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	return c.Route(msg, sessionID)
}

func (m *CacheServer) onApplicationMessageRequest(msg applicationmessagerequest.ApplicationMessageRequest, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	// Put Session identifier
	v, err := msg.GetApplReqID()
	if err != nil {
		fmt.Printf("Error getting ApplReqId for  %v, %v\n", v, err)
	}

	msg.SetApplReqID(v + SEP + sessionID.TargetCompID)

	m.outputChannel <- msg
	return
}

func (m *CacheServer) Start(cfgFileName string) {
	cfg, err := os.Open(cfgFileName)
	if err != nil {
		fmt.Printf("Error opening %v, %v\n", cfgFileName, err)
		return
	}
	defer cfg.Close()

	appSettings, err := quickfix.ParseSettings(cfg)
	if err != nil {
		fmt.Println("Error reading cfg,", err)
		return
	}

	logFactory := quickfix.NewScreenLogFactory()

	m.Acceptor, err = quickfix.NewAcceptor(m, quickfix.NewMemoryStoreFactory(), appSettings, logFactory)
	if err != nil {
		fmt.Printf("Unable to create Acceptor: %s\n", err)
		return
	}

	err = m.Acceptor.Start()
	if err != nil {
		fmt.Printf("Unable to start Acceptor: %s\n", err)
		return
	}
	fmt.Println("FIXServer: Acceptor started.")

	// Waiting for messages coming from Proxy
	m.waitForProxyMessages()
}

func (m *CacheServer) waitForProxyMessages() {
	for msg := range m.inputChannel {
		resp := new(quickfix.Message)
		resp = msg.ToMessage()
		fmt.Println("FIXServer: Start handling: " + resp.String())
		var vArray []string
		if resp.Body.Has(tag.ApplReqID) {
			vTag, _ := resp.Body.GetString(tag.ApplReqID)
			vArray = strings.Split(vTag, SEP)
			resp.Body.SetString(tag.ApplReqID, vArray[0])
		}

		resp.Header.SetString(tag.SenderCompID, m.senderCompId)
		if len(vArray) > 1 {
			resp.Header.SetString(tag.TargetCompID, vArray[1])
		} else {
			fmt.Println("FIXServer: Not setting TargetCompID. Tag 56 not found!!!")
		}

		err := quickfix.Send(resp)
		fmt.Println("FIXServer: Sent to client: " + resp.String())
		if err != nil {
			fmt.Printf("FIXServer: Unable to send message to client: %s\n", err)
		}
	}
}

func (m *CacheServer) Stop() {
	m.Acceptor.Stop()
}
