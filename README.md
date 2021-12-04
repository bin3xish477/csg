# csg
csg ("Credential Storage with Go") - a tool to organize the storage of credentials found during a CTF or Pentest.

![image](https://user-images.githubusercontent.com/44281620/144725235-a6d945b7-f225-49d5-815f-a3ef5921188f.png)

Check out my [blog](https://bin3xish477.medium.com/csg-credential-storage-for-ctf-pentests-e12d48383c8f) on `csg`  for more information.

### Setup

#### Installation (Linux)

1. With Go Installed
  - `go install github.com/bin3xish477/csg@latest`
2. From Release Page
  ```bash
  wget https://github.com/bin3xish477/csg/releases/download/v1.0/csg-linux-amd64.tar.gz
  tar zxvf csg-linux-amd64.tar.gz 
  sudo install csg /usr/local/bin/
  ```
3. From Source (requred Go to be installed)
  ```bash
  git clone https://github.com/bin3xish477/csg.git
  cd csg
  go install

  ```

#### Bash Auto-Completion 

Linux

```bash
csg completion bash > /etc/bash_completion.d/csg && source ~/.bashrc
```

MacOS

```bash
csg completion bash > /usr/local/etc/bash_completion.d/csg && source ~/.bashrc
```

### Usage Examples

```bash
csg add -i 10.10.10.1 -a wordpress -u admin -p 'SuperSecurePassword'
csg get -u admin
csg update --id 1 -u john -a mailcube
csg delete -i 1
csg purge -t 10.10.10.1
csg export -f all-creds.csv
```
