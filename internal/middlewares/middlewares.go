package middlewares

import (
	"context"
	"net/http"

	"github.com/paulsonlegacy/go-social/internal/app"
)


// Context key type to avoid collisions
type contextKey string

const appContextKey contextKey = "app"

// Middleware to inject `*app.Application` into request context
func InjectApp(app *app.Application) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Attach `app` to the request context
			ctx := context.WithValue(r.Context(), appContextKey, app)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

/*

	Middlewares in Go **must conform to `func(http.Handler) http.Handler`**.  
	This pattern allows **chaining** middleware together seamlessly.

	### **How It Works Step by Step**

	```go
	func InjectApp(app *app.Application) func(http.Handler) http.Handler {
	```

	👆 This means **`InjectApp` takes `app *app.Application` and returns another function**  
	🔹 That returned function must match **`func(http.Handler) http.Handler`**, so Go's router can use it.  

	---

	### **Breaking Down the Inner Function**
	```go
	return func(next http.Handler) http.Handler {
	```
	✅ This is the function that **middleware chains together**.  
	✅ It takes a `next http.Handler`—which is just **the next middleware or handler** in the pipeline.  
	✅ It must **return another `http.Handler`** so Go can continue processing requests.

	---

	### **Final Inner Function - The Actual Request Handling**
	```go
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), appContextKey, app)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
	```
	✅ This function actually runs when a request comes in.  
	✅ It **injects `app` into the request context** before calling `next.ServeHTTP()`.  
	✅ `next.ServeHTTP(w, r.WithContext(ctx))` **passes control to the next middleware/handler**, with `app` now inside the context.

	---

	Think of it as:  

	1️⃣ **Outer function** gets `app`.  
	2️⃣ **Middle function** takes the next handler.  
	3️⃣ **Innermost function** runs on every request, injecting `app` and calling the next handler.

	---  

*/

// Helper function to retrieve `*app.Application` from context
func GetAppFromContext(r *http.Request) *app.Application {
	if app, ok := r.Context().Value(appContextKey).(*app.Application); ok {
		return app
	}
	return nil
}
