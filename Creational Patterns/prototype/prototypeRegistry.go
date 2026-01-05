// the prototypeRegistry.go file will be used to store various concrete product Enemy as a key value pair
package main

type EnemyRegistry struct {
	prototypes map[string]EnemyPrototype
}

func NewEnemyRegistry() *EnemyRegistry {
	return &EnemyRegistry{
		prototypes: make(map[string]EnemyPrototype),
	}
}

//The reason EnemyPrototype is used as the value type in the map is to allow storing any concrete type that implements the EnemyPrototype interface.
// in this case Enemy struct implements EnemyPrototype interface hence we can store Enemy pointers in the map
// also can pass EnemyPrototype as arguments and return them as well.
// cause the client only needs to know about the interface not the concrete type
func (r *EnemyRegistry) Register(key string, p EnemyPrototype) {
	r.prototypes[key] = p
}

func (r *EnemyRegistry) Get(key string) EnemyPrototype {
	proto, ok := r.prototypes[key]
	if ok {
		return proto.clone() // return a clone of the prototype
	}
	return nil
}
