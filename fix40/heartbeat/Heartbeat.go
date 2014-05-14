//Package heartbeat msg type = 0.
package heartbeat

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a Heartbeat wrapper for the generic Message type
type Message struct {
	message.Message
}

//TestReqID is a non-required field for Heartbeat.
func (m Message) TestReqID() (*field.TestReqIDField, errors.MessageRejectError) {
	f := &field.TestReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTestReqID reads a TestReqID from Heartbeat.
func (m Message) GetTestReqID(f *field.TestReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds Heartbeat messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for Heartbeat.
func Builder() MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("0"))
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX40, "0", r
}
