# 🚀 Learn Go (Golang)

A beginner-friendly Go learning repository covering essential fundamentals, variable declarations, multi-file package structures, functions, and standard conventions.

---

## 📁 Project Structure

```text
learn-go/
│
├── practice/
│   ├── go.mod             # Go module definition
│   ├── main.go            # Application entry point (simplified)
│   ├── variable.go        # Variable types, zero values, and formatting demos
│   ├── function.go        # Function declarations and execution examples
│   ├── if-else-switch.go  # Conditional statements (if, else, switch) examples
│   ├── for-loop.go        # For loop, break, and continue problem examples
│   ├── array.go           # Array vs Slice, length, and capacity examples
│   └── struct.go          # Struct, methods, and constructor function examples
│
└── readme.md           # Documentation, commands, and important Go concepts
```

---

## 💻 Essential Go Terminal Commands

Here are the most frequently used commands when developing with Go:

### 1. Version & Environment
```bash
# Check installed Go version
go version

# View Go environment variables (GOPATH, GOROOT, etc.)
go env
```

### 2. Module Management
```bash
# Initialize a new Go module in the current directory
go mod init <module-name>

# Add missing modules and remove unused dependencies
go mod tidy

# Download dependencies listed in go.mod
go mod download

# Verify dependencies have expected content
go mod verify
```

### 3. Running & Compiling Code
```bash
# Run all Go files in the current directory (Recommended for multi-file package main)
go run .

# Run specific files together
go run main.go variable.go function.go

# Compile into an executable binary
go build

# Compile with a custom executable name
go build -o myapp.exe .
```

> ⚠️ **Important Note for Multi-File Programs:**  
> If your `main.go` calls functions defined in other files (like `variable.go` or `function.go`) under `package main`, running `go run main.go` alone will result in `undefined: ...` errors.  
> Always use **`go run .`** inside that directory!

### 4. Code Formatting & Quality
```bash
# Automatically format all Go files according to standard Go style
go fmt ./...

# Examine code for suspicious constructs and common mistakes
go vet ./...
```

### 5. Testing & Packages
```bash
# Run all tests in the current project
go test ./...

# Run tests with verbose output
go test -v ./...

# Add or update an external dependency
go get github.com/gin-gonic/gin

# Clean build artifacts and temporary files
go clean
```

---

## 🔑 Important Go Concepts & Notes

### 1. Package & Entry Point
- Every Go file must declare a package at the very top: `package <name>`.
- Executable programs require `package main` and a `func main()` function inside.
- All files residing in the **same directory** must belong to the **same package**.
- Functions, structs, and variables declared at package level in one file are **directly accessible** in any other file of the same package without needing an `import`.

### 2. Exported vs Unexported (Visibility)
Go does not use `public` or `private` keywords:
- **Exported (Public):** Starts with an **uppercase** letter (e.g., `SayHello`, `User`). Accessible outside the package.
- **Unexported (Package-Private):** Starts with a **lowercase** letter (e.g., `makeCoffee`, `age`). Accessible only within the same package.

### 3. Variable Declarations
```go
// 1. Standard explicit declaration
var age int = 25

// 2. Type inference
var name = "Alice"

// 3. Short declaration operator (Inside functions only)
country := "Bangladesh"

// 4. Grouped declaration
var (
    city    string = "Dhaka"
    pincode int    = 1200
)
```

### 4. Default Zero Values
If a variable is declared without an initial value, Go automatically assigns its "zero value":

| Data Type | Zero Value |
| :--- | :--- |
| `int`, `float64`, `uint`, etc. | `0` / `0.0` |
| `string` | `""` (empty string) |
| `bool` | `false` |
| Pointer, Slice, Map, Channel, Function, Interface | `nil` |
| Array | Array with zero values for each element (e.g., `[0 0 0]`) |

### 5. Common Printf Format Specifiers (Verbs)

| Verb | Description | Example Output |
| :--- | :--- | :--- |
| `%v` | Default format representation | `[0 0 0]` or `Alice` |
| `%+v` | Struct with field names | `{Name:Alice Age:18}` |
| `%T` | Type of the value | `string` or `int` |
| `%t` | Boolean (`true` or `false`) | `false` |
| `%d` | Base-10 integer | `18` |
| `%f` / `%.2f` | Floating-point number | `3.14` |
| `%s` | String characters | `Alice` |
| `%q` | Double-quoted string safe | `"Alice"` |
| `\n` | Newline character | Line break |

---

## 🏃 How to Run This Project

1. Open terminal in the `practice` directory:
   ```bash
   cd practice
   ```
2. Run the program:
   ```bash
   go run .
   ```

