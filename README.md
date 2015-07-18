# golang-udp-test
golangのdhcp4の挙動を検証するためのやつ
dhcp4がeth1でパケットを受け取っても、eth0からofferを返しててわからーんっとなったので。

送信側は `nc` 使うと便利

#### 受信側
```
pc$ vagrant ssh receiver
$ /vagrant/bin/linux-amd6/golang-udp-test -i enp0s8
```

#### 送信側
```
pc$ vagrant ssh sender
$ nc -u 192.168.0.12 10718
```

#### 検証方法
1. sender -> receiver へパケット送信
2. receiver -> Broadcast
3. 2の時にsenderのenp0s8以外のIFでパケットを受け取らなければおｋ
