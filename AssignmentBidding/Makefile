export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)

export ETHOSROOT=server/rootfs
export MINIMALTDROOT=server/minimaltdfs

userspace.myRpc.dir :=
patsubst := 


#remote procedure interface between client and server
userspace.myRpc.types =$(patsubst %, $(userspace.myRpc.dir)%, myRpc.t)

#client build definitions
userspace.myRpcClient.go = myRpcClient.go myRpc.go
userspace.myRpcClient.src = $(patsubst %, $(userspace.myRpc.dir)%, $(userspace.myRpcClient.go))
userspace.myRpcClient.target = myRpcClient

#server  build definitions
userspace.myRpcService.go = myRpcService.go myRpc.go
userspace.myRpcService.src = $(patsubst %, $(userspace.myRpc.dir)%, $(userspace.myRpcService.go)) 
userspace.myRpcService.target = myRpcService

#install these executables
INSTALL_ETHOSROOT_PROGRAMS += $(userspace.myRpcClient.target)
INSTALL_ETHOSROOT_SYSTEMPROGRAMS += $(userspace.myRpcService.target)

.PHONY: all install clean
all: myRpcClient myRpcService

myRpc.go: myRpc.t
	$(ETN2GO) . myRpc main $^

myRpcService: myRpcService.go myRpc.go
	ethosGo $^

myRpcClient: myRpcClient.go myRpc.go
	ethosGo $^

#installtypes,service,
install:
	sudo rm -rf server/
	(ethosParams server && cd server/ && ethosBuilder && minimaltdBuilder)
	ethosTypeInstall myRpc
	ethosDirectoryInstall /user/nobody $(ETHOSROOT)/types/spec/myRpc/Assignmentlist all
	ethosDirectoryInstall /user/nobody/assignment $(ETHOSROOT)/types/spec/myRpc/Assignmentlist all
	cp myRpcService $(ETHOSROOT)/programs
	cp myRpcClient 	$(ETHOSROOT)/programs
	ethosDirCreate  $(ETHOSROOT)/services/myRpc $(ETHOSROOT)/types/spec/myRpc/MyRpc all
	ethosServiceInstall myRpc MyRpc
	ethosStringEncode /programs/myRpcService > $(ETHOSROOT)/etc/init/services/myRpcService
	ethosStringEncode /programs/myRpcClient > $(ETHOSROOT)/etc/init/services/myRpcClient

#remove build artifacts
clean:
	rm -rf myRpc/ myRpcIndex/
	rm -f myRpc.go
	rm -f myRpcClient
	rm -f myRpcService
	rm -f myRpcClient.goo.ethos
	rm -f myRpcService.goo.ethos
