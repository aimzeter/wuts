package fake

import (
	"sort"

	zone "github.com/aimzeter/wuts/three"
)

type sStore map[string]zone.Student

type FakeStudentStore struct {
	stores sStore
}

func NewFakeStudentStore() *FakeStudentStore {
	stores := make(sStore, 5)
	return &FakeStudentStore{stores}
}

func (f *FakeStudentStore) Seed(docs []zone.Student) {
	for _, d := range docs {
		f.stores[d.PersonalInfo.NIK] = d
	}
}

func (f *FakeStudentStore) Create(d zone.Student) error {
	f.stores[d.PersonalInfo.NIK] = d
	return nil
}

func (f *FakeStudentStore) Get(nik string) zone.Student {
	return f.stores[nik]
}

func (f *FakeStudentStore) Update(d zone.Student) error {
	if _, found := f.stores[d.PersonalInfo.NIK]; found {
		f.stores[d.PersonalInfo.NIK] = d
	}
	return nil
}

func (f *FakeStudentStore) Delete(nik string) error {
	delete(f.stores, nik)
	return nil
}

func (f *FakeStudentStore) CountAll() int {
	return len(f.stores)
}

func (f *FakeStudentStore) RemoveAll() error {
	f.stores = make(sStore, 5)
	return nil
}

func (f *FakeStudentStore) All() []zone.Student {
	var docs []zone.Student
	for _, d := range f.stores {
		docs = append(docs, d)
	}

	sort.Slice(docs, func(i, j int) bool {
		return docs[i].PersonalInfo.NIK < docs[j].PersonalInfo.NIK
	})

	return docs
}
