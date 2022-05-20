import csv
from urllib.request import urlopen
from bs4 import BeautifulSoup

csvFile = open("./csvFiles/test.csv","w+")
try:
    writer = csv.writer(csvFile)
    writer.writerow(('number','number plus 2','number times 2'))
    for i in range(10):
        writer.writerow((i,2+i,i*2))
finally:
    csvFile.close()


# another task
html = urlopen("http://en.wikipedia.org/wiki/Comparison_of_text_editors")
bsObj = BeautifulSoup(html)
# 主对比表格是当前页面上的第一个表格
table = bsObj.findAll("table",{"class":"wikitable"})[0]
rows = table.findAll("tr")

csvFile = open("./csvFiles/editors.csv","wt",newline="\n",encoding='utf-8')
writer = csv.writer(csvFile)
try:
    for row in rows:
        csvRow = []
        for cell in row.findAll(['td','th']):
            csvRow.append(cell.get_text())
            writer.writerow(csvRow)
finally:
    csvFile.close()