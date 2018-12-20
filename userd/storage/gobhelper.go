package userstorage

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"log"

	"github.com/theryecatcher/chirper/userd/userdpb"
)

func init() {
	gob.Register(userdpb.User{})
}

// ToGOB64 binary encoder
func ToGOB64(m *userdpb.User) string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		log.Println(`failed gob Encode`, err)
	}
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

// FromGOB64 binary decoder
func FromGOB64(str string) *userdpb.User {
	m := userdpb.User{}
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println(`failed base64 Decode`, err)
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&m)
	if err != nil {
		log.Println(`failed gob Decode`, err)
	}
	return &m
}
