package decks_repo

import (
	"fmt"
)

type RepositoryStore interface {
	WriteRecord()
}

type RepositoryBehaviorFactory[TEntry any, TBehavior any] interface {
	GetBehavior(identity EntryIdentifier, entry TEntry) TBehavior
}

type IRepository interface {
}

type Pair struct {
	a, b interface{}
}

type Repository[TEntry any, TBehavior any] struct {
	store           *RepositoryStore
	behaviorFactory *RepositoryBehaviorFactory[TEntry, TBehavior]
	// basic storage for repo objects
	entries map[EntryIdentifier]RepositoryEntry[TEntry, TBehavior]

	entriesByNamespace map[string]map[string]*RepositoryEntry[TEntry, TBehavior]
}

func NewRepository[TEntry any, TBehavior any](store RepositoryStore, behaviorFactory RepositoryBehaviorFactory[TEntry, TBehavior]) *Repository[TEntry, TBehavior] {
	rep := Repository[TEntry, TBehavior]{store: &store, behaviorFactory: &behaviorFactory}
	rep.entries = make(map[EntryIdentifier]RepositoryEntry[TEntry, TBehavior])
	rep.entriesByNamespace = make(map[string]map[string]*RepositoryEntry[TEntry, TBehavior])

	return &rep
}

// base implementation
func (repo Repository[TEntry, TBehavior]) AddEntry(re RepositoryEntry[TEntry, TBehavior]) {
	ns := re.identity.GetNamespace()

	name := re.identity.GetName()
	repo.entries[re.identity] = re

	if repo.entriesByNamespace[ns] == nil {
		repo.entriesByNamespace[ns] = make(map[string]*RepositoryEntry[TEntry, TBehavior])
	}

	repo.entriesByNamespace[ns][name] = &re
}

func (repo Repository[TEntry, TBehavior]) AddEntryById(identifier EntryIdentifier, entry TEntry) {
	re := RepositoryEntry[TEntry, TBehavior]{identity: identifier, entry: entry}
	re.behavior = (*repo.behaviorFactory).GetBehavior(re.identity, re.entry)
	repo.AddEntry(re)
}

func (repo Repository[TEntry, TBehavior]) AddEntryByName(namespace string, name string, entry TEntry) {
	var identifier = SimpleIdentifier{namespace: namespace, name: name}
	repo.AddEntryById(identifier, entry)
}

//func (repo Repository[TEntryIdentifier, TEntry]) GetEntryById(identity TEntryIdentifier) (RepositoryEntry[TEntryIdentifier, TEntry], error) {
//	return repo.GetEntryByName(identity.GetNamespace(), identity.GetName())
//}

func (repo Repository[TEntry, TBehavior]) GetEntryById(identity EntryIdentifier) (*RepositoryEntry[TEntry, TBehavior], error) {
	return repo.GetEntryByName(identity.GetNamespace(), identity.GetName())
}

func (repo Repository[TEntry, TBehavior]) GetEntryByName(namespace string, name string) (*RepositoryEntry[TEntry, TBehavior], error) {
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
