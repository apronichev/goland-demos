# Go 1.25: Cleaner Panic Output

Recovered and re-panicked errors now produce more readable output.

## The Change

### Before Go 1.25

```console
panic: DATABASE ERROR [recovered]
panic: DATABASE ERROR [recovered]
panic: DATABASE ERROR
```

### Go 1.25 and Later
```console
panic: DATABASE ERROR [recovered, repanicked]
```