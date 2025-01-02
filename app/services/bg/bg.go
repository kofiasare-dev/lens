package bg

import (
	"sync"

	"github.com/kofiasare-dev/background"
)

type Instance struct{ *background.Background }

var instance *Instance
var once sync.Once

func GetInstance() *Instance {
	once.Do(func() {

		bg := background.New(background.DefaultConfig())

		instance = &Instance{bg}
	})

	return instance
}
