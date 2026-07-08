<div>
    <img src="toolshed.png" width="120" alt="toolshed Logo">
</div>

This is where I dump small tools I build while learning new things.

**Tools**

| File | Language | What it does |
| ------ | ---------- | -------------- |
| [`portscanner.py`](https://github.com/Shaivarth/toolshed/blob/main/portscanner.py) | Python | Scans TCP ports on a host |
| [`filehashgenerator.go`](https://github.com/Shaivarth/toolshed/blob/main/filehashgenerator.go) | Go | Computes MD5, SHA-1, SHA-256, SHA-512 hashes for a file |
| [`dnslookup.py`](https://github.com/Shaivarth/toolshed/blob/main/dnslookup.py) | Python | Performs DNS lookups for a domain |
| [`caesar.rb`](https://github.com/Shaivarth/toolshed/blob/main/caesar.rb) | Ruby | Encodes and decodes text using a Caesar cipher |
| [`counter.rb`](https://github.com/Shaivarth/toolshed/blob/main/counter.rb) | Ruby | Counts lines, words, and characters in a file |
| [`passwordgenerator.py`](https://github.com/Shaivarth/toolshed/blob/main/passwordgenerator.py) | Python | Generates secure random passwords |

**usage**

Each tool is self-contained. Run with `--help` for usage details.

- `python3` [`portscanner.py`](https://github.com/Shaivarth/toolshed/blob/main/portscanner.py) `scanme.nmap.org 1 65535`
- `python3` [`dnslookup.py`](https://github.com/Shaivarth/toolshed/blob/main/dnslookup.py) `shaivarth.com`
- `go run` [`filehashgenerator.go`](https://github.com/Shaivarth/toolshed/blob/main/filehashgenerator.go) `/path/to/file`
- `ruby` [`caesar.rb`](https://github.com/Shaivarth/toolshed/blob/main/caesar.rb) `encode 3 "Hello elephant"`  
  *(use any value in place of 3, even -ve values will give diff results)*
- `ruby` [`counter.rb`](https://github.com/Shaivarth/toolshed/blob/main/counter.rb) `notes.txt`  
  *(use any of your files 🙂)*
- `python3` [`passwordgenerator.py`](https://github.com/Shaivarth/toolshed/blob/main/passwordgenerator.py)