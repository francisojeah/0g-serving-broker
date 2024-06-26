package proxy

var (
	servicePrefix = "/v1/proxy"

	requestMetaData = map[string]struct{}{
		"Address":                     {},
		"Created-At":                  {},
		"Name":                        {},
		"Nonce":                       {},
		"Previous-Output-Token-Count": {},
		"Previous-Signature":          {},
		"Signature":                   {},
		"Token-Count":                 {},
	}
)
