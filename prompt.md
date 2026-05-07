Create a production-quality Golang repository named `go-code-katas` for practicing programming problems, algorithms, data structures, concurrency exercises, and general coding katas.

The repository must be generic and should NOT reference any external coding platforms.

The repository should be scalable, maintainable, beginner-friendly, and professional enough to showcase on GitHub.

Use the latest stable Go version and standard Go project conventions.

---

# Repository Structure

Create the following structure:

go-code-katas/
├── README.md
├── go.mod
├── .gitignore
├── Makefile
├── problems/
│   ├── arrays/
│   ├── strings/
│   ├── linkedlist/
│   ├── trees/
│   ├── graphs/
│   ├── heap/
│   ├── stack/
│   ├── queue/
│   ├── dynamic-programming/
│   ├── greedy/
│   ├── backtracking/
│   ├── sliding-window/
│   ├── binary-search/
│   ├── concurrency/
│   ├── system-design/
│   └── miscellaneous/
│
├── internal/
│   ├── ds/
│   │   ├── heap/
│   │   ├── trie/
│   │   ├── linkedlist/
│   │   ├── graph/
│   │   └── queue/
│   │
│   └── algorithms/
│       ├── bfs/
│       ├── dfs/
│       ├── dijkstra/
│       ├── unionfind/
│       └── sorting/
│
├── templates/
│   ├── bfs.go
│   ├── dfs.go
│   ├── heap.go
│   ├── union_find.go
│   └── binary_search.go
│
├── snippets/
│   ├── goroutines/
│   ├── channels/
│   ├── mutex/
│   ├── context/
│   └── generics/
│
├── benchmarks/
│   └── README.md
│
└── docs/
    ├── complexity-cheatsheet.md
    ├── go-tips.md
    └── problem-solving-patterns.md

---

# Problem Structure

Each problem must have its own directory.

Example:

problems/arrays/two-sum/
├── README.md
├── solution.go
├── solution_test.go
└── notes.md

The README for each problem should contain:

- problem summary
- approach
- complexity analysis
- edge cases
- learning notes

---

# Root README Requirements

Generate a professional README.md including:

- repository purpose
- learning goals
- folder structure explanation
- how to run tests
- how to run benchmarks
- how to add new problems
- coding conventions
- Go version used

Include badges if appropriate.

---

# Makefile

Create a Makefile with the following commands:

- test
- benchmark
- lint
- format
- tidy
- vet

Example usage:

make test
make benchmark

---

# Sample Implementations

Create at least 3 fully working sample problems:

1. Array/hashmap problem
2. Graph traversal problem
3. Concurrency problem using goroutines/channels

Each sample must include:

- implementation
- unit tests
- benchmark tests where appropriate
- README
- complexity analysis

---

# Coding Standards

Follow these rules:

- idiomatic Go
- table-driven tests
- readable code
- proper comments
- small focused packages
- avoid overengineering
- prefer simplicity and maintainability

---

# Testing Requirements

Use:

- Go standard testing package
- table-driven tests
- benchmark tests

Include:

- edge case tests
- performance benchmark examples

---

# Documentation

Generate starter documentation for:

- common algorithm patterns
- time complexity cheatsheet
- Go interview tips
- concurrency patterns
- problem-solving strategies

---

# Additional Requirements

- Keep repository generic
- Do not mention external coding challenge platforms
- Design for long-term extensibility
- Ensure scalability for hundreds of problems
- Use clean naming conventions
- Include starter reusable templates
- Include reusable data structure implementations
- Ensure all generated code compiles successfully

---

# Final Output

Generate:

- full repository structure
- starter files
- working Go code
- tests
- benchmark examples
- README files
- documentation
- Makefile
- reusable templates

All code should be production-quality and immediately usable.
