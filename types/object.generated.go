// Code generated by go generate internal/cmd; DO NOT EDIT.
package types

import (
	"fmt"
	"sync"
	"time"
)

type Object struct {
	contentMd5  string
	contentType string
	etag        string
	// ID is the unique key in service.
	ID string
	// Name is the relative path towards service's WorkDir.
	Name string
	size int64
	// Type should be one of `file`, `stream`, `dir` or `invalid`.
	Type      ObjectType
	updatedAt time.Time

	// client is the client in which Object is alive.
	client Storager
	// m stores storage related metadata.
	meta map[string]interface{}

	bit  uint64
	done uint32
	m    sync.Mutex
}

func (o *Object) GetContentMD5() (string, bool) {
	o.stat()

	if o.bit&(1<<0) == 1 {
		return o.contentMd5, true
	}
	return "", false
}

func (o *Object) MustGetContentMD5() string {
	o.stat()

	if o.bit&(1<<0) != 1 {
		panic(fmt.Sprintf("object content-md5 is not set"))
	}
	return o.contentMd5
}

func (o *Object) SetContentMD5(v string) *Object {
	o.contentMd5 = v
	o.bit |= 1 << 0
	return o
}

func (o *Object) GetContentType() (string, bool) {
	o.stat()

	if o.bit&(1<<1) == 1 {
		return o.contentType, true
	}
	return "", false
}

func (o *Object) MustGetContentType() string {
	o.stat()

	if o.bit&(1<<1) != 1 {
		panic(fmt.Sprintf("object content-type is not set"))
	}
	return o.contentType
}

func (o *Object) SetContentType(v string) *Object {
	o.contentType = v
	o.bit |= 1 << 1
	return o
}

func (o *Object) GetETag() (string, bool) {
	o.stat()

	if o.bit&(1<<2) == 1 {
		return o.etag, true
	}
	return "", false
}

func (o *Object) MustGetETag() string {
	o.stat()

	if o.bit&(1<<2) != 1 {
		panic(fmt.Sprintf("object etag is not set"))
	}
	return o.etag
}

func (o *Object) SetETag(v string) *Object {
	o.etag = v
	o.bit |= 1 << 2
	return o
}
func (o *Object) GetID() string {
	return o.ID
}

func (o *Object) SetID(v string) *Object {
	o.ID = v
	return o
}
func (o *Object) GetName() string {
	return o.Name
}

func (o *Object) SetName(v string) *Object {
	o.Name = v
	return o
}

func (o *Object) GetSize() (int64, bool) {
	o.stat()

	if o.bit&(1<<5) == 1 {
		return o.size, true
	}
	return 0, false
}

func (o *Object) MustGetSize() int64 {
	o.stat()

	if o.bit&(1<<5) != 1 {
		panic(fmt.Sprintf("object size is not set"))
	}
	return o.size
}

func (o *Object) SetSize(v int64) *Object {
	o.size = v
	o.bit |= 1 << 5
	return o
}
func (o *Object) GetType() ObjectType {
	return o.Type
}

func (o *Object) SetType(v ObjectType) *Object {
	o.Type = v
	return o
}

func (o *Object) GetUpdatedAt() (time.Time, bool) {
	o.stat()

	if o.bit&(1<<7) == 1 {
		return o.updatedAt, true
	}
	return time.Time{}, false
}

func (o *Object) MustGetUpdatedAt() time.Time {
	o.stat()

	if o.bit&(1<<7) != 1 {
		panic(fmt.Sprintf("object updated_at is not set"))
	}
	return o.updatedAt
}

func (o *Object) SetUpdatedAt(v time.Time) *Object {
	o.updatedAt = v
	o.bit |= 1 << 7
	return o
}
