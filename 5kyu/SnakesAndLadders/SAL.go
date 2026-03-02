package fiveKyu

import "fmt"

// SnakesLadders - Structure that the game is played on
type SnakesLadders struct {
	player1     map[string]interface{}
	player2     map[string]interface{}
	player1Turn bool
	winner      string
	snakes      map[int]int
	ladders     map[int]int
}

// NewGame - Method that starts a new game for the SnakesLadders struct
func (sl *SnakesLadders) NewGame() {
	sl.player1 = map[string]interface{}{"name": "Player 1", "position": 0}
	sl.player2 = map[string]interface{}{"name": "Player 2", "position": 0}
	sl.player1Turn = true
	sl.winner = ""
	sl.snakes = map[int]int{
		17: 7, 54: 34, 62: 19, 64: 60,
		87: 24, 93: 73, 95: 75, 99: 78,
	}
	sl.ladders = map[int]int{
		1: 38, 4: 14, 9: 31, 20: 38,
		28: 84, 40: 59, 51: 67, 63: 81,
		71: 91, 80: 100,
	}
}

// Play - Method that makes a turn given a dice roll from die1 and die2 for the SnakesLadders struct
// Player that dice is for should automatically be determined based on rules
func (sl *SnakesLadders) Play(die1, die2 int) string {
	// Step-8: If there's already a winner, the game is over
	if sl.winner != "" {
		return "Game over!"
	}

	// Step-2: Determine whose turn it is
	var current map[string]interface{}
	if sl.player1Turn {
		current = sl.player1
	} else {
		current = sl.player2
	}

	name := current["name"].(string)
	position := current["position"].(int)

	// Step-3: Move player by the sum of both dice
	position += die1 + die2

	// Step-6: Bounce rule — if position exceeds 100, bounce back
	if position > 100 {
		overflow := position - 100
		position = 100 - overflow
	}

	// Step-5: Check for ladder (promotion) — climb up
	if dest, ok := sl.ladders[position]; ok {
		position = dest
	}

	// Step-5: Check for snake (demotion) — slide down
	if dest, ok := sl.snakes[position]; ok {
		position = dest
	}

	// Update the player's position on their map
	current["position"] = position

	// Step-7: Check if the player has won
	if position == 100 {
		sl.winner = name
		return fmt.Sprintf("%s Wins!", name)
	}

	// Step-4 & Step-9: If doubles, same player goes again; otherwise switch turns
	if die1 != die2 {
		sl.player1Turn = !sl.player1Turn
	}

	return fmt.Sprintf("%s is on square %d", name, position)
}