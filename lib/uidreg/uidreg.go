package uidreg

import (
    "encoding"
    "fmt"
    "testing"
    )

// The type UidReg (Unique Identifier Register) is the principal type in the
// uidreg package. It provides a service of issueing new unique Ids on demand.
// The identifiers live in categories (of your choice), and their uniqueness is 
// preserved only inside their own category. In order to support the typical 
// use case, where state must be persisted, the type implements the TextMarshaler
// and TextUnmarshaler from Go's encoding package.

type UidReg struct {
    // Keeps track of the largest integer issued by category.
    used map[string]int64
}

// The function NewEmpty() makes an initialized, empty, Uid register.
func NewEmpty() *UidReg {
    return &UidReg{used: make(map[string]int64)}
}

// The function NewId() issues a new UID for the given category.
func (uidReg *UidReg) NewId(category string) (id string) {
    // Use a value one-larger than that previously used
    prev := uidReg.used[category] // copes with absence
    next := prev + 1
    uidReg.used[category] = next
    return fmt.Sprintf("%d", next)
}

func (uidReg *UidReg) UnmarshalText(text []byte) error {
    return nil
}

func (uidReg *UidReg) MarshalText() (text []byte, err error) {
    text = []byte{}
    return text, nil
}

