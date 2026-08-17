local file = io.open("security.log", "r")

if file then

    for line in file:lines() do

        if line:find("LOGIN_SUCCESS") then

            local username = line:match("user=(%w+)")
            print("[+] Successful login: " .. username)

        elseif line:find("LOGIN_FAILED") then

            local username = line:match("user=(%w+)")
            print("[!] Failed login: " .. username)

        end

    end

    file:close()

else
    print("Could not open security.log")
end