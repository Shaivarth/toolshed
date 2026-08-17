local file = io.open("security.log", "r")

if file then

    for line in file:lines() do

        if line:find("LOGIN_SUCCESS") then
            print("[+] Successful login")
        elseif line:find("LOGIN_FAILED") then
            print("[!] Failed login")
        end

    end

    file:close()

else
    print("Could not open security.log")
end