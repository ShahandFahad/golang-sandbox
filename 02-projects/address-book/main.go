package main

import (
	"fmt"
	"slices"
)

type contact struct {
	name   string
	number string
}

type AddressBook struct {
	contacts []contact
}

func (addressBook *AddressBook) add(name, number string) {

	addressBook.contacts = append(addressBook.contacts, contact{
		name:   name,
		number: number,
	})

}

func (addressBook *AddressBook) viewByName(name string) (contact, error) {

	if name == "" {
		return contact{}, fmt.Errorf("Please provide name first")
	}

	for _, contact := range addressBook.contacts {

		if contact.name == name {
			return contact, nil
		}
	}

	return contact{}, fmt.Errorf("Contact not found")
}

func (addressBook *AddressBook) viewByNumber(number string) (contact, error) {

	if number == "" {
		return contact{}, fmt.Errorf("Please provide number first")
	}

	for _, contact := range addressBook.contacts {

		if contact.number == number {
			return contact, nil
		}
	}

	return contact{}, fmt.Errorf("Contact not found")
}

func (addressBook *AddressBook) delete(name string, number string) error {

	if name == "" && number == "" {
		return fmt.Errorf("Please provide name or number")
	}

	for i := range addressBook.contacts {

		c := addressBook.contacts[i]
		if c.name == name || c.number == number {
			addressBook.contacts = slices.Delete(addressBook.contacts, i, i+1)
			return nil
		}

	}

	return fmt.Errorf("Contact to delete not found")
}

func (addressBook *AddressBook) displayAll() {
	for _, contact := range addressBook.contacts {
		fmt.Println("Contact: ", contact)
	}
}

func main() {
	addressBook := AddressBook{contacts: make([]contact, 0)}

	addressBook.add("John Doe", "123456789")
	addressBook.add("Jane Doe", "123456789")
	addressBook.add("Iron Man", "123456789")

	addressBook.displayAll()

	result, _ := addressBook.viewByName("Iron Man")
	fmt.Println("Found By Name: ", result)
	result, _ = addressBook.viewByNumber("123456789")
	fmt.Println("Found By Number: ", result)
	_, err := addressBook.viewByName("")
	fmt.Println(err)
	_, err = addressBook.viewByNumber("")
	fmt.Println(err)

	_ = addressBook.delete("Jane Doe", "")
	addressBook.displayAll()
}
