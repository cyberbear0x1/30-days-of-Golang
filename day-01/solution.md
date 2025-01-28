## Day 1: Setting Up My Golang Environment 🛠️

Today, I kicked off my 30 Days of Golang by setting up my development environment on Kali Linux. 

* **Installation:** 
    * Updated the package lists: `sudo apt update`
    * Installed Go using apt: `sudo apt install golang`

* **Verification:** 
    * Ran `go version` to confirm successful installation.

* **Workspace Setup:**
    * Created a workspace directory: `mkdir -p ~/go/src`
    * Set the GOPATH environment variable:
        * **Temporarily:** `export GOPATH=~/go`
        * **Permanently (add to `.bashrc` or `.zshrc`):** 
          ```bash
          export GOROOT=/usr/lib/go
          export GOPATH=$HOME/go
          export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
          ```