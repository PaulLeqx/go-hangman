package hangman

import "fmt"

func YourWelcome() {
	fmt.Println(`
         _       _    _                   _             _              _   _         _                   _          
        / /\    / /\ / /\                /\ \     _    /\ \           /\_\/\_\ _    / /\                /\ \     _  
       / / /   / / // /  \              /  \ \   /\_\ /  \ \         / / / / //\_\ / /  \              /  \ \   /\_\
      / /_/   / / // / /\ \            / /\ \ \_/ / // /\ \_\       /\ \/ \ \/ / // / /\ \            / /\ \ \_/ / /
     / /\ \__/ / // / /\ \ \          / / /\ \___/ // / /\/_/      /  \____\__/ // / /\ \ \          / / /\ \___/ / 
    / /\ \___\/ // / /  \ \ \        / / /  \/____// / / ______   / /\/________// / /  \ \ \        / / /  \/____/  
   / / /\/___/ // / /___/ /\ \      / / /    / / // / / /\_____\ / / /\/_// / // / /___/ /\ \      / / /    / / /   
  / / /   / / // / /_____/ /\ \    / / /    / / // / /  \/____ // / /    / / // / /_____/ /\ \    / / /    / / /    
 / / /   / / // /_________/\ \ \  / / /    / / // / /_____/ / // / /    / / // /_________/\ \ \  / / /    / / /     
/ / /   / / // / /_       __\ \_\/ / /    / / // / /______\/ / \/_/    / / // / /_       __\ \_\/ / /    / / /      
\/_/    \/_/ \_\___\     /____/_/\/_/     \/_/ \/___________/          \/_/ \_\___\     /____/_/\/_/     \/_/       
																												
	`)
}

func Draw(g *Game, guess string) {
	drawTurn(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurn(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
		`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
		`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
		`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
		`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
		`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good Guess!!!")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad Guess!!! '%s' isn't in the word\n", guess)
		fmt.Printf("Turns left: %d\n", g.TurnsLeft)
	case "lost":
		fmt.Printf("You lost, the word was: \n")
		drawLetters(g.Letters)
	case "won":
		fmt.Printf("YOU WON!!! The word was: \n")
		drawLetters(g.Letters)
	}
	
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}