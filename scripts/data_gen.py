"""
generate data
"""
import json
import random
import os
# detail1 := details.Detail{
# 		ProductID: req.GetId(),
# 		Author:    "William Shakespeare",
# 		Year:      1595,
# 		Type:      "paperback",
# 		Pages:     200,
# 		Publisher: "PublisherA",
# 		Language:  "English",
# 		ISBN10:    "1234567890",
# 		ISBN13:    "123-1234567890",
# 	}
def gen_detail(n, fd):
    details = []
    for i in range(n):
        detail = {
            "ProductId": i + 1,
            "Author": "William Shakespeare",
            "Year": 1596 + i,
            "Type": "paperback",
            "Publisher" : "PublisherA",
            "Language":  "English",
            "ISBN10":    "1234567890",
            "ISBN13":    "123-1234567890",
        }
        details.append(detail)
    json.dump(details, fd, indent=4)
    

# review1 := reviews.Review{
# 		ProductID: productID,
# 		Reviewer:  "reviewer1",
# 		Text:      "An extremely entertaining play by Shakespeare. The slapstick humour is refreshing!",
# 	}

def gen_review(n, fd):
    reviews = []
    # 3 review for 1 product
    for i in range(n * 3):
        review = {
            "ProductId": i // 3 + 1,
  		    "Reviewer":  "reviewer" + str(i + 1),
            "Text": "Absolutely fun and entertaining. The play lacks thematic depth when compared to other plays by Shakespeare."
        }
        reviews.append(review)
    json.dump(reviews, fd, indent=4)
    

def gen_rating(n, fd):
    ratings = []
    for i in range(n):
        ratings.append({
            "ProductId": i + 1,
            "Ratings": random.randint(1, 10)
        })
    json.dump(ratings, fd, indent=4)
    
if __name__ == "__main__":
    n = 20
    print("generating data with n = {}".format(n))
    os.system("mkdir -p ./data")
    for i in ["details", "reviews", "ratings"]:
        fullpath = "./data/{}.json".format(i)
        with open(fullpath, "w") as f:
            if i == "details":
                gen_detail(n, f)
            elif i == "reviews":
                gen_review(n, f)
            elif i == "ratings":
                gen_rating(n, f)
    print("done")