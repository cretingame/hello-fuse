package main

import (
	"context"
	"log"
	"syscall"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

type JoplinRegularFile struct {
	fs.MemRegularFile
}

func (j *JoplinRegularFile) Read(ctx context.Context, fh fs.FileHandle, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {
	log.Println("read")
	f := &j.MemRegularFile
	return f.Read(ctx, fh, dest, off)
}

func (j *JoplinRegularFile) Write(ctx context.Context, fh fs.FileHandle, data []byte, off int64) (uint32, syscall.Errno) {
	log.Println("write")
	f := &j.MemRegularFile
	return f.Write(ctx, fh, data, off)
}
