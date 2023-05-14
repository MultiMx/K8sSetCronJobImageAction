package kube

import (
	"bytes"
	"fmt"
)

func (a Kube) SetImage(image string) error {
	res, e := a.Request("PATCH", &Request{
		Url:  a.Conf.DeploymentUrl(),
		Body: bytes.NewBuffer([]byte(fmt.Sprintf(`[{"op": "replace", "path": "/spec/jobTemplate/spec/template/spec/containers/%d/image", "value": "%s"}]`, a.Conf.Container, image))),
	})
	if e != nil {
		return e
	}
	_ = res.Body.Close()
	return nil
}
