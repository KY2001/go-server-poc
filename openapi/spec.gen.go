// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RUTW/bSAz9KwPuHmXJyS7QQLc2/Uh7ChoXPaQ+0CPaGkeamXAoIx/Qfy9mZMd14gRF",
	"kADtxaYs8vH5kXy3oF3rnSUrAcpbCLqmFlP4gdlxDDw7TyyG0s8thYALimFFQbPxYpyFEk66Fu2ICSuc",
	"NaQolqt1dg4ZyLUnKCEIG7uAPoMgKF04dtU+sMnkVA0JSrsqAdAVtr4hKP8fH9zBGSu0II54YqTZA3VW",
	"O5Y1nQilHCtTkRUzN8Q7wPDNYie1Y3ND1UPKfQZMl51hqqA8/5V/difL9K7KzZakJRL7RHJC2Ej9lYJ3",
	"NtBDVQesPeSJV8QKbaXev1sLskvZXfwm0ceonRq7eDaxl2e02Ij1tnuOYLHqxVnFZSXdsZHrs3ggA4UZ",
	"IRPHhtunj45bFCjhy/cJZMM5RaTh7ZZBLeKhj8DGzl2sX28vOE8WvVFCrW9Q4m6tiMPw7w7ycT6OKq2z",
	"oIT/8oN8DBl4lDrRKuLHguShNsc16QslNSlcoWlwZhoj18rNVUjTzH9YSNCMseJzBeVmQSAqNcwiNTkc",
	"j+OXdlbIplbofWN0KiyWIfbbuEmM/mWaQwn/FFu7KdZeU9zfwSTLLvPQaU1UQXoxx66RF+s+2Nyenp2l",
	"K09aqBrMI+b0GRR12s5ni5xuuULBGQZ6RPHhAF5Z83uW9FeoXmD3hPQTCrJX+blhinKrjbsnyk+Kn876",
	"FQew3+b+3CFsHRDK813vO5/2037IiAseUsIuYOM0NorsyrCzbSSeQcfN2gjLokgJtQtSHo2PosPdR/Ds",
	"qk7Hh0dhQlkUCzcaWIy806NDxjfVcnkjlyO0Oebc2Ry9h37a/wwAAP//HVzQRfkIAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
