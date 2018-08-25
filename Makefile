getdep:
	which dep > /dev/null 2&>1 || go get -u github.com/golang/dep/cmd/dep

dep: getdep
	dep ensure

depup: getdep
	dep ensure -update

.PHONY: getdep dep depup
