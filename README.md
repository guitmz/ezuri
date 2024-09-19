# Ezuri
A Simple Linux ELF Runtime Crypter.
An unpacker by [f0wl](https://github.com/f0wl) can be found at [f0wl/ezuri_unpack](https://github.com/f0wl/ezuri_unpack).

# Download
```shell
curl -SsfL https://github.com/guitmz/ezuri/releases/latest/download/ezuri-v0.0.1-linux-amd64.tar.gz | tar xfvz -
```

# Compile from source:
Clone this repo.
```shell
git clone -b master --depth 1 https://github.com/guitmz/ezuri.git
cd ezuri
```

Build with
```shell
go mod init ezuri
go mod tidy
go build -o ezuri
```

# References
- https://www.guitmz.com/running-elf-from-memory/
- https://github.com/guitmz/memrun
- https://cybersecurity.att.com/blogs/labs-research/malware-using-new-ezuri-memory-loader
- https://www.bleepingcomputer.com/news/security/linux-malware-authors-use-ezuri-golang-crypter-for-zero-detection/
