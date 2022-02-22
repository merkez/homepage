Data and code for my homepage living at [reisinge.net](https://reisinge.net). Changes inside `data` folder are automatically reflected in the homepage.

When you change the code

1) Test locally: `make run` and go to http://localhost:5001

2) Release to dockerhub: `make release`

NOTE: in case of problems with some missing files (like CSS), exec to homepage container and run `git pull` in /tmp/homepage.