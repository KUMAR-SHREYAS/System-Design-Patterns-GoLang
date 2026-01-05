package main

type ak47 struct {
	Gun
}

func newAK47() IGun {
	return &ak47{
		Gun{
			name:  "AK47",
			power: 4,
		},
	}
}
