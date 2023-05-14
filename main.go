package main

import (
	"github.com/MultiMx/K8sSetCronJobImageAction/controllers"
	"github.com/MultiMx/K8sSetCronJobImageAction/global"
	"github.com/MultiMx/K8sSetCronJobImageAction/pkg/kube"
	log "github.com/sirupsen/logrus"
)

func main() {
	api := kube.New(&global.Config.Config)

	e := controllers.SetImage(api)
	if e != nil {
		log.Fatalln(e)
	}
}
