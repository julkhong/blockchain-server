package handlers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	pb "github.com/julkhong/blockchain-server"
	log "github.com/sirupsen/logrus"
)

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return privateKey, &privateKey.PublicKey
}

//func createRequest(id int, fromID, toID string, amount float64, privateKey *rsa.PrivateKey) Request {
//	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
//	params := []float64{amount}
//	key := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
//	signature := generateSignature(id, timestamp, params, key, privateKey)
//
//	return Request{
//		ID:        id,
//		Timestamp: timestamp,
//		Params:    params,
//		Key:       key,
//		Signature: signature,
//	}
//}

func generateSignature(id int, timestamp int64, params []float64, key string, privateKey *rsa.PrivateKey) string {
	str := fmt.Sprintf("%d%d", id, timestamp)
	for _, param := range params {
		str += fmt.Sprintf("%f", param)
	}
	str += key

	hash := sha256.Sum256([]byte(str))
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	return base64.StdEncoding.EncodeToString(signature)
}

func validateSignature(request *pb.ChainCodeRequest, params map[string]interface{}) bool {
	// Decode the PEM-encoded string to obtain the DER-encoded bytes
	block, _ := pem.Decode([]byte(request.Key))
	if block == nil {
		log.WithFields(log.Fields{"request": request}).Error("unable to decode public key")
		return false
	}
	// Parse the DER-encoded bytes to extract the RSA public key
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.WithFields(log.Fields{"request": request}).Error("failed to parse RSA public key")
		return false
	}
	// Build the string representation of the request fields
	str := fmt.Sprintf("%d%d", request.Id, request.TimeStamp)
	for _, val := range params {
		str += fmt.Sprintf("%f", val)
	}
	str += request.Key

	// Decode the request signature
	signature, err := base64.StdEncoding.DecodeString(request.Signature)
	if err != nil {
		log.WithFields(log.Fields{"request": request}).Error("unable to decode signature")
		return false
	}

	// Verify the request signature
	hash := sha256.Sum256([]byte(str))
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		log.WithFields(log.Fields{"request": request}).Error("unable to verify signature")
		return false
	}
	return true
}

func validateChainCodeRequest(request *pb.ChainCodeRequest) bool {
	var params map[string]interface{}
	if request.Id == 0 || request.TimeStamp == nil || request.Params == nil || request.Key == "" || request.Signature == "" {
		log.WithFields(log.Fields{"request": request}).Error("Bad request from client")
		return false
	}
	_, err := base64.StdEncoding.DecodeString(request.Key)
	if err != nil {
		log.WithFields(log.Fields{"request": request}).Error("Invalid public key from client")
		return false
	}
	if request.Params.SenderID != "" {
		params["senderID"] = request.Params.SenderID
	}
	if request.Params.AccountID != "" {
		params["accountID"] = request.Params.AccountID
	}
	if request.Params.Amount != 0 {
		params["amount"] = request.Params.Amount
	}
	if request.Params.ReceiverID != "" {
		params["receiverID"] = request.Params.ReceiverID
	}
	return validateSignature(request, params)
}
