//Package rfqrequest msg type = AH.
package rfqrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a RFQRequest wrapper for the generic Message type
type Message struct {
	message.Message
}

//RFQReqID is a required field for RFQRequest.
func (m Message) RFQReqID() (*field.RFQReqIDField, errors.MessageRejectError) {
	f := &field.RFQReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRFQReqID reads a RFQReqID from RFQRequest.
func (m Message) GetRFQReqID(f *field.RFQReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for RFQRequest.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from RFQRequest.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for RFQRequest.
func (m Message) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from RFQRequest.
func (m Message) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrivateQuote is a non-required field for RFQRequest.
func (m Message) PrivateQuote() (*field.PrivateQuoteField, errors.MessageRejectError) {
	f := &field.PrivateQuoteField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrivateQuote reads a PrivateQuote from RFQRequest.
func (m Message) GetPrivateQuote(f *field.PrivateQuoteField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RFQRequest.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RFQRequest.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds RFQRequest messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for RFQRequest.
func Builder(
	rfqreqid *field.RFQReqIDField,
	norelatedsym *field.NoRelatedSymField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("AH"))
	builder.Body().Set(rfqreqid)
	builder.Body().Set(norelatedsym)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "AH", r
}
