2010-2015 collected via:


======================================
2000-2009 collected via:

for i in $(echo "01 02 03 04 05 06 07 08 09" && for ((i=10; i<=56; i++)); do echo $i; done); do curl -s http://www.census.gov/popest/data/counties/asrh/2007/files/cc-est2007-alldata-$i.csv > $i.csv ; done

rm $(grep \< * | cut -d':' -f1 | uniq)
