// This is the client code that uses the prototype registry to create enemies
package main

func main() {
	// EnemyRegistry to store all the different enemy prototypes
	registry := NewEnemyRegistry()

	// Register some enemy prototypes
	registry.Register("flying", &Enemy{"FlyingEnemy", 100, 12.0, false, "Laser"})
	registry.Register("armored", &Enemy{"ArmoredEnemy", 300, 6.0, true, "Cannon"})

	// Create clones of the registered prototypes
	enemy1 := registry.Get("flying")
	enemy2 := registry.Get("flying")

	// Modify the clone if needed using type assertion
	if enemyInstance, ok := enemy2.(*Enemy); ok {
		enemyInstance.Health = 80
	}
	enemy3 := registry.Get("armored")

	enemy1.printStats()
	enemy2.printStats()
	enemy3.printStats()

}
