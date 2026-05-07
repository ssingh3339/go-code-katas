# go-code-katas

[![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A production-quality Go repository for practicing programming problems, algorithms, data structures, concurrency exercises, and general coding katas. This repository is designed to be scalable, maintainable, and beginner-friendly.

## 🎯 Purpose

This repository serves as:

- **Learning Platform:** Practice Go programming and computer science fundamentals
- **Interview Preparation:** Common coding interview questions and patterns
- **Reference Guide:** Reusable implementations of data structures and algorithms
- **Best Practices:** Examples of idiomatic Go code and testing strategies
- **Portfolio Project:** Professional codebase suitable for showcasing on GitHub

## 🏗️ Repository Structure

```
go-code-katas/
├── problems/              # Problem implementations by category
│   ├── arrays/           # Array and hash map problems
│   ├── strings/          # String manipulation problems
│   ├── linkedlist/       # Linked list problems
│   ├── trees/            # Binary tree and BST problems
│   ├── graphs/           # Graph traversal and algorithms
│   ├── heap/             # Heap-based problems
│   ├── stack/            # Stack problems
│   ├── queue/            # Queue problems
│   ├── dynamic-programming/  # DP problems
│   ├── greedy/           # Greedy algorithm problems
│   ├── backtracking/     # Backtracking problems
│   ├── sliding-window/   # Sliding window pattern
│   ├── binary-search/    # Binary search problems
│   ├── concurrency/      # Goroutines and channels
│   ├── system-design/    # System design exercises
│   └── miscellaneous/    # Other problems
│
├── internal/             # Reusable internal packages
│   ├── ds/              # Data structures
│   │   ├── heap/        # Min/Max heap implementations
│   │   ├── trie/        # Trie (prefix tree)
│   │   ├── linkedlist/  # Linked list utilities
│   │   ├── graph/       # Graph representations
│   │   └── queue/       # Queue and deque
│   │
│   └── algorithms/      # Algorithm implementations
│       ├── bfs/         # Breadth-first search
│       ├── dfs/         # Depth-first search
│       ├── dijkstra/    # Dijkstra's algorithm
│       ├── unionfind/   # Union-find (disjoint set)
│       └── sorting/     # Sorting algorithms
│
├── templates/           # Reusable algorithm templates
│   ├── bfs.go          # BFS template
│   ├── dfs.go          # DFS template
│   ├── heap.go         # Heap template
│   ├── union_find.go   # Union-find template
│   └── binary_search.go # Binary search template
│
├── snippets/            # Concurrency and Go-specific patterns
│   ├── goroutines/     # Goroutine patterns
│   ├── channels/       # Channel patterns
│   ├── mutex/          # Mutex and sync patterns
│   ├── context/        # Context usage
│   └── generics/       # Generic programming examples
│
├── benchmarks/          # Performance benchmarks
├── docs/               # Documentation
│   ├── complexity-cheatsheet.md
│   ├── go-tips.md
│   └── problem-solving-patterns.md
│
├── go.mod              # Go module file
├── .gitignore          # Git ignore rules
├── Makefile            # Build and test automation
└── README.md           # This file
```

## 📚 Problem Structure

Each problem follows a consistent structure:

```
problems/category/problem-name/
├── README.md           # Problem description, approach, complexity
├── solution.go         # Implementation
├── solution_test.go    # Unit tests and benchmarks
└── notes.md           # Learning notes and insights
```

### Example Problem

See [problems/arrays/two-sum/](problems/arrays/two-sum/) for a complete example.

## 🚀 Getting Started

### Prerequisites

- Go 1.23 or higher
- Make (optional, but recommended)

### Installation

```bash
# Clone the repository
git clone https://github.com/satyendrasingh/go-code-katas.git
cd go-code-katas

# Initialize Go modules
go mod tidy

# Verify installation
make test
```

## 🧪 Running Tests

```bash
# Run all tests
make test

# Run tests for a specific package
go test ./problems/arrays/two-sum/

# Run tests with coverage
make test  # includes coverage report

# Run tests with race detector
go test -race ./...

# Run tests verbosely
go test -v ./...
```

## 📊 Running Benchmarks

```bash
# Run all benchmarks
make benchmark

# Run benchmarks for a specific package
go test -bench=. ./problems/arrays/two-sum/

# Run benchmarks with memory statistics
go test -bench=. -benchmem ./problems/arrays/two-sum/

# Compare benchmark results
go test -bench=. ./problems/arrays/two-sum/ > old.txt
# Make changes...
go test -bench=. ./problems/arrays/two-sum/ > new.txt
benchstat old.txt new.txt
```

## 🛠️ Development Workflow

### Adding a New Problem

1. **Create directory structure:**
   ```bash
   mkdir -p problems/category/problem-name
   ```

2. **Create required files:**
   - `README.md` - Problem description, approach, complexity analysis
   - `solution.go` - Implementation
   - `solution_test.go` - Tests and benchmarks
   - `notes.md` - Learning notes

3. **Follow naming conventions:**
   - Package name: lowercase, matches directory name
   - Test functions: `TestFunctionName`
   - Benchmark functions: `BenchmarkFunctionName`

4. **Write tests first (TDD):**
   - Use table-driven tests
   - Cover edge cases
   - Include benchmarks

5. **Document thoroughly:**
   - Clear problem statement
   - Approach explanation
   - Time and space complexity
   - Edge cases

### Code Quality

```bash
# Format code
make format

# Run linter (requires golangci-lint)
make lint

# Run go vet
make vet

# Run all checks
make check
```

## 📖 Documentation

- **[Complexity Cheatsheet](docs/complexity-cheatsheet.md)** - Big-O notation reference
- **[Go Tips](docs/go-tips.md)** - Idiomatic Go patterns and best practices
- **[Problem Solving Patterns](docs/problem-solving-patterns.md)** - Common algorithmic patterns

## 🎓 Learning Goals

This repository helps you:

- ✅ Master Go programming language fundamentals
- ✅ Understand data structures and their implementations
- ✅ Practice common algorithms (sorting, searching, graph traversal)
- ✅ Learn algorithmic patterns (two pointers, sliding window, etc.)
- ✅ Develop testing and benchmarking skills
- ✅ Write concurrent programs using goroutines and channels
- ✅ Analyze time and space complexity
- ✅ Prepare for technical interviews

## 📝 Coding Conventions

### General Principles

1. **Idiomatic Go:** Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
2. **Simplicity:** Prefer clear, simple code over clever solutions
3. **Readability:** Code is read more often than written
4. **Testing:** All functions have comprehensive tests
5. **Documentation:** Exported functions are documented

### Naming

- **Packages:** Short, lowercase, single word
- **Functions:** CamelCase for exported, camelCase for unexported
- **Variables:** Short names in small scopes, descriptive in larger scopes
- **Constants:** MixedCaps or UPPER_CASE for exported

### Testing

- **Table-driven tests:** Use for multiple test cases
- **Test names:** Descriptive, use `t.Run()` for subtests
- **Benchmarks:** Include for performance-critical code
- **Coverage:** Aim for high coverage, but quality over quantity

### Comments

```go
// Package doc comment
package example

// PublicFunc does something important.
// It takes x and returns y.
func PublicFunc(x int) int {
    // Implementation comments when needed
    return x * 2
}
```

## 🤝 Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-problem`)
3. Follow the project structure and conventions
4. Add tests and documentation
5. Ensure all tests pass (`make test`)
6. Format code (`make format`)
7. Commit changes (`git commit -m 'Add new problem'`)
8. Push to branch (`git push origin feature/new-problem`)
9. Open a Pull Request

## 📦 Dependencies

This repository uses only the Go standard library. No external dependencies are required.

## 🔧 Makefile Commands

| Command | Description |
|---------|-------------|
| `make test` | Run all tests with coverage |
| `make benchmark` | Run all benchmarks |
| `make lint` | Run golangci-lint |
| `make format` | Format all Go files |
| `make tidy` | Tidy go.mod and go.sum |
| `make vet` | Run go vet |
| `make check` | Run vet, format check, and tests |
| `make clean` | Clean build artifacts |
| `make help` | Display available commands |

## 🌟 Sample Problems

The repository includes fully working implementations:

1. **[Two Sum](problems/arrays/two-sum/)** - Hash map pattern, O(n) solution
2. **[Number of Islands](problems/graphs/number-of-islands/)** - Graph traversal with DFS/BFS/Union-Find
3. **[Web Crawler](problems/concurrency/web-crawler/)** - Concurrent programming with goroutines and channels

## 📊 Performance

All solutions include:

- Time complexity analysis
- Space complexity analysis
- Benchmark tests
- Comparison with alternative approaches

## 🎯 Goals and Roadmap

- [x] Set up repository structure
- [x] Add core data structures
- [x] Implement algorithm templates
- [x] Create concurrency examples
- [x] Add sample problems
- [x] Write comprehensive documentation
- [ ] Add 50+ problems across categories
- [ ] Include system design examples
- [ ] Add visualization tools
- [ ] Create video tutorials

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by various coding practice platforms
- Built with Go 1.23
- Follows Go community best practices

## 📞 Contact

For questions, suggestions, or discussions:

- Open an issue
- Submit a pull request
- Star the repository if you find it useful!

---

**Happy Coding! 🚀**
