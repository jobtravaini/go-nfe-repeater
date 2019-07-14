package arquivei

import "go-nfe-repeater/nfe"

func MapNfe(data NfeData) nfe.Nfe {
	return nfe.Nfe{AccessKey: data.AccessKey, XmlValue: data.Xml}
}
