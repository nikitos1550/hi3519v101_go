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

More about facility structure you can read in corresponding [readme file](./facility).

### Deploy development enviroiment
After you logged into ssh, suppose 

```console
foo@build-hisi:~$ git clone https://github.com/nikitos1550/hi3519v101_go -b testing
foo@build-hisi:~$ cd hi3519v101_go
foo@build-hisi:~$ cp Makefile.user.params.example Makefile.user.params
foo@build-hisi:~$ make prepare
```

Take a note, that if you have several copies or repo (for example you are working on several branches simultaneously),
you should create `Makefile.user.params` and `make prepare` in each repo instance.

## 🎈 Usage <a name="usage"></a>
Development ifrastructure are built with makefiles, bash scripts and several python3 utils. 
All these things are tied with facility server.
Your entry point to development enviroiment is **repo's root makefile**.
Let's see what command are exposed to us.

```console
foo@build-hisi:~/hi3519v101_go$ make
Help:
  - make prepare            - prepare; MUST be done before anything
  - make deploy-app         - build&deploy application onto particular board
  - make deploy-app-control - build&deploy application, then attach serial console onto particular board
  - make control            - attach serial console onto particular board
  - make rootfs.squashfs    - build application and pack it within RootFS image
  - make kernel             - build board kernel
  - make cleanall           - remove all artifacts
```

Software can't be run in an abstract, it can run only on some specified hardware.
There are number of supported hadwares (more you can read in corresponding [readme file](./boards)).
Exact hardware profile you are working with is setuped in you `Makefile.user.params`
Each hardware profile will require it's own kernel and rootfs build.

```make
...
CAMERA  ?= 1
BOARD   ?= jvt_hi3519v101_imx274
...
```

Build system is hierarchical, all makefile targets will invoke all corresponding underlayer targets (exception is prepare target).
If you will run `make deploy-app-control` for first time, build system will first build toolchain, then kernel and rootfs, then app, ...
You advised to `make rootfs.squashfs` and `make kernel` for target hardware profile in advance.

## 🚀 Deploy application to camera <a name = "deployment"></a>
To deploy and test application on a real hardware you should make following steps:
1. Lock some camera from a pool.
2. Tune your `Makefile.user.params`: choose application target and board profile.
3. Run deploy application target in main make file.

**TODO** *Lock some camera for your individual use.*

There are two commands that you can use to deploy software to camera:
* `make deploy-app`
* `make deploy-app-control`

Command are similar but second one will automaticly open camera's main console and let you control camera.
You are advised always use `make deploy-app-control` and monitor what is going on.

Deploy will always build application from scratch. 
More about application build process you can read in corresponding [readme file](./application).

There are several targets, be default tester application will be built.
If you want to build main version of application uncomment following line to your `Makefile.user.params`
```make
APP_TARGET := daemon
```

## 📁 Repo structure and further study <a name="repo_structure"></a>
Each dir contains it's own README.md, that expand the topic.

```
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
- **Golang**, **C** - programming languages for application
- **Python3**, **bash/sh**, **make**  - Tools, build automation and facility
- **...**
