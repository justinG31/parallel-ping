#  parallel-ping

parallel-ping will take three websites as inputs, and will create three go-routines per website. Each go-routine will utilize singlePing() to ping a website 100 times. 

# How to Run

Clone the repository in the terminal:
git clone https://github.com/justinG31/parallel-ping.git

Run the main program :
go run main.go


###  Websites to ping in analysis
    - google.com
    - bc.edu
    - gov.uk


### Output and analysis
The output consists of the runtime for 100 pings of each website three times for a total of 9 go-routines. Additonally, the output includes the runtime of the ovearall program excluding the user-input time. 

For analysis, we run the program numerous times at different points of the day changing the variable for the number of processors. The evaluation is done comparing the average run time in relation to the the amount of processors utilized. 


