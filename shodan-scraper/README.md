# Shodan Scraper

Being a Security researcher Shodan is a really useful tool for various tasks, especially finding vulnerabilities and disclosing them. I found myself querying, testing and reporting these vulnerabilities all with the help of Shodan.

While Shodan is good in and of itself I've often needed some automation on top of the Shodan search for further investigation, oftentimes confirmation and to get more details for the report I'm writing.

Enter Shodan scraper! A simple tool where my goal is to create easy to use functions and protocol handlers to interact with the hosts I'm investigating. It will mostly be information gathering, but I'm hoping that the functions could be usefull in other bigger projects I'm writing as well as for the reconnaissance part of the security research, where I'm exploring all of the data on Shodan.

- [ ] Implement a concept of projects, where almost everything that's not a simple search query has to use a project
- [ ] Provide stats for a project, especially for the handler and it's actual interaction with the service, error Count etc.
- [x] Ability to search based on query
- [ ] Proxy support
- [ ] Rewrite as a daemon, that can be queried from an API

