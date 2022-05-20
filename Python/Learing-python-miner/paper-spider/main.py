from urllib.request import urlopen
from bs4 import BeautifulSoup
from urllib.error import HTTPError

def getBsObj(url):
    try:
        html = urlopen(url)
    except HTTPError as e:
        print(e)
        return None
    bsObj = BeautifulSoup(html,features="lxml")
    return bsObj

def getPapers(url):
    papers = {}
    bsObj = getBsObj(url)
    themeList = bsObj.body.find("div",{"id":"main"}).findAll("ul",{"class":"publ-list"})
    for theme in themeList:
        itemList  = theme.findAll("li",{"class":"entry inproceedings"})
        for item in itemList:
            authors = ""
            for author in item.cite.findAll("span",{"itemprop":"author"}):
                authors += "{},".format(author.a.span.get_text())
            title = item.cite.find("span",{"class":"title"}).get_text()
            papers[authors]=title
    return papers

url =  "https://dblp.org/db/conf/dac/dac2019.html"
# url = "https://dblp.org/db/conf/micro/micro2019.html"
print(getPapers(url))