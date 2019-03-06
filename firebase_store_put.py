import requests
import json

responses = []
names = []
venueIDs = []


def searchStore(name):
  URL = "https://api.foursquare.com/v2/venues/search?client_id=4TCOBRNYJSHUVUA21XNC1NGYF1EFSMKSSCPTLPHXE1KITWME&client_secret=M4JZCLSXBFPDKS0G52OYPI1J13KIYDIYKLS5J0ASRCDKDVTE&near=La%20Jolla%2C%20CA&categoryId=4bf58dd8d48988d118951735&v=20180101&fbclid=IwAR0Q7Xexj1M-x9TL2B5_ne0z21DRYh1YgT0L-Ka12fj9lherx1TlcTS7YBs"
  responses.append(requests.get(URL).json().get("response").get("venues"))

def getNameAndId(responsesToFind):
  for response in responsesToFind:
    names.append(response.get("name"))
    venueIDs.append(response.get("id"))

def getTime(venueIDsToSearch):
#for id in venueIDsToSearch:
    id = venueIDsToSearch[0]
    URL = "https://api.foursquare.com/v2/venues/" + id + "/hours?" + "client_id=4TCOBRNYJSHUVUA21XNC1NGYF1EFSMKSSCPTLPHXE1KITWME&client_secret=M4JZCLSXBFPDKS0G52OYPI1J13KIYDIYKLS5J0ASRCDKDVTE&v=20180101"
    HOURS = requests.get(URL).json().get("response").get("popular").get("timeframes")
    print(HOURS)

    

searchStore("VONS")
getNameAndId(responses[0])
getTime(venueIDs)
