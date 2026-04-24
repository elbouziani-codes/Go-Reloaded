# Go-Reloaded

A text processing and auto-correction tool written in Go.
This program reads a text file, applies multiple transformations based on specific rules, and outputs the corrected result into another file.

---

## 🚀 Features

* Convert numbers:

  * `(hex)` → hexadecimal to decimal
  * `(bin)` → binary to decimal

* Text transformations:

  * `(up)` → uppercase
  * `(low)` → lowercase
  * `(cap)` → capitalize
  * Support multiple words: `(up, 2)`, `(low, 3)`, `(cap, n)`

* Smart punctuation formatting:

  * Fix spacing before/after punctuation
  * Handle grouped punctuation like `...`, `!?`
  * Proper formatting of `'quotes'`

* Grammar correction:

  * Automatically converts `a` → `an` before vowels or `h`

---

## 📂 Project Structure

```
go-reloaded/
├── main.go
├── utils/
├── input.txt
├── result.txt
```

---

## 🛠️ Usage

### Run the program:

```bash
go run . input.txt output.txt
```

---

## 📌 Examples

### Input:

```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

### Output:

```
Simply add 66 and 2 and you will see the result is 68.
```

---

### Input:

```
There is no greater agony than bearing a untold story inside you.
```

### Output:

```
There is no greater agony than bearing an untold story inside you.
```

---

### Input:

```
Punctuation tests are ... kinda boring ,what do you think ?
```

### Output:

```
Punctuation tests are... kinda boring, what do you think?
```

---

## 🧠 How It Works

The program parses the input text and applies transformations sequentially:

1. Detect and process special tags `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)`
2. Apply transformations on previous word(s)
3. Normalize punctuation spacing and grouping
4. Fix grammar rules like `a → an`

---

## 🧪 Testing

It is recommended to create custom test files:

```bash
go run . sample.txt result.txt
cat result.txt
```

---

## ⚙️ Requirements

* Go (Golang)
* Standard Go packages only

---

## 🎯 Learning Objectives

This project helps you understand:

* File handling in Go (fs API)
* String manipulation
* Text parsing techniques
* Algorithmic thinking
* Clean code practices

---

## 👤 Author

Mohammed Amine El Bouziani
📍 Oujda, Morocco
💼 GitHub: https://github.com/elbouziani-codes

---

## 📄 License

This project is part of the Zone01 Oujda curriculum.
