// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testReceipts(t *testing.T) {
	t.Parallel()

	query := Receipts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testReceiptsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Receipts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReceiptsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Receipts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Receipts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReceiptsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ReceiptSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Receipts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReceiptsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ReceiptExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Receipt exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ReceiptExists to return true, but got false.")
	}
}

func testReceiptsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	receiptFound, err := FindReceipt(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if receiptFound == nil {
		t.Error("want a record, got nil")
	}
}

func testReceiptsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Receipts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testReceiptsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Receipts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testReceiptsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	receiptOne := &Receipt{}
	receiptTwo := &Receipt{}
	if err = randomize.Struct(seed, receiptOne, receiptDBTypes, false, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}
	if err = randomize.Struct(seed, receiptTwo, receiptDBTypes, false, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = receiptOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = receiptTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Receipts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testReceiptsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	receiptOne := &Receipt{}
	receiptTwo := &Receipt{}
	if err = randomize.Struct(seed, receiptOne, receiptDBTypes, false, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}
	if err = randomize.Struct(seed, receiptTwo, receiptDBTypes, false, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = receiptOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = receiptTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Receipts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func receiptBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Receipt) error {
	*o = Receipt{}
	return nil
}

func receiptAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Receipt) error {
	*o = Receipt{}
	return nil
}

func receiptAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Receipt) error {
	*o = Receipt{}
	return nil
}

func receiptBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Receipt) error {
	*o = Receipt{}
	return nil
}

func receiptAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Receipt) error {
	*o = Receipt{}
	return nil
}

func receiptBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Receipt) error {
	*o = Receipt{}
	return nil
}

func receiptAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Receipt) error {
	*o = Receipt{}
	return nil
}

func receiptBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Receipt) error {
	*o = Receipt{}
	return nil
}

func receiptAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Receipt) error {
	*o = Receipt{}
	return nil
}

func testReceiptsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Receipt{}
	o := &Receipt{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, receiptDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Receipt object: %s", err)
	}

	AddReceiptHook(boil.BeforeInsertHook, receiptBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	receiptBeforeInsertHooks = []ReceiptHook{}

	AddReceiptHook(boil.AfterInsertHook, receiptAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	receiptAfterInsertHooks = []ReceiptHook{}

	AddReceiptHook(boil.AfterSelectHook, receiptAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	receiptAfterSelectHooks = []ReceiptHook{}

	AddReceiptHook(boil.BeforeUpdateHook, receiptBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	receiptBeforeUpdateHooks = []ReceiptHook{}

	AddReceiptHook(boil.AfterUpdateHook, receiptAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	receiptAfterUpdateHooks = []ReceiptHook{}

	AddReceiptHook(boil.BeforeDeleteHook, receiptBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	receiptBeforeDeleteHooks = []ReceiptHook{}

	AddReceiptHook(boil.AfterDeleteHook, receiptAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	receiptAfterDeleteHooks = []ReceiptHook{}

	AddReceiptHook(boil.BeforeUpsertHook, receiptBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	receiptBeforeUpsertHooks = []ReceiptHook{}

	AddReceiptHook(boil.AfterUpsertHook, receiptAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	receiptAfterUpsertHooks = []ReceiptHook{}
}

func testReceiptsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Receipts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReceiptsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(receiptColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Receipts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReceiptToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Receipt
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, receiptDBTypes, false, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ReceiptSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Receipt)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testReceiptToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Receipt
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, receiptDBTypes, false, strmangle.SetComplement(receiptPrimaryKeyColumns, receiptColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Receipts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testReceiptsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testReceiptsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ReceiptSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testReceiptsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Receipts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	receiptDBTypes = map[string]string{`ID`: `int`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `DeletedAt`: `timestamp`, `Content`: `varchar`, `UserID`: `int`}
	_              = bytes.MinRead
)

func testReceiptsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(receiptPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(receiptAllColumns) == len(receiptPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Receipts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testReceiptsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(receiptAllColumns) == len(receiptPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Receipt{}
	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Receipts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, receiptDBTypes, true, receiptPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(receiptAllColumns, receiptPrimaryKeyColumns) {
		fields = receiptAllColumns
	} else {
		fields = strmangle.SetComplement(
			receiptAllColumns,
			receiptPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ReceiptSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testReceiptsUpsert(t *testing.T) {
	t.Parallel()

	if len(receiptAllColumns) == len(receiptPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLReceiptUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Receipt{}
	if err = randomize.Struct(seed, &o, receiptDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Receipt: %s", err)
	}

	count, err := Receipts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, receiptDBTypes, false, receiptPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Receipt struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Receipt: %s", err)
	}

	count, err = Receipts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
