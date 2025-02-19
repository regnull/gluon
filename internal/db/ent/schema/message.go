package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/ProtonMail/gluon/imap"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", imap.NewInternalMessageID()).Unique().Immutable(),
		field.String("RemoteID").Optional().Unique().GoType(imap.MessageID("")),
		field.Time("Date"),
		field.Int("Size"),
		field.String("Body"),
		field.String("BodyStructure"),
		field.String("Envelope"),
		field.Bool("Deleted").Default(false),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		// Apply message has many flags.
		edge.To("flags", MessageFlag.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),

		// Apply message has many UIDs.
		edge.From("UIDs", UID.Type).Ref("message"),
	}
}

func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("RemoteID"),
	}
}
