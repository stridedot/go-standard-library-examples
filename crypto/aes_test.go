package crypto_test

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"testing"
)

// AES 要求 key 的长度只能是 16, 24, 32 个字节
func TestAesKeySizeError(t *testing.T) {
	key := []byte("12345678901234567890123456")
	_, err := aes.NewCipher(key)
	if err == nil {
		t.Fatal("Got nil, want err")
	}
}

// AES GCM 模式加密解密
func TestAesGCMEncryptDecrypt(t *testing.T) {
	// 步骤一：加密
	key := []byte("12345678901234567890123456")
	plaintext := []byte("exampleplaintext")
	block, err := aes.NewCipher(key[:aes.BlockSize])
	if err != nil {
		t.Fatalf("aes.NewCipher err: %v", err)
	}

	// 创建 GCM 模式的 AEAD
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatalf("cipher.NewGCM err: %v", err)
	}

	// 创建随机数,这里在实际应用中让它只生成一次.不然每次都需要进行修改
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		t.Fatalf("io.ReadFull err: %v", err)
	}

	// 生成密文
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	t.Logf("ciphertext = %s\n", base64.RawURLEncoding.EncodeToString(ciphertext))

	// 步骤二：解密
	block, err = aes.NewCipher(key[:aes.BlockSize])
	if err != nil {
		t.Fatalf("aes.NewCipher err: %v", err)
	}
	aesgcm, err = cipher.NewGCM(block)
	if err != nil {
		t.Fatalf("cipher.NewGCM err: %v", err)
	}

	// 生成明文内容
	plaintext, err = aesgcm.Open(nil, nonce, ciphertext, nil)
	t.Logf("plaintext = %s\n", plaintext)
}

func TestAes(t *testing.T) {

}
