package embed

import _ "embed"

//go:embed countries.yml
var CountriesData []byte
