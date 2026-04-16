# helper

Reusable HTTP response helpers for Go APIs.

## Install

```bash
go get github.com/your-github-username/httphelper@latest
```

## Usage

```go
package main

import (
	"net/http"

	helper "github.com/your-github-username/httphelper"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
	})
}
```

## Optional custom logger

```go
helper.SetErrorLogger(func(err error, context string) {
	// Hook into your logging stack
})
```

## Publish to GitHub

1. Create a GitHub repository (for example: `httphelper`).
2. Update `go.mod` module path to match your repository:
   - `module github.com/<your-username>/<repo-name>`
3. Commit and push:

```bash
git init
git add .
git commit -m "Initial reusable helper package"
git branch -M main
git remote add origin https://github.com/<your-username>/<repo-name>.git
git push -u origin main
```

4. (Recommended) Create semantic tags for stable versions:

```bash
git tag v1.0.0
git push origin v1.0.0
```

Then import from other projects with:

```bash
go get github.com/<your-username>/<repo-name>@latest
```
