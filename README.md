# TaHB EMR Webapp
This directory and its subdirectories contain the source code for a functioning base product for the use of creating and submitting EMR (Electronic Medical Records) onto a AWS EC2 hosted database.

Built using Go for the backend it communicates to the hosted database created by PostgreSQL to a web application created with Javascript.

The project allows the hosting of the webapp along with communications with a created database that can also be hosted as well.
Users are able to create and send records, or save drafts to their accounts for later use.

### Install required packages for local testing
* Download and install the latest release of Go ( at least Go 1.13 ) [here.](https://golang.org)

The code is dockerized and is able to be used on multiple OS platforms such as Windows, MAC, and Linux. 
Other operating systems that can run Go should be able to run as well.
