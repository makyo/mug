package game

type State struct {
	CharacterLocations map[DBRef]DBRef
	ObjectLocations    map[DBRef]DBRef
	RoomContents       map[DBRef][]DBRef
	CharacterContents  map[DBRef][]DBRef
	//connectedCharacters map[DBRef]
	//compiledPrograms map[DBRef]string
}

func (s *State) Move(what, from, to DBRef) error {
}

func (s *State) Get(what, who DBRef) error {
}

func (s *State) Drop(what, who DBRef) error {
}

func (s *State) Connect(who DBRef) error {
}

func (s *State) Disconnect(who DBRef) error {
}

func (s *State) Notify(who DBRef) error {
}

func (s *State) ONotify(who, DBRef) error {
}
