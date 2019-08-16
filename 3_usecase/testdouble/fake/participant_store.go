package fake

import (
	"sort"

	zone "github.com/aimzeter/wuts/3_usecase"
)

type pStore map[string]zone.Participant

type FakeParticipantStore struct {
	stores pStore
}

func NewFakeParticipantStore() *FakeParticipantStore {
	stores := make(pStore, 5)
	return &FakeParticipantStore{stores}
}

func (f *FakeParticipantStore) Seed(docs []zone.Participant) {
	for _, d := range docs {
		f.stores[d.PersonalInfo.NIK] = d
	}
}

func (f *FakeParticipantStore) Create(d zone.Participant) error {
	f.stores[d.PersonalInfo.NIK] = d
	return nil
}

func (f *FakeParticipantStore) Get(nik string) zone.Participant {
	return f.stores[nik]
}

func (f *FakeParticipantStore) Update(d zone.Participant) error {
	if _, found := f.stores[d.PersonalInfo.NIK]; found {
		f.stores[d.PersonalInfo.NIK] = d
	}
	return nil
}

func (f *FakeParticipantStore) Delete(nik string) error {
	delete(f.stores, nik)
	return nil
}

func (f *FakeParticipantStore) CountAll() int {
	return len(f.stores)
}

func (f *FakeParticipantStore) RemoveAll() error {
	f.stores = make(pStore, 5)
	return nil
}

func (f *FakeParticipantStore) All() []zone.Participant {
	var docs []zone.Participant
	for _, d := range f.stores {
		docs = append(docs, d)
	}

	sort.Slice(docs, func(i, j int) bool {
		return docs[i].PersonalInfo.NIK < docs[j].PersonalInfo.NIK
	})

	return docs
}
