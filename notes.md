First run was a base line and just processed the files data in a single thread format. Doing it this way the measurements.txt 
was fully processed in:
1 minute 37 seconds

The first improvement I made was to read the file in chunks. I broke the file up by the amount of cores my laptop has (12)
and then done the processing through Go funcs, passing the data across channels. This drastically improved the performance
and got the data parsing down to:
29 seconds



