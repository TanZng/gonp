package main

type MovementSystem struct {
}

func (m *MovementSystem) Update(player *Player, world *World, input *Input) {
	dir := input.Direction()
	vy := dir * player.Speed
	nextY := player.Y + vy
	if nextY >= 0 && nextY <= world.Height-player.Height {
		player.Y += vy
	}
}
