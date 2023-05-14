package controllers

import (
	"errors"
	"github.com/MultiMx/K8sSetCronJobImageAction/global"
	"github.com/MultiMx/K8sSetCronJobImageAction/pkg/kube"
	log "github.com/sirupsen/logrus"
	"time"
)

func SetImage(api *kube.Kube) error {
	var counter uint8
	for {
		e := api.SetImage(global.Config.Image)
		if e == nil {
			log.Infoln("success")
			return nil
		}
		counter++
		log.Errorln("request set image api failed: ", e)
		if counter >= 5 {
			return errors.New("failed: maximum number of attempts reached")
		}
		time.Sleep(time.Second)
	}
}
