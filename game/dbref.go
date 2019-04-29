package game

type Flag rune

const (
	CHARACTER Flag = 'C'
	ROOM      Flag = 'R'
	OBJECT    Flag = 'O'
	ACTION    Flag = 'A'
	PROGRAM   Flag = 'P'

	// Character flags
	GUEST      = 'g' // Whether the character is a guest.
	BUILDER    = 'b' // Whether the character can build rooms.
	PROGRAMMER = 'p' // Whether the character can create programs.
	MODERATOR  = 'm' // Whether the character is a moderator.
	WIZARD     = 'w' // Whether the character is a wizard.
	FLYING     = 'f' // Whether the character can fly.
	JUMPABLE   = 'j' // Whether someone can jump to a character
	KILLABLE   = 'k' // Whether the character can be killed.

	// Object flags
	ZOMBIE = 'z' // Whether an object can be a zombie.
	STICKY = 's' // Whether an object returns home if dropped.

	// Room flags
	DARK  = 'd' // Whether a room's contents are hidden.
	HAVEN = 'h' // Whether a room allows killing.
	ABODE = 'a' // Whether a room can be a home.
)

type DBRef int
