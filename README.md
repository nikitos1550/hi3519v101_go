<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="docs/images/gopher200.png" alt="OpenHisiIpcam"></a>
</p>

<h3 align="center">OpenHisiIpCam</h3>

---

<p align="center"> Ip camera firmware
    <br> 
</p>

## 📝 Table of Contents
- [About](#about)
- [Target hardware](#target_hardware)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Deploy application to camera](#deployment)
- [Repo structure and further study](#repo_structure)
- [Tech stack](#tech_stack)

## 👓 About <a name = "about"></a>
Project target is to make open customizable scriptable embedded software for HiSilicon based ip cameras.

TODO

## 📷 Target hardware <a name="target_hardware"></a>
A few words about target hardware...

## 🏁 Getting Started <a name="getting_started"></a>
These instructions will get you a copy of the project up and running on remote facility machine for development and testing purposes. 
Development enviroiment deployment on local machines is beyond the scope of this document. 

This repo designed

*Later when project will be moved into mature state, we will split it for several repos.*

### Remote facility
Ip address of the remote facility is 213.141.129.12. 
You can ssh via 2223 port or establish vpn with facility network and ssh 192.168.10.2.

More about facility structure you can read in corresponding [readne file](./facility).

### Deploy development enviroiment
A step by step series of examples that tell you how to get a development env running.


```
$ git clone https://github.com/nikitos1550/hi3519v101_go -b testing
$ cd hi3519v101_go
$ cp Makefile.user.params.example Makefile.user.params
$ make prepare
```

## 🎈 Usage <a name="usage"></a>
Add notes about how to use the system.

## 🚀 Deploy application to camera <a name = "deployment"></a>
Add additional notes about how to deploy this on a live system.

## 📁 Repo structure and further study <a name="repo_structure"></a>
Each dir contains it`s own README.md, that expand it`s topic.

```bash
.
├── ... - git repo files
├── Makefile - main makefile, this is entry point for development enviroiment
├── Makefile.user.params.example - custom dev env settings example
├── README.md - this document
├── application - target application
├── boards - camera hardware profiles
├── buildroot-2019.08-patch - patch files for vanilla buildroot
├── burner - tool for deployment firmware to camera via u-boot
├── docs - documentation that didn`t find home in other dirs
├── facility - remote development server related files, configuraions, etc
├── hi3516av100 - will be moved to ./hisilicon
├── hi3516av200 - will be moved to ./hisilicon
├── hi3516cv100 - will be moved to ./hisilicon
├── hi3516cv200 - will be moved to ./hisilicon
├── hi3516cv300 - will be moved to ./hisilicon
├── hi3516cv500 - will be moved to ./hisilicon
├── hi3516ev200 - will be moved to ./hisilicon
├── hi3519av100 - will be moved to ./hisilicon
├── hi3559av100 - will be moved to ./hisilicon
├── hisilicon - ???
├── output - build time artifacts
├── rootfs - ???
└── scripts - useful tools for development
```

## ⛏️  Tech stack <a name="tech_stack"></a>
- **Golang** - Main programming language for application
- **Python3**, **bash/sh**, **make**  - Tools, build automation and facility
- **...**
