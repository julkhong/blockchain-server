## Blockchain server
A backend server that interacts with smart contract (hyperledege fabric with Golang).


## Architecture overview
![tech_spec drawio](https://github.com/julkhong/blockchain-server/assets/70477671/3ebbc8bc-b508-4463-a54c-4ce49e375a5f)

## Tech/framework used
Windows: VsCode, Ubuntu for cli, Docker, hyperledger fabric test net 

Mac: Goland, zsh 

## Features
The repo consists of two separate core functions. 
- Backend server
- Smart contract

Backend server has following apis:
1) POST /echo
2) POST /balance
3) GET /balance

Smart contract has following methods:
1) Init 
2) InitAccounts
3) SendBalance
4) GetBalance

## Project checklist
DONE:
1) Created a backend server with Golang
2) Created handlers and apis for Send() and GetBalance() methods
3) Created request validations accordingly
4) Created a smart contract that does Send() and GetBalance() method
5) Created unit tests for smart contract
6) Created benchmarking tests for smart contract
7) Deployed smart contract to test net
8) Tested smart contract in test net using cli

PENDING:
1) Unable to interact with smart contract in backend server
2) Units test for handlers

## Installation and dependencies
Please refer to official hyperledger fabric guide 
https://hyperledger-fabric.readthedocs.io/en/latest/getting_started.html 

## Tests
- Smart contract includes units tests and benchmarking test
- All the codes have passed lint checks


## Contact
Created by Jun Ming Khong (Julian). 
Feel free to send me an email!
 officialkhong@gmail.com 

## TODO
1) To complete the documentation, that includes:
    - API spec
    - Steps to run the project
    - Sequence diagram for the flows
    - more
