# Stack Transformation

You are extending a small generic `Stack` implementation in `stack.go`.

## Your Task

Implement support for transforming the values stored in a `Stack` using a caller-provided function.

The transformation should:

1. Visit every element in the stack exactly once.
2. Apply the provided function to each element.
3. Return the transformed values.

**Important Requirements:**
- The transformation must work for any output type.
- The original stack must not be modified.
