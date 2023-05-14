# :car: Golang Mini Project - Vehicle Rent :motorcycle:

## Introduction

The main purpose of creating a rental application is to make it easy for customers to search and rent the vehicles they need. The rental application can connect customers with online vehicle rental service providers, allowing them to easily choose the desired vehicle and quickly and efficiently make reservations. The reason for creating this application is based on personal experience of renting vehicles using personal chats through social media which, in my opinion, is less efficient and less organized. With the rental application, the vehicle rental process becomes more structured and well-organized. Additionally, the rental application can provide security and convenience for customers in conducting transactions, as they can view complete information about the vehicle and make online payments.

## Use Case Diagram

![Use Case Diagram](https://github.com/hklfach/go_mini-project/blob/master/docs/usecase.jpg?raw=true)

## Entity-Relationship Diagram

![ERD](https://github.com/hklfach/go_mini-project/blob/master/docs/erd.jpg?raw=true)

## Application Flow 

Use postman or other API tools to test the API
1. Admin register and login (Use the token for authorization)
2. Admin creates, edits, or deletes vehicle (available car and motorcyle only)
3. Customer register and login (Use the token for authorization)
4. Customer finds the vehicle by their name
5. Customer creates the order
6. Customer pays their order
7. Customer gives rate to their order
8. Customer gets their order history
9. Admin register and login (Use the token for authorization)
10. Admin gets order list
11. Admin updates the rating of a vehicle based on orders

## How to Run Application in Docker

***NOTES: Make sure docker is installed***

1. Run docker compose in root folder
    ```
    run docker-compose up -d
    ```

2. Check running container
    ```
    docker ps
    ```

3. Check app container logs
    ```
    docker logs {$container_id}
    ```

***NOTES: Make sure the logs shows echo framework logo and the port***

![echo](https://github.com/hklfach/go_mini-project/blob/master/docs/echo.jpg?raw=true)

4. Test the app using API tools
    - Use localhost with port 8000
    - Endpoint available in [./routes/route.go](https://github.com/hklfach/go_mini-project/blob/master/routes/routes.go)
    
    Example:
    ```
    http://localhost:8000/login
    ```
    ```
    http://localhost:8000/auth/users
    ```

## How to Deploy Application in AWS

***NOTES: Make sure to create EC2 instances first***

1. Open terminal (You can use CMD if using windows)

2. Connect to EC2 instance using SSH (Different based on instance configure)

    Formula command
    ```
    ssh -i "public_key" username@Public_IPv4_DNS
    ```

    Example command
    ```
    ssh -i "keys.pem" ec2-user@ec2-18-136-126-223.ap-southeast-1.compute.amazonaws.com
    ```

3. Install docker, docker-compose, git, mysql server, and nginx

    Get into super user mode
    ```
    sudo su
    ```

    Update package
    ```
    yum update -y
    ```

    Install git
    ```
    yum install git -y
    ```
    ```
    git -v
    ```

    Install docker
    ```
    yum install -y docker
    ```
    ```
    docker -v
    ```

    Install docker-compose
    ```
    curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/bin/docker-compose && chmod +x /usr/bin/docker-compose
    ```
    ```
    docker-compose -v
    ```

    Install nginx
    ```
    amazon-linux-extras install nginx1.12
    ```
    ```
    nginx -v
    ```

    Install mysql server
    ```
    amazon-linux-extras install epel -y 
    ```
    ```
    yum install https://dev.mysql.com/get/mysql80-community-release-el7-5.noarch.rpm 
    ```
    ```
    yum install mysql-community-server
    ```

4. Clone repository
    ```
    git clone --single-branch --branch master https://github.com/hklfach/go_mini-project.git
    ```

5. Configure nginx for reverse proxy


    Go to nginx directory
    ```
    cd /etc/nginx/
    ```

    Edit nginx.conf using nano
    ```
    nano nginx.conf
    ```

    Comment root variable in line 42 and add proxy_pass inside location {}
    ```
    ...
    server {
        listen       80 default_server;
        listen       [::]:80 default_server;
        server_name  _;
        # root         /usr/share/nginx/html;

        # Load configuration files for the default server block.
        include /etc/nginx/default.d/*.conf;

        location / {
                proxy_pass http://127.0.0.1:8000;
        }

        error_page 404 /404.html;
            location = /40x.html {
        }

        error_page 500 502 503 504 /50x.html;
            location = /50x.html {
        }
    }
    ...
    ```

6. Use Docker and NGINX

    Start docker
    ```
    service docker start
    ```

    Get into go-learn-deploy directory
    ```
    cd go-learn-deploy
    ```

    Use docker compose
    ```
    docker-compose up -d
    ```

    Start NGINX
    ```
    systemctl start nginx
    ```

7. Open program in public IP Adress and Test API using Postman

    Example
    ```
    http://18.136.126.223/
    ```
## (Optional) Check database using mysql service
    
1. Use docker-compose exec to get into container service
    ```
    docker-compose exec mysql bash
    ```

2. Login to mysql (Password: password)
    ```
    mysql -u root -p -P3307 mini_project_db
    ```

3. Run any sql command
    ```
    SHOWS DATABASES;
    ```

## Notes for Future Update

- Create Continous Deployment using [Github Actions](https://docs.github.com/en/actions) to deploy in AWS or else
- Using path/url to store payment receipt image instead of using blobs
- Add reservation feature
- Frontend integration
- Fully tested unit testing (94.1% Coverage right now)

