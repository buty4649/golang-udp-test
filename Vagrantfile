# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure(2) do |config|
  config.vm.box = "base"
  config.vm.synced_folder ".", "/vagrant"

  def define_virtualbox(c, opt = {})
    ## network
    c.vm.network :private_network, ip: opt[:nic1_ip] if opt.has_key?(:nic1_ip)
    c.vm.network :private_network, ip: opt[:nic2_ip] if opt.has_key?(:nic2_ip)

    # ref http://vboxmania.net/content/vboxmanage-modifyvm%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89
    c.vm.provider :virtualbox do |vbox|
      # IPv6とDNSでのネットワーク遅延対策で追記
      vbox.customize ["modifyvm", :id, "--natdnsproxy1", "off"]
      vbox.customize ["modifyvm", :id, "--natdnshostresolver1", "off"]
    end

    ## spec
    opt[:memory] ||= 512
    opt[:cpus]   ||= 2
    c.vm.provider :virtualbox do |vbox|
      vbox.customize ["modifyvm", :id, "--memory", opt[:memory]]
      vbox.customize ["modifyvm", :id, "--cpus", opt[:cpus]]
    end
  end

  config.vm.define :receiver do |c|
    c.vm.box      = "puppetlabs/centos-7.0-64-puppet"
    c.vm.hostname = "receiver"
    define_virtualbox c, nic1_ip: "192.168.0.11", nic2_ip: "192.168.69.11"
  end

  config.vm.define :sender do |c|
    c.vm.box      = "puppetlabs/centos-7.0-64-puppet"
    c.vm.hostname = "sender"
    define_virtualbox c
    define_virtualbox c, nic1_ip: "192.168.0.12", nic2_ip: "192.168.69.12"
  end

end
