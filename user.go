package main

import (
	"crypto/rand"
	"encoding/binary"

	"github.com/duo-labs/webauthn/webauthn"
)

// User represents the user model
type User struct {
	id          uint64
	name        string
	displayName string
	credentials []webauthn.Credential
}

// NewUser creates and returns a new User
func NewUser(name string, displayName string) *User {

	user := &User{}
	user.id = randomUint64()
	user.name = name
	user.displayName = displayName
	// user.credentials = []webauthn.Credential{}

	return user
}

func randomUint64() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}

// WebAuthnID returns the user's ID
func (u User) WebAuthnID() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, uint64(u.id))
	return buf
}

// WebAuthnName returns the user's username
func (u User) WebAuthnName() string {
	return u.name
}

// WebAuthnDisplayName returns the user's display name
func (u User) WebAuthnDisplayName() string {
	return u.displayName
}

// WebAuthnIcon() is not (yet) implemented
func (u User) WebAuthnIcon() string {
	return ""
}

// TODO: should 'cred' be a pointer?
func (u *User) AddCredential(cred webauthn.Credential) {
	u.credentials = append(u.credentials, cred)
}

// WebAuthnCredentials returns credentials owned by the user
func (u User) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}
