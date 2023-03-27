package packages

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
func BufferExample() {
	var b bytes.Buffer

	b.Write([]byte("Hello "))
	fmt.Fprint(&b, "world!")
	b.WriteTo(os.Stdout)
}

func TestBuffer() {
	BufferExample()
}

// Protocol buffers is a mechanism that allows to generate sourse code for structures.
// Also it is lightweight

// to create go file from proto: protoc --go_out=. .\packages\*.proto
func ProtoBuf() error {
	Lucius := &Person{
		Name: "Lucius",
		Age: 22,
		Socialfollowers: &SocialFollowers{
			Twitter: 1000,
			Youtube: 2000,
		},
	}

	data, err := proto.Marshal(Lucius)
	if err != nil {
		return err
	}

	GetLucius := &Person{}

	proto.Unmarshal(data, GetLucius)

	fmt.Printf("%v\n", Lucius)
	return nil
}

func TestProtoBuf() {
	if err := ProtoBuf(); err != nil {
		log.Fatal(err)
	}
}