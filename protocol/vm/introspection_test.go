package vm

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestCheckOutput(t *testing.T) {
	assetID2 := [32]byte{2}
	var (
		index     uint64 = 4
		data             = mustDecodeHex("1f2a05f881ed9fa0c9068a84823677409f863891a2196eb55dbfbb677a566374")
		amount    uint64 = 7
		assetID          = assetID2[:]
		vmVersion uint64 = 1
		code             = []byte("controlprog")
	)

	checkOutput := func(index1 uint64, data1 []byte, amount1 uint64, assetID1 []byte, vmVersion1 uint64, code1 []byte) (bool, error) {
		if !(index1 == index) {
			t.Fatalf("requested index = %v want %v", index1, index)
		}
		if !(bytes.Equal(data1, data)) {
			t.Fatalf("requested data = %v want %v", data1, data)
		}
		if !(amount1 == amount) {
			t.Fatalf("requested amount = %v want %v", amount1, amount)
		}
		if !(bytes.Equal(assetID1, assetID)) {
			t.Fatalf("requested assetID = %v want %v", assetID1, assetID)
		}
		if !(vmVersion1 == vmVersion) {
			t.Fatalf("requested vmVersion = %v want %v", vmVersion1, vmVersion)
		}
		if !(bytes.Equal(code1, code)) {
			t.Fatalf("requested code = %v want %v", code1, code)
		}
		return true, nil
	}

	vm := &virtualMachine{
		runLimit: 50000,
		context:  &Context{CheckOutput: checkOutput},
		program: cat(
			OP_4, // index
			PushdataBytes(data),
			OP_7, // amount
			PushdataBytes(assetID),
			OP_1, // vm version
			PushdataBytes([]byte("controlprog")),
			OP_CHECKOUTPUT,
		),
	}
	err := vm.run()
	if err != nil {
		t.Fatalf("vm.run() = %v, want nil error", err)
	}

	want := [][]byte{{1}}
	if !reflect.DeepEqual(vm.dataStack, want) {
		t.Errorf("final stack = %v, want %v", vm.dataStack, want)
	}
}

func TestCheckOutputFalse(t *testing.T) {
	assetID2 := [32]byte{2}
	assetID := assetID2[:]

	checkOutput := func(uint64, []byte, uint64, []byte, uint64, []byte) (bool, error) {
		return false, nil
	}

	vm := &virtualMachine{
		runLimit: 50000,
		context:  &Context{CheckOutput: checkOutput},
		program: cat(
			OP_4,
			PushdataBytes(mustDecodeHex("1f")),
			OP_7,
			PushdataBytes(assetID),
			OP_1,
			PushdataBytes([]byte("controlprog")),
			OP_CHECKOUTPUT,
		),
	}
	err := vm.run()
	if err != nil {
		t.Fatalf("vm.run() = %v, want nil error", err)
	}

	want := [][]byte{{}}
	if !reflect.DeepEqual(vm.dataStack, want) {
		t.Errorf("final stack = %v, want %v", vm.dataStack, want)
	}
}

func TestCheckOutputError(t *testing.T) {
	assetID2 := [32]byte{2}
	assetID := assetID2[:]

	wantErr := errors.New("foo")

	checkOutput := func(uint64, []byte, uint64, []byte, uint64, []byte) (bool, error) {
		return false, wantErr
	}

	vm := &virtualMachine{
		runLimit: 50000,
		context:  &Context{CheckOutput: checkOutput},
		program: cat(
			OP_4,
			PushdataBytes(mustDecodeHex("1f")),
			OP_7,
			PushdataBytes(assetID),
			OP_1,
			PushdataBytes([]byte("controlprog")),
			OP_CHECKOUTPUT,
		),
	}
	err := vm.run()
	if err != wantErr {
		t.Errorf("vm.run() = %v, want %v", err, wantErr)
	}
}

// x can be byte, []byte, or string;
// they will be concatenated into a flat []byte
func cat(x ...interface{}) []byte {
	var b []byte
	for _, x := range x {
		switch x := x.(type) {
		case Op:
			b = append(b, byte(x))
		case byte:
			b = append(b, x)
		case []byte:
			b = append(b, x...)
		case string:
			b = append(b, x...)
		default:
			panic(fmt.Sprintf("unsupported type %T", x))
		}
	}
	return b
}
