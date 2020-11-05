# gin.static

Template to run a static web server as Lyrid serverless.

We are still working on this, but in essence, how it works are as follow:

1. Git clone the code   

2. Copy all your static content into the public folder

3. Run lc code submit to submit your package to Lyrid platform, and create serverless artifacts

4. Web server is accessible by calling:

```
curl -H "Authorization: Basic <$your_access_key_secret>" https://api.lyrid.io/x/gin/static/latest/fnx/
```

- tip: changing app/module name in lyrid-definition before submit will change the URL 

(Optional) run a reverse-proxy that automatically inject your Authorization header on your domain-managed proxy server to be able to use your own domain name. 