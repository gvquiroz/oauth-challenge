# OAuth Challenge
[![codecov](https://codecov.io/gh/gvquiroz/oauth-challenge/branch/master/graph/badge.svg)](https://codecov.io/gh/gvquiroz/oauth-challenge) [![Build Status](https://travis-ci.org/gvquiroz/oauth-challenge.svg?branch=master)](https://travis-ci.org/gvquiroz/oauth-challenge) [![Maintainability](https://api.codeclimate.com/v1/badges/98c015dc86ede613332e/maintainability)](https://codeclimate.com/github/gvquiroz/oauth-challenge/maintainability)
# How to run with docker

1. build container by executing **docker build -t oauthapi/oauth-challenge:latest .**
2. run cointainer by executing **docker run -itd --name CharCounterAPI -p 4567:4567 oauthapi/oauth-challenge:latest**
3. clean container **docker rm -f CharCounterAPI**

# How to run with go (1.9.2) environment

1. To install deps execute **make deps**
1. To run tests and build execute **make**
# Challenge instructions

Develop a webApi to make the test script located in the test directory to pass. You can not modify the tests, your solution has to make these test pass as they are.

You can use any of these languages to develop your solution: JavaScript, Java, CSharp, Python, Ruby, Groovy, Smalltalk or Php. If you prefer to use something different please check with us before starting.

To run the test script you will need to:

1. Install Node JS (version 8, check the corresponding installation procedure for your OS)
2. Install dependencies by executing _npm Install_
3. Run tests be execting _npm test_

Once you have completed the solution, please zip it and upload it to GDrive/Dropbox. Then share it with us by sending an email to eng.bootcamp@auth0.com including the link to the uploaded zip file.


