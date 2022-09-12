# Faber

## Faber 环境配置

- Docker、Docker-Compose
- Golang
- Fabric
- Other: Wget, Git, GCC, Make
- Environment

## Faber 目录结构

```

- FaberRoot ("Defined in const.")
  - faber
    - asserts
      - fabric
      - fabric-ca
      - fabric-samples
      - go.tar.gz
    - bin
      - ...fabric binary files
    - config
    - data
      - configtx
        - configtx.yaml
      - docker
        - ...docker-compose.yaml 🆗 controller/docker/ca.go, orderer.go, peer.go, tools.go
      - genesis-block
      - organizations
        - cryptogen
          - cryptogen-config.yaml (for each peer organization) 🆗 controller/config/crypto-config.go
          - cryptogen-config.yaml (for each orderer organization) 🆗 controller/config/crypto-config.go
        - fabric-ca
        - ordererOrganizations
          - ordererOrg
            - msp
              - config.yaml
        - peerOrganizations
          - commonOrg
            - msp
              - config.yaml
    - go
      - bin
        - ...go binary files
    - log (TODO)

```


## Network.sh创建网络的步骤

- 检查镜像、软件版本和环境变量
- 检查预设配置文件 ../config
- 创建组织信息和创世块
- 


