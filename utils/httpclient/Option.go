package httpclient

type Option struct {
	Host string `toml:"host"`
}

type PkgOption struct {
	Metrics string `json:"metrics"`
}
