import requests
import json

responses = []
names = []
venueIDs = []
locations = []
hours = []


def searchStore(name):
  URL = "https://api.foursquare.com/v2/venues/search?client_id=4TCOBRNYJSHUVUA21XNC1NGYF1EFSMKSSCPTLPHXE1KITWME&client_secret=M4JZCLSXBFPDKS0G52OYPI1J13KIYDIYKLS5J0ASRCDKDVTE&near=La%20Jolla%2C%20CA&categoryId=4bf58dd8d48988d118951735&v=20180101&fbclid=IwAR0Q7Xexj1M-x9TL2B5_ne0z21DRYh1YgT0L-Ka12fj9lherx1TlcTS7YBs"
  responses.append(requests.get(URL).json().get("response").get("venues"))

#print(json.dumps(responses[0], indent=4))

def getNameIdAndCatID(responsesToFind):
  for response in responsesToFind:
    names.append(response.get("name"))
    venueIDs.append(response.get("id"))
    
def getLocation(locationsToSearch):
  for location in locationsToSearch:
    LOCATION_DICT = {
         "lat": "",
         "lng": ""
    }

    LOCATION_DICT["lat"] = location.get("location").get("lat")
    LOCATION_DICT["lng"] = location.get("location").get("lng")

    locations.append(LOCATION_DICT)
   
#  print(json.dumps(locations, indent=4))

def getTime(venueIDsToSearch):
#  for id in venueIDsToSearch:
    id = venueIDsToSearch[0]
    URL = "https://api.foursquare.com/v2/venues/" + id + "/hours?" + "client_id=4TCOBRNYJSHUVUA21XNC1NGYF1EFSMKSSCPTLPHXE1KITWME&client_secret=M4JZCLSXBFPDKS0G52OYPI1J13KIYDIYKLS5J0ASRCDKDVTE&v=20180101"
    HOURS = requests.get(URL).json().get("response").get("popular").get("timeframes")
    HOUR_DICT = {
         "1": {
                  "close": "",
                  "open": ""
               },

         "2": {
                  "close": "",
                  "open": ""
               },

         "3": {
                  "close": "",
                  "open": ""
               },

         "4": {
                  "close": "",
                  "open": ""
               },

         "5": {
                  "close": "",
                  "open": ""
               },

         "6": {
                  "close": "",
                  "open": ""
               },

         "7": {
                  "close": "",
                  "open": ""
               },
         }

    for same_hour_group in HOURS:
      for days in same_hour_group["days"]:
        HOUR_DICT[str(days)]["close"] = same_hour_group.get("open")[0].get("end")
        HOUR_DICT[str(days)]["open"] = same_hour_group.get("open")[0].get("start")

    hours.append(HOUR_DICT)

allParsed = []
def parseAll():
#  if len(names) != len(venueIDs) or len(names) != len(hours) or len(names) != len(locations):
#    print("Length of dicts don't match")
#  else:
#    for i in range(len(names)):
      STORE_DICT = {
        venueIDs[0]: {
                      "categories": "",
                      "hours": hours[0],
                      "location": locations[0],
                      "name": names[0],
                 }
      }
      
      allParsed.append(STORE_DICT)
#print(json.dumps(allParsed, indent=4))
      
#print(json.dumps(allParsed, indent=4))



searchStore("VONS")
getNameIdAndCatID(responses[0])
getTime(venueIDs)
getLocation(responses[0])

#yay we finally donezo :)
parseAll()
print(allParsed[0])
r = requests.post(url = "https://honey4groceries.firebaseio.com/stores.json", data=json.dumps(allParsed[0]).replace("\'", "\"")).json()
print(r)
