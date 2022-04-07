# *** This repository has been deprecated and is no longer supported. We are now supporting [LACNet's branch](https://github.com/LACNetNetworks/rotation-validator). ***

# Rotation Mechanism for Validator Nodes

This component is a client used to rotate the validator nodes in LACChain networks. 

## Prerequisites

* Being a validator node in LACChain network
* Go 1.13+ installation or later
* **GOPATH** environment variable is set correctly

## Install

```
$ git clone https://github.com/lacchain/rotation-validator

$ cd rotation-validator/client
$ go build
```

## Run

Execute the executable file generated previously in a Validator node

```
$ ./client
```

## Know More

* [In depth overview of the rotation mechanism](https://github.com/lacchain/rotation-validator/blob/main/docs/OVERVIEW.md).
* [Architecture](docs/Architecture.md).

## Copyright 2021 LACChain

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
