Name: openlan-serve
Version: 0.0.3
Release: 1%{?dist}
Summary: OpenLan's Server Software
Group: Applications/Communications
License: Apache 2.0
URL: https://github.com/danieldin95/openlan-ui

BuildRequires: go

%define _source_dir ${RPM_SOURCE_DIR}/openlan-ui-%{version}

%description
OpenLan's Project Software

%build
cd %_source_dir && make all

%install
mkdir -p %{buildroot}/usr/bin
cp %_source_dir/backend/serve %{buildroot}/usr/bin/openlan-serve

mkdir -p %{buildroot}/etc/sysconfig
cp %_source_dir/packaging/serve.cfg %{buildroot}/etc/sysconfig/openlan-serve.cfg
mkdir -p %{buildroot}/etc/openlan-serve

mkdir -p %{buildroot}/usr/lib/systemd/system
cp %_source_dir/packaging/openlan-serve.service %{buildroot}/usr/lib/systemd/system

mkdir -p %{buildroot}/var/openlan-serve/dist
cp -R %_source_dir/dist %{buildroot}/var/openlan-serve
cp -R %_source_dir/packaging/ca %{buildroot}/var/openlan-serve

%pre
firewall-cmd --permanent --zone=public --add-port=10088/tcp || {
  echo "You need allowed TCP port 10088 manually."
}
firewall-cmd --reload || :

%files
%defattr(-,root,root)
/etc/*
/usr/bin/*
/usr/lib/systemd/system/*
/var/openlan-serve
