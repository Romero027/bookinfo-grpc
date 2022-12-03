-- ./wrk/wrk -t1 -c1 -d 10s http://10.96.88.88:8080 -L -s ./scripts/lua/bookinfo.lua


local function req()
    local method = "GET"
    local path = "http://10.96.88.88:80/"
    local headers = {}
    return wrk.format(method, path, headers, nil)
  end
  
  request = function()
      return req()
  end