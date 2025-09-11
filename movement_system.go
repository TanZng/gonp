package main

type MovementSystem struct {
}

func (m *MovementSystem) Player(player *Player, world *World, input *Input) {
	dir := input.Direction()
	vy := dir * player.Speed
	nextY := player.Y + vy
	if nextY >= 0 && nextY <= world.Height-player.Height {
		player.Y += vy
	}
}

func (m *MovementSystem) Ball(ball *Ball, world *World, players []*Player) {
	ball.X += ball.DirX * ball.Speed
	ball.Y += ball.DirY * ball.Speed

	if ball.X < 0 || ball.X > world.Width {
		ball.Respawn(world.Width/2, world.Height/2)
	}

	if ball.Y < 0 || ball.Y > world.Height {
		ball.DirY = -ball.DirY
		ball.Update()
	}

	for _, player := range players {
		if checkCollision(ball, player) {
			ball.DirX = -ball.DirX
			ball.Update()
		}
	}
}

func checkCollision(ball *Ball, player *Player) bool {
	return ball.X < player.X+player.Width &&
		ball.X+ball.Radius > player.X &&
		ball.Y < player.Y+player.Height &&
		ball.Y+ball.Radius > player.Y
}
