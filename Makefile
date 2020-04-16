#
# github.com/danieldin95/openlan-ui
#

.PHONY: serve dist test

all: dist

# compile
dist:
	npm run build

#
linux/rpm:
	@./packaging/auto.sh
	rpmbuild -ba packaging/openlan-ui.spec
	@cp -rvf ~/rpmbuild/RPMS/x86_64/openlan-ui-*.rpm .
