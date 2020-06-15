package main

import (
	"fmt"
	"io/ioutil"
	"log"
	personpb "pjserol/tuto-proto3-go/src/person"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

const fname = "person.bin"

func main() {

	person := getPerson()
	fmt.Println("---Display proto person---")
	fmt.Println(person)

	fmt.Println("---Write to file---")
	writeToFile(fname, person)

	fmt.Println("---Read file---")
	p, err := readFromFile(fname)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Result after read file:", p)

	fmt.Println("---To JSON---")
	pStr := toJSON(person)
	fmt.Println(pStr)

	fmt.Println("---From JSON---")
	pJSON, err := fromJSON(pStr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pJSON)

}

func getPerson() *personpb.Person {
	p := personpb.Person{
		FirstName: "Tic",
		LastName:  "Pong",
		Age:       25,
		Hidden:    true,
		Hobbies:   []string{"books", "sports"},
		Gender:    personpb.Gender_MALE,
		Professions: []*personpb.Profession{
			&personpb.Profession{
				Year:  "2009-2010",
				Title: "Developer",
			},
			&personpb.Profession{
				Year:  "2010-2011",
				Title: "Senior Developer",
			},
		},
	}

	p.FirstName = "Ping"

	//fmt.Println(p)

	// better to use the getter to get the value to avoid nil pointer
	//fmt.Println(p.GetFirstName())

	return &p
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Couldn't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Couldn't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string) (*personpb.Person, error) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Couldn't read the file", err)
		return nil, err
	}

	p := personpb.Person{}
	if err := proto.Unmarshal(in, &p); err != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err)
		return nil, err
	}

	return &p, nil
}

func toJSON(pb proto.Message) string {

	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Couldn't convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON(str string) (*personpb.Person, error) {
	p := &personpb.Person{}
	if err := jsonpb.UnmarshalString(str, p); err != nil {
		log.Fatalln("Couldn't convert the string into the protocol buffers struct", err)
		return nil, err
	}

	return p, nil
}
