# Yapper
A new twitter like social media site. Start yappin' now!

## How to Run
It is very simple to run, however, some prereqs are needed.

### Step 1. Ensure postgres is installed
This requires the use of postgres. Install it using the instructions for your operating system!

### Step 2. Install go
This is using Go as the majority language. It uses go templates and the Echo framework to generate html server-side.

### Step 3. Install tailwind via npm
This is rather straight forward. Ensure npm is installed and install tailwind
```
npm install -D tailwindcss
```

### Step 4. Run 
Run the server.go file and it should run on localhost:1323
```
go run server.go
```

## Plans
Here is a set of plans which I have for this app

### Docker containers
I plan to use docker containers to manage the backend. I have an algorithm in mind for showing users new yaps

### Yaps and Topics
A yap is simply a post, which contains a limited-sized message, image, or video, along with a topic(s). A topic is like a hashtag on twitter.

### Algorithm
When users post a yap, a topic (or topics) must be included. This will then be used to tell the algorithm what topics a user is interested in.
The more a user yaps about say *fortnite*, then the higher the user scores on that topic; thus, the user will see more yaps of that topic.

### Redis
I plan to use redis to store user sessions (user auth with JWT is already done, however, I am learning how to use redis since redis is more than just user sessions).
I would like to use Redis to cache various activities for users (whatever those activites may be).

### Circles
Circles are like subreddits, where one can write much longer posts and even use markdown. These posts are not affected/seen by the algorithm.
Instead, a user must yap about (aka, share) a post with a topic in order for that post to reach the audience. This means that if someone's post isn't yapped about,
then why should it be shown to anybody?
