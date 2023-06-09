package storage

import (
	"github.com/dolthub/swiss"
	"os"
	"sync"
)

type value struct {
	value      string
	name       string
	next       *swiss.Map[string, ivalue]
	mutex      *sync.RWMutex
	isIndex    bool
	attachedTo map[string]bool
}

type ivalue interface {
	GetValue() (val string)
	Get(name string) (val string, err error)
	Set(key, val string)
	SetValueUnique(val string) error
	CreateIndex(name string)
	RecreateIndex()
	IsIndex() bool
	IsEmpty() bool
	NextOrCreate(name string) (val ivalue)
	Next(name string) (val ivalue, ok bool)
	Size() int
	Iter(func(k string, v ivalue) bool)
	AttachIndex(src ivalue) error
	DeleteIndex()
	Delete(key string) error
	Has(key string) bool
	IsAttached(name string) bool
	setAttached(name string)
	save(path string) error
	Name() string
}

func NewIndex(name string) *value {
	return &value{mutex: &sync.RWMutex{}, next: swiss.NewMap[string, ivalue](100), isIndex: true, name: name}
}

func (v *value) GetValue() (val string) {
	v.mutex.RLock()
	val = v.value
	v.mutex.RUnlock()
	return val
}

func (v *value) Name() string {
	return v.name
}

func (v *value) IsEmpty() bool {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	return v.next == nil
}

func (v *value) Get(name string) (string, error) {
	v.mutex.RLock()
	val, ok := v.next.Get(name)
	v.mutex.RUnlock()
	if !ok {
		return "", ErrNotFound
	}

	return val.GetValue(), nil
}

func (v *value) Size() int {
	v.mutex.RLock()
	size := v.next.Count()
	v.mutex.RUnlock()
	return size
}

func (v *value) IsIndex() bool {
	return v.isIndex
}

func (v *value) IsAttached(name string) bool {
	return v.attachedTo[name]
}

func (v *value) setAttached(name string) {
	v.mutex.Lock()
	v.attachedTo[name] = true
	v.mutex.Unlock()
}

func (v *value) AttachIndex(src ivalue) (err error) {
	v.mutex.Lock()
	defer func() {
		v.mutex.Unlock()
		if err == nil {
			src.setAttached(v.Name())
		}
	}()
	if v.next == nil {
		v.next = swiss.NewMap[string, ivalue](10)
		v.next.Put(src.Name(), src)
		return nil
	}
	if v.next.Has(src.Name()) {
		return nil
	}
	v.isIndex = true
	v.next.Put(src.Name(), src)
	return nil
}

func (v *value) Iter(f func(k string, v ivalue) bool) {
	v.next.Iter(f)
}

func (v *value) DeleteIndex() {
	v.mutex.Lock()
	v.next = nil
	v.mutex.Unlock()
}

func (v *value) NextOrCreate(name string) ivalue {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	val, ok := v.next.Get(name)
	if !ok {
		blank := NewIndex(name)
		v.next.Put(name, blank)
		return blank
	}

	return val
}

func (v *value) Next(name string) (ivalue, bool) {
	v.mutex.RLock()
	val, ok := v.next.Get(name)
	v.mutex.RUnlock()
	if !ok {
		return nil, false
	}

	return val, true
}

func (v *value) Set(key, val string) {
	v.mutex.Lock()
	v.next.Put(key, &value{value: val, mutex: &sync.RWMutex{}})
	v.mutex.Unlock()
}

func (v *value) SetValueUnique(val string) error {
	v.mutex.Lock()
	if v.next.Has(val) {
		v.mutex.Unlock()
		return ErrAlreadyExists
	}
	v.value = val
	v.mutex.Unlock()
	return nil
}

func (v *value) CreateIndex(name string) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	if v.next.Has(name) {
		return
	}

	v.next.Put(name, &value{next: swiss.NewMap[string, ivalue](100), mutex: &sync.RWMutex{}})
}

func (v *value) RecreateIndex() {
	v.mutex.Lock()
	v.next = swiss.NewMap[string, ivalue](10000)
	v.mutex.Unlock()
}

func (v *value) Delete(key string) error {
	v.mutex.Lock()
	ok := v.next.Has(key)

	if !ok {
		return ErrNotFound
	}

	v.next.Delete(key)
	v.mutex.Unlock()
	return nil
}

func (v *value) Has(key string) bool {
	v.mutex.RLock()
	ok := v.next.Has(key)
	v.mutex.RUnlock()
	return ok
}

func (v *value) save(path string) (err error) {

	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}

	var f *os.File
	v.next.Iter(func(key string, v ivalue) bool {
		name := path + "/" + key
		if v.IsIndex() {
			err = v.save(name)
			if err != nil {
				return true
			}
			return false
		}

		f, err = os.OpenFile(name, os.O_CREATE, os.ModePerm)
		if err != nil {
			return true
		}

		_, err = f.WriteString(v.GetValue())
		if err != nil {
			return true
		}
		f.Close()

		return false
	})

	return err
}
