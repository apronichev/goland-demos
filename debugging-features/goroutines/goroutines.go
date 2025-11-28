package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Restaurant simulation with multiple goroutines

type Order struct {
	ID       int
	Customer string
	Dish     string
}

type PreparedOrder struct {
	Order Order
	Cook  string
}

func main() {
	fmt.Println("🍽️  Restaurant Simulation Started")
	fmt.Println("================================\n")

	// Channels for communication between goroutines
	orders := make(chan Order, 10)
	cooking := make(chan PreparedOrder, 10)
	ready := make(chan PreparedOrder, 10)
	done := make(chan bool)

	var wg sync.WaitGroup

	// Goroutine 1: Host - Takes orders from customers
	wg.Add(1)
	go host(orders, &wg)

	// Goroutines 2-4: Three Cooks - Prepare the food
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go cook(i, orders, cooking, &wg)
	}

	// Goroutine 5: Quality Control - Checks the food
	wg.Add(1)
	go qualityControl(cooking, ready, &wg)

	// Goroutine 6: Waiter - Delivers the food
	wg.Add(1)
	go waiter(ready, &wg)

	// Goroutine 7: Manager - Monitors restaurant activity
	wg.Add(1)
	go manager(done, &wg)

	// Wait for manager to signal completion
	<-done

	// Close channels to signal goroutines to finish
	close(orders)

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("\n================================")
	fmt.Println("🏁 Restaurant Closed for the Day")
}

// Goroutine 1: Host takes orders from customers
func host(orders chan<- Order, wg *sync.WaitGroup) {
	defer wg.Done()

	customers := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	dishes := []string{"Pizza", "Pasta", "Burger", "Salad", "Steak", "Soup"}

	for i := 1; i <= 10; i++ {
		customer := customers[rand.Intn(len(customers))]
		dish := dishes[rand.Intn(len(dishes))]

		order := Order{
			ID:       i,
			Customer: customer,
			Dish:     dish,
		}

		fmt.Printf("📋 Host: Order #%d taken - %s ordered %s\n", order.ID, order.Customer, order.Dish)
		orders <- order

		time.Sleep(time.Millisecond * 300) // Simulate time between orders
	}

	fmt.Println("📋 Host: No more orders for today")
}

// Goroutines 2-4: Cooks prepare the food
func cook(id int, orders <-chan Order, cooking chan<- PreparedOrder, wg *sync.WaitGroup) {
	defer wg.Done()
	cookName := fmt.Sprintf("Cook-%d", id)

	for order := range orders {
		fmt.Printf("👨‍🍳 %s: Preparing %s for %s (Order #%d)\n",
			cookName, order.Dish, order.Customer, order.ID)

		// Simulate cooking time
		cookingTime := time.Millisecond * time.Duration(500+rand.Intn(500))
		time.Sleep(cookingTime)

		prepared := PreparedOrder{
			Order: order,
			Cook:  cookName,
		}

		cooking <- prepared
		fmt.Printf("✅ %s: Finished %s (Order #%d)\n", cookName, order.Dish, order.ID)
	}

	fmt.Printf("👨‍🍳 %s: Shift ended\n", cookName)
}

// Goroutine 5: Quality Control checks the food
func qualityControl(cooking <-chan PreparedOrder, ready chan<- PreparedOrder, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ready)

	count := 0
	for prepared := range cooking {
		count++
		fmt.Printf("🔍 QC: Inspecting Order #%d (%s)... ",
			prepared.Order.ID, prepared.Order.Dish)

		time.Sleep(time.Millisecond * 200) // Simulate inspection

		// Randomly approve/reject (90% approval rate)
		if rand.Float32() < 0.9 {
			fmt.Println("APPROVED ✓")
			ready <- prepared
		} else {
			fmt.Println("REJECTED ✗ (sent back)")
			// In a real scenario, would send back to kitchen
		}
	}

	fmt.Printf("🔍 QC: Inspected %d orders, shift ended\n", count)
}

// Goroutine 6: Waiter delivers the food
func waiter(ready <-chan PreparedOrder, wg *sync.WaitGroup) {
	defer wg.Done()

	delivered := 0
	for prepared := range ready {
		fmt.Printf("🚶 Waiter: Delivering %s to %s (Order #%d)\n",
			prepared.Order.Dish, prepared.Order.Customer, prepared.Order.ID)

		time.Sleep(time.Millisecond * 200) // Simulate delivery

		fmt.Printf("🎉 Customer %s: Received %s - Bon Appétit!\n",
			prepared.Order.Customer, prepared.Order.Dish)
		delivered++
	}

	fmt.Printf("🚶 Waiter: Delivered %d orders, shift ended\n", delivered)
}

// Goroutine 7: Manager monitors the restaurant
func manager(done chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("👔 Manager: Monitoring restaurant operations...")

	// Monitor for a certain duration
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	duration := time.Second * 5
	timeout := time.After(duration)

	ticks := 0
	for {
		select {
		case <-ticker.C:
			ticks++
			fmt.Printf("👔 Manager: Status check #%d - All systems running\n", ticks)
		case <-timeout:
			fmt.Println("👔 Manager: Closing time! Signaling shutdown...")
			done <- true
			return
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
