package wallet

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/justindwlee/bitcoinClone/utils"
)

const (
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
	privateKey string = "30770201010420b550812c6c3ab97911935474f0cb2e2b10bbb5fd01aa738a2f31c9b4b584b4cca00a06082a8648ce3d030107a144034200040a91210b1b9906f7fb69d8d8805af919a3546bb4b89583b2c91097cc2c30a53e715aa5d48bce1c4447987a843203278b261c6084f9b9cb6777f5e410a73a76da"
	signature string = "546d20df9e98670dad790cc57956865eca2431215318558dc63e4022162bd1a25ad048b55e67da4f1b1bac88775b99a5bc79e53db38ea16813020c953da6cf8b"
)

func Start(){
	privByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	private, err := x509.ParseECPrivateKey(privByte)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	utils.HandleErr(err)

	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]
	
	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	hashBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	ok := ecdsa.Verify(&private.PublicKey, hashBytes, &bigR, &bigS)
	fmt.Println(ok)
}