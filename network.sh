#!/bin/bash
cd /home/julkhong/go/src/github.com/julkhong/fabric-samples/test-network
./network.sh down
./network.sh up createChannel
cd /home/julkhong/go/src/github.com/julkhong/blockchain-server/smartcontract
GO111MODULE=on go mod vendor
cd /home/julkhong/go/src/github.com/julkhong/fabric-samples/test-network
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
peer version
peer lifecycle chaincode package basic.tar.gz --path /home/julkhong/go/src/github.com/julkhong/blockchain-server/smartcontract --lang golang --label basic_1.0 
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
peer lifecycle chaincode install basic.tar.gz
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
peer lifecycle chaincode install basic.tar.gz

