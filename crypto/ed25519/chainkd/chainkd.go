package chainkd

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"io"

	"chain/crypto/ed25519"
	"chain/crypto/ed25519/internal/edwards25519"
)

type (
	// TODO(bobg): consider making these types opaque. See https://github.com/chain/chain/pull/1875#discussion_r80577736
	XPrv [64]byte
	XPub [64]byte
)

// NewXPrv takes a source of random bytes and produces a new XPrv. If
// r is nil, crypto/rand.Reader is used.
func NewXPrv(r io.Reader) (xprv XPrv, err error) {
	if r == nil {
		r = rand.Reader
	}
	var entropy [32]byte
	_, err = io.ReadFull(r, entropy[:])
	if err != nil {
		return xprv, err
	}
	return RootXPrv(entropy[:]), nil
}

// RootXPrv takes a seed binary string and produces a new XPrv.
func RootXPrv(seed []byte) (xprv XPrv) {
	h := hmac.New(sha512.New, []byte("Root"))
	h.Write(seed)
	h.Sum(xprv[:0])
	modifyRootScalar(xprv[:32])
	return
}

func (xprv XPrv) XPub() (xpub XPub) {
	var buf [32]byte
	copy(buf[:], xprv[:32])

	var P edwards25519.ExtendedGroupElement
	edwards25519.GeScalarMultBase(&P, &buf)
	P.ToBytes(&buf)

	copy(xpub[:32], buf[:])
	copy(xpub[32:], xprv[32:])

	return
}

func (xprv XPrv) Child(sel []byte, hardened bool) XPrv {
	if hardened {
		return xprv.HardenedChild(sel)
	} else {
		return xprv.NonhardenedChild(sel)
	}
}

func (xprv XPrv) HardenedChild(sel []byte) (res XPrv) {
	h := hmac.New(sha512.New, xprv[:32])
	h.Write([]byte("H"))
	h.Write(xprv[32:])
	h.Write(sel)
	h.Sum(res[:0])
	modifyRootScalar(res[:32])
	return
}

func (xprv XPrv) NonhardenedChild(sel []byte) (res XPrv) {
	xpub := xprv.XPub()

	h := hmac.New(sha512.New, xpub[32:])
	h.Write([]byte("N"))
	h.Write(xpub[32:])
	h.Write(sel)
	h.Sum(res[:0])

	modifyFactorScalar(res[:32])

	var carry int
	carry = 0
	for i := 0; i < 32; i++ {
		sum := int(xprv[i]) + int(res[i]) + carry
		res[i] = byte(sum & 0xff)
		carry = (sum >> 8)
	}
	if carry != 0 {
		panic("sum does not fit in 256-bit int")
	}
	return
}

func (xpub XPub) Child(sel []byte) (res XPub) {
	var f [32]byte
	var F edwards25519.ExtendedGroupElement

	h := hmac.New(sha512.New, xpub[32:])
	h.Write([]byte("N"))
	h.Write(xpub[32:])
	h.Write(sel)
	h.Sum(res[:0])

	modifyFactorScalar(res[:32])

	copy(f[:], res[:32])
	edwards25519.GeScalarMultBase(&F, &f)

	var (
		pubkey [32]byte
		P      edwards25519.ExtendedGroupElement
	)
	copy(pubkey[:], xpub[:32])
	P.FromBytes(&pubkey)

	var (
		P2 edwards25519.ExtendedGroupElement
		R  edwards25519.CompletedGroupElement
		Fc edwards25519.CachedGroupElement
	)
	F.ToCached(&Fc)
	edwards25519.GeAdd(&R, &P, &Fc)
	R.ToExtended(&P2)

	P2.ToBytes(&pubkey)

	copy(res[:32], pubkey[:])

	return res
}

func (xprv XPrv) Derive(path [][]byte) XPrv {
	res := xprv
	for _, p := range path {
		res = res.Child(p, false)
	}
	return res
}

func (xpub XPub) Derive(path [][]byte) XPub {
	res := xpub
	for _, p := range path {
		res = res.Child(p)
	}
	return res
}

func (xprv XPrv) Sign(msg []byte) []byte {
	return Ed25519InnerSign(xprv.ExpandedPrivateKey(), msg)
}

func (xpub XPub) Verify(msg []byte, sig []byte) bool {
	return ed25519.Verify(xpub.PublicKey(), msg, sig)
}

func (xprv XPrv) ExpandedPrivateKey() ExpandedPrivateKey {
	var res [64]byte
	h := hmac.New(sha512.New, []byte("Expand"))
	h.Write(xprv[:])
	h.Sum(res[:0])
	copy(res[:32], xprv[:32])
	return res[:]
}

// PublicKey extracts the ed25519 public key from an xpub.
func (xpub XPub) PublicKey() ed25519.PublicKey {
	return ed25519.PublicKey(xpub[:32])
}

// s must be >= 32 bytes long and gets rewritten in place.
// This is NOT the same pruning as in Ed25519: it additionally clears the third
// highest bit to ensure subkeys do not overflow the second highest bit.
func modifyRootScalar(s []byte) {
	s[0] &= 248
	s[31] &= 31 // clear top 3 bits
	s[31] |= 64 // set second highest bit
}

// Clears lowest 3 bits and highest 23 bits of `f`.
func modifyFactorScalar(f []byte) {
	f[0] &= 248 // clear bottom 3 bits
	f[29] &= 1  // clear 7 high bits
	f[30] = 0   // clear 8 bits
	f[31] = 0   // clear 8 bits
}
