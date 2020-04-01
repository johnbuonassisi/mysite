# MySite - John Buonassisi's Website

## My simple, personal, mostly static website, served by a Go HTTP server. The
purpose of the project was to introduce myself to HTML, CSS, and Go's http
package. The site consists of a main page that acts as an overview about myself
and a bit of a resume. I intend to also implement a blog page that will render
blog posts I create in mardown format.

## To run locally:
1. cd mysite
2. go build
3. ./mysite
4. Got to web browser and enter localhost as the url

## To run in AWS EC2 Instance:
1. SSH into AWS EC2 Instance
1. Git clone mysite project from github, https://github.com/johnbuonassisi/mysite.git
2. cd mysite
3. Build the http server, go build -v
4. Run a screen session, screen
5. Run the Go binary ./mysite
6. Exit screen session to keep mysite running in the background. ctrl+a, ctrl+d
7. Exit SSH session with EC2 Instance

## AWS EC2 Instance Setup:
1. Run a AWS Linux EC2 Instance, I used a free nano tier
2. Add security policies such that inbound http connections are allowed
3. Create an Elastic IP and associate it with the AWS EC2 Instance
4. Buy a domain name, I bought johnbuonassisi.ca from AWS Route 53
5. Create a Registered Domain for johnbuonassisi.ca from AWS Route 53
6. Wait a while (up to 2 days) for DNS propogation
7. With mysite running in the EC2 instance, navigate to the domain in a web 
   browser. Like www.johnbuonassisi.ca.

**Note that I have added User Data to be AWS instance so steps 1-3 are executed 
each time my website instance is started.**

## To stop in AWS EC2 Instance:
1. Run screen -r to enter the screen session running mysite
2. Do ctrl-c to stop mysite
3. Exit screen session, ctrl+a, ctrl+d
