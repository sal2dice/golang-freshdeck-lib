package decks_repo

import (
	"testing"
)

type MockRepositoryStore struct {
}

func (store *MockRepositoryStore) WriteRecord() {

}

type MockBehaviorFactory struct {
}

func (m MockBehaviorFactory) GetBehavior(identity EntryIdentifier, entry string) string {
	//TODO implement me
	//	panic("implement me")
	return "a"
}

//type NewType = RepositoryEntry[SimpleIdentifier, string]

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestEntryAssignment(t *testing.T) {
	var rs = MockRepositoryStore{}
	var bh = MockBehaviorFactory{}

	var repo = NewRepository[string, string](&rs, &bh)

	if repo.store == nil {
		t.Fatalf("NO!")
	}

	var entry = "foo"
	idObj := SimpleIdentifier{"ns", "name"}

	repo.AddEntry(RepositoryEntry[string, string]{idObj, entry, "behavior"})
	repo.AddEntryById(idObj, entry)
	repo.AddEntryByName("ns", "myName2", entry)

	_, exists := repo.GetEntryById(idObj)

	if exists != nil {
		t.Fatalf("NO!")
	}

}
