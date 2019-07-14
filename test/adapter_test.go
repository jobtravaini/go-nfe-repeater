package test

import (
	"go-nfe-repeater/arquivei"
	"testing"
)

func testAdapterMapping(t *testing.T) {
	nfeData := arquivei.NfeData{AccessKey:"1", Xml: "<nfeProc versao=\"3.10\" xmlns=\"http://www.portalfiscal.inf.br/nfe\"></nfeProc>"}

	modelNfe := arquivei.MapNfe(nfeData)

	if nfeData.AccessKey != modelNfe.AccessKey || nfeData.Xml != modelNfe.XmlValue {
		t.Fatalf("The adapter mapping is incorrect")
	}
}