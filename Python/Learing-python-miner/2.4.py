from bs4 import BeautifulSoup
from urllib.request import urlopen
from urllib.error import HTTPError
import re

def getBsObj(url):
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

url = "http://www.pythonscraping.com/pages/page3.html"
bsObj = getBsObj(url)
if(bsObj != None):
    images = bsObj.findAll("img",{"src":re.compile("\.\.\/img\/gifts/img.*\.jpg")})
    for image in images:
        print(image)
        print(image["src"])
else:
    print("something wrong happen")