# status_checker
Checks the status of websites
Supports checking the status of multiple websites in parallel

To check the status of a specific website, add its domain to the links.txt file in a new line.

## Production Files
windows : main.exe

linux : main

## Steps:
1. Open the working directory in the terminal and run the production file.
2. The value after the command determines whether to keep checking the website statues continuously.
(eg: true )
3. If choosen true in previous step, the user can also add the time interval((in seconds) between checking the status of the same domain.
(eg: 5) (default value: 3)

## Examples:

1. To run once at default:
    ./main
    
2. To run continuously:
    ./ main true
    
3. To run continuously with interval of 6 seconds:
    ./ main true 6
