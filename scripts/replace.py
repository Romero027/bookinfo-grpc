"""
This is an ad hoc solution for go and docker usernames.
The script iterates all files in this repo, doing find-and-replace job, so one doesn't need to manually change them.
 
Currently, it will replace all "$docker_old" by "$docker_new" in all `*.yaml` configs. 
Similarly, it will replace $go_old by $go_new in all `*.go` files and `Dockfile`.
"""
import sys
import os

docker_old = "livingshade"
docker_new = "xzhu0027"
go_old = "livingshade"
go_new = "Romero027"

def replace(path, old, new):
    with open(path, "r") as f:
        content = f.read()    
    content = content.replace(old, new)
    with open(path, "w") as f:
        f.write(content)

for root, dirs, files in os.walk(".", topdown=False):
    for file in files: 
        path = str(root + "/" + file)
        name = file.split('.')
        if name == "Dockerfile":
            replace(path, go_old, go_new)
        elif len(name) == 2:
            if name[1] == "go" or name[1] == "yaml":
                print("replacing {}".format(path))                
                if name[1] == "yaml":
                    replace(path, docker_old, docker_new)
                if name[1] == "go":
                    replace(path, go_old, go_new)
               
                
                    
        
