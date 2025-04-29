# Go-Django Integration Demo

This application demonstrates the integration between a Go program and a Django server. The Go program manages the Django server lifecycle and communicates with its API.

## Prerequisites

- Go 1.24 or later
- Python 3.x – https://www.python.org/downloads/
- pip3 (Python package manager) 
- virtualenv (recommended for Python environment management)

## Setup

### 1. Python Environment Setup

```bash
# Create a virtual environment
python3 -m venv .venv
# Activate the virtual environment
# On macOS/Linux:
source .venv/bin/activate
# On Windows:
.venv\Scripts\activate
# Install Django
pip3 install django
``` 

### 2. Environment Variables

The application requires the following environment variables to be set:
```bash
# Get your Python path
which python3 # Note this path
# Check if Django is installed and where
pip3 show django
``` 

Replace the paths in `python-in-goland/main.go` (line 41).
```python
env = append(env, "PYTHONPATH=/Users/jetbrains/myProjects/goland-demos/.venv/lib/python3.13/site-packages:"+os.Getenv("PYTHONPATH"))
``` 

### 3. Django Setup
```bash 
cd python-in-goland/pyapi 
python3 manage.py migrate
``` 

## Running the Application

1. Open `python-in-goland/main.go` and run the `main` function.

The application will:
- Start the Django server
- Display server startup progress
- Show a message when the server is running 
 
Click the **Stop** icon in the Run tool window to stop the server.

## API Endpoints

The Django server provides the following endpoints:

- `http://localhost:8000/api/hello/` - Returns a JSON response with a greeting message
- `http://localhost:8000/api/debug/` - (Debug endpoint) Shows all registered URLs

## Stopping the Application

Click the **Stop** icon in the Run tool window to stop the server.

## Development

To modify the API:
1. Add new views in `pyapi/api/views.py`.
2. Register URLs in `pyapi/api/urls.py`.
3. Restart the application to apply changes.