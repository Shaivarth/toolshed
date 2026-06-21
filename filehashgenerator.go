package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
	"time"
)

type algo struct {
	name    string
	newHash func() hash.Hash
}

var algorithms = []algo{
	{"MD5     ", md5.New},
	{"SHA-1   ", sha1.New},
	{"SHA-256 ", sha256.New},
	{"SHA-512 ", sha512.New},
}

func main() {
	path, err := parseArgs(os.Args[1:])
	if err != nil {
		if errors.Is(err, errHelp) {
			printUsage()
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "error: %v\n\nRun 'filehash --help' for usage.\n", err)
		os.Exit(1)
	}

	result, err := hashFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", friendlyErr(path, err))
		os.Exit(2)
	}

	printResults(result)
}

var errHelp = errors.New("help requested")

func parseArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("no file path provided")
	}

	if args[0] == "--help" || args[0] == "-h" {
		return "", errHelp
	}

	if len(args) > 1 {
		return "", fmt.Errorf("too many arguments; expected exactly one file path")
	}

	return args[0], nil
}

type hashResult struct {
	path    string
	size    int64
	elapsed time.Duration
	digests []digest
}

type digest struct {
	name string
	hex  string
}

func hashFile(path string) (hashResult, error) {
	f, err := os.Open(path)
	if err != nil {
		return hashResult{}, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return hashResult{}, err
	}

	hashers := make([]hash.Hash, len(algorithms))
	for i, a := range algorithms {
		hashers[i] = a.newHash()
	}

	multi := io.MultiWriter(toWriters(hashers)...)

	start := time.Now()
	if _, err := io.Copy(multi, f); err != nil {
		return hashResult{}, fmt.Errorf("reading file: %w", err)
	}
	elapsed := time.Since(start)

	digests := make([]digest, len(algorithms))
	for i, h := range hashers {
		digests[i] = digest{
			name: algorithms[i].name,
			hex:  fmt.Sprintf("%x", h.Sum(nil)),
		}
	}

	return hashResult{
		path:    path,
		size:    info.Size(),
		elapsed: elapsed,
		digests: digests,
	}, nil
}

func toWriters(hashers []hash.Hash) []io.Writer {
	writers := make([]io.Writer, len(hashers))
	for i, h := range hashers {
		writers[i] = h
	}
	return writers
}

const (
	colWidth   = 72
	borderChar = "─"
)

func printResults(r hashResult) {
	labelWidth := 9
	totalWidth := labelWidth + colWidth + 5

	sep := strings.Repeat(borderChar, totalWidth)

	fmt.Println()
	fmt.Printf("  File    : %s\n", r.path)
	fmt.Printf("  Size    : %s\n", formatBytes(r.size))
	fmt.Printf("  Elapsed : %s\n", r.elapsed.Round(time.Millisecond))
	fmt.Println()
	fmt.Println(sep)
	fmt.Printf("  %-*s  %s\n", labelWidth, "Algorithm", "Hash Digest")
	fmt.Println(sep)

	for _, d := range r.digests {
		fmt.Printf("  %-*s  %s\n", labelWidth, d.name, d.hex)
	}

	fmt.Println(sep)
	fmt.Println()
}

func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func friendlyErr(path string, err error) error {
	if os.IsNotExist(err) {
		return fmt.Errorf("file not found: %q", path)
	}
	if os.IsPermission(err) {
		return fmt.Errorf("permission denied: %q", path)
	}
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		return fmt.Errorf("invalid path %q: %v", path, pathErr.Err)
	}
	return err
}

func printUsage() {
	fmt.Print(`
filehash — compute cryptographic hashes for a file

USAGE
  filehash <file>
  filehash --help

ARGUMENTS
  <file>     Path to the file you want to hash.

FLAGS
  --help, -h  Show this help message and exit.

HASHES COMPUTED
  MD5, SHA-1, SHA-256, SHA-512

EXAMPLES
  filehash /etc/hosts
  filehash ~/Downloads/ubuntu.iso

EXIT CODES
  0   Success
  1   Bad arguments / usage error
  2   File I/O error (not found, permission denied, etc.)

`)
}