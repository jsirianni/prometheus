describe package('prometheus') do
    it { should be_installed }
end

[
    '/usr/bin/prometheus',
    '/usr/bin/promtool',
].each do |bin|
    describe file(bin) do
        its('mode') { should cmp '0755' }
        its('owner') { should eq 'root' }
        its('group') { should eq 'root' }
        its('type') { should cmp 'file' }
    end
end

describe file('/usr/lib/systemd/system/prometheus.service') do
    its('mode') { should cmp '0640' }
    its('owner') { should eq 'root' }
    its('group') { should eq 'root' }
    its('type') { should cmp 'file' }
end

[
    '/etc/prometheus',
    '/var/lib/prometheus',
    '/var/lib/prometheus/console_libraries',
    '/var/lib/prometheus/consoles',
    '/var/lib/prometheus/tsdb',
].each do |dir|
    describe file(dir) do
        its('mode') { should cmp '0750' }
        its('owner') { should eq 'prometheus' }
        its('group') { should eq 'prometheus' }
        its('type') { should cmp 'directory' }
    end
end

[
    '/etc/prometheus/prometheus.yml',
    '/etc/prometheus/web.yml',
    '/etc/prometheus/rules.yml',
    '/var/lib/prometheus/consoles/node-overview.html',
    '/var/lib/prometheus/consoles/node-disk.html',
    '/var/lib/prometheus/consoles/node-cpu.html',
    '/var/lib/prometheus/consoles/node.html',
    '/var/lib/prometheus/consoles/prometheus.html',
    '/var/lib/prometheus/consoles/index.html.example',
    '/var/lib/prometheus/consoles/prometheus-overview.html',
    '/var/lib/prometheus/console_libraries/menu.lib',
    '/var/lib/prometheus/console_libraries/prom.lib'
].each do |config|
    describe file(config) do
        its('mode') { should cmp '0640' }
        its('owner') { should eq 'prometheus' }
        its('group') { should eq 'prometheus' }
        its('type') { should cmp 'file' }
    end
end

describe file('/usr/share/doc/prometheus') do
    its('mode') { should cmp '0755' }
    its('owner') { should eq 'root' }
    its('group') { should eq 'root' }
    its('type') { should cmp 'directory' }
end

[
    '/usr/share/doc/prometheus/NOTICE',
    '/usr/share/doc/prometheus/LICENSE',
    '/usr/share/doc/prometheus/npm_licenses.tar.bz2'
].each do |file|
    describe file(file) do
        its('mode') { should cmp '0644' }
        its('owner') { should eq 'root' }
        its('group') { should eq 'root' }
        its('type') { should cmp 'file' }
    end
end

describe user('prometheus') do
    it { should exist }
    its('group') { should eq 'prometheus' }
    its('lastlogin') { should eq nil }
    its('shell') { should eq '/sbin/nologin' }
end

describe group('prometheus') do
    it { should exist }
end

# After starting the service, the process should be running and 
# allow these tests to pass.

describe systemd_service('prometheus') do
    it { should be_installed }
    it { should_not be_enabled }
end

# after starting Prometheus, these files and directories should
# be created and proves that the process has permission to write
# to the tsdb directory.
[
    '/var/lib/prometheus/tsdb/lock',
    '/var/lib/prometheus/tsdb/queries.active',
    '/var/lib/prometheus/tsdb/wal/00000000'
].each do |file|
    describe file(file) do
        its('mode') { should cmp '0644' }
        its('owner') { should eq 'prometheus' }
        its('group') { should eq 'prometheus' }
        its('type') { should cmp 'file' }
    end
end

# Default install listens on all interfaces because
# it is intended to be used by remote bindplane
# instances.
describe port(9090) do
    it { should be_listening }
    its('protocols') { should include 'tcp' }
    its('addresses') { should_not include '127.0.0.1' }
    its('addresses') { should include '0.0.0.0' }
    its('processes') {should include 'prometheus'}
end

[
    'http://localhost:9090/graph',
    'http://localhost:9090/metrics',
    'http://localhost:9090/status'
].each do |url|
    describe http(url) do
        its('status') { should cmp 200 }
    end
end

describe processes('prometheus') do
    its('users') { should eq ['prometheus'] }
end
