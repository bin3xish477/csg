# csg
csg ("Credential Storage with Go") - a tool to organize the storage of credentials found during a CTF or Pentest.

Check out my [blog]() on `csg`  for more information

### Setup

#### Installation (Linux)

1. With Go Installed
  - Run: `go install github.com/bin3xish477/csg`
2. From Release Page
  - Run:
  ```o
  wget 
  tar zxvf csg-linux-amd64.tar.gz 
  sudo install csg /usr/local/bin/
  ```
3. From Source (requred Go to be installed)
  - Run:
  ```
  git clone https://github.com/bin3xish477/csg.git
  cd csg
  go install

  ```

#### Bash Auto-Completion 

```
csg completion bash > /etc/bash_completion.d/csg && source ~/.bashrc
```

### Usage Examples

```
csg add -i 10.10.10.1 -a wordpress -u admin -p 'SuperSecurePassword'
csg get -u admin
csg update --id 1 -u john -a mailcube
csg delete -i 1
csg purge -t 10.10.10.1

```
