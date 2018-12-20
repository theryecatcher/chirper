[![Build Status](https://www.travis-ci.com/theryecatcher/chirper.svg?token=shEHxWnqj4LyQxYpERrT&branch=master)](https://www.travis-ci.com/theryecatcher/chirper)

<h1 align="center">chirper</h1>
Twitter clone for Ditsributed Systems Course

### Environment
You need to have Go Environment setup on your local machine.

### Get the code
Just do a git pull on the repo or better you could do a go get any one of the below commands should suffice.
```
git clone https://github.com/theryecatcher/chirper/
go get github.com/theryecatcher/chirper
```

### Have Fun!!!
The code is designed as an in Memory application with Persistent storage provided by the RAFT Consensus Protocol. The Radt Wrapper, Content Wrapper and the User Wrapper run as daemons and are connected as microservices for the main web server. 

The application can be started using the startall.sh bash script as below. The web server doesn't run as a daemon and is a live shell command in the terminal that runs the startall script (Ctrl + C kills it).
```
cd chirper
./startall.sh
```
Note that as the web server uses the HTTP sockets module so you will need to give sudo access to your system which is present in the script just need to provide the sudo password after you start the script.

You can also run them separately to have a functional application. Please follow the below sequence of commands to achieve the same.

You can append & (Run in the background) to each of the processes. Or else For Debugging Purposes please run in different Terminals:
```
cd $GOPATH
src/github.com/theryecatcher/chirper/cmd/backendCntD/backendCntD
src/github.com/theryecatcher/chirper/cmd/backendUsrD/backendUsrD
sudo src/github.com/theryecatcher/chirper/cmd/web/web
```

For killing any of the nodes or daemons just run the commands in kill.txt.

As seen above the commands start the daemons and also the main web application. Until this is hosted it is configured to run on localhost.

You can now browse to "http://localhost/" to start playing with the application.

### Disclaimer!!!
The application is still under active development so kindly let us know if you face any bugs contact details are provided in the About Page.
