# K8sSetImageAction

更新 Rancher Deployment 镜像地址 Action

## 使用

```yaml
- name: Update Cronjob
  uses: MultiMx/K8sSetCronJobImageAction@v0.2
  with:
    backend: 'https://some.rancher.com'
    token: ${{ secrets.CATTLE_TOKEN }} # rancher api bearer token
    namespace: 'control'
    cronjob: 'apicenter'
    image: image.url:version
    container: 1 # optional, container index number, default 0
    cluster: some_cluster # optional, cluster name, default local
```
