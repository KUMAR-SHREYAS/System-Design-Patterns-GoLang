package main

type EnemyPrototype interface {
	clone() EnemyPrototype
	printStats()
}
