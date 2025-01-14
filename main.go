package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Request struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	MsListProduct int    `json:"msListProduct"`
}

type Response struct {
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	ResponseData        string `json:"ResponseData"`
	ResponseException   string `json:"ResponseException"`
}

type EncryptedData struct {
	IV      string `json:"iv"`
	Content string `json:"content"`
}

func main() {
	http.HandleFunc("/api/product", handleProduct)
	log.Printf("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleProduct(w http.ResponseWriter, r *http.Request) {

	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	reqEncrypt := req.Username + "|" + req.Password + "|"
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	keyLength := 32

	key, err := generateRandomKey(keyLength)
	if err != nil {
		log.Fatal(err)
	}

	// key := []byte("thisis32bitlongpassphraseimusing")

	plaintext := []byte(reqEncrypt)

	encryptedData, err := Encrypt(key, plaintext)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}

	encryptedJSON, _ := json.MarshalIndent(encryptedData, "", "  ")
	fmt.Println("Encrypted Data (JSON):")
	fmt.Println(string(encryptedJSON))

	response := Response{
		ResponseCode:        "00",
		ResponseDescription: "Success",
		ResponseData:        "http://localhost:8081/?pr=" + encryptedData.IV + "%S4|" + encryptedData.Content,
		ResponseException:   "",
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func generateRandomKey(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func Encrypt(key []byte, plaintext []byte) (*EncryptedData, error) {
	// Pastikan panjang key yang digunakan adalah valid (16, 24, atau 32 byte untuk AES)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("key length must be 16, 24, or 32 bytes")
	}

	// Membuat IV (Initialization Vector) acak
	iv := make([]byte, aes.BlockSize) // AES Block size adalah 16 byte
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	// Membuat cipher block dengan kunci
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Menyusun plaintext untuk panjang yang sesuai dengan blok AES (multiple dari BlockSize)
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	paddedText := append(plaintext, make([]byte, padding)...)

	// Mengenkripsi data
	cbc := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(paddedText))
	cbc.CryptBlocks(ciphertext, paddedText)

	// Mengonversi hasil enkripsi dan IV ke dalam format hex
	ivHex := hex.EncodeToString(iv)
	ciphertextHex := hex.EncodeToString(ciphertext)

	// Mengembalikan struct EncryptedData
	return &EncryptedData{
		IV:      ivHex,
		Content: ciphertextHex,
	}, nil
}
