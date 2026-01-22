// observer interface
package main

type FitnessDataObserver interface {
	Update(data FitnessData)
}
