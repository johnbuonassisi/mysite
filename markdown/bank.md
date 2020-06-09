During the past year I finally decided to start learning more about personal finance to understand how I could better structure my finances
to achieve my desired goals. To analyze my expenses I decided to create a Mint account. For those not familiar with Mint, it is an app and 
website that connects to your financial accounts and displays them all in one place. It also has nifty features like the ability to auto-categorize transactions, which is what I was most interested in.  

There was one problem, after connecting all my accounts to Mint I was only seeing the last 3 months of transactions when I expected to see many years worth. I quickly found
[this](https://help.mint.com/Accounts-and-Transactions/888963071/How-can-I-get-transactions-older-than-90-days.htm), financial institutions Mint connects to only provide 3 
months of previous transactions to Mint. This was unfortunate because the only transactions I could see were starting in March 2020, when COVID-19 took grasp of the world
and forced many countries into isolation. So naturally, my expenses for these 3 months were not indicative
of my regular expenses. 

I decided to investigate how I could get older transactions into Mint. I quickly found out there was no good way to do this.

### The Problem with Mint

Mint allow usere to [add transactions manually](https://help.mint.com/Accounts-and-Transactions/888960761/How-do-I-add-manual-transactions.htm)
but I had hundred of transactions to add, I wasn't going to manually enter all of them. Mint provides no way to upload a batch of transactions.

However, when you manually create a transaction in Mint, a simple HTTP POST request with a cookie, form-data, and a token is sent by the browser to Mint.
I could write a script to automatically do this for a bunch of my transactions...

### The Problem with my Bank

Even if Mint did have a way to upload a batch of transactions, how would I get these transactions from my bank in the first place? Looking at my 
online banking account, it too only provided the last 3 months worth of transactions. Older transactions were contained in statements that I could
download in PDF format. Unfortunately PDFs are notoriously difficult for computer programs to read as they do not contain structured data, like
HTML or JSON.

However, if I could convert the PDF statements to CSV files, then I could write a script that would read the CSV file and for each line send a 
request to Mint to create a transaction.


### The Solution

Enter a service called [DocParser](www.docparser.com)! After a lot of searching I found this platform which provides the ability to parse PDFs into structured data
and send that data to services of your choice. In my case, I could upload a bank statement, have the transactions parsed into a table, then have a webhook send a
POST request to Mint for each transaction. Here is how I did it.

### Step 1 - Create Initial Document Parser

The way DocPaser works is that you upload a template PDF document and then use this document as a template to setup parsing rules. After setting up your parser you are able
to upload PDF documents that follow the same format and have them parsed using the configured parser.

Create a new parser

Select the type of document you want to parse, in my case I selected "Bank Statements"

Upload a sample bank statement

Create your first parsing rule by creating the column boundaries in your bank statement. Docparser will use these boundaries to convert the statement from its PDF form to a table.

Now you can add filters to refine the produced table. In my case I first filtered the table to start when column 2 contains "Opening Balance" and when column 2 does not have a value.
This removed all the rows that did not correspond to the transactions in my statement. I then fixed up the date column by setting the value of empty cells to the value of the previous row. 
Then, I formated the date column to be m/d (11/06) instead of d M (6 Nov), this would make the date easier to use when it came time to create the webhook. The last filtering step was to
add column titles to each column; Date, Description, Withdrawls, Deposits, and Balance.

### Step 2 - Refine Document Parser

Unfortunately, after creating the initial document parser there were still some issues. First, my bank statement did not include the year in the data of each transaction. I would need to know
the year when creating the webhook. Second, the statement I uploaded contained 2 pages, and the transaction table on the second page did not line up with the first page, so no transactions
on the second page were parsed.

First, the year of the statement could be extracted from the first page of the statement in the top right corner. I created a second parsing rule to extract just the year from the statement.

Second, getting page 2 transactions proved to be simple. I create a third parsing rule almost identical to the first, but had different column boundaries, and different start and end conditions.
I setup the column boundaries to follow the transaction table on the second page. Then I updated the table filter to start when column 1 contains "Date" and end when column 2 contains "Closing".

Third, I created a final parsing rule to merge the two tables produced by the parsing rules I created for pages 1 and 2.

Each parsing rule creates an output that is usable when creating the webhook. So I could now use the data contained in the merged parsing rule and the year parsing rule in the webhook.

### Step 2 - Create Webhook

Create a new integration in DocParser and select "Advanced Webhook"

By capturing the manual transaction creation request in Google Chrome, I found the request I needed to send to Mint followed the curl request below

```
curl 'https://mint.intuit.com/updateTransaction.xevent' \
  -H 'authority: mint.intuit.com' \
  -H 'pragma: no-cache' \
  -H 'cache-control: no-cache' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36' \
  -H 'x-requested-with: XMLHttpRequest' \
  -H 'content-type: application/x-www-form-urlencoded; charset=UTF-8' \
  -H 'accept: */*' \
  -H 'origin: https://mint.intuit.com' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-dest: empty' \
  -H 'referer: https://mint.intuit.com/transaction.event' \
  -H 'accept-language: en-US,en;q=0.9' \
  mtCashSplit=on&
  mtCheckNo=&
  tag1597252=2&
  tag1572535=0&
  tag1572536=0&
  tag1572537=0&
  task=txnadd&
  txnId=%3A0&
  mtType=pending-other&
  mtAccount=4889566&
  symbol=&
  note=&
  isInvestment=false&
  catId=20&
  category=Uncategorized&
  merchant=Test%20Transaction&
  date=06%2F09%2F2020&
  amount=12.12&
  mtIsExpense=true&
  mtCashSplitPref=1&
  --compressedd
```

So I setup the Advanced Webhook to follow this request. By setting "Repeating Data Behaviour" to "One request per row" the integration would send one
request to Mint for every row in the parsed table. Then all I had to do was set merchant, date, and amount fields in the form data to data that was parsed
from my bank statements. Data from the merged table produced by my parsing rules could be accessed by using the following syntax, {{merged.<field name>}} and
the year I parsed from the statement could be accessed through {{year_of_statement}}. The initial body payload in the webhook describes all the data values
that are accessible when setting up the webhook.

The easiest field to set was "merchant" which is equivalent to the "description" column in the parsed tables. Setting "merchant" to "{{merge.description}}" would
do the trick.

Setting the date filed involved using the "date" column of the parsed tables, and also the year parsed from the statement. Setting date to {{merge.date}}%2F{{year_of_statement}}
would produce a result that looked like 01/02/2019. The %2F is the encoded version of /.

Finally, the amount field could be set. Unfortunately, we have 2 columns to choose from in the parsed merged table, either deposits or withdrawls. This is further complicated by the fact that
the form has an additional field, "mtIsExpense", that should be false for deposits and true for withdrawls. So I decided to create two webhooks, one for withdrawls that would set
"amount" to {{merged.withdrawls}} and "mtIsExpense" to true and another one for deposits that would set "amount" to {{merged.deposits}} and "mtExpense" to false.

### Step 3 - Upload Documents for Processing

Now came the moment of truth, uploading some other statements and seeing if the parsing and webhook would cause my bank statement transactions to show up in Mint.
