# Fixes Applied to Resolve Pod Restart Loop Issue

## Date: October 20, 2025

## Problems Identified

1. **Checksum Mismatch Error**: The `go.sum` file had an incorrect checksum for `github.com/munnerz/goautoneg`
2. **HTTP Server Synchronization Issue**: The HTTP server was started in a goroutine without verifying it successfully bound to port 8080
3. **Fatal Error Handling**: The original code used `log.Fatalf` in the HTTP server goroutine, which would crash the entire process if the server failed to start
4. **Potential Panic in getDeploymentInfo**: The code accessed array elements without checking if they exist first

## Solutions Applied

### 1. Fixed Go Module Checksum (go.sum)

**File**: `go.sum`  
**Line**: 45

**Changed**:
```diff
- github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822/go.mod h1:+n7T8mK8HuQTcFX0m3SeQ17CVfieQyqhduHQiwPtd5Q=
+ github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822/go.mod h1:+n7T8mK8HuQTcFwEeznm/DIxMOiR9yIdICNftLE1DvQ=
```

**Why**: The checksum didn't match what Go was downloading from the module proxy, causing security errors during build.

### 2. Added HTTP Server Synchronization (main.go)

**Changes to `K8sDemo` struct**:
```go
type K8sDemo struct {
    clientset   *kubernetes.Clientset
    namespace   string
    serverReady chan bool  // NEW: Channel to signal server is ready
}
```

**Changes to `main()` function**:
- Added a channel-based synchronization mechanism to wait for the HTTP server to be ready
- The main function now waits for the server to bind successfully before continuing
- Added timeout protection (10 seconds) to prevent hanging if server fails
- Proper error channel to catch server startup failures

**Why**: This ensures the HTTP server is fully operational before Kubernetes health probes start checking, preventing premature restarts.

### 3. Improved HTTP Server Error Handling

**File**: `main.go`  
**Function**: `startHTTPServer()`

**Changed**:
```go
func (d *K8sDemo) startHTTPServer() error {  // Now returns error instead of calling log.Fatalf
    // ... handlers setup ...
    
    // Create a listener to check if port binding succeeds
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        return fmt.Errorf("failed to bind to port 8080: %w", err)
    }
    
    log.Println("HTTP server successfully bound to :8080")
    
    // Signal that server is ready
    d.serverReady <- true
    
    // Start serving
    log.Println("HTTP server is now accepting connections")
    server := &http.Server{}
    if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
        return fmt.Errorf("HTTP server error: %w", err)
    }
    return nil
}
```

**Why**: 
- Returns errors instead of crashing with `log.Fatalf`
- Explicitly binds to the port first to verify it's available
- Signals readiness only after successful binding
- Provides better error messages for debugging

### 4. Added Nil Checks in getDeploymentInfo

**File**: `main.go`  
**Function**: `getDeploymentInfo()`

**Changed**:
```go
replicas := int32(0)
if deploy.Spec.Replicas != nil {
    replicas = *deploy.Spec.Replicas
}
// ... print replicas ...
if len(deploy.Spec.Template.Spec.Containers) > 0 {
    fmt.Printf("    Image: %s\n", deploy.Spec.Template.Spec.Containers[0].Image)
}
```

**Why**: Prevents potential panics when accessing nil pointers or empty arrays.

### 5. Added net Import

**File**: `main.go`  
**Line**: 8

**Added**:
```go
import (
    // ... other imports ...
    "net"  // NEW: Required for net.Listen
    // ... other imports ...
)
```

**Why**: Required for the `net.Listen` function used to bind to the port.

## Additional Improvements

### 1. Created Debug Script (debug-pod.sh)

A comprehensive troubleshooting script that displays:
- Pod status and events
- Current and previous container logs
- RBAC configuration
- Docker image status
- Deployment details
- Troubleshooting tips

**Usage**:
```bash
./debug-pod.sh
```

### 2. Enhanced README (README.md)

Added comprehensive troubleshooting section including:
- Steps to diagnose CrashLoopBackOff
- Common causes of restart loops
- Manual debugging commands
- Quick fix commands
- Go module checksum error resolution

## Testing Recommendations

After applying these fixes, test the deployment with:

```bash
# Clean slate
make delete
make clean

# Rebuild and deploy
make deploy

# Monitor the deployment
watch kubectl get pods -n k8s-demo

# Check logs
make logs

# Or use the debug script
./debug-pod.sh
```

## Expected Behavior After Fixes

1. ✅ Docker build completes successfully without checksum errors
2. ✅ HTTP server binds to port 8080 before health checks start
3. ✅ Pods transition from Pending → Running without restarts
4. ✅ Readiness and liveness probes succeed
5. ✅ Application logs show successful Kubernetes API operations
6. ✅ No CrashLoopBackOff status

## Root Cause Analysis

The original restart loop was likely caused by:

1. **Primary**: HTTP server not being ready when health probes started checking
   - The server was started asynchronously without waiting for it to bind
   - If port binding took longer than expected, probes would fail
   - Failed probes → container restart → repeat

2. **Secondary**: Potential for fatal errors during server startup
   - Any error would call `log.Fatalf` and crash the entire process
   - This would trigger immediate restarts by Kubernetes

3. **Contributing**: Missing nil checks could cause panics in certain scenarios

## Files Modified

1. ✅ `go.sum` - Fixed checksum
2. ✅ `main.go` - Complete rewrite of server startup and error handling
3. ✅ `debug-pod.sh` - New diagnostic script
4. ✅ `README.md` - Enhanced troubleshooting documentation
5. ✅ `FIXES.md` - This file

## Verification Steps

After deploying, verify the fix worked:

```bash
# Pods should be running
kubectl get pods -n k8s-demo
# Expected: STATUS = Running, RESTARTS = 0

# Check logs for successful startup
kubectl logs -n k8s-demo -l app=k8s-demo
# Expected: "HTTP server successfully bound to :8080"
# Expected: "HTTP server is now accepting connections"
# Expected: "Application is running"

# Test health endpoints
kubectl port-forward -n k8s-demo service/k8s-demo 8080:8080 &
curl http://localhost:8080/health
# Expected: OK
```

## Prevention

To prevent similar issues in the future:

1. **Always synchronize critical services**: Wait for servers to be ready before continuing
2. **Return errors instead of using log.Fatal in goroutines**: Better error propagation
3. **Add comprehensive logging**: Log each step of initialization
4. **Test with realistic conditions**: Simulate slow startups, port conflicts, etc.
5. **Use debug tools**: Create scripts to quickly diagnose issues

## Contact

If you still experience restart loops after these fixes, run `./debug-pod.sh` and check the output for specific error messages.


