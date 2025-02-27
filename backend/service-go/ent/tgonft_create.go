// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"toucan-leaderboard/ent/tgonft"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TGoNFTCreate is the builder for creating a TGoNFT entity.
type TGoNFTCreate struct {
	config
	mutation *TGoNFTMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetWalletPub sets the "wallet_pub" field.
func (tnc *TGoNFTCreate) SetWalletPub(s string) *TGoNFTCreate {
	tnc.mutation.SetWalletPub(s)
	return tnc
}

// SetNillableWalletPub sets the "wallet_pub" field if the given value is not nil.
func (tnc *TGoNFTCreate) SetNillableWalletPub(s *string) *TGoNFTCreate {
	if s != nil {
		tnc.SetWalletPub(*s)
	}
	return tnc
}

// SetRankType sets the "rank_type" field.
func (tnc *TGoNFTCreate) SetRankType(i int) *TGoNFTCreate {
	tnc.mutation.SetRankType(i)
	return tnc
}

// SetNillableRankType sets the "rank_type" field if the given value is not nil.
func (tnc *TGoNFTCreate) SetNillableRankType(i *int) *TGoNFTCreate {
	if i != nil {
		tnc.SetRankType(*i)
	}
	return tnc
}

// SetRankYear sets the "rank_year" field.
func (tnc *TGoNFTCreate) SetRankYear(i int) *TGoNFTCreate {
	tnc.mutation.SetRankYear(i)
	return tnc
}

// SetNillableRankYear sets the "rank_year" field if the given value is not nil.
func (tnc *TGoNFTCreate) SetNillableRankYear(i *int) *TGoNFTCreate {
	if i != nil {
		tnc.SetRankYear(*i)
	}
	return tnc
}

// SetRankSeason sets the "rank_season" field.
func (tnc *TGoNFTCreate) SetRankSeason(i int) *TGoNFTCreate {
	tnc.mutation.SetRankSeason(i)
	return tnc
}

// SetNillableRankSeason sets the "rank_season" field if the given value is not nil.
func (tnc *TGoNFTCreate) SetNillableRankSeason(i *int) *TGoNFTCreate {
	if i != nil {
		tnc.SetRankSeason(*i)
	}
	return tnc
}

// SetRank sets the "rank" field.
func (tnc *TGoNFTCreate) SetRank(i int) *TGoNFTCreate {
	tnc.mutation.SetRank(i)
	return tnc
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (tnc *TGoNFTCreate) SetNillableRank(i *int) *TGoNFTCreate {
	if i != nil {
		tnc.SetRank(*i)
	}
	return tnc
}

// SetMintTx sets the "mint_tx" field.
func (tnc *TGoNFTCreate) SetMintTx(s string) *TGoNFTCreate {
	tnc.mutation.SetMintTx(s)
	return tnc
}

// SetNillableMintTx sets the "mint_tx" field if the given value is not nil.
func (tnc *TGoNFTCreate) SetNillableMintTx(s *string) *TGoNFTCreate {
	if s != nil {
		tnc.SetMintTx(*s)
	}
	return tnc
}

// SetMtime sets the "mtime" field.
func (tnc *TGoNFTCreate) SetMtime(t time.Time) *TGoNFTCreate {
	tnc.mutation.SetMtime(t)
	return tnc
}

// SetNillableMtime sets the "mtime" field if the given value is not nil.
func (tnc *TGoNFTCreate) SetNillableMtime(t *time.Time) *TGoNFTCreate {
	if t != nil {
		tnc.SetMtime(*t)
	}
	return tnc
}

// SetCtime sets the "ctime" field.
func (tnc *TGoNFTCreate) SetCtime(t time.Time) *TGoNFTCreate {
	tnc.mutation.SetCtime(t)
	return tnc
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (tnc *TGoNFTCreate) SetNillableCtime(t *time.Time) *TGoNFTCreate {
	if t != nil {
		tnc.SetCtime(*t)
	}
	return tnc
}

// SetID sets the "id" field.
func (tnc *TGoNFTCreate) SetID(u uint64) *TGoNFTCreate {
	tnc.mutation.SetID(u)
	return tnc
}

// Mutation returns the TGoNFTMutation object of the builder.
func (tnc *TGoNFTCreate) Mutation() *TGoNFTMutation {
	return tnc.mutation
}

// Save creates the TGoNFT in the database.
func (tnc *TGoNFTCreate) Save(ctx context.Context) (*TGoNFT, error) {
	var (
		err  error
		node *TGoNFT
	)
	tnc.defaults()
	if len(tnc.hooks) == 0 {
		if err = tnc.check(); err != nil {
			return nil, err
		}
		node, err = tnc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TGoNFTMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tnc.check(); err != nil {
				return nil, err
			}
			tnc.mutation = mutation
			if node, err = tnc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tnc.hooks) - 1; i >= 0; i-- {
			if tnc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tnc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tnc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TGoNFT)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TGoNFTMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tnc *TGoNFTCreate) SaveX(ctx context.Context) *TGoNFT {
	v, err := tnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tnc *TGoNFTCreate) Exec(ctx context.Context) error {
	_, err := tnc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tnc *TGoNFTCreate) ExecX(ctx context.Context) {
	if err := tnc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tnc *TGoNFTCreate) defaults() {
	if _, ok := tnc.mutation.WalletPub(); !ok {
		v := tgonft.DefaultWalletPub
		tnc.mutation.SetWalletPub(v)
	}
	if _, ok := tnc.mutation.RankType(); !ok {
		v := tgonft.DefaultRankType
		tnc.mutation.SetRankType(v)
	}
	if _, ok := tnc.mutation.RankYear(); !ok {
		v := tgonft.DefaultRankYear
		tnc.mutation.SetRankYear(v)
	}
	if _, ok := tnc.mutation.RankSeason(); !ok {
		v := tgonft.DefaultRankSeason
		tnc.mutation.SetRankSeason(v)
	}
	if _, ok := tnc.mutation.Rank(); !ok {
		v := tgonft.DefaultRank
		tnc.mutation.SetRank(v)
	}
	if _, ok := tnc.mutation.MintTx(); !ok {
		v := tgonft.DefaultMintTx
		tnc.mutation.SetMintTx(v)
	}
	if _, ok := tnc.mutation.Mtime(); !ok {
		v := tgonft.DefaultMtime
		tnc.mutation.SetMtime(v)
	}
	if _, ok := tnc.mutation.Ctime(); !ok {
		v := tgonft.DefaultCtime
		tnc.mutation.SetCtime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tnc *TGoNFTCreate) check() error {
	if _, ok := tnc.mutation.WalletPub(); !ok {
		return &ValidationError{Name: "wallet_pub", err: errors.New(`ent: missing required field "TGoNFT.wallet_pub"`)}
	}
	if _, ok := tnc.mutation.RankType(); !ok {
		return &ValidationError{Name: "rank_type", err: errors.New(`ent: missing required field "TGoNFT.rank_type"`)}
	}
	if _, ok := tnc.mutation.RankYear(); !ok {
		return &ValidationError{Name: "rank_year", err: errors.New(`ent: missing required field "TGoNFT.rank_year"`)}
	}
	if _, ok := tnc.mutation.RankSeason(); !ok {
		return &ValidationError{Name: "rank_season", err: errors.New(`ent: missing required field "TGoNFT.rank_season"`)}
	}
	if _, ok := tnc.mutation.Rank(); !ok {
		return &ValidationError{Name: "rank", err: errors.New(`ent: missing required field "TGoNFT.rank"`)}
	}
	if _, ok := tnc.mutation.MintTx(); !ok {
		return &ValidationError{Name: "mint_tx", err: errors.New(`ent: missing required field "TGoNFT.mint_tx"`)}
	}
	if _, ok := tnc.mutation.Mtime(); !ok {
		return &ValidationError{Name: "mtime", err: errors.New(`ent: missing required field "TGoNFT.mtime"`)}
	}
	if _, ok := tnc.mutation.Ctime(); !ok {
		return &ValidationError{Name: "ctime", err: errors.New(`ent: missing required field "TGoNFT.ctime"`)}
	}
	return nil
}

func (tnc *TGoNFTCreate) sqlSave(ctx context.Context) (*TGoNFT, error) {
	_node, _spec := tnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tnc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (tnc *TGoNFTCreate) createSpec() (*TGoNFT, *sqlgraph.CreateSpec) {
	var (
		_node = &TGoNFT{config: tnc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tgonft.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tgonft.FieldID,
			},
		}
	)
	_spec.OnConflict = tnc.conflict
	if id, ok := tnc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tnc.mutation.WalletPub(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgonft.FieldWalletPub,
		})
		_node.WalletPub = value
	}
	if value, ok := tnc.mutation.RankType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankType,
		})
		_node.RankType = value
	}
	if value, ok := tnc.mutation.RankYear(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankYear,
		})
		_node.RankYear = value
	}
	if value, ok := tnc.mutation.RankSeason(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRankSeason,
		})
		_node.RankSeason = value
	}
	if value, ok := tnc.mutation.Rank(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tgonft.FieldRank,
		})
		_node.Rank = value
	}
	if value, ok := tnc.mutation.MintTx(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tgonft.FieldMintTx,
		})
		_node.MintTx = value
	}
	if value, ok := tnc.mutation.Mtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgonft.FieldMtime,
		})
		_node.Mtime = value
	}
	if value, ok := tnc.mutation.Ctime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tgonft.FieldCtime,
		})
		_node.Ctime = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TGoNFT.Create().
//		SetWalletPub(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TGoNFTUpsert) {
//			SetWalletPub(v+v).
//		}).
//		Exec(ctx)
//
func (tnc *TGoNFTCreate) OnConflict(opts ...sql.ConflictOption) *TGoNFTUpsertOne {
	tnc.conflict = opts
	return &TGoNFTUpsertOne{
		create: tnc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TGoNFT.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tnc *TGoNFTCreate) OnConflictColumns(columns ...string) *TGoNFTUpsertOne {
	tnc.conflict = append(tnc.conflict, sql.ConflictColumns(columns...))
	return &TGoNFTUpsertOne{
		create: tnc,
	}
}

type (
	// TGoNFTUpsertOne is the builder for "upsert"-ing
	//  one TGoNFT node.
	TGoNFTUpsertOne struct {
		create *TGoNFTCreate
	}

	// TGoNFTUpsert is the "OnConflict" setter.
	TGoNFTUpsert struct {
		*sql.UpdateSet
	}
)

// SetWalletPub sets the "wallet_pub" field.
func (u *TGoNFTUpsert) SetWalletPub(v string) *TGoNFTUpsert {
	u.Set(tgonft.FieldWalletPub, v)
	return u
}

// UpdateWalletPub sets the "wallet_pub" field to the value that was provided on create.
func (u *TGoNFTUpsert) UpdateWalletPub() *TGoNFTUpsert {
	u.SetExcluded(tgonft.FieldWalletPub)
	return u
}

// SetRankType sets the "rank_type" field.
func (u *TGoNFTUpsert) SetRankType(v int) *TGoNFTUpsert {
	u.Set(tgonft.FieldRankType, v)
	return u
}

// UpdateRankType sets the "rank_type" field to the value that was provided on create.
func (u *TGoNFTUpsert) UpdateRankType() *TGoNFTUpsert {
	u.SetExcluded(tgonft.FieldRankType)
	return u
}

// AddRankType adds v to the "rank_type" field.
func (u *TGoNFTUpsert) AddRankType(v int) *TGoNFTUpsert {
	u.Add(tgonft.FieldRankType, v)
	return u
}

// SetRankYear sets the "rank_year" field.
func (u *TGoNFTUpsert) SetRankYear(v int) *TGoNFTUpsert {
	u.Set(tgonft.FieldRankYear, v)
	return u
}

// UpdateRankYear sets the "rank_year" field to the value that was provided on create.
func (u *TGoNFTUpsert) UpdateRankYear() *TGoNFTUpsert {
	u.SetExcluded(tgonft.FieldRankYear)
	return u
}

// AddRankYear adds v to the "rank_year" field.
func (u *TGoNFTUpsert) AddRankYear(v int) *TGoNFTUpsert {
	u.Add(tgonft.FieldRankYear, v)
	return u
}

// SetRankSeason sets the "rank_season" field.
func (u *TGoNFTUpsert) SetRankSeason(v int) *TGoNFTUpsert {
	u.Set(tgonft.FieldRankSeason, v)
	return u
}

// UpdateRankSeason sets the "rank_season" field to the value that was provided on create.
func (u *TGoNFTUpsert) UpdateRankSeason() *TGoNFTUpsert {
	u.SetExcluded(tgonft.FieldRankSeason)
	return u
}

// AddRankSeason adds v to the "rank_season" field.
func (u *TGoNFTUpsert) AddRankSeason(v int) *TGoNFTUpsert {
	u.Add(tgonft.FieldRankSeason, v)
	return u
}

// SetRank sets the "rank" field.
func (u *TGoNFTUpsert) SetRank(v int) *TGoNFTUpsert {
	u.Set(tgonft.FieldRank, v)
	return u
}

// UpdateRank sets the "rank" field to the value that was provided on create.
func (u *TGoNFTUpsert) UpdateRank() *TGoNFTUpsert {
	u.SetExcluded(tgonft.FieldRank)
	return u
}

// AddRank adds v to the "rank" field.
func (u *TGoNFTUpsert) AddRank(v int) *TGoNFTUpsert {
	u.Add(tgonft.FieldRank, v)
	return u
}

// SetMintTx sets the "mint_tx" field.
func (u *TGoNFTUpsert) SetMintTx(v string) *TGoNFTUpsert {
	u.Set(tgonft.FieldMintTx, v)
	return u
}

// UpdateMintTx sets the "mint_tx" field to the value that was provided on create.
func (u *TGoNFTUpsert) UpdateMintTx() *TGoNFTUpsert {
	u.SetExcluded(tgonft.FieldMintTx)
	return u
}

// SetMtime sets the "mtime" field.
func (u *TGoNFTUpsert) SetMtime(v time.Time) *TGoNFTUpsert {
	u.Set(tgonft.FieldMtime, v)
	return u
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *TGoNFTUpsert) UpdateMtime() *TGoNFTUpsert {
	u.SetExcluded(tgonft.FieldMtime)
	return u
}

// SetCtime sets the "ctime" field.
func (u *TGoNFTUpsert) SetCtime(v time.Time) *TGoNFTUpsert {
	u.Set(tgonft.FieldCtime, v)
	return u
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *TGoNFTUpsert) UpdateCtime() *TGoNFTUpsert {
	u.SetExcluded(tgonft.FieldCtime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TGoNFT.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tgonft.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TGoNFTUpsertOne) UpdateNewValues() *TGoNFTUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(tgonft.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TGoNFT.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TGoNFTUpsertOne) Ignore() *TGoNFTUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TGoNFTUpsertOne) DoNothing() *TGoNFTUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TGoNFTCreate.OnConflict
// documentation for more info.
func (u *TGoNFTUpsertOne) Update(set func(*TGoNFTUpsert)) *TGoNFTUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TGoNFTUpsert{UpdateSet: update})
	}))
	return u
}

// SetWalletPub sets the "wallet_pub" field.
func (u *TGoNFTUpsertOne) SetWalletPub(v string) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetWalletPub(v)
	})
}

// UpdateWalletPub sets the "wallet_pub" field to the value that was provided on create.
func (u *TGoNFTUpsertOne) UpdateWalletPub() *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateWalletPub()
	})
}

// SetRankType sets the "rank_type" field.
func (u *TGoNFTUpsertOne) SetRankType(v int) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetRankType(v)
	})
}

// AddRankType adds v to the "rank_type" field.
func (u *TGoNFTUpsertOne) AddRankType(v int) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.AddRankType(v)
	})
}

// UpdateRankType sets the "rank_type" field to the value that was provided on create.
func (u *TGoNFTUpsertOne) UpdateRankType() *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateRankType()
	})
}

// SetRankYear sets the "rank_year" field.
func (u *TGoNFTUpsertOne) SetRankYear(v int) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetRankYear(v)
	})
}

// AddRankYear adds v to the "rank_year" field.
func (u *TGoNFTUpsertOne) AddRankYear(v int) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.AddRankYear(v)
	})
}

// UpdateRankYear sets the "rank_year" field to the value that was provided on create.
func (u *TGoNFTUpsertOne) UpdateRankYear() *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateRankYear()
	})
}

// SetRankSeason sets the "rank_season" field.
func (u *TGoNFTUpsertOne) SetRankSeason(v int) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetRankSeason(v)
	})
}

// AddRankSeason adds v to the "rank_season" field.
func (u *TGoNFTUpsertOne) AddRankSeason(v int) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.AddRankSeason(v)
	})
}

// UpdateRankSeason sets the "rank_season" field to the value that was provided on create.
func (u *TGoNFTUpsertOne) UpdateRankSeason() *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateRankSeason()
	})
}

// SetRank sets the "rank" field.
func (u *TGoNFTUpsertOne) SetRank(v int) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetRank(v)
	})
}

// AddRank adds v to the "rank" field.
func (u *TGoNFTUpsertOne) AddRank(v int) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.AddRank(v)
	})
}

// UpdateRank sets the "rank" field to the value that was provided on create.
func (u *TGoNFTUpsertOne) UpdateRank() *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateRank()
	})
}

// SetMintTx sets the "mint_tx" field.
func (u *TGoNFTUpsertOne) SetMintTx(v string) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetMintTx(v)
	})
}

// UpdateMintTx sets the "mint_tx" field to the value that was provided on create.
func (u *TGoNFTUpsertOne) UpdateMintTx() *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateMintTx()
	})
}

// SetMtime sets the "mtime" field.
func (u *TGoNFTUpsertOne) SetMtime(v time.Time) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetMtime(v)
	})
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *TGoNFTUpsertOne) UpdateMtime() *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateMtime()
	})
}

// SetCtime sets the "ctime" field.
func (u *TGoNFTUpsertOne) SetCtime(v time.Time) *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetCtime(v)
	})
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *TGoNFTUpsertOne) UpdateCtime() *TGoNFTUpsertOne {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateCtime()
	})
}

// Exec executes the query.
func (u *TGoNFTUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TGoNFTCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TGoNFTUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TGoNFTUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TGoNFTUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TGoNFTCreateBulk is the builder for creating many TGoNFT entities in bulk.
type TGoNFTCreateBulk struct {
	config
	builders []*TGoNFTCreate
	conflict []sql.ConflictOption
}

// Save creates the TGoNFT entities in the database.
func (tncb *TGoNFTCreateBulk) Save(ctx context.Context) ([]*TGoNFT, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tncb.builders))
	nodes := make([]*TGoNFT, len(tncb.builders))
	mutators := make([]Mutator, len(tncb.builders))
	for i := range tncb.builders {
		func(i int, root context.Context) {
			builder := tncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TGoNFTMutation)
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
					_, err = mutators[i+1].Mutate(root, tncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, tncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tncb *TGoNFTCreateBulk) SaveX(ctx context.Context) []*TGoNFT {
	v, err := tncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tncb *TGoNFTCreateBulk) Exec(ctx context.Context) error {
	_, err := tncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tncb *TGoNFTCreateBulk) ExecX(ctx context.Context) {
	if err := tncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TGoNFT.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TGoNFTUpsert) {
//			SetWalletPub(v+v).
//		}).
//		Exec(ctx)
//
func (tncb *TGoNFTCreateBulk) OnConflict(opts ...sql.ConflictOption) *TGoNFTUpsertBulk {
	tncb.conflict = opts
	return &TGoNFTUpsertBulk{
		create: tncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TGoNFT.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tncb *TGoNFTCreateBulk) OnConflictColumns(columns ...string) *TGoNFTUpsertBulk {
	tncb.conflict = append(tncb.conflict, sql.ConflictColumns(columns...))
	return &TGoNFTUpsertBulk{
		create: tncb,
	}
}

// TGoNFTUpsertBulk is the builder for "upsert"-ing
// a bulk of TGoNFT nodes.
type TGoNFTUpsertBulk struct {
	create *TGoNFTCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TGoNFT.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tgonft.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TGoNFTUpsertBulk) UpdateNewValues() *TGoNFTUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(tgonft.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TGoNFT.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TGoNFTUpsertBulk) Ignore() *TGoNFTUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TGoNFTUpsertBulk) DoNothing() *TGoNFTUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TGoNFTCreateBulk.OnConflict
// documentation for more info.
func (u *TGoNFTUpsertBulk) Update(set func(*TGoNFTUpsert)) *TGoNFTUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TGoNFTUpsert{UpdateSet: update})
	}))
	return u
}

// SetWalletPub sets the "wallet_pub" field.
func (u *TGoNFTUpsertBulk) SetWalletPub(v string) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetWalletPub(v)
	})
}

// UpdateWalletPub sets the "wallet_pub" field to the value that was provided on create.
func (u *TGoNFTUpsertBulk) UpdateWalletPub() *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateWalletPub()
	})
}

// SetRankType sets the "rank_type" field.
func (u *TGoNFTUpsertBulk) SetRankType(v int) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetRankType(v)
	})
}

// AddRankType adds v to the "rank_type" field.
func (u *TGoNFTUpsertBulk) AddRankType(v int) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.AddRankType(v)
	})
}

// UpdateRankType sets the "rank_type" field to the value that was provided on create.
func (u *TGoNFTUpsertBulk) UpdateRankType() *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateRankType()
	})
}

// SetRankYear sets the "rank_year" field.
func (u *TGoNFTUpsertBulk) SetRankYear(v int) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetRankYear(v)
	})
}

// AddRankYear adds v to the "rank_year" field.
func (u *TGoNFTUpsertBulk) AddRankYear(v int) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.AddRankYear(v)
	})
}

// UpdateRankYear sets the "rank_year" field to the value that was provided on create.
func (u *TGoNFTUpsertBulk) UpdateRankYear() *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateRankYear()
	})
}

// SetRankSeason sets the "rank_season" field.
func (u *TGoNFTUpsertBulk) SetRankSeason(v int) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetRankSeason(v)
	})
}

// AddRankSeason adds v to the "rank_season" field.
func (u *TGoNFTUpsertBulk) AddRankSeason(v int) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.AddRankSeason(v)
	})
}

// UpdateRankSeason sets the "rank_season" field to the value that was provided on create.
func (u *TGoNFTUpsertBulk) UpdateRankSeason() *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateRankSeason()
	})
}

// SetRank sets the "rank" field.
func (u *TGoNFTUpsertBulk) SetRank(v int) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetRank(v)
	})
}

// AddRank adds v to the "rank" field.
func (u *TGoNFTUpsertBulk) AddRank(v int) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.AddRank(v)
	})
}

// UpdateRank sets the "rank" field to the value that was provided on create.
func (u *TGoNFTUpsertBulk) UpdateRank() *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateRank()
	})
}

// SetMintTx sets the "mint_tx" field.
func (u *TGoNFTUpsertBulk) SetMintTx(v string) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetMintTx(v)
	})
}

// UpdateMintTx sets the "mint_tx" field to the value that was provided on create.
func (u *TGoNFTUpsertBulk) UpdateMintTx() *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateMintTx()
	})
}

// SetMtime sets the "mtime" field.
func (u *TGoNFTUpsertBulk) SetMtime(v time.Time) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetMtime(v)
	})
}

// UpdateMtime sets the "mtime" field to the value that was provided on create.
func (u *TGoNFTUpsertBulk) UpdateMtime() *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateMtime()
	})
}

// SetCtime sets the "ctime" field.
func (u *TGoNFTUpsertBulk) SetCtime(v time.Time) *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.SetCtime(v)
	})
}

// UpdateCtime sets the "ctime" field to the value that was provided on create.
func (u *TGoNFTUpsertBulk) UpdateCtime() *TGoNFTUpsertBulk {
	return u.Update(func(s *TGoNFTUpsert) {
		s.UpdateCtime()
	})
}

// Exec executes the query.
func (u *TGoNFTUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TGoNFTCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TGoNFTCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TGoNFTUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
