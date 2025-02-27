// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/ProtonMail/gluon/imap"
	"github.com/ProtonMail/gluon/internal/db/ent/mailboxflag"
)

// MailboxFlag is the model entity for the MailboxFlag schema.
type MailboxFlag struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "Value" field.
	Value         string `json:"Value,omitempty"`
	mailbox_flags *imap.InternalMailboxID
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MailboxFlag) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mailboxflag.FieldID:
			values[i] = new(sql.NullInt64)
		case mailboxflag.FieldValue:
			values[i] = new(sql.NullString)
		case mailboxflag.ForeignKeys[0]: // mailbox_flags
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MailboxFlag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MailboxFlag fields.
func (mf *MailboxFlag) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mailboxflag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mf.ID = int(value.Int64)
		case mailboxflag.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Value", values[i])
			} else if value.Valid {
				mf.Value = value.String
			}
		case mailboxflag.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mailbox_flags", values[i])
			} else if value.Valid {
				mf.mailbox_flags = new(imap.InternalMailboxID)
				*mf.mailbox_flags = imap.InternalMailboxID(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MailboxFlag.
// Note that you need to call MailboxFlag.Unwrap() before calling this method if this MailboxFlag
// was returned from a transaction, and the transaction was committed or rolled back.
func (mf *MailboxFlag) Update() *MailboxFlagUpdateOne {
	return (&MailboxFlagClient{config: mf.config}).UpdateOne(mf)
}

// Unwrap unwraps the MailboxFlag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mf *MailboxFlag) Unwrap() *MailboxFlag {
	_tx, ok := mf.config.driver.(*txDriver)
	if !ok {
		panic("ent: MailboxFlag is not a transactional entity")
	}
	mf.config.driver = _tx.drv
	return mf
}

// String implements the fmt.Stringer.
func (mf *MailboxFlag) String() string {
	var builder strings.Builder
	builder.WriteString("MailboxFlag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mf.ID))
	builder.WriteString("Value=")
	builder.WriteString(mf.Value)
	builder.WriteByte(')')
	return builder.String()
}

// MailboxFlags is a parsable slice of MailboxFlag.
type MailboxFlags []*MailboxFlag

func (mf MailboxFlags) config(cfg config) {
	for _i := range mf {
		mf[_i].config = cfg
	}
}
