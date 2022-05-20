from urllib.request import urlopen
from bs4 import BeautifulSoup
from urllib.error import  HTTPError

def getHtml(url):
    try:
        html = urlopen(url)
    except HTTPError as e1:
        print(e1)
        return None
    try:
        bsObj = BeautifulSoup(html.read(),features="html.parser")
    except AttributeError as e2:
        print(e2)
        return None
    return bsObj


url = "http://www.pythonscraping.com/pages/warandpeace.html"
bsObj = getHtml(url)
namelist = bsObj.findAll("span",{"class":"green"})
for name in namelist:
    print(name.get_text())
# print(bsObj)