// Package resources init all resources here
package resources

import (
	"fmt"
	"os"

	"github.com/kdpujie/log4go"
)

var (
	resources = make(map[string]Closer)
)

// Closer close interface for resources
type Closer interface {
	Close() error
}

// Register resource for release
func Register(name string, r Closer) error {
	if _, ok := resources[name]; ok {
		log4go.Warn("resource[%v] update", name)
	}
	resources[name] = r
	return nil
}

// Init the resources
func Init() error {
	// kafka producer init
	_, err := initProducer()
	if err != nil {
		return fmt.Errorf("[resources-init-%v] err:%v", ResourceNameProducer, err.Error())
	}
	return nil
}

// Close destroy or release resources
func Close() {
	log4go.Debug("resources destroy, pid:%v", os.Getpid())
	for name, r := range resources {
		err := r.Close()
		if err != nil {
			log4go.Error("resources[%s] destroy failed:%s", name, err.Error())
		} else {
			log4go.Info("resources[%s] destroy finish", name)
		}
	}
}
