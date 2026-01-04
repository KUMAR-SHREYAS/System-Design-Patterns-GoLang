package main

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun{
			name:  "Musket",
			power: 1,
		},
	}
}
