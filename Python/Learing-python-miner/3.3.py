from urllib.parse import urlparse
from urllib.request import urlopen
from bs4 import BeautifulSoup
import re
import random
from datetime import datetime

pages = set()
random.seed(datetime.now())

#获取页面所有内链的列表
def getInternalLinks(bsObj,includeUrl):
    includeUrl = '{}://{}'.format(urlparse(includeUrl).scheme, urlparse(includeUrl).netloc)
    print('includeUrl: {}'.format(includeUrl))
    InternalLinks = []
    # 找出所有以“/"开头的链接
    for link in bsObj.find_all('a',href = re.compile('^(/|.*'+includeUrl+')')):
        if link.attrs['href'] is not None:
            if link.attrs['href'] not in InternalLinks:
                # 看看是不是相对路径 ， 还是一个完整的链接
                if link.attrs['href'].startswith('/'):
                    InternalLinks.append(includeUrl + link.attrs['href'])
                else:
                    InternalLinks.append(link.attrs['href'])
    return InternalLinks

#获取所有外链的列表
def getExternalLinks(bsObj,excludeUrl):
    externalLinks = []
    #找出所有以“http” or “www"开头且不包含当前URL的链接
    for link in bsObj.find_all("a",href = re.compile("^(htttp|www)((?!"+excludeUrl+").)*$")):
        if link.attrs['href'] is not None:
            if link.attrs['href'] not in externalLinks:
                externalLinks.append(link.attrs['href'])
    return externalLinks


def getRandomExternalLink(startingPage):
    html = urlopen(startingPage)
    bsObj = BeautifulSoup(html,"html.parser")
    externalLinks = getExternalLinks(bsObj,urlparse(startingPage).netloc)
    if len(externalLinks) == 0:
        print("No external links, looing around the site for one")
        domain = '{}://{}'.format(urlparse(startingPage).scheme,urlparse(startingPage).netloc)
        internalLinks = getInternalLinks(bsObj,domain)
        nextPage =  internalLinks[random.randint(0,len(internalLinks)-1)]
        print(nextPage)
        return getRandomExternalLink(nextPage)
    else:
        return externalLinks[random.randint(0,len(externalLinks)-1)]

def followExternalOnly(startingSite):
    exteralLink = getRandomExternalLink(startingSite)
    print('Random external link is:{}'.format(exteralLink))
    followExternalOnly(exteralLink)

followExternalOnly('https://oreilly.com')