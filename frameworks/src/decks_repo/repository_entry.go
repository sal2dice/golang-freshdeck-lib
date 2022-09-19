package decks_repo

type EntryIdentifier interface {
	GetNamespace() string
	GetName() string
}

type SimpleIdentifier struct {
	namespace string
	name      string
}

func (s SimpleIdentifier) GetNamespace() string {
	return s.namespace
}

func (s SimpleIdentifier) GetName() string {
	return s.name
}

type RepositoryEntry[TEntry any] struct {
	identity EntryIdentifier
	entry    TEntry
}
