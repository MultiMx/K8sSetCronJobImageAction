package models

import "github.com/MultiMx/K8sSetCronJobImageAction/pkg/kube"

type Config struct {
	kube.Config
	Image string `env:"IMAGE,required"`
}
