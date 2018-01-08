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

.PHONY: all install
all: myProgram

install: myProgram
	sudo rm -rf server/
	(ethosParams server && cd server/ && ethosBuilder && minimaltdBuilder)
	ethosTypeInstall myType
	ethosDirectoryInstall /user/nobody $(ETHOSROOT)/types/spec/myType/MyType all
	ethosDirectoryInstall /user/nobody/boxes $(ETHOSROOT)/types/spec/myType/MyType all
	cp myProgram $(ETHOSROOT)/programs
	ethosStringEncode /programs/myProgram > $(ETHOSROOT)/etc/init/services/myProgram

myType.go: myType.t
	$(ETN2GO) . myType main $^

myProgram: myProgram.go myType.go
	ethosGo $^ 

clean:
	rm -rf myType/ myTypeIndex/
	rm -f myType.go
	rm -f myProgram
	rm -f myProgram.goo.ethos
