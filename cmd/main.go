package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"

	"github.com/cbergoon/merkletree"
)

type Element struct {
	Payload []byte
}

func (e Element) CalculateHash() ([]byte, error) {
	h := sha256.New()
	_, err := h.Write(e.Payload)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func (e Element) Equals(other merkletree.Content) (bool, error) {
	currentHash, err := e.CalculateHash()
	if err != nil {
		return false, nil
	}

	otherHash, err := other.CalculateHash()
	if err != nil {
		return false, nil
	}

	return bytes.Equal(currentHash, otherHash), nil
}

func main() {

	content := createContent(4)

	//Create a new Merkle Tree from the list of Content
	tree, err := merkletree.NewTree(content)
	if err != nil {
		log.Fatal(err)
	}

	//Get the Merkle Root of the tree
	mr := tree.MerkleRoot()
	log.Println(mr)

	path, index, err := tree.GetMerklePath(content[3])
	if err != nil {
		log.Fatal(err)
	}

	log.Println("----------")
	log.Println(path)
	log.Println("----------")
	log.Println(index)

	verificationResult, err := merkletree.VerifyContentWithPath(mr, content[2], path, index)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(verificationResult)

}

func createContent(count int) []merkletree.Content {
	var elements []merkletree.Content

	for i := 0; i < count; i++ {
		payload := make([]byte, 4)
		_, err := rand.Read(payload)
		if err != nil {
			panic(err)
		}

		element := Element{Payload: payload}
		elements = append(elements, element)
	}

	return elements
}
