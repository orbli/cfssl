// Package sign implements the HTTP handler for the certificate signing command.
package sign

import (
	"net/http"

	"github.com/orbli/cfssl/api/signhandler"
	"github.com/orbli/cfssl/config"
	"github.com/orbli/cfssl/log"
	"github.com/orbli/cfssl/signer/universal"
)

// NewHandler generates a new Handler using the certificate
// authority private key and certficate to sign certificates. If remote
// is not an empty string, the handler will send signature requests to
// the CFSSL instance contained in remote by default.
func NewHandler(caFile, caKeyFile string, policy *config.Signing) (http.Handler, error) {
	root := universal.Root{
		Config: map[string]string{
			"cert-file": caFile,
			"key-file":  caKeyFile,
		},
	}
	s, err := universal.NewSigner(root, policy)
	if err != nil {
		log.Errorf("setting up signer failed: %v", err)
		return nil, err
	}

	return signhandler.NewHandlerFromSigner(s)
}

// NewAuthHandler generates a new AuthHandler using the certificate
// authority private key and certficate to sign certificates. If remote
// is not an empty string, the handler will send signature requests to
// the CFSSL instance contained in remote by default.
func NewAuthHandler(caFile, caKeyFile string, policy *config.Signing) (http.Handler, error) {
	root := universal.Root{
		Config: map[string]string{
			"cert-file": caFile,
			"key-file":  caKeyFile,
		},
	}
	s, err := universal.NewSigner(root, policy)
	if err != nil {
		log.Errorf("setting up signer failed: %v", err)
		return nil, err
	}

	return signhandler.NewAuthHandlerFromSigner(s)
}
