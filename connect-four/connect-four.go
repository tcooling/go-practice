package main

import (
	"fmt"
	"math/rand"
)

// ANSI escape codes for red, yellow and reset
const reset = "\033[0m"

var colourCodes = map[Colour]string{
	Yellow: "\033[33m",
	Red:    "\033[31m",
}

const boardWidth int = 7
const boardHeight int = 6

type Row [boardWidth]CellState
type Board [boardHeight]Row

type GameState struct {
	Fsm            FSM
	PlayerColour   Colour
	ComputerColour Colour
	CurrentTurn    Turn
	Board          Board
}

type FSM int

const (
	Start FSM = iota
	PrintGameboard
	CheckForWinOrLoss
	InputGuess
	Win
	Draw
	Lose
	Exit
)

type Colour int

const (
	Yellow Colour = iota
	Red
)

type Coordinate struct {
	Col int
	Row int
}

// TODO: improve model, cell fill should be same as colour
type CellState struct {
	CellFill   CellFill
	Coordinate Coordinate
}

type CellFill int

const (
	EmptyCell CellFill = iota
	YellowCell
	RedCell
)

type Turn int

const (
	Player Turn = iota
	Computer
)

func (t Turn) String() string {
	return [...]string{"Player", "Computer"}[t]
}

// Really annoys me having the default? Can the compiler not infer?
func (c CellFill) String() string {
	switch c {
	case EmptyCell:
		return "E"
	case RedCell:
		return colourString(Red, "R")
	case YellowCell:
		return colourString(Yellow, "Y")
	default:
		return "E"
	}
}

func (t Turn) Next() Turn {
	if t == Player {
		return Computer
	}
	return Player
}

func (c Colour) CellFill() CellFill {
	if c == Yellow {
		return YellowCell
	}
	return RedCell
}

func main() {
	goingFirst := randomTurn()

	// Init the board with pre set coordinates
	var board Board
	for rowIdx, row := range board {
		for colIdx := range row {
			board[rowIdx][colIdx] = CellState{
				CellFill: EmptyCell,
				Coordinate: Coordinate{
					Col: colIdx,
					Row: rowIdx,
				},
			}
		}
	}

	state := GameState{
		Fsm:            Start,
		PlayerColour:   getColour(goingFirst, Player),
		ComputerColour: getColour(goingFirst, Computer),
		CurrentTurn:    goingFirst,
		Board:          board,
	}

	for state.Fsm != Exit {
		state = fsm(state)
	}
}

func fsm(state GameState) GameState {
	switch state.Fsm {
	case Start:
		return start(state)
	case PrintGameboard:
		return printGameboard(state)
	case CheckForWinOrLoss:
		return checkForWinOrLoss(state)
	case InputGuess:
		return inputGuess(state)
	case Win:
		return win(state)
	case Draw:
		return draw(state)
	case Lose:
		return lose(state)
	default:
		return exit(state)
	}
}

func colourString(colour Colour, str string) string {
	return colourCodes[colour] + str + reset
}

func start(gameState GameState) GameState {
	fmt.Println("Starting a game of Connect 4...")
	fmt.Printf("%s is going first\n", colourString(Yellow, gameState.CurrentTurn.String()))

	gameState.Fsm = PrintGameboard
	return gameState
}

func printGameboard(gameState GameState) GameState {
	printBoard(gameState.Board)
	gameState.Fsm = CheckForWinOrLoss
	return gameState
}

func printBoard(board Board) {
	// Print the x index
	for index := range board[0] {
		fmt.Printf(" %d ", index+1)
	}
	fmt.Println()

	// Print the board
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("[%s]", cell.CellFill.String())
		}
		fmt.Println()
	}
}

func (board Board) IsFull() bool {
	for _, row := range board {
		if !row.IsFull() {
			return false
		}
	}

	return true
}

func (row Row) IsFull() bool {
	for _, cell := range row {
		if cell.CellFill == EmptyCell {
			return false
		}
	}
	return true
}

// TODO: print strike through of line when won?
func checkForWinOrLoss(gameState GameState) GameState {
	if gameState.Board.IsFull() {
		gameState.Fsm = Draw
		return gameState
	}

	var colour Colour
	if gameState.CurrentTurn == Computer {
		colour = gameState.ComputerColour
	} else {
		colour = gameState.PlayerColour
	}
	cellFillColour := colour.CellFill()

	fmt.Println("CHECK_FOR_WIN_OR_LOSS for player 🎉", gameState.CurrentTurn.String())

	var coords []Coordinate
	for _, row := range gameState.Board {
		for _, cell := range row {
			if cell.CellFill == cellFillColour {
				coords = append(coords, cell.Coordinate)
			}
		}
	}
	fmt.Println("Coords", coords)

	gameState.Fsm = InputGuess
	gameState.CurrentTurn = gameState.CurrentTurn.Next()
	return gameState
}

// BUG: what if players have filled in the whole board?
func inputGuess(gameState GameState) GameState {
	colour := getCurrentTurnColour(gameState)
	fmt.Printf("%s place disk: ", colourString(colour, gameState.CurrentTurn.String()))
	var discXIndex int
	_, err := fmt.Scanf("%d", &discXIndex)
	if err != nil {
		fmt.Println("Error reading input:", err)
		gameState.Fsm = PrintGameboard
		return gameState
	}

	validIndex := validDiscXIndex(discXIndex)
	validColumn := columnNotFull(discXIndex, gameState.Board[0])

	if validIndex && validColumn {
		colour := getCurrentTurnColour(gameState)
		gameState.DropDisc(discXIndex-1, colour)
		gameState.Fsm = PrintGameboard
		return gameState
	} else {
		var reason string
		if !validIndex {
			reason = "please choose a column between 1 and 7"
		} else {
			reason = "please choose a column which is not full"
		}
		fmt.Printf("Invalid disk position [%d], %s\n", discXIndex, reason)
		gameState.Fsm = PrintGameboard
		return gameState
	}
}

func validDiscXIndex(discXIndex int) bool {
	return discXIndex > 0 && discXIndex < (boardWidth+1)
}

func columnNotFull(discXIndex int, topRow [7]CellState) bool {
	if discXIndex < 1 || discXIndex > boardWidth {
		return false
	}
	return topRow[discXIndex-1].CellFill == EmptyCell
}

func win(gameState GameState) GameState {
	fmt.Println("WIN")
	gameState.Fsm = Exit
	return gameState
}

func draw(gameState GameState) GameState {
	fmt.Println("The board is full and nobody has won, the result is a draw.")
	gameState.Fsm = Exit
	return gameState
}

func lose(gameState GameState) GameState {
	fmt.Println("LOSE")
	gameState.Fsm = Exit
	return gameState
}

func exit(gameState GameState) GameState {
	fmt.Println("Exiting...")
	gameState.Fsm = Exit
	return gameState
}

func randomTurn() Turn {
	if rand.Intn(2) == 0 {
		return Player
	}
	return Computer
}

// The player going first is always yellow
func getColour(goingFirst Turn, turn Turn) Colour {
	if goingFirst == turn {
		return Yellow
	}
	return Red
}

func getCurrentTurnColour(gameState GameState) Colour {
	if gameState.CurrentTurn == Player {
		return gameState.PlayerColour
	}
	return gameState.ComputerColour
}

// Already validated the inputs
func (gameState *GameState) DropDisc(colunn int, colour Colour) {
	for row := len(gameState.Board) - 1; row >= 0; row-- {
		if gameState.Board[row][colunn].CellFill == EmptyCell {
			gameState.Board[row][colunn].CellFill = colour.CellFill()
			break
		}
	}
}
