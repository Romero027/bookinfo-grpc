-- ./wrk/wrk -t1 -c1 -d 10s http://10.96.88.88:8080 -L -s ./scripts/lua/bookinfo.lua


local function productpage()
    local method = "GET"
    local path = "http://10.96.88.88:8080/productpage"
    local headers = {}
    return wrk.format(method, path, headers, nil)
end

local function login()
    local method = "GET"
    local path = "http://10.96.88.88:8080/login"
    local headers = {}
    return wrk.format(method, path, headers, nil)
end
  
request = function()
    return productpage()
end