#/bin/bash

sysctl net.ipv4.ip_forward=1

#iptables -t nat -F

    #    $EXT_R_IP - внешний IP роутера
   #     $LOCAL_IP - внутренний "фэйковый" адрес машины, которую надо "выкидывать" наружу
  #      $PORT1 - Порт, на который будут заходить извне и попадать на локальную машину
 #       $PORT2 - Порт, который "выбрасывается" наружу(например, 80 - http, либо 21 - ftp)


#   На роутере говорим следующие команды(от рута)

iptables -t nat -A PREROUTING -p tcp -d 192.168.0.2 --dport 9001 -j DNAT --to-destination 192.169.0.101:2345
iptables -A FORWARD -i eth0 -d 192.169.0.101 -p tcp --dport 2345 -j ACCEPT

#iptables -t nat -A PREROUTING -p tcp --dport 9001 -j DNAT --to-destination 192.169.0.101:2345
#iptables -t nat -A POSTROUTING -p tcp -d 192.169.0.101 --dport 2345 -j SNAT --to-source 192.168.0.2
#iptables -t nat -L -n

#iptables -t nat -A PREROUTING -i enp3s0 -p tcp -m tcp --dport 9001 -j DNAT --to-destination 192.169.0.101:2345
#iptables -t nat -A POSTROUTING -o enp3s0 -j SNAT --to-source 192.168.0.2

#iptables -t nat -A PREROUTING -p tcp -d 192.168.0.2 --dport 9001 -j DNAT --to-destination 192.169.0.101:2345
#iptables -t nat -A PREROUTING -p tcp -d 192.168.0.2 --dport 8001 -j DNAT --to-destination 192.169.0.101:80

#iptables -t nat -A POSTROUTING -j MASQUERADE

iptables -n -t nat -L
