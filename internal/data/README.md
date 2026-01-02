# MRDS csv parser

In order for this program to work you must include the ``mrds.csv`` file in this directory. Since it is too big to include on the repo, it can be found through the following link:

https://mrdata.usgs.gov/mrds/

The correct download is the ``CSV (single)`` format option.

### DepID primary key

DepID makes a great primary key as proven by the ``validation_test::TestDepIDUnique`` test case. 