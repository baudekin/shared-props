package shared_props

type PropertyMap struct {
	Props map[string]string
}

func Updater() chan PropertyMap {
	for {
		select {}
	}

}

func User(in <-chan PropertyMap) {

}
