from selenium import webdriver
from selenium.webdriver.firefox.options import Options
import time

options = Options()
options.profile = 'jenny'
options.add_argument('-headless')
browser = webdriver.Firefox(options=options)
browser.get('https://app.slack.com/client')
time.sleep(26)
browser.quit()
