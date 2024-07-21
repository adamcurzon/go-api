# Go API ⚙️

This is an API built in GO and deployed with Nginx & Github actions.

## Setup

1. Create `go-api` user

   ```bash
   adduser go-api
   ```

2. Clone the repo into `/home/go-api/web`

3. Run the first build of the api

   ```bash
   go build main.go
   ```

4. Create a system service to run the go binary as the go-api user

   ```bash
   sudo nano /lib/systemd/system/go-api.service
   ```

   ```yaml
   [Unit]
   Description=go-api


   [Service]
   Type=simple
   Restart=always
   RestartSec=5s
   User=go-api
   Group=go-api
   ExecStart=/home/go-api/web/main

   [Install]
   WantedBy=multi-user.target
   ```

5. Start the service

   ```bash
   service go-api start
   ```

6. Check the service is running and the local api is accessable

   ```bash
    service go-api status
    curl localhost:9990
   ```

7. Update permissions to allow the `go-api` user to restart the service for future deployments

   ```bash
   sudo visudo
   ```

   ```yaml
   go-api ALL=(ALL) NOPASSWD: /usr/sbin/service go-api restart
   ```

8. Create Nginx proxy to forward the local port 9990 to the web

   ```
   sudo nano /etc/nginx/sites-available/go-api
   ```

   ```bash
   server {
       server_name go.adamcurzon.co.uk;

       location / {
           proxy_pass http://localhost:9990;
       }
   }
   ```

9. Enable the site in Nginx

   ```
   ln -s /etc/nginx/sites-available/go-api /etc/nginx/sites-enabled/
   ```

10. Restart Nginx

    ```bash
    service nginx restart
    ```

11. Check the live site

    ```bash
    curl https://go.adamcurzon.co.uk
    ```

12. Create a `deploy.sh` script in the `/home/go-api/devops` folder

    ```bash
    git -C ~/web pull origin main
    cd ~/web; go build main.go
    sudo service go-api restart
    ```

13. Add required secrets for the `main.yml` deployment action in GitHub
14. Push a new commit to the `main` branch and check the actions has successfully ran
15. The api is live with ci/cd 🎉
