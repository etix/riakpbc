package main

import (
	"fmt"
	"github.com/mrb/riakpbc"
	"log"
	"os/exec"
	"runtime"
	"time"
)

type Data struct {
	Data string `json:"data"`
}

func main() {
	runtime.GOMAXPROCS(7)
	cluster := []string{"127.0.0.1:8087", "127.0.0.1:8088", "127.0.0.1:8089", "127.0.0.1:8090"}
	riak := riakpbc.NewClient(cluster)

	err := riak.Dial()
	if err != nil {
		log.Print(err)
	}

	var actionEnd time.Time
	actionBegin := time.Now()

	c := make(chan bool)

	for g := 0; g < 7; g++ {
		go func(which int) {
			log.Print("<", which, "> Loaded")
			var times int
			var errs int
			for {
				actionBegin := time.Now()

				times = times + 1
				_, err := riak.StoreObject("bucket", "data", "{'ok':'ok'}")
				if err != nil {
					errs = errs + 1
				}

				_, err = riak.SetClientId("coolio")
				if err != nil {
					errs = errs + 1
				}

				_, err = riak.GetClientId()
				if err != nil {
					errs = errs + 1
				}

				data, err := riak.FetchObject("bucket", "data")
				if err != nil {
					break
				}
				if string(data.GetContent()[0].GetValue()) != "{'ok':'ok'}" {
					log.Fatal("!!!")
				}

				_, err = riak.StoreObject("bucket", "moreData", "stringData")
				if err != nil {
					errs = errs + 1
				}

				_, err = riak.FetchObject("bucket", "moreData")
				if err != nil {
					errs = errs + 1
				}

				actionDuration := time.Now().Sub(actionBegin)
				log.Printf("%s !<%#v> %s", riak.Pool(), errs, actionDuration)
			}
		}(g)
	}

	progress := 0
	for {
		select {
		case <-c:
			actionEnd = time.Now()
			actionDuration := actionEnd.Sub(actionBegin)
			log.Print("Ran for ", actionDuration)
			riak.Close()
		case <-time.After(5000 * time.Millisecond):
			progress = progress + 1
			var err error
			switch progress {
			case 1:
				StartFW(5, []int{8087, 8088})
			case 2:
				DeleteFW()
			}
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func StartFW(kb int, ports []int) {
	cmd := exec.Command("sudo", "ipfw", "delete 1")
	cmd.Run()
	configString := fmt.Sprintf("pipe 1 config bw \"%#v\"KByte/s", kb)
	cmd = exec.Command("sudo", "ipfw", configString)
	for _, port := range ports {
		throttleString := fmt.Sprintf("add 1 pipe 1 src-port %#v", port)
		cmd := exec.Command("sudo", "ipfw", throttleString)
		cmd.Run()
	}
}

func DeleteFW() {
	cmd := exec.Command("sudo", "ipfw", "delete 1")
	cmd.Run()
}
