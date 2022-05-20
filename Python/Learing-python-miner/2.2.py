from urllib.request import urlopen
from bs4 import BeautifulSoup
from urllib.error import HTTPError

def getHtml(url):
    try:
        html = urlopen(url)
    except HTTPError as e1:
        print(e1)
        return None
    try:
        bsObj = BeautifulSoup(html.read())
    except AttributeError as e2:
        print(e2)
        return None
    return bsObj

url = "http://www.pythonscraping.com/pages/page3.html"
bsObj = getHtml(url)

# # test chidren
# if(bsObj != None):
#     for child in bsObj.find("table",{"id":"giftList"}).children:
#         print(child)

# # test next_siblings
# if(bsObj != None):
#     for sibling in bsObj.find("table",{"id":"giftList"}).tr.next_siblings:
#         print(sibling)

if(bsObj != None):
    print(bsObj.find("img",{"src": "../img/gifts/img1.jpg"}).parent.previous_sibling.get_text())

