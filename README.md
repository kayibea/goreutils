# goreutils

A tiny collection of minimal coreutils written in Go.

## Usage

```bash
# monolith style
goreutils echo hello
goreutils cat file.txt

# symlink style
ln -s goreutils echo
./echo hello
```

> Who needs Rust coreutils anyway? 😎
