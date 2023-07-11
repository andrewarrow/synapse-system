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
    local_storage_data = file.read()

# Load session cookies from the file (optional)
with open('session_cookies.txt', 'r') as file:
    session_cookies = []
    for line in file:
        name, value = line.strip().split('=')
        cookie = {'name': name, 'value': value}
        session_cookies.append(cookie)


browser.execute_script(f"window.localStorage.clear();")
browser.execute_script(f"window.localStorage.setItem('yourDataKey', '{local_storage_data}');")

for cookie in session_cookies:
    browser.add_cookie(cookie)

browser.get('https://app.slack.com/client')
