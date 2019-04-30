package game

import (
	"encoding/json"
	"io/ioutil"
)

type Game struct {
	saveFile string
	Defaults struct {
		Character Object
		Room      Object
		Object    Object
		Action    Object
		Program   Object
	}
	Objects map[DBRef]Object
	State   State
}

func LoadGame(saveFile string) (*Game, error) {
	g := &Game{
		saveFile: saveFile,
	}
	data, err := ioutil.ReadFile(saveFile)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, g); err != nil {
		return nil, err
	}
	for id, obj := range g.Objects {
		obj.id = id
		obj.iid, err = parseDBRef(id)
		if err != nil {
			return nil, fmt.Errorf("%s; corrupted save file?", err)
		}
	}
	return g
}

func (g *Game) SaveGame() error {
	data, err := json.Marshal(g)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(g.saveFile, data, 0644)
	return err
}
