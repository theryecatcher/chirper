[![Build Status](https://www.travis-ci.com/theryecatcher/chirper.svg?token=shEHxWnqj4LyQxYpERrT&branch=master)](https://www.travis-ci.com/theryecatcher/chirper)

<h1 align="center">chirper</h1>
Twitter clone for Ditsributed Systems Course

### Environment
You need to have Go Environment setup on you local machine.

### Get the code
Just do a git pul on the repo or better you could do a go get any one of the below commands should suffice.
```
git clone https://github.com/theryecatcher/chirper/
go get github.com/theryecatcher/chirper
cd chirper/
```

### Have Fun!!!
The code is designed as a in Memory application for the 1st Phase. The Content Database and the User Database run as daemons and aren't part of the main application. You would need to run them seperately to have a functional application. Please follow the below sequence of commands to acheive the same.

You can append & (Run in background) to each of the processes. Or else For Debugging Purposes please run in different Terminals:
```
cd $GOPATH
src/github.com/theryecatcher/chirper/cmd/backendCntD/backendCntD
src/github.com/theryecatcher/chirper/cmd/backendUsrD/backendUsrD
sudo src/github.com/theryecatcher/chirper/cmd/web/web
```

As seen above the commands start the daemons and also the main wen application. Until this is hosted it is configured to run on localhost.

You can now browse to "localhost/" to start playing with the application.

### Disclaimer!!!
The application is still under active development so kindly let us know if you face any bugs contact details are provided in the About Page.
