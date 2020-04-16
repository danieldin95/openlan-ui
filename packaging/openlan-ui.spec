Name: openlan-ui
Version: 0.0.3
Release: 1%{?dist}
Summary: OpenLan's UI Software
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

mkdir -p %{buildroot}/var/openlan/ui/dist
cp -R %_source_dir/dist %{buildroot}/var/openlan/ui/dist

%files
%defattr(-,root,root)
/var/openlan/*
