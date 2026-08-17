local file = io.open("security.log", "r")

if file then
    local data = file:read("*a")
    print(data)
    file:close()
else
    print("Could not open security.log")
end