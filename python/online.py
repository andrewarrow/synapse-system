from selenium import webdriver
from selenium.webdriver.firefox.options import Options
import time
import sys

path = sys.argv[1]
timeString = sys.argv[2]
options = Options()
options.add_argument("-profile")
options.add_argument(path)
options.add_argument('-headless')
browser = webdriver.Firefox(options=options)
browser.get('https://app.slack.com/client')
time.sleep(int(timeString))
browser.quit()
