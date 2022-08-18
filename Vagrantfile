# -*- mode: ruby -*-
# vi: set ft=ruby :

$docker = <<-SCRIPT

docker compose -f /vagrant_data/docker-compose.yaml up

# docker build -t docker-go-alpine /vagrant_data/echo_crud

SCRIPT

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.

class VagrantPlugins::ProviderVirtualBox::Action::Network
  def dhcp_server_matches_config?(dhcp_server, config)
    true
  end
end

Vagrant.configure("2") do |config|

  config.vm.define :dockervm do |dockervm|

    dockervm.vm.box = "ubuntu/bionic64"
    dockervm.vm.hostname = "dockervm"

    dockervm.vm.network "private_network", ip: "192.168.33.2"

    dockervm.vm.synced_folder "../data", "/vagrant_data", create: true

    dockervm.vm.provider "virtualbox" do |vb|
    dockervm.vagrant.plugins = "vagrant-docker-compose"

      vb.memory = "3096"
    end

    dockervm.vm.provision :docker_compose
    dockervm.vm.provision "docker" do |d|
      d.post_install_provision "shell", inline: $docker
    end
    
  end
end

