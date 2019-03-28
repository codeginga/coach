package coach

// Player wraps play func, player is responsible to run play
type Player interface {
	Play() error
}

// PlayFunc is for supporting func to satisfy Player
type PlayFunc func() error

// Play satisfies Player
func (p PlayFunc) Play() error {
	return p()
}

// ErrHandler handles error
type ErrHandler func(err error) error

// Captain binds Player
func Captain(players ...Player) Player {
	player := func() error {
		for _, player := range players {
			if err := player.Play(); err != nil {
				return err
			}
		}

		return nil
	}

	return PlayFunc(player)
}

// CaptainErrHandle binds Player with error check
func CaptainErrHandle(errHandler ErrHandler, players ...Player) Player {
	player := func() error {
		for _, player := range players {
			err := player.Play()
			if err == nil {
				continue
			}

			if err = errHandler(err); err != nil {
				return err
			}
		}

		return nil
	}

	return PlayFunc(player)
}
