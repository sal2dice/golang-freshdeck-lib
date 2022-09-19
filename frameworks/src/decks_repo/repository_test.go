package decks_repo

import (
	"testing"
)

type MockRepositoryStore struct {
}

func (store *MockRepositoryStore) WriteRecord() {

}

//type NewType = RepositoryEntry[SimpleIdentifier, string]

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestEntryAssignment(t *testing.T) {
	var rs = MockRepositoryStore{}

	var repo = NewRepository[SimpleIdentifier, string](&rs)

	if repo.store == nil {
		t.Fatalf("NO!")
	}

	var entry = "foo"

	var idObj = SimpleIdentifier{namespace: "ns", name: "myName"}

	repo.AddEntry(RepositoryEntry[string]{idObj, entry})
	repo.AddEntryById(idObj, entry)
	repo.AddEntryByName("ns", "myName2", entry)

	_, exists := repo.GetEntryById(idObj)

	if exists != nil {
		t.Fatalf("NO!")
	}

}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestEntryAssignment2(t *testing.T) {
	var rs = MockRepositoryStore{}

	var repo = NewRepository[SimpleIdentifier, string](&rs)

	if repo.store == nil {
		t.Fatalf("NO!")
	}

	var entry = "foo"

	repo.AddEntry(RepositoryEntry[string]{SimpleIdentifier{namespace: "ns", name: "myName1"}, entry})
	repo.AddEntryById(SimpleIdentifier{namespace: "ns", name: "myName"}, entry)
	//	repo.AddValue("ns", "myName2", entry)
}
