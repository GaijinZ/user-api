# -*- mode: ruby -*-
# vi: set ft=ruby :

$docker = <<-SCRIPT

docker compose -f /vagrant_data/docker-compose.yaml up

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

  config.vm.network "private_network", ip: "192.168.33.2"

  # config.vm.synced_folder ".", "/vagrant_data", create: true

#   config.vm.define "psqlVM" do |psqlVM|

#     psqlVM.vm.box = "ubuntu/trusty64"
#     psqlVM.vm.hostname = "psqlVM"

#     psqlVM.vm.network "private_network", ip: "192.168.33.30"

#     psqlVM.vm.provision "shell", inline: <<-SHELL
#     apt-get update

#     sudo apt-get install postgresql postgresql-contrib -y
#     echo "<------ CREATING USERAPI DATABASE AND USER ------>"
    
#     sudo -u postgres psql -c "CREATE USER userapi WITH PASSWORD 'userapi' CREATEDB LOGIN;"
#     sudo -u postgres psql -c "CREATE database userapi WITH OWNER = userapi;"
#     sudo -u postgres psql < /vagrant_data/init-scripts/postgres_init.sql

#     sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE userapi TO userapi;"
#     sudo -u postgres psql -d userapi -c "GRANT ALL PRIVILEGES ON ALL tables IN SCHEMA public TO userapi" 
#     sudo -u postgres psql -d userapi  -c "GRANT ALL PRIVILEGES ON ALL sequences IN SCHEMA public TO userapi" 
    
#     echo "<------ POSTGRES CONFIGURATION ------>"
#     sudo sed -i "s/#listen_address.*/listen_addresses '*'/" /etc/postgresql/9.3/main/postgresql.conf
#     sudo cp /vagrant_data/pg_hba.conf /etc/postgresql/9.3/main/ 
#     sudo service postgresql restart
#   SHELL

# end

# config.vm.define "cassandraVM" do |cassandraVM|

    
#   cassandraVM.vm.provider "virtualbox" do |vb|
    
#     # Customize the amount of memory on the VM:
#     vb.memory = "2048"
#   end
  
#   cassandraVM.vm.hostname = "cassandraVM"
#   cassandraVM.vm.box = "ubuntu/focal64"

#   cassandraVM.vm.network "private_network", ip: "192.168.33.40"

#   cassandraVM.vm.provision "shell", inline: <<-SHELL
#     apt-get update

#     curl -fsSL https://get.docker.com -o get-docker.sh
#     DRY_RUN=1 sudo sh ./get-docker.sh
#     sudo docker pull cassandra
#     sudo docker network create cassandra
#     sudo mkdir -p /cassandra/data
#     sudo docker run --rm -d --name cassandra --hostname cassandra -v /cassandra/data:/var/lib/cassandra -p 9042:9042 -p 9160:9160 --network cassandra cassandra
#     sudo docker run --rm --network cassandra -v "/vagrant_data/init-scripts/cassandra_init.cql:/cassandra_init.cqll" -e CQLSH_HOST=cassandra -e CQLSH_PORT=9042 -e CQLVERSION=3.4.5 nuvo/docker-cqlsh
#     SHELL

#   cassandraVM.vm.provision "shell", run: "always", inline: <<-SHELL
#     sudo docker run --rm -d --name cassandra --hostname cassandra -v /cassandra/data:/var/lib/cassandra -p 9042:9042 -p 9160:9160 --network cassandra cassandra
#     SHELL
  
# end

# config.vm.define "dockerVM" do |dockerVM|

    
#   dockerVM.vm.provider "virtualbox" do |vb|
    
#     # Customize the amount of memory on the VM:
#     vb.memory = "4096"
#     vb.cpus = 4
#   end
  
#   dockerVM.vm.hostname = "dockerVM"
#   dockerVM.vm.box = "ubuntu/focal64"

#   dockerVM.vm.synced_folder "../", "/vagrant_data"
#   dockerVM.vm.synced_folder "../../consumer", "/consumer_data"
#   dockerVM.vm.network "private_network", ip: "192.168.33.50"

#   dockerVM.vm.provision "shell", inline: <<-SHELL
#     apt-get update
#     echo "<--- DOCKER INSTALL --->"
#     curl -fsSL https://get.docker.com -o get-docker.sh
#     DRY_RUN=1 sudo sh ./get-docker.sh
#     sudo usermod -aG docker vagrant && newgrp docker
#     sudo docker network create usernet
    
#     echo "<--- CREATE DIRECTORIES --->"
#     sudo mkdir -p /cassandra/data
#     sudo mkdir -p /postgres/data
#     sudo mkdir -p /zookeeper
#     sudo mkdir -p /kafka/broker-1
#     sudo mkdir -p /kafka/broker-2
#     sudo mkdir -p /kafka/broker-3
#     sudo chmod 777 -R /zookeeper/
#     sudo chmod 777 -R /kafka/
    
#     echo "<--- IMPORT SAMPLE DATABASE TO CASSANDRA --->"
#     sudo docker network create cassandra
#     sudo docker pull cassandra
#     sudo docker run --rm -d --name cassandra --hostname cassandra -v /cassandra/data:/var/lib/cassandra -p 9042:9042 -p 9160:9160 -e CASSANDRA_CLUSTER_NAME=userapi --network cassandra cassandra
#     sleep 60
#     sudo docker run --rm --network cassandra -v "/vagrant_data/init-scripts/cassandra_init.cql:/scripts/data.cql" -e CQLSH_HOST=cassandra -e CQLSH_PORT=9042 -e CQLVERSION=3.4.5 nuvo/docker-cqlsh
#     sudo docker kill cassandra
#     sudo docker image rm nuvo/docker-cqlsh
#     sudo docker network rm cassandra
#     SHELL
    
#     dockerVM.vm.provision "shell", run: "always", inline: <<-SHELL
#     echo "<--- RUNNING SERVICE --->"
#     sudo docker compose -f /vagrant_data/docker-compose.yml build
#     sudo docker compose -f /vagrant_data/docker-compose.yml up -d
#     SHELL
  
# end

config.vm.define "minikube" do |minikube|

    
  minikube.vm.provider "virtualbox" do |vb|
    # Customize the amount of memory on the VM:
    vb.memory = "8096"
    vb.cpus = 4
  end
  
  minikube.vm.hostname = "minikube"
  minikube.vm.box = "ubuntu/focal64"

  minikube.vm.synced_folder ".", "/vagrant_data", create: true
  minikube.vm.synced_folder "../../consumer", "/consumer_data", create: true
  minikube.vm.network "private_network", ip: "192.168.33.60"
  
  minikube.vm.provision "shell", inline: <<-SHELL
    apt-get update
    sudo mkdir -p /bitnami/cassandra
    sudo mkdir -p /bitnami/redis
    sudo mkdir -p /postgres/data
    sudo chmod 777 -R /kafka/
    sudo chmod 777 -R /bitnami/
    echo "<--- DOCKER INSTALL --->"
    curl -fsSL https://get.docker.com -o get-docker.sh
    DRY_RUN=1 sudo sh ./get-docker.sh
    sudo usermod -aG docker vagrant && newgrp docker
    curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
    sudo install minikube-linux-amd64 /usr/local/bin/minikube
    minikube docker-env
    eval $(minikube -p minikube docker-env)

    curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
    chmod 700 get_helm.sh
    ./get_helm.sh
    helm repo add bitnami https://charts.bitnami.com/bitnami
    helm install -f /vagrant_data/k8s/cassandra-values.yaml cassandra bitnami/cassandra --version 9.4.2
    helm install -f /vagrant_data/k8s/redis-values.yaml redis bitnami/redis
    helm install kafka -f /vagrant_data/k8s/kafka-values.yaml bitnami/kafka
    minikube kubectl -- create -f  /vagrant_data/k8s/postgres.yaml
    minikube kubectl -- create -f  /vagrant_data/k8s/userapi.yaml
    minikube kubectl -- create -f  /vagrant_data/k8s/consumer.yaml
    SHELL

    minikube.vm.provision "shell",  inline: <<-SHELL
    minikube start --driver=none --memory 6144 --cpus 2
    alias kubectl="minikube kubectl --"
    eval $(minikube -p minikube docker-env)
    docker build --tag userapi:latest /vagrant_data
    ./vagrant_data/k8s/restart.sh
    docker build --tag consumer:latest /consumer_data
    # minikube kubectl -- port-forward service/userapi 8200:8200 8201:8201 --address 0.0.0.0
    # minikube start
    SHELL
end

config.vm.define "minikube2" do |minikube|

    
  minikube.vm.provider "virtualbox" do |vb|
    
    # Customize the amount of memory on the VM:
    vb.memory = "4096"
    vb.cpus = 2
  end
  
  minikube.vm.hostname = "minikube"
  minikube.vm.box = "ilionx/ubuntu2004-minikube"

  minikube.vm.synced_folder "../", "/vagrant_data"
  minikube.vm.synced_folder "../../consumer", "/consumer_data"
  minikube.vm.network "private_network", ip: "192.168.33.70"
  minikube.vm.network "forwarded_port", guest: 7080, host: 8111

  end
end