from playwright import sync_playwright

with sync_playwright() as p:
    iphone_11 = p.devices['iPhone 11 Pro']
    browser = p.webkit.launch(headless=False)
    context = browser.newContext(
        **iphone_11,
        locale='zh-CN'
    )
    page = context.newPage()
    page.goto('https://www.baidu.com/')
    page.click('#logo')
    page.screenshot(path='colosseum-iphone.png')
    browser.close()