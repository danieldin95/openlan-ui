.PHONY: serve dist test

PKG = github.com/danieldin95/lightstar/libstar

LDFLAGS += -X $(PKG).Commit=$$(git rev-list -1 HEAD)
LDFLAGS += -X $(PKG).Date=$$(date +%FT%T%z)
LDFLAGS += -X $(PKG).Version=$$(cat ../VERSION)


all: serve dist


serve:
	cd backend && go build -mod=vendor -ldflags "$(LDFLAGS)" -o serve.bin main.go


dist:
	npm run build


linux/rpm: linux/rpm/serve


linux/rpm/serve:
	./packaging/auto.sh
	rpmbuild -ba packaging/openlan-serve.spec
	cp -rvf ~/rpmbuild/RPMS/x86_64/openlan-serve-*.rpm .


test:
	go test -v -mod=vendor -bench=. github.com/danieldin95/openlan-ui/backend/http
