package main

type MovementSystem struct {
}

func (m *MovementSystem) Player(player *Player, world *World, input *Input) {
	dir := input.Player1()
	vy := dir * player.Speed
	nextY := player.Y + vy
	if nextY >= 0 && nextY <= world.Height-player.Height {
		player.Y += vy
	}
}

func (m *MovementSystem) Enemy(player *Player, world *World, input *Input) {
	dir := input.Player2()
	vy := dir * player.Speed
	nextY := player.Y + vy
	if nextY >= 0 && nextY <= world.Height-player.Height {
		player.Y += vy
	}
}

func (m *MovementSystem) Ball(ball *Ball, world *World, players []*Player) {
	ball.X += ball.DirX * ball.Speed
	ball.Y += ball.DirY * ball.Speed

	if ball.X > world.Width {
		ball.Respawn(world.Width/2, world.Height/2)
		players[0].Score += 1
	}

	if ball.X < 0 {
		ball.Respawn(world.Width/2, world.Height/2)
		players[1].Score += 1
	}

	if ball.Y < 0 || ball.Y > world.Height {
		ball.DirY = -ball.DirY
	}

	for _, player := range players {
		if checkCollision(ball, player) {
			ball.DirX = -ball.DirX
			ball.IncreaseSpeed()
		}
	}
}

func checkCollision(ball *Ball, player *Player) bool {
	return ball.X < player.X+player.Width &&
		ball.X+ball.Radius > player.X &&
		ball.Y < player.Y+player.Height &&
		ball.Y+ball.Radius > player.Y
}
