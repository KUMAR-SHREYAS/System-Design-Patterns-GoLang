// This is a concrete product
package main

import "fmt"

type Enemy struct {
	Type    string
	Health  int
	Speed   float64
	Armored bool
	Weapon  string
}

// Enemy struct implements EnemyPrototype
// the fact that an EnemyProtoype interface can be returned wrapped around a
// a concrete Enemmy pointer , tellls tha client doesnt need
// to know the concrete type of the object being cloned
// just the fact that it implements the EnemyPrototype interface

func (e *Enemy) clone() EnemyPrototype {
	return &Enemy{
		Type:    e.Type,
		Health:  e.Health,
		Speed:   e.Speed,
		Armored: e.Armored,
		Weapon:  e.Weapon,
	}
}

func (e *Enemy) printStats() {
	fmt.Printf("%s [Health: %d, Speed: %.1f, Armored: %t, Weapon: %s]\n",
		e.Type, e.Health, e.Speed, e.Armored, e.Weapon)
}
