# ApiServer

# How to do Unit Tests

01. clone the repo
02. go to the repo root directory
03. execute: go get -v
04. now go to /repo-root/tests/unit/running
05. copy the path of "keys" directory under the root directory
06. open .env file inside the unit/running directory (not in the root directory) and change the KEYS_DIR to the copied path
07. execute: go test -v