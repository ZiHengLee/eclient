package crypto

import (
	"fmt"
	"testing"
)

const skey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAxOm9GalC/a4MiH1QxXDZO4ENiyvEhlgmz6Dz7u2I3Gp2Ly34
SVxw4WMwrpGP/5JkcPA9XwEs1yhCgAgWSKQActRqDiknbYIIJMwnTSnl8Q2cels8
pgwm2btm65HV1uBA8/WWBIpWEvAsiL5nboGj7I6EJ449+4h/ptrcZzFeqQTCpdPK
cJ/1pv9FMAw82AiZYq/rpi/oNHXSUFauCys56OY+4Qoa/MwDwhKhHlFaZgcAUXqv
SMWxQ9KkpHyBpxQW1h6vyNPYtXrIIGHGUpqHyoWDYxZCxJtBnrN1LuS9L0gKaDgE
RSkHjg83apglK4w+ykhq3Y6GUk6zamDYB7KgRQIDAQABAoIBAQCgK3BBu0heiIBx
VV2p/Ez29dZKaeOFU5beNJG2u0gj4gUYA1B+e87lxoUGuihjPnSs2P1SleYfbQQK
mIlntN8YiGdr9VW44Zg2NkmRno8HMIcjj6sZdbD0Ulc5Lnq+OdMbnZvVPaEd7naL
FDWV0aXA+XvNoQoR28nZsxSsa6AP4fzDDnXm8TCDcMWAu9D9Ushi7szx6abkojfG
sJwTsUV08e18RT+a4eXcaEfUII0ZjDlt047W5Eh+upjTsHujCoBH9Q2Zs9VsZ4hW
QwIbKvcyIDO4vViKX1NHzAbYo84JxK+B9DmsGZKppuRzpF+UdIkiIcvQO5T7bLFq
z/CVjx7RAoGBAPTIsMjgXL5K3TxMMngyPdgYnnNfjEdWT7GsIqVUzBchyVguBfy5
dEJxLfU/gzh4Ro1FV+Z37paedPjALmvJdTqRRiKNySq67mUG3F/OXQJ9D3n96Fbp
r7KTvxnhdjS5Zjmd8M5peDSMVJOdR1WPYFA9Lluwxtz967/DGew202cfAoGBAM3v
hgszJB6J5vFj2prmleIg8gZrhwWqkS2QlsFx8Tc5kFYaBYB6iufMzQ/tv4Is8YOx
HklLOc2y7pFpITw/cT30aTBcSSvMHYcWxv9y/EGHDYyli0WBaAG5uPMIOk9OR2z8
bzmLdTR/LSeQfuDRQ9vZD0/hdr8s2JqGhNqxgUAbAoGAS/8++67EvIPLuhvCE5Ut
pDjIgPNqPv+qEHRr8T6peokP1OO9EiCV2sb+yMSmPdMBvHyJ9NDJb1PU5yHrh6Vt
UvileLbifPS6bP07UDdmow4mzu4ow0scNGowi513MJlbBoplqAAQxOfInmfXLNYt
xduo1+jDZPskxU2Sb8b5zWsCgYBSvhaBkX7FtmrQOmqvKk62mL4lzjZmFG5YfGvW
Nc3RfpN7xODJCnOdRzBtYf08dIple1jL4inLeEVar37nEiaR2g4ZnDraGh0TFhfG
s5CfU4AiLrGSr7f2WYTyQTAMhUs1gwe9e0cQXLjr2Kbh9bLqNoBsfA6WKZKuI7Zl
mERyWQKBgQCmua4C54rr82VLGW+rbqRnPdBHjUff9Gl9p98s0iPlOOU+PPEZbQOL
NKyRILEeI6i6zW2RSi2TmBNvK4CiQKxblxP4vJZ3+ZSbOAPgdsloWZVFutViMNe8
QnPjxPEOwjzjcmiVciRwDvLhSG3R+Bv3tLXUy7acj99n0VaOpxPqGA==
-----END RSA PRIVATE KEY-----`

func TestFieldCipher(t *testing.T) {
	c := NewFieldCipher()
	c.AddKey("coin", []byte(skey))
	src := "ungWv48Bz+pBQUDeXa4iI7ADYaOWF3qctBD/YfIAFa0="
	aesKeySrc, aesKey, ckey, err := c.GenKey()
	if err != nil {
		t.Fatalf("gen key err:%v", err)
	}
	fmt.Printf("aesKeySrc:%v\n", aesKeySrc)
	fmt.Printf("ckey:%v\n", ckey)
	res, err := c.Encrypt([]byte(src), aesKey)
	if err != nil {
		t.Fatalf("encrypt err:%v", err)
	}
	fmt.Printf("encrypt:%v\n", res)
	aesKeySrc, aesKey, err = c.ParseKey(ckey)
	if err != nil {
		t.Fatalf("parse key err:%v", err)
	}
	fmt.Printf("aesKeySrc:%v\n", aesKeySrc)
	dat, err := c.Decrypt(res, aesKey)
	if err != nil {
		t.Fatalf("decrypt err:%v", err)
	}
	fmt.Printf("decrypt:%v\n", string(dat))
}
