# DevOps

> [维基百科](https://zh.wikipedia.org/zh-cn/DevOps)

DevOps 是一种重视 `软件开发人员` 和 `IT 运维技术人员` 之间沟通合作的文化、运动或惯例。透过自动化“软件交付”和“架构变更”的流程，来使得构建、测试、发布软件能够更加地快捷、频繁和可靠。 传统的软件组织将开发、IT 运营和质量保障设为各自分离的部门，在这种环境下如何采用新的开发方法，是一个重要的课题。

## CI/CD

CI/CD 的核心概念是持续集成、持续交付和持续部署。

### CI 持续集成（Continuous Integration）

### CD 持续交付（Continuous Delivery）

### CD 持续部署（Continuous Deployment）

## 工具文档

### Go

Go 是 Google 开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。 罗伯特·格瑞史莫，罗勃·派克及肯·汤普逊于 2007 年 9 月开始设计 Go，稍后 Ian Lance Taylor、Russ Cox 加入项目。Go 是基于 Inferno 操作系统所开发的。

### Hugo

Hugo 是一个用 Go 编写的静态网站生成器，2013 由 Steve Francia 原创，自 v0.14 由 Bjørn Erik Pedersen 主力开发，并由全球各地的开发者和用户提交贡献。Hugo 以 Apache License 2.0 授权的开放源代码项目。

[Hugo](https://hugo-docs.netlify.com)

#### 主题

[slate](https://github.com/slatedocs/slate)-ruby

Beautiful static documentation for your API https://spectrum.chat/slate

[docuapi](https://github.com/bep/docuapi)

### Docker

Docker 是一个开放平台，用于开发应用、交付应用、运行应用。 Docker 允许用户将基础设施中的应用单独分割出来，形成更小的颗粒，从而提高交付软件地速度。 Docker 容器 与虚拟机类似，但原理上，容器是将操作系统层虚拟化，虚拟机则是虚拟化硬件，因此容器更具有便携性、高效地利用服务器。

### Kubernetes

Kubernetes 是用于自动部署、扩展和管理 `容器化应用程序` 的开源系统。该系统由 Google 设计并捐赠给 Cloud Native Computing Foundation 来使用。 它旨在提供 `跨主机集群的自动部署、扩展以及运行应用程序容器的平台`。 它支持一系列容器工具, 包括 Docker 等。

[Kubernetes](https://k8s-docs.netlify.com/)

## Istio

云平台令使用它们的公司受益匪浅。但不可否认的是，上云会给 DevOps 团队带来压力。为了可移植性，开发人员必须使用微服务来构建应用，同时运维人员也正在管理着极端庞大的混合云和多云的部署环境。 Istio 允许您连接、保护、控制和观察服务。

从较高的层面来说，Istio 有助于降低这些部署的复杂性，并减轻开发团队的压力。它是一个完全开源的服务网格，作为透明的一层接入到现有的分布式应用程序里。它也是一个平台，拥有可以集成任何日志、遥测和策略系统的 API 接口。Istio 多样化的特性使您能够成功且高效地运行分布式微服务架构，并提供保护、连接和监控微服务的统一方法。

[Istio](https://istio-docs.netlify.com/)

## Knative

Knative 最初由 Google 打造，现在有 50 多家不同公司向其贡献过代码。它提供了一组必备组件，用于在 Kubernetes 上构建和运行无服务器应用。Knative 为 Kubernetes 上的云原生应用提供缩容到零、自动扩缩、集群内构建以及事件框架等功能。无论是在本地、云端还是第三方数据中心，Knative 都可以应用来源于实践中基于 Kubernetes 的成功框架的标准化最佳做法。最重要的是，Knative 使开发者能够专注于编写代码，而无需担心构建、部署和管理应用等“单调而棘手”的工作。

[Knative](https://knative-docs.netlify.com/)

### [Helm](https://helm.sh/)

Helm 是 Kubernetes 的软件包管理员。由于操作系统软件包管理器可以轻松在 OS 上安装工具，因此 Helm 可以轻松将应用程序和资源安装到 Kubernetes 群集中。

虽然 Helm 是该项目的名称，命令行客户端也被命名 helm。按照惯例，当谈到这个项目时，Helm 被大写。在谈到客户时，helm 是小写的。

[中文文档](https://whmzsu.github.io/helm-doc-zh-cn/)

### [Tekton](https://cloud.google.com/tekton)

Tekton 是一个强大且灵活的 Kubernetes 原生开源框架，可用于创建持续集成和交付 (CI/CD) 系统。该框架可让您跨多个云服务商或本地系统进行构建、测试和部署，而无需操心基础实现详情。

### Jenkins

Jenkins 是一款由 Java 编写的开源的持续集成工具。在与 Oracle 发生争执后，项目从 Hudson 项目复刻。 Jenkins 提供了软件开发的持续集成服务。它运行在 Servlet 容器中。

### Jenkins X

Jenkins X 为 Kubernetes 提供自动化 CI + CD,预览环境使用 Tekton, Knative, Prow, Skaffold 和 Helm 的引入请求

### Skaffold

Skaffold 是一款命令行工具，旨在促进 Kubernetes 应用的持续开发。你可以在本地迭代应用源码，然后将其部署到本地或者远程 Kubernetes 集群中。Skaffold 会处理构建、上传和应用部署方面的工作流。它通用可以在自动化环境中使用，例如 CI/CD 流水线，以实施同样的工作流，并作为将应用迁移到生产环境时的工具

### Prow

Prow 是 Google 发起的适应云原生开源项目的 ChatOps 系统。Kubernetes、Istio 等项目都使用 Prow 实现开源协同。

### Fluentd

Fluentd 是最初由 Treasure Data 开发的跨平台开源数据收集软件项目。它主要是用 Ruby 编程语言编写的。

### Kind

kind 是用于运行于本地 Kubernetes 集群工具，使用 Docker container “nodes”.
kind 是主要设计用于测试 Kubernetes 本身, 也可用于本地开发或 CI。

### CloudEvents 规范

[CloudEvents](https://cloudevents.io/)是一种以通用方式描述事件数据的规范。CloudEvents 旨在简化跨服务，平台及其他方面的事件声明和发送。CloudEvents 　最初由　 CNCF Severless 工作组提出。
