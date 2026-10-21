# CLI Tranclate Training Session

cli programm to study words and practice it

## Features
- Add word and Tranclate
- Delete words pair
- Check all words
- Make a training session to practice your words knowledges
- Choose amount of words to training it 
- Practice only this amount of word, which you decided before
- Check results history of every training session

## How to run
terminal:
``` bash
go run ./cmd/app
```

## Project structure
```text 
/cmd/app - application entry point
/internal - all business logic
/internal/cli - program working scenary 
/internal/handler - program handler, part of program which is working with user
/internal/service - part of program, which is kinda skretch handler, business logic and storage
```

## Scenary

1. User go to this platform
2. User enter some word translate pairs
3. User turn on training session in spite of his free time. If it is not a lot, user may choose less words
4. When user recognized he know this word, he can delete this word with tranclate
5. User can check his training sessions history


 
