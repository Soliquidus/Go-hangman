package hangman

import "testing"

func TestLetterInWord(t *testing.T) {
	word := []string{"b", "o", "b", "e", "t", "t", "e"}
	guess := "b"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. Got %v", word, guess, hasLetter)
	}
}

func TestLetterNotInWord(t *testing.T) {
	word := []string{"b", "o", "b", "e", "t", "t", "e"}
	guess := "a"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does not contain letter %s. Got %v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Errorf("Error should be returned whes using an invalid word like ''")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "Bob")
	g.MakeAGuess("b")
	validState(t, "goodGuess", g.State)
}

func TestGameAlreadyGuessed(t *testing.T) {
	g, _ := New(3, "Bob")
	g.MakeAGuess("b")
	g.MakeAGuess("b")
	validState(t, "alreadyGuessed", g.State)
}

func TestGameBadGuess(t *testing.T) {
	g, _ := New(3, "Bob")
	g.MakeAGuess("a")
	validState(t, "badGuess", g.State)
}

func TestGameWon(t *testing.T) {
	g, _ := New(3, "Bob")
	g.MakeAGuess("B")
	g.MakeAGuess("o")
	g.MakeAGuess("B")
	validState(t, "won", g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := New(3, "Bob")
	g.MakeAGuess("a")
	g.MakeAGuess("p")
	g.MakeAGuess("z")
	validState(t, "lost", g.State)
}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("state should be '%v'. Got %v", expectedState, actualState)
		return false
	}
	return true
}
