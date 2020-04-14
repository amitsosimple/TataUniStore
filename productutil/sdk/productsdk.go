/**** Amit Chatter (amitsosimple@gmail.com) ****/

package sdk

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"sync"
	"time"

	"../config"
	u "../log"
)

type Sdk struct {
	Products []config.Product
	closeChan   chan struct{}
	lock        sync.RWMutex
}

var configFile string

func (sdk *Sdk)GetProductFile() string {
	if &configFile == nil || configFile == "" {
		usr, _ := user.Current()
		configFile = filepath.Join(usr.HomeDir, "ProductionCatalog", "ProductCatalog.json")
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		os.Create(configFile)
		_ = ioutil.WriteFile(configFile, []byte("[]"), 0644)
	}

	return configFile
}

func NewProductClient() (productSDK *Sdk, sdkErr error) {

	ret := &Sdk{
		closeChan:          make(chan struct{}),
		lock:               sync.RWMutex{},
	}

	file, _ := ioutil.ReadFile(ret.GetProductFile())
	json.Unmarshal(file, &ret.Products)

	// A goroutines which monitor local configuration file changes and update the config
	go ret.localConfigChange()

	return ret, nil
}

func (sdk *Sdk)localConfigChange() {
	updateFrequency := 1 * time.Microsecond
	initialStat, err := os.Stat(sdk.GetProductFile())
	if err != nil {
		return
	}

	ticker := time.NewTicker(updateFrequency)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-sdk.closeChan:
			ticker.Stop()
			return
		default:
			stat, err := os.Stat(sdk.GetProductFile())
			if err != nil {
				continue
			}

			if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
				initialStat = stat
				sdk.lock.Lock()
				file, _ := ioutil.ReadFile(sdk.GetProductFile())
				json.Unmarshal(file, &sdk.Products)
				sdk.lock.Unlock()
				u.GeneralLogger.Println("Local file changed, Production information is reread and updated")
			}
		}
	}
}

func (sdk *Sdk) GetProducts() []config.Product{
	return sdk.Products
}

func (sdk *Sdk) UpdateCatalog(products []config.Product) error {
	u.Exit(u.Enter())
	info, _ := json.Marshal(products)
	configFile := sdk.GetProductFile()

	sdk.lock.Lock()
	error := ioutil.WriteFile(configFile, info, 0644)
	sdk.lock.Unlock()

	return error
}
