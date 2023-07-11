from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.firefox.options import Options
import sys
import time
import json

options = Options()
options.add_argument('-headless')
browser = webdriver.Firefox(options=options)

print("read")

with open('local_storage_data.txt', 'r') as file:
    json_data = file.read()

local_storage_data = json.loads(json_data)

browser.execute_script(f"window.localStorage.clear();")
for key, value in local_storage_data.items():
    script = f"window.localStorage.setItem('{key}', '{value}');"
    browser.execute_script(script)

with open('session_cookies.txt', 'r') as file:
    session_cookies = []
    for line in file:
        name, value = line.strip().split('=')
        cookie = {'name': name, 'value': value}
        session_cookies.append(cookie)

for cookie in session_cookies:
    browser.add_cookie(cookie)

browser.get('https://app.slack.com/client')
