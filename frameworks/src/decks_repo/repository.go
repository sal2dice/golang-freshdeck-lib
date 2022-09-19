package decks_repo

import (
	"fmt"
)

type RepositoryStore interface {
	WriteRecord()
}

type IRepository interface {
}

type Pair struct {
	a, b interface{}
}

type Repository[TEntry any] struct {
	store *RepositoryStore
	// basic storage for repo objects
	entries map[EntryIdentifier]RepositoryEntry[TEntry]

	entriesByNamespace map[string]map[string]*RepositoryEntry[TEntry]
}

func NewRepository[TEntryIdentifier EntryIdentifier, TEntry any](store RepositoryStore) *Repository[TEntry] {
	rep := Repository[TEntry]{store: &store}
	rep.entries = make(map[EntryIdentifier]RepositoryEntry[TEntry])
	rep.entriesByNamespace = make(map[string]map[string]*RepositoryEntry[TEntry])
	return &rep
}

// base implementation
func (repo Repository[TEntry]) AddEntry(re RepositoryEntry[TEntry]) {
	ns := re.identity.GetNamespace()

	name := re.identity.GetName()
	repo.entries[re.identity] = re

	if repo.entriesByNamespace[ns] == nil {
		repo.entriesByNamespace[ns] = make(map[string]*RepositoryEntry[TEntry])
	}

	repo.entriesByNamespace[ns][name] = &re
}

func (repo Repository[TEntry]) AddEntryById(identifier EntryIdentifier, entry TEntry) {
	r := RepositoryEntry[TEntry]{identity: identifier, entry: entry}
	repo.AddEntry(r)
}

func (repo Repository[TEntry]) AddEntryByName(namespace string, name string, entry TEntry) {
	//	var identifier = SimpleIdentifier{namespace: namespace, name: name}
	//	r := RepositoryEntry[TEntryIdentifier, TEntry]{identity: TEntryIdentifier(identifier, entry: entry}
	//	repo.AddEntry(r)
}

//func (repo Repository[TEntryIdentifier, TEntry]) GetEntryById(identity TEntryIdentifier) (RepositoryEntry[TEntryIdentifier, TEntry], error) {
//	return repo.GetEntryByName(identity.GetNamespace(), identity.GetName())
//}

func (repo Repository[TEntry]) GetEntryById(identity EntryIdentifier) (*RepositoryEntry[TEntry], error) {
	return repo.GetEntryByName(identity.GetNamespace(), identity.GetName())
}

func (repo Repository[TEntry]) GetEntryByName(namespace string, name string) (*RepositoryEntry[TEntry], error) {
	findNs, nsExists := repo.entriesByNamespace[namespace]

	if !nsExists {
		//		var identity TEntryIdentifier = SimpleIdentifier{namespace: namespace, name: name}
		// is this really how this is supposed to work?!?1
		return nil, fmt.Errorf("namespace %v does not exist", namespace)
	}

	findByName, entryExists := findNs[name]

	if !entryExists {
		return nil, fmt.Errorf("namespace %v name %v does not exist", namespace, name)
	}

	return findByName, nil
}
