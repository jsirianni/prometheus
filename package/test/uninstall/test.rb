describe package('prometheus') do
    it { should_not be_installed }
end

[
    '/usr/bin/prometheus',
    '/usr/bin/promtool',
].each do |bin|
    describe file(bin) do
        it { should_not exist }
    end
end

describe file('/usr/lib/systemd/system/prometheus.service') do
    it { should_not exist }
end

[
    '/etc/prometheus',
    '/var/lib/prometheus/console_libraries',
    '/var/lib/prometheus/consoles',
].each do |dir|
    describe file(dir) do
        it { should_not exist }
    end
end

# Won't get cleaned up to do tsdb directory having
# files not managed by the package.
[
    '/var/lib/prometheus/tsdb',
].each do |dir|
    describe file(dir) do
        it { should exist }
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
        it { should_not exist }
    end
end

describe file('/usr/share/doc/prometheus') do
    it { should_not exist }
end

[
    '/usr/share/doc/prometheus/NOTICE',
    '/usr/share/doc/prometheus/LICENSE'
].each do |file|
    describe file(file) do
        it { should_not exist }
    end
end

# We do not cleanup the user, so it should still exist.
describe user('prometheus') do
    it { should exist }
    its('group') { should eq 'prometheus' }
    its('lastlogin') { should eq nil }
    its('shell') { should eq '/sbin/nologin' }
end

describe group('prometheus') do
    it { should exist }
end

describe systemd_service('prometheus') do
    it { should_not be_installed }
end

describe port(9090) do
    it { should_not be_listening }
end

describe processes('prometheus') do
    it { should_not exist }
end
