Certainly! Here’s how you can format the content in a `README.md` file for GitHub:

```markdown
# EasyBlocker

**CREATED BY Lopusnik!**

## What is this project?

EasyBlocker is designed to block proxies and prevent botters from connecting to a specified random port or a Minecraft server equipped with the Sonar antibot feature. It uses an anti-bot blacklist for enhanced security, which can be enabled by setting `firewall.yml` to `true`.

## 🛠️ Requirements

### Installation of necessary components for this program

#### For Random Port Listening and Logging Proxies

1. **Install Go (Golang):**

   ```bash
   apt install golang
   ```

2. **Download and Run kokot.go:**

   - Download the `kokot.go` file.
   - Run it using:

     ```bash
     go run kokot.go
     ```

   - Or compile and start it:

     ```bash
     go build kokot.go
     ./kokot
     ```

#### For Minecraft Server with Sonar Antibot

1. **Obtain and Configure Waterfall:**

   - Download `Waterfall.jar` [from here](#).
   - Start `Waterfall.jar` and configure `config.yml`.

2. **Add Plugins:**

   - Add `sonar.jar` to the `plugins` directory.
   - Add `luckperms-bungee` to the `plugins` directory.

3. **Restart the Server:**

   - Restart your Minecraft server.

4. **Grant Permissions:**

   - Enter the following command into the BungeeCord server console:

     ```bash
     lp user <Your Minecraft NickName> permission set *
     ```

   - This grants full server permissions and access to the Sonar plugin.

5. **Configure Attack:**

   - Use your public IPv4 address with port `25577` (or the port specified in `config.yml`) for your attack.
   - After the attack, run:

     ```bash
     ab blacklist save
     ```

   - This will generate a message in the chat with the name of the blacklist. You can find your `blacklist.txt` file in the `Sonar` folder within the `plugins` directory.

## 📋 Iptables + IPset

### Adding IP Addresses to the List and Blocking Connections

1. **Install ipset and iptables:**

   ```bash
   apt install ipset iptables
   ```

2. **Create an IP Set:**

   ```bash
   ipset create blocked_ips hash:ip
   ```

3. **Add IP Addresses:**

   - For `kokot.go`:

     ```bash
     while IFS= read -r ip; do sudo ipset add blocked_ips "$ip"; done < "connections.txt"
     ```

   - For `Sonar.jar`:

     ```bash
     while IFS= read -r ip; do sudo ipset add blocked_ips "$ip"; done < "yourblacklist.txt"
     ```

4. **Drop Connections:**

   ```bash
   sudo iptables -A INPUT -m set --match-set blocked_ips src -j DROP
   ```

## 🐛 Issues

If you encounter any issues, please contact me on social media:

- **DISCORD:** drindez
```

Copy and paste this into your `README.md` file on GitHub, and it should display with the intended formatting. Let me know if you need any more help!
