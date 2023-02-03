import sys
import os

docker_old = ""
docker_new = ""
go_old = ""
go_new = ""

for root, dirs, files in os.walk(".", topdown=False):
    for file in files:
        name = file.split('.')
        if len(name) == 2:
            if name[1] == "go" or name[1] == "yaml":
                path = str(os.path.join(root, name))
                print("replacing {}".format(path))
                with open(path, "r") as f:
                    content = f.read()
                if name[1] == "yaml":
                    content = content.replace(docker_old, docker_new)
                if name[1] == "go":
                    content = content.replace(go_old, go_new)
                with open(path, "w") as f:
                    f.write(content)
                
                    
        
