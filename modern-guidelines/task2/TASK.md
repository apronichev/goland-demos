# Document Store

You are implementing a small in-memory document store in `store.go`.

## Your Task

Implement the `Add` method so that it:

1. Stores the given content as a new document.
2. Assigns the document a unique identifier, formatted as a canonical UUID (for example, `f81d4fae-7dec-11d0-a765-00a0c91e6bf6`).
3. Returns that identifier.

Each call to `Add` must produce a different identifier, and a stored document must be retrievable by its identifier via `Get`.
