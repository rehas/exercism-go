package bowling

import "fmt"

// Define the Game type here.
type Game struct {
	frames []*Frame
}
type Frame struct {
	scores [3]int
}

func NewGame() *Game {
	return &Game{
		frames: make([]*Frame, 0),
	}
}

func NewFrame() *Frame {
	return &Frame{
		scores: [3]int{-1, -1, -1},
	}
}

func (bg *Game) Roll(pinsHit int) error {
	lastFrame := bg.getLastFrame()

	if err := bg.validateGameContinues(lastFrame); err != nil {
		return err
	}

	if err := bg.validateRollWithFrame(pinsHit, lastFrame); err != nil {
		return err
	}

	if bg.isFirstThrow(lastFrame) {
		lastFrame.scores[0] = pinsHit
	} else if bg.isSecondThrow(lastFrame) {
		lastFrame.scores[1] = pinsHit
	} else {
		lastFrame.scores[2] = pinsHit
	}
	return nil
}

func (bg *Game) validateGameContinues(lastFrame *Frame) error {
	if len(bg.frames) < 10 {
		return nil
	}

	if len(bg.frames) == 10 && (lastFrame.scores[0] == -1 || lastFrame.scores[1] == -1) {
		return nil
	}

	if len(bg.frames) == 10 && lastFrame.scores[0] == 10 && lastFrame.scores[2] == -1 {
		return nil
	}
	if len(bg.frames) == 10 && lastFrame.scores[0]+lastFrame.scores[1] == 10 && lastFrame.scores[2] == -1 {
		return nil
	}
	return fmt.Errorf("Cannot roll after game is over")
}

func (bg *Game) isFirstThrow(lastFrame *Frame) bool {
	return lastFrame.scores[0] == -1
}

func (bg *Game) isSecondThrow(lastFrame *Frame) bool {
	return lastFrame.scores[0] != -1 && lastFrame.scores[1] == -1
}

func (bg *Game) getLastFrame() *Frame {
	if len(bg.frames) == 0 {
		emptyFrame := NewFrame()
		bg.frames = append(bg.frames, emptyFrame)
		return emptyFrame
	}

	lastFrame := bg.frames[len(bg.frames)-1]

	if len(bg.frames) == 10 {
		return lastFrame
	}

	if lastFrame.scores[0] == 10 || lastFrame.scores[1] > -1 {
		emptyFrame := NewFrame()
		bg.frames = append(bg.frames, emptyFrame)
		return emptyFrame
	}

	return lastFrame
}

func (bg *Game) validateRollWithFrame(pinsHit int, lastFrame *Frame) error {
	if pinsHit > 10 {
		return fmt.Errorf("Pin count exceeds pins on the lane")
	}
	if pinsHit < 0 {
		return fmt.Errorf("Negative roll is invalid")
	}
	// if it's first roll return
	if bg.isFirstThrow(lastFrame) {
		return nil
	}

	isTenthFrame := len(bg.frames) == 10
	if !isTenthFrame && bg.isSecondThrow(lastFrame) {
		if pinsHit+lastFrame.scores[0] > 10 {
			return fmt.Errorf("Pin count exceeds pins on the lane")
		}
	}
	// check tenth frame
	if isTenthFrame {
		// can be all strikes or two strikes and a less than strike as the current
		if lastFrame.scores[0] == 10 && lastFrame.scores[1] == 10 {
			return nil
		}
		// if first one is a strike and second one is not than, last one can only be max remaining pins
		if lastFrame.scores[0] == 10 && lastFrame.scores[1] > -1 && lastFrame.scores[1]+pinsHit > 10 {
			return fmt.Errorf("Pin count exceeds pins on the lane")
		}
	}
	return nil
}

func (g *Game) Score() (int, error) {
	if err := g.validateFrames(); err != nil {
		return 0, err
	}
	var totalScore int
	for i := 0; i < len(g.frames); i++ {
		totalScore += g.getScoreForFrame(i)
	}
	return totalScore, nil
}

func (g *Game) validateFrames() error {
	if len(g.frames) < 10 {
		return fmt.Errorf("Score cannot be taken until the end of the game")
	}
	// check last round if it's complete
	lastRound := g.frames[len(g.frames)-1]
	if isSpare(lastRound) && lastRound.scores[2] == -1 {
		return fmt.Errorf("Score cannot be taken until the end of the game")
	}
	if isStrike(lastRound) && (lastRound.scores[1] < 0 || lastRound.scores[2] < 0) {
		return fmt.Errorf("Score cannot be taken until the end of the game")
	}
	return nil
}

func isSpare(currentFrame *Frame) bool {
	return currentFrame.scores[0]+currentFrame.scores[1] == 10
}

func isStrike(currentFrame *Frame) bool {
	return currentFrame.scores[0] == 10
}

func (g *Game) getScoreForFrame(index int) int {
	// Check strike
	currentFrame := g.frames[index]

	if isStrike(currentFrame) {
		return g.getScoreForStrike(index)
	}
	if isSpare(currentFrame) {
		return g.getScoreForSpare(index)
	}
	// Check spare

	// return normal score
	return g.frames[index].scores[0] + g.frames[index].scores[1]
}

func (g *Game) getScoreForSpare(index int) int {
	if index == 9 {
		return g.getFrameScore(index)
	}
	return 10 + g.getNextHitScore(index)
}

func (g *Game) getScoreForStrike(index int) int {
	if index == 9 {
		return g.getFrameScore(index)
	}
	return 10 + g.getNextTwoHitScores(index)
}

func (g *Game) getFrameScore(index int) int {
	currentFrame := g.frames[index]
	var total int
	for _, v := range currentFrame.scores {
		if v > 0 {
			total += v
		}
	}
	return total
}

func (g *Game) getNextTwoHitScores(index int) int {
	nextFrame := g.frames[index+1]
	if !isStrike(nextFrame) || index == 8 {
		// not strike on nextFrame or nextFrame is lastFrame
		return nextFrame.scores[0] + nextFrame.scores[1]
	} else {
		return 10 + g.getNextHitScore(index+1)
	}
}

func (g *Game) getNextHitScore(index int) int {
	nextFrame := g.frames[index+1]
	return nextFrame.scores[0]
}
