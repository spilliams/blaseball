package model

// Completable represents an object that can be asked if it's complete or not.
type Completable interface {
	Incomplete() bool
}

// Unforbiddable represents an object which may contain forbidden knowledge,
// and can be unforbidden to show that knowledge.
type Unforbiddable interface {
	Unforbid()
}
