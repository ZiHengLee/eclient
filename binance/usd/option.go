package usd

type Option struct {
	CryptoKeys map[string]string
}

var gOpt *Option

func Init(opt *Option) (err error) {
	gOpt = opt
	return
}
