### toolshed

This is where I dump small tools I build while learning new things.

**Tools**

| File | Language | What it does |
| ------ | ---------- | -------------- |
| `filehashgenerator.go` | Go | Computes MD5, SHA-1, SHA-256, SHA-512 hashes for a file |
| `caesar.rb` | Ruby | Encodes and decodes text using a Caesar cipher |
| `counter.rb` | Ruby | Counts lines, words, and characters in a file |
| `passwordgenerator.py` | Python | Generates secure random passwords |

**usage**

Each tool is self-contained. Run with `--help` for usage details.

```bash
go run filehashgenerator.go /path/to/file

ruby caesar.rb encode 3 "Hello elephant" 
[use any value in place of 3, even -ve values will give diff results]

ruby counter.rb notes.txt 
[use any of your filess] :)

python passwordgenerator.py
```
