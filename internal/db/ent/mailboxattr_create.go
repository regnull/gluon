// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ProtonMail/gluon/internal/db/ent/mailboxattr"
)

// MailboxAttrCreate is the builder for creating a MailboxAttr entity.
type MailboxAttrCreate struct {
	config
	mutation *MailboxAttrMutation
	hooks    []Hook
}

// SetValue sets the "Value" field.
func (mac *MailboxAttrCreate) SetValue(s string) *MailboxAttrCreate {
	mac.mutation.SetValue(s)
	return mac
}

// Mutation returns the MailboxAttrMutation object of the builder.
func (mac *MailboxAttrCreate) Mutation() *MailboxAttrMutation {
	return mac.mutation
}

// Save creates the MailboxAttr in the database.
func (mac *MailboxAttrCreate) Save(ctx context.Context) (*MailboxAttr, error) {
	var (
		err  error
		node *MailboxAttr
	)
	if len(mac.hooks) == 0 {
		if err = mac.check(); err != nil {
			return nil, err
		}
		node, err = mac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MailboxAttrMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mac.check(); err != nil {
				return nil, err
			}
			mac.mutation = mutation
			if node, err = mac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mac.hooks) - 1; i >= 0; i-- {
			if mac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MailboxAttr)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MailboxAttrMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mac *MailboxAttrCreate) SaveX(ctx context.Context) *MailboxAttr {
	v, err := mac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mac *MailboxAttrCreate) Exec(ctx context.Context) error {
	_, err := mac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mac *MailboxAttrCreate) ExecX(ctx context.Context) {
	if err := mac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mac *MailboxAttrCreate) check() error {
	if _, ok := mac.mutation.Value(); !ok {
		return &ValidationError{Name: "Value", err: errors.New(`ent: missing required field "MailboxAttr.Value"`)}
	}
	return nil
}

func (mac *MailboxAttrCreate) sqlSave(ctx context.Context) (*MailboxAttr, error) {
	_node, _spec := mac.createSpec()
	if err := sqlgraph.CreateNode(ctx, mac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mac *MailboxAttrCreate) createSpec() (*MailboxAttr, *sqlgraph.CreateSpec) {
	var (
		_node = &MailboxAttr{config: mac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mailboxattr.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mailboxattr.FieldID,
			},
		}
	)
	if value, ok := mac.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mailboxattr.FieldValue,
		})
		_node.Value = value
	}
	return _node, _spec
}

// MailboxAttrCreateBulk is the builder for creating many MailboxAttr entities in bulk.
type MailboxAttrCreateBulk struct {
	config
	builders []*MailboxAttrCreate
}

// Save creates the MailboxAttr entities in the database.
func (macb *MailboxAttrCreateBulk) Save(ctx context.Context) ([]*MailboxAttr, error) {
	specs := make([]*sqlgraph.CreateSpec, len(macb.builders))
	nodes := make([]*MailboxAttr, len(macb.builders))
	mutators := make([]Mutator, len(macb.builders))
	for i := range macb.builders {
		func(i int, root context.Context) {
			builder := macb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MailboxAttrMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, macb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, macb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, macb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (macb *MailboxAttrCreateBulk) SaveX(ctx context.Context) []*MailboxAttr {
	v, err := macb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (macb *MailboxAttrCreateBulk) Exec(ctx context.Context) error {
	_, err := macb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (macb *MailboxAttrCreateBulk) ExecX(ctx context.Context) {
	if err := macb.Exec(ctx); err != nil {
		panic(err)
	}
}
