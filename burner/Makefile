deb-dep:
	sudo apt-get install python-serial python-pip
	sudo pip install netifaces
	sudo apt-get remove brltty

tftpy-install:
	test -e tftpy || git clone https://github.com/msoulier/tftpy
	cd tftpy; sudo python ./setup.py install
	sudo rm -rf ./tftpy

#tftpy-install:
#	sudo apt-get install python-tftpy python-pip
#	sudo pip install --upgrade tftpy
