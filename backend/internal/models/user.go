package models

import (
    "time"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID              int       `json:"id"`
    Email           string    `json:"email"`
    PasswordHash    string    `json:"-"`
    DateOfBirth     time.Time `json:"date_of_birth"`
    TermsAccepted   bool      `json:"terms_accepted"`
    DeviceFingerprint string  `json:"-"`
    IPAddress       string    `json:"-"`
    EmailVerified   bool      `json:"email_verified"`
    VerificationToken string  `json:"-"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

func (u *User) HashPassword(password string) error {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.PasswordHash = string(hash)
    return nil
}

func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
    return err == nil
}
