//go:build tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.0.1
// 	protoc               v3.21.5
// source: examples/known-types/known/known.proto

package known

import (
	context "context"
	wasm "github.com/knqyf263/go-plugin/wasm"
)

var wellKnown WellKnown

func RegisterWellKnown(p WellKnown) {
	wellKnown = p
}

//export well_known_diff
func _well_known_diff(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req DiffRequest
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := wellKnown.Diff(context.Background(), req)
	if err != nil {
		return 0
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}