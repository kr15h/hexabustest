#GoLang hexabus test

Simple test application for testing the hexabus switches. Uses Raspberry Pi as main environment for now. Probably works on any other Linux machine. Actually should work everywhere.

##Plug data

| Name 			| Value	
|---			|---
| Antenna ID:	| IPv6: fe80::50:c4ff:fe04:02BD
| Plug 1: 		| IPv6: fe80::50:c4ff:fe04:0390

##Setup

This guide is based on []the one found in hexabus wiki](https://github.com/mysmartgrid/hexabus/wiki/Connect-PC-Directly). Added a bit more detail on solving problems that I had.

 1. Plug in the USB antenna into one of your Pi's USB ports. Wait till the green LED goes on.
 
 2. Type `lsusb` in your terminal to see if the device is attached.
 
 3. Type `sudo ifconfig` to see the device listed as a network interface. In my case it's `usb0`.
 
 4. Type `sudo ifconfig usb0 add fafa::50:c4ff:fe04:02bd/64` and add the following to `/etc/network/interfaces`:  
 ```
auto usb0
allow-hotplug usb0
iface usb0 inet6 static
        ip6addr fafa::50:c4ff:fe04:02bd
        netmask 64
 ```  
 This should give the `usb0` interface an IPv6 IP address on reboot. Replace the `fafa` prefix with anything else, use only HEX numbers.
 
 5. Type `sudo ifconfig` to see if an IPv6 IP has been applied. You should see output similar to this:  
 ```
usb0      Link encap:Ethernet  HWaddr 02:50:c4:04:02:bd
          inet6 addr: fe80::50:c4ff:fe04:2bd/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1284  Metric:1
          RX packets:10 errors:0 dropped:0 overruns:0 frame:0
          TX packets:8 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:460 (460.0 B)  TX bytes:648 (648.0 B)
 ```  
 
 6. If you do not have the `inet6 addr` line, probably it's because the `ipv6` kernel module is not enabled. To enable it, type `sudo modprobe ipv6` and add `ipv6` to `/etc/modules` so it is being enabled after the next reboot.
 
 7. Pair the USB plug with the switch. Press and hold the button on the plug until the LED starts to blink red. Then press and hold the button on the power switch for a couple of seconds - it will blink red and then green on release. The LED on the USB plug should turn green as well. If both LEDs (the one on the USB switch and on the power plug) are green - the devices have been paired. 
 
 8. Try to ping the plug:  
 ```
 ping6 -Iusb0 fe80::50:c4ff:fe04:8390
 ```  
 Use the original plug address here.
 
 9. Install `radvd` with `sudo apt-get install radvd`
 
 10. Config `radvd` by opening the file `/etc/radvd.conf`. Add the following contents there:
 ```
 interface usb0
{
  AdvSendAdvert on;
  AdvLinkMTU 1280;
  AdvCurHopLimit 128;
  AdvReachableTime 360000;
  MinRtrAdvInterval 100;
  MaxRtrAdvInterval 150;
  AdvDefaultLifetime 200;
  prefix fafa::/64
  {
    AdvOnLink on;
    AdvAutonomous on;
    AdvPreferredLifetime 4294967295;
    AdvValidLifetime 4294967295;
  };
};
 ```  
 Again - remember to change the fafa with your prefix. Run `sudo radvd` to start the daemon.
 
 11. Ping your plugs with global addresses:  
 ```
 ping6 fafa::50:c4ff:fe04:8390
 ```  
 As you see - we don't need the -Iusb0 anymore.
 
