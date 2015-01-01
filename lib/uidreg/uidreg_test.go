package uidreg

import (
    "encoding"
    "testing"
    )

func TestShowsExampleUsage(t *testing.T) {
    // Typically we create Uid registers by de-serializing a previously
    // persisted one, to discover the details of previously issued Uids. 
    mockSerialized := []byte(`"used": {"apple": "1"}`)
    var uidReg UidReg
    if uidReg.UnmarshalText(mockSerialized) != nil {
        t.Errorf("UnmarshallText() failed.")
    }

    // Now we can ask for a new Id to be issued, and can expect its value
    // to be sensible w.r.t. those previously issued.
    appleId := uidReg.NewId("apple");
    if appleId != "2" {
        t.Errorf("Wrong Id. Expected 2, but got %s", appleId)
    }
    pearId := uidReg.NewId("pear");
    if pearId != "1" {
        t.Errorf("Wrong Id. Expected 1, but got %s", pearId)
    }

    // And we typically re-serialise the updated Uid register in readiness to
    // persist it again.
    updatedSerialized, err := uidReg.MarshalText()
    if err != nil {
        t.Errorf("MarshallText() failed.")
    }
    if updatedSerialized != thisfoothing {
        t.Errorf("MarshalText() produced wrong text. Expected <%s>, but got <%s>",
                thisfoothing, updatedSerialized)
    }
}

func TestZeroValue(t *testing.T) {
    var uidReg UidReg
    appleId := uidReg.NewId("apple");
    if appleId != "1" {
        t.Errorf("Wrong Id. Expected 1, but got %s", appleId)
    }
}

func TestStringsAreShort(t *testing.T) {
    // The implmementation strives to make the Ids short by encoding the
    // integers it uses internally into base64 encoded ASCII. So we check this
    // encoding here.
    mockedSerialized := `"used": {"apple": "9"}`
    uidReg, err := UidReg.UnmarshallText(mockSerialized)
    if err != nil {
        t.Errorf("UnmarshallText() failed.")
    }
    appleId := uidReg.NewId("apple");
    if appleId != "xyzbase64encoded" {
        t.Errorf("Wrong Id. Expected %s, but got %s", xyzbaseencoded, appleId)
    }
}
