# EasyBlocker

+ MADE BY LOOPH!

## Whats is this project?
This project si for Blocking proxies from every botter you send to the specifed random port its listening or minecraft server using sonar antibot with ab blacklist save or firewall.yml set to true


🛠️ | Requirements:

(( Installing the needed things fro this program ))

* For Random port listening for tcp connectio that  will type every connection from proxy to connections.txt:
  - ```apt install go-golang```
  - download the kokot.go
  - ```go run kokot.go``` or ```go build kokot.go```, for go build kokot.go aftert hat you need to type ```./kokot``` to start it.
   
* For Minecraft serevr with sonar antibot:
  - Waterfall.jar from [here](https://papermc.io/software/waterfall)
  - start the WaterFall.jar and configure config.yml.
  - now add sonar.jar to plugins with luckperms-bungee.
  - after you added sonar and luckperms restart the server.
  - when you join the bungeecord server type this command into console: ```lp user <Your Minecraft NickName> permission set *```, this command will grand you full permissions for the serevr and sonar plugin.
  - after that get your public ipv4 ip with port 25577 if you didnt change it in the config.yml and start your attack there after the attack type ab blacklist save, you will get message into chat what name it has.
  - after that go to your plugins folder and find sonar folder and after you get in there will be your blacklist.txt file!.
  
📋 | Iptables + IPset: 

(( For adding the ips to list and dropping their connection ))

 - At first you need to install ipset and iptables using command: apt install ipset, apt install iptables
 - after you install them you will create your ipset using command: ```ipset create blocked_ips hash:ip```
 - after that you will add your ips from connections.txt or blacklsit.txt from sonar to the ipset using this commands:
  - For kokot.go: ```while IFS= read -r ip; do     sudo ipset add blocked_ips "$ip"; done < "connections.txt"```
  - For Sonar.jar: ```while IFS= read -r ip; do     sudo ipset add blocked_ips "$ip"; done < "yourblacklist.txt"```
 - After you are done adding this into the ipset you need to drop the ips from ipset using this iptables command: ```sudo iptables -A INPUT -m set --match-set blocked_ips src -j DROP```
 
 🐛 | Issues:
 
 (( If you have any Issues with this contact me on social media ))
  - DISCORD: **Looph#2695**
