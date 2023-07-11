from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.firefox.options import Options
import sys

email = sys.argv[1]

options = Options()
options.add_argument('-headless')
#options.headless = True
browser = webdriver.Firefox(options=options)
browser.get('https://manypw.slack.com/')

elem = browser.find_element(By.NAME, 'email')
elem.send_keys(email + Keys.RETURN)
