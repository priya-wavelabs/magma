// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/propertytype"
)

// PropertyTypeDelete is the builder for deleting a PropertyType entity.
type PropertyTypeDelete struct {
	config
	predicates []predicate.PropertyType
}

// Where adds a new predicate to the delete builder.
func (ptd *PropertyTypeDelete) Where(ps ...predicate.PropertyType) *PropertyTypeDelete {
	ptd.predicates = append(ptd.predicates, ps...)
	return ptd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ptd *PropertyTypeDelete) Exec(ctx context.Context) (int, error) {
	return ptd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (ptd *PropertyTypeDelete) ExecX(ctx context.Context) int {
	n, err := ptd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ptd *PropertyTypeDelete) sqlExec(ctx context.Context) (int, error) {
	spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: propertytype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: propertytype.FieldID,
			},
		},
	}
	if ps := ptd.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ptd.driver, spec)
}

// PropertyTypeDeleteOne is the builder for deleting a single PropertyType entity.
type PropertyTypeDeleteOne struct {
	ptd *PropertyTypeDelete
}

// Exec executes the deletion query.
func (ptdo *PropertyTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ptdo.ptd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{propertytype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ptdo *PropertyTypeDeleteOne) ExecX(ctx context.Context) {
	ptdo.ptd.ExecX(ctx)
}
