package server

import (
	"fmt"
	"github.com/cyantarek/ghost-clone-project/config"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

func (s *Server) SetupJwt() error {
	fmt.Println("JWT enabled")
	signBytes, err := ioutil.ReadFile(config.PRIVATE_KEY)
	if err != nil {
		return err
	}

	s.SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}

	verifyBytes, err := ioutil.ReadFile(config.PUBLIC_KEY)
	if err != nil {
		return err
	}

	s.VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}
	return nil
}
