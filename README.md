# Sentence Analyzer in Go 📚✂️

This Go project analyzes two text files by:
- Counting the number of words in the **first 100 sentences** of each file.
- Displaying the results as **stem-and-leaf plots** in the terminal.

## 📁 Project Structure

```
author-style-detector/
├── cmd/
│   └── main.go
└── internal/
    └── analyzer/
        ├── analyzer.go
        └── plot.go
```

## 🚀 How to Use

1. **Clone the repository**:
   ```bash
   git clone https://github.com/meruyert4/author-style-detector.git
   cd author-style-detector
   ```

2. **Add your text files** (e.g., `sample1.txt`, `sample2.txt`) to the root directory.

3. **Run the program**:
   ```bash
   go run cmd/main.go sample1.txt sample2.txt
   ```

4. The program will output two stem-and-leaf plots — one for each file — in the terminal.

## 📝 Example Output

```
Analyzing: sample1.txt
 2 | 3 5 7
 3 | 0 2 4 8
 4 | 1 1 5 9
...

Analyzing: sample2.txt
 1 | 9
 2 | 0 2 8
 3 | 1 6
...
```

## 📦 Requirements

- Go 1.18 or newer

## 🧠 How It Works

- The text from each file is split into sentences using `.`, `?`, and `!` as delimiters.
- The first 100 sentences are selected.
- Each sentence is split into words, and the word count is recorded.
- A stem-and-leaf plot groups the counts by tens (e.g., 23 → stem: 2, leaf: 3).
