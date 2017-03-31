package vm_test

import (
	"bytes"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/kr/pretty"

	"chain/crypto/sha3pool"
	"chain/errors"
	"chain/protocol/bc"
	. "chain/protocol/vm"
	"chain/testutil"
)

func TestNextProgram(t *testing.T) {
	prog, err := Assemble("NEXTPROGRAM 0x010203 EQUAL")
	if err != nil {
		t.Fatal(err)
	}
	context := &Context{
		VMVersion:            1,
		Code:                 []byte{0xcd, 0x3, 0x1, 0x2, 0x3, 0x87},
		Arguments:            nil,
		BlockHash:            &[]byte{0x90, 0xf6, 0x1f, 0xd7, 0x35, 0xcd, 0x28, 0x7c, 0x78, 0xf, 0x43, 0xdc, 0xa7, 0xe7, 0xaf, 0x63, 0xad, 0x83, 0xa, 0x63, 0xcf, 0x56, 0x65, 0x7, 0x21, 0xcd, 0xb4, 0x54, 0xa3, 0xee, 0xf9, 0x2f},
		BlockTimeMS:          new(uint64),
		NextConsensusProgram: &[]byte{1, 2, 3},
	}

	vm := &VirtualMachine{
		RunLimit: 50000,
		Program:  prog,
		Context:  context,
	}
	_, err = vm.Run()
	if err != nil {
		t.Errorf("got error %s, expected none", err)
	}

	prog, err = Assemble("NEXTPROGRAM 0x0102 EQUAL")
	if err != nil {
		t.Fatal(err)
	}
	vm = &VirtualMachine{
		RunLimit: 50000,
		Program:  prog,
		Context:  context,
	}
	_, err = vm.Run()
	if err == nil && vm.FalseResult() {
		err = ErrFalseVMResult
	}
	switch err {
	case nil:
		t.Error("got ok result, expected failure")
	case ErrFalseVMResult:
		// ok
	default:
		t.Errorf("got error %s, expected ErrFalseVMResult", err)
	}
}

func TestBlockTime(t *testing.T) {
	prog, err := Assemble("BLOCKTIME 3263827 NUMEQUAL")
	if err != nil {
		t.Fatal(err)
	}
	context := &Context{
		VMVersion:            1,
		Code:                 []byte{0xce, 0x3, 0x53, 0xcd, 0x31, 0x9c},
		Arguments:            nil,
		BlockHash:            &[]byte{0xc5, 0x5d, 0x64, 0xec, 0x18, 0xec, 0x3e, 0xb6, 0x30, 0xb5, 0x69, 0xd4, 0xb1, 0x23, 0xb9, 0x2c, 0x80, 0x64, 0xce, 0x8d, 0x7d, 0x7c, 0xc0, 0xbf, 0xaf, 0x57, 0x8, 0xb0, 0x62, 0x71, 0xa3, 0xed},
		BlockTimeMS:          u64(0x31cd53),
		NextConsensusProgram: &[]byte{},
	}

	vm := &VirtualMachine{
		RunLimit: 50000,
		Program:  prog,
		Context:  context,
	}
	_, err = vm.Run()
	if err != nil {
		t.Errorf("got error %s, expected none", err)
	}

	prog, err = Assemble("BLOCKTIME 3263826 NUMEQUAL")
	if err != nil {
		t.Fatal(err)
	}
	vm = &VirtualMachine{
		RunLimit: 50000,
		Program:  prog,
		Context:  context,
	}
	_, err = vm.Run()
	if err == nil && vm.FalseResult() {
		err = ErrFalseVMResult
	}
	switch err {
	case nil:
		t.Error("got ok result, expected failure")
	case ErrFalseVMResult:
		// ok
	default:
		t.Errorf("got error %s, expected ErrFalseVMResult", err)
	}
}

func TestOutputIDAndNonceOp(t *testing.T) {
	var emptyHash bc.Hash
	sha3pool.Sum256(emptyHash[:], nil)

	anchorID := []byte{0xc4, 0xa6, 0xe6, 0x25, 0x6d, 0xeb, 0xfc, 0xa3, 0x79, 0x59, 0x5e, 0x44, 0x4b, 0x91, 0xaf, 0x56, 0x84, 0x63, 0x97, 0xe8, 0x0, 0x7e, 0xa8, 0x7c, 0x40, 0xc6, 0x22, 0x17, 0xd, 0xd1, 0x3f, 0xf7}
	outputID := []byte{0xa, 0x60, 0xf9, 0xb1, 0x29, 0x50, 0xc8, 0x4c, 0x22, 0x10, 0x12, 0xa8, 0x8, 0xef, 0x77, 0x82, 0x82, 0x3b, 0x7e, 0x16, 0xb7, 0x1f, 0xe2, 0xba, 0x1, 0x81, 0x1c, 0xda, 0x96, 0xa2, 0x17, 0xdf}
	prog := []byte{uint8(OP_OUTPUTID)}
	vm := &VirtualMachine{
		RunLimit: 50000,
		Program:  prog,
		Context:  &Context{SpentOutputID: &outputID},
	}
	gotVM, err := vm.Step()
	if err != nil {
		t.Fatal(err)
	}

	expectedStack := [][]byte{outputID}
	if !testutil.DeepEqual(gotVM.DataStack, expectedStack) {
		t.Errorf("expected stack %v, got %v; vm is:\n%s", expectedStack, gotVM.DataStack, spew.Sdump(vm))
	}

	prog = []byte{uint8(OP_OUTPUTID)}
	vm = &VirtualMachine{
		RunLimit: 50000,
		Program:  prog,
		Context:  &Context{SpentOutputID: nil},
	}
	_, err = vm.Step()
	if err != ErrContext {
		t.Errorf("expected ErrContext, got %v", err)
	}

	prog = []byte{uint8(OP_NONCE)}
	vm = &VirtualMachine{
		RunLimit: 50000,
		Program:  prog,
		Context:  &Context{AnchorID: nil},
	}
	_, err = vm.Step()
	if err != ErrContext {
		t.Errorf("expected ErrContext, got %v", err)
	}

	prog = []byte{uint8(OP_NONCE)}
	vm = &VirtualMachine{
		RunLimit: 50000,
		Program:  prog,
		Context:  &Context{AnchorID: &anchorID},
	}
	gotVM, err = vm.Step()
	if err != nil {
		t.Fatal(err)
	}

	expectedStack = [][]byte{anchorID}
	if !testutil.DeepEqual(gotVM.DataStack, expectedStack) {
		t.Errorf("expected stack %v, got %v", expectedStack, gotVM.DataStack)
	}
}

func TestIntrospectionOps(t *testing.T) {
	tx := bc.NewTx(bc.TxData{
		ReferenceData: []byte("txref"),
		Inputs: []*bc.TxInput{
			bc.NewSpendInput(nil, bc.Hash{}, bc.AssetID{1}, 5, 1, []byte("spendprog"), bc.Hash{}, []byte("ref")),
			bc.NewIssuanceInput(nil, 6, nil, bc.Hash{}, []byte("issueprog"), nil, nil),
		},
		Outputs: []*bc.TxOutput{
			bc.NewTxOutput(bc.AssetID{3}, 8, []byte("wrongprog"), nil),
			bc.NewTxOutput(bc.AssetID{3}, 8, []byte("controlprog"), nil),
			bc.NewTxOutput(bc.AssetID{2}, 8, []byte("controlprog"), nil),
			bc.NewTxOutput(bc.AssetID{2}, 7, []byte("controlprog"), nil),
			bc.NewTxOutput(bc.AssetID{2}, 7, []byte("controlprog"), []byte("outref")),
		},
		MinTime: 0,
		MaxTime: 20,
	})

	entry0 := tx.TxEntries.TxInputs[0]
	context0 := &Context{
		EntryID:    []byte{0x2e, 0x68, 0xd7, 0x8c, 0xde, 0xaa, 0x98, 0x94, 0x4c, 0x12, 0x51, 0x2c, 0xf9, 0xc7, 0x19, 0xeb, 0x48, 0x81, 0xe9, 0xaf, 0xb6, 0x1e, 0x4b, 0x76, 0x6d, 0xf5, 0xf3, 0x69, 0xae, 0xe6, 0x39, 0x2c},
		NumResults: u64(5),
		AssetID:    &[]byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		Amount:     u64(5),
		MinTimeMS:  new(uint64),
		MaxTimeMS:  u64(20),
		EntryData:  &[]byte{0x44, 0xbe, 0x5e, 0x14, 0xce, 0x21, 0x6f, 0x4b, 0x2c, 0x35, 0xa5, 0xeb, 0xb, 0x35, 0xd0, 0x78, 0xbd, 0xa5, 0x5c, 0xf0, 0x5b, 0x5d, 0x36, 0xee, 0xe, 0x7a, 0x1, 0xfb, 0xc6, 0xef, 0x62, 0xb7},
		TxData:     &[]byte{0x3e, 0x51, 0x90, 0xf2, 0x69, 0x1e, 0x6d, 0x45, 0x1c, 0x50, 0xed, 0xf9, 0xa9, 0xa6, 0x6a, 0x7a, 0x67, 0x79, 0xc7, 0x87, 0x67, 0x64, 0x52, 0x81, 0xd, 0xbf, 0x4f, 0x6e, 0x40, 0x53, 0x68, 0x2c},
		DestPos:    new(uint64),
		CheckOutput: func(index uint64, data []byte, amount uint64, assetID []byte, vmVersion uint64, code []byte) (bool, error) {
			checkEntry := func(e bc.Entry) (bool, error) {
				check := func(prog bc.Program, value bc.AssetAmount, dataHash bc.Hash) bool {
					return (prog.VMVersion == vmVersion &&
						bytes.Equal(prog.Code, code) &&
						bytes.Equal(value.AssetID[:], assetID) &&
						value.Amount == amount &&
						(len(data) == 0 || bytes.Equal(dataHash[:], data)))
				}

				switch e := e.(type) {
				case *bc.Output:
					return check(e.Body.ControlProgram, e.Body.Source.Value, e.Body.Data), nil

				case *bc.Retirement:
					return check(bc.Program{}, e.Body.Source.Value, e.Body.Data), nil
				}

				return false, ErrContext
			}

			checkMux := func(m *bc.Mux) (bool, error) {
				if index >= uint64(len(m.Witness.Destinations)) {
					return false, errors.Wrapf(ErrBadValue, "index %d >= %d", index, len(m.Witness.Destinations))
				}
				return checkEntry(m.Witness.Destinations[index].Entry)
			}

			switch e := entry0.(type) {
			case *bc.Mux:
				return checkMux(e)

			case *bc.Issuance:
				if m, ok := e.Witness.Destination.Entry.(*bc.Mux); ok {
					return checkMux(m)
				}
				if index != 0 {
					return false, errors.Wrapf(ErrBadValue, "index %d >= 1", index)
				}
				return checkEntry(e.Witness.Destination.Entry)

			case *bc.Spend:
				if m, ok := e.Witness.Destination.Entry.(*bc.Mux); ok {
					return checkMux(m)
				}
				if index != 0 {
					return false, errors.Wrapf(ErrBadValue, "index %d >= 1", index)
				}
				return checkEntry(e.Witness.Destination.Entry)
			}

			return false, ErrContext
		},
	}

	type testStruct struct {
		op      Op
		startVM *VirtualMachine
		wantErr error
		wantVM  *VirtualMachine
	}
	cases := []testStruct{{
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				{4},
				mustDecodeHex("1f2a05f881ed9fa0c9068a84823677409f863891a2196eb55dbfbb677a566374"),
				{7},
				append([]byte{2}, make([]byte, 31)...),
				{1},
				[]byte("controlprog"),
			},
			Context: &Context{},
		},
		wantVM: &VirtualMachine{
			RunLimit:     50101,
			DeferredCost: -117,
			DataStack:    [][]byte{{1}},
			Context:      context0,
		},
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				{3},
				mustDecodeHex("1f2a05f881ed9fa0c9068a84823677409f863891a2196eb55dbfbb677a566374"),
				{7},
				append([]byte{2}, make([]byte, 31)...),
				{1},
				[]byte("controlprog"),
			},
			Context: context0,
		},
		wantVM: &VirtualMachine{
			RunLimit:     50102,
			DeferredCost: -118,
			DataStack:    [][]byte{{}},
			Context:      context0,
		},
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				{0},
				[]byte{},
				{1},
				append([]byte{9}, make([]byte, 31)...),
				{1},
				[]byte("missingprog"),
			},
			Context: context0,
		},
		wantVM: &VirtualMachine{
			RunLimit:     50070,
			DeferredCost: -86,
			DataStack:    [][]byte{{}},
			Context:      context0,
		},
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			Context: context0,
		},
		wantErr: ErrDataStackUnderflow,
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				[]byte("controlprog"),
			},
			Context: context0,
		},
		wantErr: ErrDataStackUnderflow,
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				append([]byte{2}, make([]byte, 31)...),
				{1},
				[]byte("controlprog"),
			},
			Context: context0,
		},
		wantErr: ErrDataStackUnderflow,
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				{7},
				append([]byte{2}, make([]byte, 31)...),
				{1},
				[]byte("controlprog"),
			},
			Context: context0,
		},
		wantErr: ErrDataStackUnderflow,
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				mustDecodeHex("1f2a05f881ed9fa0c9068a84823677409f863891a2196eb55dbfbb677a566374"),
				{7},
				append([]byte{2}, make([]byte, 31)...),
				{1},
				[]byte("controlprog"),
			},
			Context: context0,
		},
		wantErr: ErrDataStackUnderflow,
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				{4},
				mustDecodeHex("1f2a05f881ed9fa0c9068a84823677409f863891a2196eb55dbfbb677a566374"),
				{7},
				append([]byte{2}, make([]byte, 31)...),
				Int64Bytes(-1),
				[]byte("controlprog"),
			},
			Context: context0,
		},
		wantErr: ErrBadValue,
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				{4},
				mustDecodeHex("1f2a05f881ed9fa0c9068a84823677409f863891a2196eb55dbfbb677a566374"),
				Int64Bytes(-1),
				append([]byte{2}, make([]byte, 31)...),
				{1},
				[]byte("controlprog"),
			},
			Context: context0,
		},
		wantErr: ErrBadValue,
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				Int64Bytes(-1),
				mustDecodeHex("1f2a05f881ed9fa0c9068a84823677409f863891a2196eb55dbfbb677a566374"),
				{7},
				append([]byte{2}, make([]byte, 31)...),
				{1},
				[]byte("controlprog"),
			},
			Context: context0,
		},
		wantErr: ErrBadValue,
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			DataStack: [][]byte{
				{5},
				mustDecodeHex("1f2a05f881ed9fa0c9068a84823677409f863891a2196eb55dbfbb677a566374"),
				{7},
				append([]byte{2}, make([]byte, 31)...),
				{1},
				[]byte("controlprog"),
			},
			Context: context0,
		},
		wantErr: ErrBadValue,
	}, {
		op: OP_CHECKOUTPUT,
		startVM: &VirtualMachine{
			RunLimit: 0,
			DataStack: [][]byte{
				{4},
				mustDecodeHex("1f2a05f881ed9fa0c9068a84823677409f863891a2196eb55dbfbb677a566374"),
				{7},
				append([]byte{2}, make([]byte, 31)...),
				{1},
				[]byte("controlprog"),
			},
			Context: context0,
		},
		wantErr: ErrRunLimitExceeded,
	}, {
		op: OP_ASSET,
		startVM: &VirtualMachine{
			Context: context0,
		},
		wantVM: &VirtualMachine{
			RunLimit:     49959,
			DeferredCost: 40,
			DataStack:    [][]byte{append([]byte{1}, make([]byte, 31)...)},
			Context:      context0,
		},
	}, {
		op: OP_AMOUNT,
		startVM: &VirtualMachine{
			Context: context0,
		},
		wantVM: &VirtualMachine{
			RunLimit:     49990,
			DeferredCost: 9,
			DataStack:    [][]byte{{5}},
			Context:      context0,
		},
	}, {
		op: OP_PROGRAM,
		startVM: &VirtualMachine{
			Program: []byte("spendprog"),
			Context: NewTxVMContext(tx.TxEntries, tx.TxEntries.TxInputs[0], bc.Program{VMVersion: 1, Code: []byte("spendprog")}, nil),
		},
		wantVM: &VirtualMachine{
			RunLimit:     49982,
			DeferredCost: 17,
			DataStack:    [][]byte{[]byte("spendprog")},
			Context:      NewTxVMContext(tx.TxEntries, tx.TxEntries.TxInputs[0], bc.Program{VMVersion: 1, Code: []byte("spendprog")}, nil),
		},
	}, {
		op: OP_PROGRAM,
		startVM: &VirtualMachine{
			Program:  []byte("issueprog"),
			RunLimit: 50000,
			Context:  NewTxVMContext(tx.TxEntries, tx.TxEntries.TxInputs[1], bc.Program{VMVersion: 1, Code: []byte("issueprog")}, nil),
		},
		wantVM: &VirtualMachine{
			RunLimit:     49982,
			DeferredCost: 17,
			DataStack:    [][]byte{[]byte("issueprog")},
			Context:      NewTxVMContext(tx.TxEntries, tx.TxEntries.TxInputs[1], bc.Program{VMVersion: 1, Code: []byte("issueprog")}, nil),
		},
	}, {
		op: OP_MINTIME,
		startVM: &VirtualMachine{
			Context: context0,
		},
		wantVM: &VirtualMachine{
			RunLimit:     49991,
			DeferredCost: 8,
			DataStack:    [][]byte{[]byte{}},
			Context:      context0,
		},
	}, {
		op: OP_MAXTIME,
		startVM: &VirtualMachine{
			Context: context0,
		},
		wantVM: &VirtualMachine{
			RunLimit:     49990,
			DeferredCost: 9,
			DataStack:    [][]byte{{20}},
			Context:      context0,
		},
	}, {
		op: OP_TXDATAHASH,
		startVM: &VirtualMachine{
			Context: context0,
		},
		wantVM: &VirtualMachine{
			RunLimit:     49959,
			DeferredCost: 40,
			DataStack: [][]byte{{
				62, 81, 144, 242, 105, 30, 109, 69, 28, 80, 237, 249, 169, 166, 106, 122,
				103, 121, 199, 135, 103, 100, 82, 129, 13, 191, 79, 110, 64, 83, 104, 44,
			}},
			Context: context0,
		},
	}, {
		op: OP_DATAHASH,
		startVM: &VirtualMachine{
			Context: context0,
		},
		wantVM: &VirtualMachine{
			RunLimit:     49959,
			DeferredCost: 40,
			DataStack: [][]byte{{
				68, 190, 94, 20, 206, 33, 111, 75, 44, 53, 165, 235, 11, 53, 208, 120,
				189, 165, 92, 240, 91, 93, 54, 238, 14, 122, 1, 251, 198, 239, 98, 183,
			}},
			Context: context0,
		},
	}, {
		op: OP_INDEX,
		startVM: &VirtualMachine{
			Context: context0,
		},
		wantVM: &VirtualMachine{
			RunLimit:     49991,
			DeferredCost: 8,
			DataStack:    [][]byte{[]byte{}},
			Context:      context0,
		},
	}, {
		// The current entry is input 0
		op: OP_ENTRYID,
		startVM: &VirtualMachine{
			Context: context0,
		},
		wantVM: &VirtualMachine{
			RunLimit:     49959,
			DeferredCost: 40,
			DataStack:    [][]byte{tx.TxEntries.TxInputIDs[0][:]},
			Context:      context0,
		},
	}, {
		// The current entry is input 1
		op: OP_ENTRYID,
		startVM: &VirtualMachine{
			Context: NewTxVMContext(tx.TxEntries, tx.TxEntries.TxInputs[1], bc.Program{VMVersion: 1}, nil),
		},
		wantVM: &VirtualMachine{
			RunLimit:     49959,
			DeferredCost: 40,
			DataStack:    [][]byte{tx.TxEntries.TxInputIDs[1][:]},
			Context:      NewTxVMContext(tx.TxEntries, tx.TxEntries.TxInputs[1], bc.Program{VMVersion: 1}, nil),
		},
	}, {
		// The current entry is the internal mux node
		op: OP_ENTRYID,
		startVM: &VirtualMachine{
			Context: NewTxVMContext(tx.TxEntries, tx.TxEntries.TxInputs[0].(*bc.Spend).Witness.Destination.Entry, bc.Program{VMVersion: 1}, nil),
		},
		wantVM: &VirtualMachine{
			RunLimit:     49959,
			DeferredCost: 40,
			DataStack:    [][]byte{tx.TxEntries.TxInputs[0].(*bc.Spend).Witness.Destination.Ref[:]},
			Context:      NewTxVMContext(tx.TxEntries, tx.TxEntries.TxInputs[0].(*bc.Spend).Witness.Destination.Entry, bc.Program{VMVersion: 1}, nil),
		},
	}}

	txops := []Op{
		OP_CHECKOUTPUT, OP_ASSET, OP_AMOUNT, OP_PROGRAM,
		OP_MINTIME, OP_MAXTIME, OP_TXDATAHASH, OP_DATAHASH,
		OP_INDEX, OP_OUTPUTID,
	}

	for _, op := range txops {
		cases = append(cases, testStruct{
			op: op,
			startVM: &VirtualMachine{
				RunLimit: 0,
				Context:  context0,
			},
			wantErr: ErrRunLimitExceeded,
		})
	}

	for i, c := range cases {
		t.Logf("case %d", i)
		prog := []byte{byte(c.op)}
		vm := c.startVM
		if c.wantErr != ErrRunLimitExceeded {
			vm.RunLimit = 50000
		}
		vm.Program = prog
		gotVM, err := vm.Run()
		switch errors.Root(err) {
		case c.wantErr:
			// ok
		case nil:
			t.Errorf("case %d, op %s: got no error, want %v", i, OpName(c.op), c.wantErr)
		default:
			t.Errorf("case %d, op %s: got err = %v want %v", i, OpName(c.op), err, c.wantErr)
		}
		if c.wantErr != nil {
			continue
		}

		c.wantVM.Program = prog
		c.wantVM.PC = 1
		c.wantVM.NextPC = 1
		c.wantVM.Context = gotVM.Context

		if !testutil.DeepEqual(gotVM, c.wantVM) {
			t.Errorf("case %d, op %s: unexpected vm result\n\tgot:  %+v\n\twant: %+v\nstartVM is:\n%s", i, OpName(c.op), gotVM, c.wantVM, spew.Sdump(c.startVM))
		}
	}
}

func NewTxVMContext(tx *bc.TxEntries, entry bc.Entry, prog bc.Program, args [][]byte) *Context {
	var (
		numResults = uint64(len(tx.Results))
		txData     = tx.Body.Data[:]
		entryID    = bc.EntryID(entry) // TODO(bobg): pass this in, don't recompute it

		assetID       *[]byte
		amount        *uint64
		entryData     *[]byte
		destPos       *uint64
		anchorID      *[]byte
		spentOutputID *[]byte
	)

	switch e := entry.(type) {
	case *bc.Nonce:
		if iss, ok := e.Anchored.(*bc.Issuance); ok {
			a1 := iss.Body.Value.AssetID[:]
			assetID = &a1
			amount = &iss.Body.Value.Amount
		}

	case *bc.Issuance:
		a1 := e.Body.Value.AssetID[:]
		assetID = &a1
		amount = &e.Body.Value.Amount
		destPos = &e.Witness.Destination.Position
		d := e.Body.Data[:]
		entryData = &d
		a2 := e.Body.AnchorID[:]
		anchorID = &a2

	case *bc.Spend:
		a1 := e.SpentOutput.Body.Source.Value.AssetID[:]
		assetID = &a1
		amount = &e.SpentOutput.Body.Source.Value.Amount
		destPos = &e.Witness.Destination.Position
		d := e.Body.Data[:]
		entryData = &d
		s := e.Body.SpentOutputID[:]
		spentOutputID = &s

	case *bc.Output:
		d := e.Body.Data[:]
		entryData = &d

	case *bc.Retirement:
		d := e.Body.Data[:]
		entryData = &d
	}

	var txSigHash *[]byte
	txSigHashFn := func() []byte {
		if txSigHash == nil {
			hasher := sha3pool.Get256()
			defer sha3pool.Put256(hasher)

			hasher.Write(entryID[:])
			hasher.Write(tx.ID[:])

			var hash bc.Hash
			hasher.Read(hash[:])
			hashBytes := hash.Bytes()
			txSigHash = &hashBytes
		}
		return *txSigHash
	}

	checkOutput := func(index uint64, data []byte, amount uint64, assetID []byte, vmVersion uint64, code []byte) (bool, error) {
		checkEntry := func(e bc.Entry) (bool, error) {
			check := func(prog bc.Program, value bc.AssetAmount, dataHash bc.Hash) bool {
				return (prog.VMVersion == vmVersion &&
					bytes.Equal(prog.Code, code) &&
					bytes.Equal(value.AssetID[:], assetID) &&
					value.Amount == amount &&
					(len(data) == 0 || bytes.Equal(dataHash[:], data)))
			}

			switch e := e.(type) {
			case *bc.Output:
				return check(e.Body.ControlProgram, e.Body.Source.Value, e.Body.Data), nil

			case *bc.Retirement:
				return check(bc.Program{}, e.Body.Source.Value, e.Body.Data), nil
			}

			return false, ErrContext
		}

		checkMux := func(m *bc.Mux) (bool, error) {
			if index >= uint64(len(m.Witness.Destinations)) {
				return false, errors.Wrapf(ErrBadValue, "index %d >= %d", index, len(m.Witness.Destinations))
			}
			return checkEntry(m.Witness.Destinations[index].Entry)
		}

		switch e := entry.(type) {
		case *bc.Mux:
			return checkMux(e)

		case *bc.Issuance:
			if m, ok := e.Witness.Destination.Entry.(*bc.Mux); ok {
				return checkMux(m)
			}
			if index != 0 {
				return false, errors.Wrapf(ErrBadValue, "index %d >= 1", index)
			}
			return checkEntry(e.Witness.Destination.Entry)

		case *bc.Spend:
			if m, ok := e.Witness.Destination.Entry.(*bc.Mux); ok {
				return checkMux(m)
			}
			if index != 0 {
				return false, errors.Wrapf(ErrBadValue, "index %d >= 1", index)
			}
			return checkEntry(e.Witness.Destination.Entry)
		}

		return false, ErrContext
	}

	result := &Context{
		VMVersion: prog.VMVersion,
		Code:      prog.Code,
		Arguments: args,

		EntryID: entryID[:],

		TxVersion: &tx.Body.Version,

		TxSigHash:     txSigHashFn,
		NumResults:    &numResults,
		AssetID:       assetID,
		Amount:        amount,
		MinTimeMS:     &tx.Body.MinTimeMS,
		MaxTimeMS:     &tx.Body.MaxTimeMS,
		EntryData:     entryData,
		TxData:        &txData,
		DestPos:       destPos,
		AnchorID:      anchorID,
		SpentOutputID: spentOutputID,
		CheckOutput:   checkOutput,
	}

	pretty.Println(result)
	return result
}

func u64(v uint64) *uint64 {
	p := new(uint64)
	*p = v
	return p
}

func sigHashFunc(v []byte) func() []byte {
	return func() []byte { return v }
}
