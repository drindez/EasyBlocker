EasyBlocker
CREATED BY Lopusnik!
What is this project?
This project is designed to block proxies and prevent botters from connecting to a specified random port or a Minecraft server equipped with the Sonar antibot feature. It employs an anti-bot blacklist for enhanced security, which can be enabled by setting firewall.yml to "true."

🛠️ | Requirements:

(( Installation of necessary components for this program ))

For random port listening for TCP connections and logging every proxy connection to connections.txt:

Install Go (golang): ```apt install go-golang```
Download the kokot.go file.
Run kokot.go using either ```go run kokot.go``` or compile it with ```go build kokot.go```. After compiling, start it with ```./kokot```.
For a Minecraft server with Sonar antibot:

Obtain Waterfall.jar from [here](https://papermc.io/software/waterfall).
Start the WaterFall.jar and configure config.yml.
Add the sonar.jar to the plugins directory along with luckperms-bungee.
After adding Sonar and Luckperms, restart the server.
When you join the BungeeCord server, enter the following command into the console: ```lp user <Your Minecraft NickName> permission set *```. This command grants you full server permissions and access to the Sonar plugin.
Afterward, use your public IPv4 address with port 25577 (if you haven't modified it in config.yml) for your attack. After the attack, run the command ab blacklist save to receive a message in the chat indicating the name of the blacklist.
To access your blacklist.txt file, navigate to the plugins folder and find the Sonar folder.
📋 | Iptables + IPset:

(( Adding IP addresses to the list and blocking their connections ))

First, install ipset and iptables using the following commands: apt install ipset and apt install iptables.
Create an IP set with the command: ipset create blocked_ips hash:ip.
Add IP addresses from connections.txt or yourblacklist.txt (for Sonar.jar) to the IP set using the following commands:
For kokot.go: ```while IFS= read -r ip; do sudo ipset add blocked_ips "$ip"; done < "connections.txt"```
For Sonar.jar: ```while IFS= read -r ip; do sudo ipset add blocked_ips "$ip"; done < "yourblacklist.txt"```
Once you have added the IP addresses to the IP set, use the following iptables command to drop connections from the IP set: ```sudo iptables -A INPUT -m set --match-set blocked_ips src -j DROP```
🐛 | Issues:

(( If you encounter any issues, please contact me on social media ))

DISCORD: lopusnik1337
