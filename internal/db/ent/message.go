// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ProtonMail/gluon/imap"
	"github.com/ProtonMail/gluon/internal/db/ent/message"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID imap.InternalMessageID `json:"id,omitempty"`
	// RemoteID holds the value of the "RemoteID" field.
	RemoteID imap.MessageID `json:"RemoteID,omitempty"`
	// Date holds the value of the "Date" field.
	Date time.Time `json:"Date,omitempty"`
	// Size holds the value of the "Size" field.
	Size int `json:"Size,omitempty"`
	// Body holds the value of the "Body" field.
	Body string `json:"Body,omitempty"`
	// BodyStructure holds the value of the "BodyStructure" field.
	BodyStructure string `json:"BodyStructure,omitempty"`
	// Envelope holds the value of the "Envelope" field.
	Envelope string `json:"Envelope,omitempty"`
	// Deleted holds the value of the "Deleted" field.
	Deleted bool `json:"Deleted,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageQuery when eager-loading is set.
	Edges MessageEdges `json:"edges"`
}

// MessageEdges holds the relations/edges for other nodes in the graph.
type MessageEdges struct {
	// Flags holds the value of the flags edge.
	Flags []*MessageFlag `json:"flags,omitempty"`
	// UIDs holds the value of the UIDs edge.
	UIDs []*UID `json:"UIDs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FlagsOrErr returns the Flags value or an error if the edge
// was not loaded in eager-loading.
func (e MessageEdges) FlagsOrErr() ([]*MessageFlag, error) {
	if e.loadedTypes[0] {
		return e.Flags, nil
	}
	return nil, &NotLoadedError{edge: "flags"}
}

// UIDsOrErr returns the UIDs value or an error if the edge
// was not loaded in eager-loading.
func (e MessageEdges) UIDsOrErr() ([]*UID, error) {
	if e.loadedTypes[1] {
		return e.UIDs, nil
	}
	return nil, &NotLoadedError{edge: "UIDs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			values[i] = new(imap.InternalMessageID)
		case message.FieldDeleted:
			values[i] = new(sql.NullBool)
		case message.FieldSize:
			values[i] = new(sql.NullInt64)
		case message.FieldRemoteID, message.FieldBody, message.FieldBodyStructure, message.FieldEnvelope:
			values[i] = new(sql.NullString)
		case message.FieldDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Message", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			if value, ok := values[i].(*imap.InternalMessageID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case message.FieldRemoteID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field RemoteID", values[i])
			} else if value.Valid {
				m.RemoteID = imap.MessageID(value.String)
			}
		case message.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Date", values[i])
			} else if value.Valid {
				m.Date = value.Time
			}
		case message.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Size", values[i])
			} else if value.Valid {
				m.Size = int(value.Int64)
			}
		case message.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Body", values[i])
			} else if value.Valid {
				m.Body = value.String
			}
		case message.FieldBodyStructure:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field BodyStructure", values[i])
			} else if value.Valid {
				m.BodyStructure = value.String
			}
		case message.FieldEnvelope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Envelope", values[i])
			} else if value.Valid {
				m.Envelope = value.String
			}
		case message.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field Deleted", values[i])
			} else if value.Valid {
				m.Deleted = value.Bool
			}
		}
	}
	return nil
}

// QueryFlags queries the "flags" edge of the Message entity.
func (m *Message) QueryFlags() *MessageFlagQuery {
	return (&MessageClient{config: m.config}).QueryFlags(m)
}

// QueryUIDs queries the "UIDs" edge of the Message entity.
func (m *Message) QueryUIDs() *UIDQuery {
	return (&MessageClient{config: m.config}).QueryUIDs(m)
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return (&MessageClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("RemoteID=")
	builder.WriteString(fmt.Sprintf("%v", m.RemoteID))
	builder.WriteString(", ")
	builder.WriteString("Date=")
	builder.WriteString(m.Date.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("Size=")
	builder.WriteString(fmt.Sprintf("%v", m.Size))
	builder.WriteString(", ")
	builder.WriteString("Body=")
	builder.WriteString(m.Body)
	builder.WriteString(", ")
	builder.WriteString("BodyStructure=")
	builder.WriteString(m.BodyStructure)
	builder.WriteString(", ")
	builder.WriteString("Envelope=")
	builder.WriteString(m.Envelope)
	builder.WriteString(", ")
	builder.WriteString("Deleted=")
	builder.WriteString(fmt.Sprintf("%v", m.Deleted))
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message

func (m Messages) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
