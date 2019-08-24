# USCIS case Tracker
This project was build accomplish the need for monitoring the status of the cases for the range provided. Currently the script is set to run every hour.
    
## Build project
```go build uscis.ho```
    
## Run
```./uscis``` 
        
## Arugments
```
-center int
    wac = 0 & eac = 1; default:0
-email string
    gmail address for sending and receiving status
-limit int
    no. of cases; default:100 (default 100)
-pass string
    gmail password for sending and receiving status
-start int
    start of the series; default:1919054100 (default 1919054100)
```
## Sample output
 ```
Change in status for wac19190xxxxx - Case Was Approved -----> Fingerprint Review Was Completed
Change in status for wac19190xxxxx - Case Was Received -----> Fingerprint Review Was Completed
-----------SUMMARY at 2019-08-24 11:57:59.848445001 -0500 CDT m=+7276.894935229---------
Fingerprint Review Was Completed - 13
Case Was Received - 79
Request for Additional Evidence Was Sent - 55
Case Was Approved And My Decision Was Emailed - 24
Case Was Approved - 76
Withdrawal Acknowledgement Notice Was Sent - 1
Case Was Transferred And A New Office Has Jurisdiction - 10
Card Was Delivered To Me By The Post Office - 7
Fee Refund Was Mailed - 1
Decision Notice Mailed - 2
Name Was Updated - 3
CASE STATUS - 1
Request For Premium Processing Services Was Received - 1
Response To USCIS' Request For Evidence Was Received - 10
Case Was Reopened - 1
Case Was Received and A Receipt Notice Was Emailed - 1
Document Was Mailed - 2
Notice Explaining USCIS Actions Was Mailed - 2
Correspondence Was Received And USCIS Is Reviewing It - 4
Card Was Mailed To Me - 1
Date of Birth Was Updated - 1
--------map size : 295---------
-----------SUMMARY---------     
 ```
