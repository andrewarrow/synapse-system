from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.firefox.options import Options
import sys
import time
import json

email = sys.argv[1]

options = Options()
#options.add_argument('-headless')
#options.headless = True
browser = webdriver.Firefox(options=options)
browser.get('https://manypw.slack.com/')

elem = browser.find_element(By.NAME, 'email')
elem.send_keys(email + Keys.RETURN)

time.sleep(26)

local_storage_data = browser.execute_script("return window.localStorage;")
session_cookies = browser.get_cookies()

browser.quit()
print("writing")

json_data = json.dumps(local_storage_data)

with open('local_storage_data.txt', 'w') as file:
    file.write(json_data)

with open('session_cookies.txt', 'w') as file:
    for cookie in session_cookies:
        file.write(f"{cookie['name']}={cookie['value']}\n")

print("done")
