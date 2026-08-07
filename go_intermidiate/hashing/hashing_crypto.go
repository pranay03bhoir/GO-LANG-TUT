
// hashing_crypto.go
// This program demonstrates in VERY DETAILED steps how to hash passwords using SHA-256/SHA-512
// in Go, and how to apply salting and encoding for secure password storage and verification.
//
// Embedded are thorough explanations of each operation for your learning!
package main

import (
	"crypto/rand"    // For generating cryptographically secure random numbers (used for salt)
	"crypto/sha256"  // For computing SHA-256 hashes
	"crypto/sha512"  // For computing SHA-512 hashes
	"encoding/base64"// For encoding binary data (hashes and salt) as printable text
	"fmt"
	"io"
)

func main() {
	// -------------------------------------------------
	// 1. The Password to Hash
	// -------------------------------------------------
	password := "password123"
	fmt.Println("=== PASSWORD HASHING & CRYPTO EXPLAINED in Go ===")
	fmt.Printf("Step 1: The original password is: %q\n\n", password)

	// -------------------------------------------------
	// 2. Hashing the Password with SHA-256 and SHA-512
	// -------------------------------------------------
	// Turn the password string into a byte slice (since hash functions work on bytes)
	fmt.Println("Step 2: Compute cryptographic hashes of the password (without salt).")
	hash256 := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Printf("   • SHA-256 hash (raw bytes):     %v\n", hash256)
	fmt.Printf("   • SHA-512 hash (raw bytes):     %v\n", hash512)
	// Printing hexadecimal representations (commonly used for hashes)
	fmt.Printf("   • SHA-256 hash (hex):           %x\n", hash256)
	fmt.Printf("   • SHA-512 hash (hex):           %x\n\n", hash512)

	// -------------------------------------------------
	// 3. Generating a Secure Random Salt
	// -------------------------------------------------
	fmt.Println("Step 3: Generate a cryptographically secure random salt (16 bytes).")
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("[!] Error generating salt:", err)
		return
	}
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Printf("   • Salt (base64-encoded):        %s\n\n", saltStr)
	fmt.Println("     A salt is random data added to the password before hashing,")
	fmt.Println("     making hash tables/rainbow table attacks much harder.")

	// -------------------------------------------------
	// 4. Hashing the PASSWORD + SALT (as in user signup)
	// -------------------------------------------------
	signUpHash := hashPassword(password, salt)
	fmt.Println("\nStep 4: Hash the password + salt, then encode the hash with base64.")
	fmt.Printf("   • Final salted password hash:    %s\n", signUpHash)
	fmt.Println("     This is what would be stored in a database along with the salt (in encoded form).")

	// -------------------------------------------------
	// 5. Simulate User Login: Verify entered password using stored salt and hash
	// -------------------------------------------------
	fmt.Println("\nStep 5: Simulate user login -- verify password using the stored hash and salt.")
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("[!] Error decoding base64 salt:", err)
		return
	}
	loginHash := hashPassword(password, decodedSalt)
	fmt.Printf("   • Recomputed hash for login:     %s\n", loginHash)
	if signUpHash == loginHash {
		fmt.Println("✅ Login successful: password is correct!")
	} else {
		fmt.Println("❌ Invalid credentials: password is incorrect.")
	}

	// -------------------------------------------------
	// 6. Summary and Learning Points
	// -------------------------------------------------
	fmt.Println("\n=== SUMMARY AND LEARNING POINTS ===")
	fmt.Println("1. Always use cryptographic hash functions, never store passwords in plain text.")
	fmt.Println("2. SHA-256 and SHA-512 are cryptographic hashes; SHA-256 is often enough for passwords (but see below).")
	fmt.Println("3. Always add a secure random salt to each password before hashing. Store the salt alongside the hash.")
	fmt.Println("4. Encode both the salt and the hash (e.g., using base64) to make them safe for storage/text transmission.")
	fmt.Println("5. At login, decode the stored salt, add it to the candidate password, hash it, and compare to the stored hash.")
	fmt.Println("6. For even better password protection, use dedicated password-hashing algorithms like bcrypt, scrypt, or Argon2 instead of a plain SHA256.")
	fmt.Println("7. Never use a fixed salt or no salt at all.")

}

/*
 * Function: generateSalt
 * Purpose: Generates a cryptographically secure random byte slice used as a salt.
 *          Salt ensures that the same password will have different hashes each time,
 *          preventing attackers from easily reverse-engineering hashes.
 */
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16) // 16 bytes = 128 bits, strong enough for most salt use cases
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err // Properly propagate any source errors
	}
	return salt, nil
}

/*
 * Function: hashPassword
 * Purpose: Combine the given salt and password, then compute a SHA-256 hash,
 *          and return its base64-encoded string (safe for storage/transmission).
 * Steps:
 *    1. Concatenate salt + password (salt first).
 *    2. Compute hash of combined bytes.
 *    3. Encode resulting hash using base64.
 */
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
