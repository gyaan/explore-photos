# My project to learn Nodejs as a front end stack and golang as a API server

Technologies used in this project :

1.Angular js 

2.Nodejs 

3.Express for nodejs

4.Golang

5.MongoDB

6.bootstrap for clean UI

Features implemented

1. Single page application using angular js routing.

2. Express routing to call Loggin, Logout, Signup and other APIs from different server

3. Infinite scroll to display images 

4. Nodejs as a backend to consume other place hosted REST APIi  

5. Golang to get the images from flickr using flickr api

6. golang to build REST APIs for other services 

<b>How to execute :<b> 

run below metioned command one by one :

go get github.com/gyaan/explorePhotos/flickr_images

go run $GOPATH/src/github.com/gyaan/explorePhotos/flickr_images/main.go

go get github.com/gyaan/explorePhotos/api

go run $GOPATH/src/github.com/gyaan/explorePhotos/api/main.go  >/dev/null 2>&1 &

cd $HOME

git clone https://github.com/gyaan/explorePhotos.git

cd $HOME/explorePhotos

npm install 

npm start  >/dev/null 2>&1 &
 
open url: http://localhost:3000/	
