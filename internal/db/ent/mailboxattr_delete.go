// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ProtonMail/gluon/internal/db/ent/mailboxattr"
	"github.com/ProtonMail/gluon/internal/db/ent/predicate"
)

// MailboxAttrDelete is the builder for deleting a MailboxAttr entity.
type MailboxAttrDelete struct {
	config
	hooks    []Hook
	mutation *MailboxAttrMutation
}

// Where appends a list predicates to the MailboxAttrDelete builder.
func (mad *MailboxAttrDelete) Where(ps ...predicate.MailboxAttr) *MailboxAttrDelete {
	mad.mutation.Where(ps...)
	return mad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mad *MailboxAttrDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mad.hooks) == 0 {
		affected, err = mad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MailboxAttrMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mad.mutation = mutation
			affected, err = mad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mad.hooks) - 1; i >= 0; i-- {
			if mad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mad *MailboxAttrDelete) ExecX(ctx context.Context) int {
	n, err := mad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mad *MailboxAttrDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: mailboxattr.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mailboxattr.FieldID,
			},
		},
	}
	if ps := mad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// MailboxAttrDeleteOne is the builder for deleting a single MailboxAttr entity.
type MailboxAttrDeleteOne struct {
	mad *MailboxAttrDelete
}

// Exec executes the deletion query.
func (mado *MailboxAttrDeleteOne) Exec(ctx context.Context) error {
	n, err := mado.mad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mailboxattr.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mado *MailboxAttrDeleteOne) ExecX(ctx context.Context) {
	mado.mad.ExecX(ctx)
}
