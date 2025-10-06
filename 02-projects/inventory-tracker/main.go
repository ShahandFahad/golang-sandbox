package main

import "fmt"

type Item struct {
	Name     string
	Quantity int
	Price    float64
}

type InventoyrManager struct {
	Items map[string]Item
}

type InventoryAction interface {
	Process(*InventoyrManager) error
}

type Restock struct {
	Name     string
	Quantity int
}

type Sell struct {
	Name     string
	Quantity int
}

func (r *Restock) Process(manager *InventoyrManager) error {
	item, ok := manager.Items[r.Name]
	if !ok {
		return fmt.Errorf("Item '%s' not found in inventory", r.Name)
	}

	if r.Quantity < 0 {
		return fmt.Errorf("Please provide correct quantity")
	}

	// update quantity
	item.Quantity += r.Quantity
	manager.Items[r.Name] = item

	return nil
}

func (s *Sell) Process(manager *InventoyrManager) error {

	item, ok := manager.Items[s.Name]
	if !ok {
		return fmt.Errorf("Item '%s' not found in inventory", s.Name)
	}

	if s.Quantity < 0 {
		return fmt.Errorf("Please provide correct quantity")
	}

	if item.Quantity < s.Quantity {
		return fmt.Errorf("Insufficient stock")
	}

	item.Quantity -= s.Quantity
	manager.Items[s.Name] = item

	return nil
}

func getByName(name string, manager *InventoyrManager) {

	item, ok := manager.Items[name]
	if ok {
		fmt.Println(item)
	} else {
		fmt.Println("Item not found in inventory: ", name)
	}

}

func main() {

	manager := InventoyrManager{Items: make(map[string]Item)}

	manager.Items["test-1"] = Item{Name: "a", Quantity: 10, Price: 10}
	manager.Items["test-2"] = Item{Name: "b", Quantity: 11, Price: 20}
	manager.Items["test-3"] = Item{Name: "c", Quantity: 12, Price: 30}
	manager.Items["test-4"] = Item{Name: "d", Quantity: 13, Price: 40}
	manager.Items["test-5"] = Item{Name: "e", Quantity: 14, Price: 50}

	fmt.Println(manager)

	getByName("test-1", &manager)
	getByName("test-2", &manager)
	getByName("test-3", &manager)
	getByName("test-4", &manager)
	getByName("test-5", &manager)

	restock := Restock{Name: "test-4", Quantity: 10}
	err := restock.Process(&manager)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Restock: ")
	getByName("test-4", &manager)

	sell := Sell{Name: "test-5", Quantity: 10}
	err = sell.Process(&manager)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Sell: ")
	getByName("test-5", &manager)

	sell = Sell{Name: "test-5", Quantity: 10}
	err = sell.Process(&manager)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Sell: ")
	getByName("test-5", &manager)

}
