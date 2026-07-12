package webhookverify

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

// VerifyFacebookSignature ກວດ header X-Hub-Signature-256 (ຮູບແບບ "sha256=<hex>") ທຽບກັບ
// HMAC-SHA256(appSecret, rawBody) — ໃຊ້ hmac.Equal ເພື່ອປ້ອງກັນ timing attack
func VerifyFacebookSignature(appSecret string, rawBody []byte, signatureHeader string) bool {
	if appSecret == "" || signatureHeader == "" {
		return false
	}
	const prefix = "sha256="
	if !strings.HasPrefix(signatureHeader, prefix) {
		return false
	}
	expectedHex := strings.TrimPrefix(signatureHeader, prefix)
	expectedSig, err := hex.DecodeString(expectedHex)
	if err != nil {
		return false
	}

	mac := hmac.New(sha256.New, []byte(appSecret))
	mac.Write(rawBody)
	computedSig := mac.Sum(nil)

	return hmac.Equal(computedSig, expectedSig)
}

// VerifyLineSignature ກວດ header X-Line-Signature (base64 ຂອງ HMAC-SHA256(channelSecret, rawBody))
func VerifyLineSignature(channelSecret string, rawBody []byte, signatureHeader string) bool {
	if channelSecret == "" || signatureHeader == "" {
		return false
	}
	expectedSig, err := base64.StdEncoding.DecodeString(signatureHeader)
	if err != nil {
		return false
	}

	mac := hmac.New(sha256.New, []byte(channelSecret))
	mac.Write(rawBody)
	computedSig := mac.Sum(nil)

	return hmac.Equal(computedSig, expectedSig)
}
