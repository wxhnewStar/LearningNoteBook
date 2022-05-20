import datetime
from urllib.request import urlopen
from bs4 import BeautifulSoup
from urllib.error import HTTPError
import re
import random

def getBsObj(url):
    try:
        html = urlopen(url)
    except HTTPError as e1:
        print(e1)
        return None
    bsObj = BeautifulSoup(html)
    return bsObj

def getLinks(articleUrl):
    bsObj = getBsObj("http://en.wikipedia.org"+articleUrl)
    return bsObj.find("div",{"id":"bodyContent"}).findAll("a",href=re.compile("^(/wiki/)((?!:).)*$"))


random.seed(datetime.datetime.now())
links = getLinks("/wiki/Kevin_Bacon")
print(type(links))
while len(links) > 0:
    newArticle = links[random.randint(0,len(links)-1)].attrs["href"]
    print(newArticle)
    links = getLinks(newArticle)