package datas

import (
	"flag"

	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
)

// DataStore provides versioned storage for noms values. Each DataStore instance represents one moment in history. Heads() returns the Commit from each active fork at that moment. The Commit() method returns a new DataStore, representing a new moment in history.
type DataStore interface {
	chunks.ChunkStore

	// MaybeHead returns the current Head Commit of this Datastore, which contains the current root of the DataStore's value tree, if available. If not, it returns a new Commit and 'false'.
	MaybeHead(datasetID string) (Commit, bool)

	// Head returns the current head Commit, which contains the current root of the DataStore's value tree.
	Head(datasetID string) Commit

	// Datasets returns the root of the datastore which is a MapOfStringToRefOfCommit where string is a datasetID.
	Datasets() MapOfStringToRefOfCommit

	// Commit updates the Commit that datasetID in this datastore points at. If the update cannot be performed, e.g., because of a conflict, error will non-nil. The newest snapshot of the datastore is always returned.
	Commit(datasetID string, commit Commit) (DataStore, error)

	// Delete removes the Dataset named datasetID from the map at the root of the DataStore. The Dataset data is not necessarily cleaned up at this time, but may be garbage collected in the future. If the update cannot be performed, e.g., because of a conflict, error will non-nil. The newest snapshot of the datastore is always returned.
	Delete(datasetID string) (DataStore, error)

	// CopyReachableChunksP copies to |sink| all chunks reachable from (and including) |r|, but that are not in the subtree rooted at |exclude|
	CopyReachableChunksP(r, exclude ref.Ref, sink chunks.ChunkSink, concurrency int)

	// CopyMissingChunksP copies to sink all chunks reachable from (and including) |r| that it does not already have
	CopyMissingChunksP(r ref.Ref, sink chunks.ChunkStore, concurrency int)
}

func NewDataStore(cs chunks.ChunkStore) DataStore {
	return newLocalDataStore(cs)
}

type Flags struct {
	ldb         chunks.LevelDBStoreFlags
	dynamo      chunks.DynamoStoreFlags
	hflags      chunks.HTTPStoreFlags
	memory      chunks.MemoryStoreFlags
	datastoreID *string
}

func NewFlags() Flags {
	return NewFlagsWithPrefix("")
}

func NewFlagsWithPrefix(prefix string) Flags {
	return Flags{
		chunks.LevelDBFlags(prefix),
		chunks.DynamoFlags(prefix),
		chunks.HTTPFlags(prefix),
		chunks.MemoryFlags(prefix),
		flag.String(prefix+"store", "", "name of datastore to access datasets in"),
	}
}

func (f Flags) CreateDataStore() (DataStore, bool) {
	var cs chunks.ChunkStore
	if cs = f.ldb.CreateStore(*f.datastoreID); cs != nil {
	} else if cs = f.dynamo.CreateStore(*f.datastoreID); cs != nil {
	} else if cs = f.memory.CreateStore(*f.datastoreID); cs != nil {
	}

	if cs != nil {
		return newLocalDataStore(cs), true
	}

	if cs = f.hflags.CreateStore(*f.datastoreID); cs != nil {
		return newRemoteDataStore(cs), true
	}

	return &LocalDataStore{}, false
}

func (f Flags) CreateFactory() (Factory, bool) {
	var cf chunks.Factory
	if cf = f.ldb.CreateFactory(); cf != nil {
	} else if cf = f.dynamo.CreateFactory(); cf != nil {
	} else if cf = f.memory.CreateFactory(); cf != nil {
	}

	if cf != nil {
		return &localFactory{cf}, true
	}

	if cf = f.hflags.CreateFactory(); cf != nil {
		return &remoteFactory{cf}, true
	}
	return &localFactory{}, false
}
