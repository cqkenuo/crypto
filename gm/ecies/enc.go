/*
Copyright Baidu Inc. All Rights Reserved.

jingbo@baidu.com
*/
package ecies

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"

	libecies "github.com/xuperchain/crypto/core/ecies/libecies"
)

func Encrypt(publicKey *ecdsa.PublicKey, msg []byte) (cypherText []byte, err error) {
	if publicKey.Curve.Params().Name != "P-256" {
		err = fmt.Errorf("curve [%v] is not supported yet.", publicKey.Curve.Params().Name)
		return nil, err
	}

	pub := libecies.ImportECDSAPublic(publicKey)

	ct, err := libecies.Encrypt(rand.Reader, pub, msg, nil, nil)
	if err != nil {
		return nil, err
	}

	return ct, nil
}

func Decrypt(privateKey *ecdsa.PrivateKey, cypherText []byte) (msg []byte, err error) {
	if privateKey.PublicKey.Curve.Params().Name != "P-256" {
		err = fmt.Errorf("curve [%v] is not supported yet.", privateKey.PublicKey.Curve.Params().Name)
		return nil, err
	}

	prv := libecies.ImportECDSA(privateKey)

	pt, err := prv.Decrypt(rand.Reader, cypherText, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return pt, nil
}
